package catalog

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/treeverse/lakefs/db"
)

const ListEntriesByLevelMaxLimit = 1000

func (c *cataloger) ListEntriesByLevel(ctx context.Context, repository, reference, prefix, after, delimiter string, limit int) ([]LevelEntryResult, bool, error) {
	if limit < 0 || limit > ListEntriesByLevelMaxLimit {
		limit = ListEntriesByLevelMaxLimit
	}
	if err := Validate(ValidateFields{
		{Name: "repository", IsValid: ValidateRepositoryName(repository)},
		{Name: "reference", IsValid: ValidateReference(reference)},
	}); err != nil {
		return nil, false, err
	}
	ref, err := ParseRef(reference)
	if err != nil {
		return nil, false, err
	}
	branchName := ref.Branch
	commitID := ref.CommitID
	markers, err := c.db.Transact(func(tx db.Tx) (interface{}, error) {
		branchID, err := getBranchID(tx, repository, branchName, LockTypeNone)
		if err != nil {
			return nil, fmt.Errorf(" get branch ID failed: %w", err)
		}
		lineage, err := getLineage(tx, branchID, commitID)
		if err != nil {
			return nil, fmt.Errorf("get lineage failed: %w", err)
		}
		prefixQuery := sqListByPrefix(prefix, after, delimiter, branchID, limit+1, commitID, lineage)
		sql, args, err := prefixQuery.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			return nil, fmt.Errorf("list by level ToSql failed : %w", err)
		}
		var markerList []LevelEntryResult
		err = tx.Select(&markerList, sql, args...)
		if err != nil {
			return nil, fmt.Errorf("list by level query failed : %w", err)
		}
		return retrieveEntries(tx, markerList, branchID, commitID, lineage, prefix)
	}, c.txOpts(ctx, db.ReadOnly())...)
	if markers == nil || err != nil {
		return nil, false, err
	}
	result := markers.([]LevelEntryResult)
	moreToRead := paginateSlice(&result, limit)
	return markers.([]LevelEntryResult), moreToRead, nil
}

func retrieveEntries(tx db.Tx, markerList []LevelEntryResult, branchID int64, commitID CommitID, lineage []lineageCommit, prefix string) (interface{}, error) {
	type entryRun struct {
		startRunIndex, runLength   int
		startEntryRun, endEntryRun string
	}
	var entryRuns []entryRun
	var inRun bool
	var previousInRun string
	var run entryRun
	for i := range markerList {
		p := markerList[i].Path
		if len(p) > 0 {
			if strings.HasSuffix(p, DirectoryTerminationChar) { // remove termination character, if present
				p = p[:len(p)-len(DirectoryTerminationChar)]
				markerList[i].Path = p
			}
		}
		if p[len(p)-1] == "/"[0] { // terminating by '/'(slash) character is an indication of a directory
			// its absence indicates a leaf entry that has to be read from DB
			if inRun {
				inRun = false
				run.endEntryRun = previousInRun
				entryRuns = append(entryRuns, run)
			}
		} else { // an entry
			previousInRun = p
			if !inRun {
				inRun = true
				run.startEntryRun = p
				run.runLength = 1
				run.startRunIndex = i
			} else {
				run.runLength++
			}
		}

	}
	if inRun {
		run.endEntryRun = previousInRun
		entryRuns = append(entryRuns, run)
	}
	entriesReader := sqEntriesLineageV(branchID, commitID, lineage)
	for _, r := range entryRuns {
		entriesList := make([]Entry, 0)
		rangeReader := sq.Select("path", "physical_address", "creation_date", "size", "checksum", "metadata").
			Where("path between ? and ?", prefix+r.startEntryRun, prefix+r.endEntryRun).FromSelect(entriesReader, "e")
		sql, args, err := rangeReader.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			return nil, fmt.Errorf("rangeReader ToSql failed : %w", err)
		}
		err = tx.Select(&entriesList, sql, args...)
		if err != nil {
			return nil, fmt.Errorf("reading entries failed : %w", err)
		}
		if len(entriesList) != r.runLength {
			errStr := fmt.Sprintf("expecte to read %d entries, got %d", r.runLength, len(entriesList)) + " : %w"
			return nil, fmt.Errorf(errStr, err)
		}
		for i := 0; i < r.runLength; i++ {
			markerList[r.startRunIndex+i].Entry = &entriesList[i]
		}
	}
	return markerList, nil
}
