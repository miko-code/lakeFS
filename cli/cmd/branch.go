package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-openapi/swag"

	"github.com/treeverse/lakefs/api/gen/models"

	"github.com/treeverse/lakefs/uri"

	"golang.org/x/xerrors"

	"github.com/jedib0t/go-pretty/table"

	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "create and manage branches within a repository",
	Long:  `Create delete and list branches within a lakeFS repository`,
}

var branchListCmd = &cobra.Command{
	Use:     "list [repository uri]",
	Short:   "list branches in a repository",
	Example: "lakectl branch list lakefs://myrepo",
	Args: ValidationChain(
		HasNArgs(1),
		IsRepoURI(0),
	),
	RunE: func(cmd *cobra.Command, args []string) error {

		amount, _ := cmd.Flags().GetInt("amount")
		after, _ := cmd.Flags().GetString("after")

		u := uri.Must(uri.Parse(args[0]))
		client, err := getClient()
		if err != nil {
			return err
		}
		response, pagination, err := client.ListBranches(context.Background(), u.Repository, after, amount)
		if err != nil {
			return err
		}

		// generate table
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Branch Name", "Commit ID"})
		for _, branch := range response {
			t.AppendRow([]interface{}{*branch.ID, *branch.CommitID})
		}
		t.Render()

		if pagination != nil && swag.BoolValue(pagination.HasMore) {
			fmt.Printf("for more results, run with `--amount %d --after \"%s\"\n", amount, pagination.NextOffset)
		}

		// done
		return nil
	},
}

var branchCreateCmd = &cobra.Command{
	Use:   "create [branch uri]",
	Short: "create a new branch in a repository",
	Args: ValidationChain(
		HasNArgs(1),
		IsBranchURI(0),
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		u := uri.Must(uri.Parse(args[0]))
		client, err := getClient()
		if err != nil {
			return err
		}
		sourceRawUri, _ := cmd.Flags().GetString("source")
		suri, err := uri.Parse(sourceRawUri)
		if err != nil {
			return fmt.Errorf("failed to parse source URI: %s", err)
		}
		if !strings.EqualFold(suri.Repository, u.Repository) {
			return fmt.Errorf("source branch must be in the same repository")
		}

		sourceBranch, err := client.GetBranch(context.Background(), u.Repository, suri.Refspec)
		if err != nil {
			return xerrors.Errorf("could not get source branch: %w", err)
		}
		fmt.Printf("got source branch '%s', using commit '%s'\n",
			suri.Refspec, *sourceBranch.CommitID)
		err = client.CreateBranch(context.Background(), u.Repository, &models.Refspec{
			CommitID: sourceBranch.CommitID,
			ID:       swag.String(u.Refspec),
		})
		return err
	},
}

var branchDeleteCmd = &cobra.Command{
	Use:   "delete [branch uri]",
	Short: "delete a branch in a repository, along with its uncommitted changes (CAREFUL)",
	Args: ValidationChain(
		HasNArgs(1),
		IsBranchURI(0),
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		sure, err := cmd.Flags().GetBool("sure")
		if err != nil || !sure {
			confirmation, err := confirm("Are you sure you want to delete branch")
			if err != nil || !confirmation {
				return fmt.Errorf("please confirm by passing the --sure | -y flag")
			}
		}
		client, err := getClient()
		if err != nil {
			return err
		}
		u := uri.Must(uri.Parse(args[0]))
		err = client.DeleteBranch(context.Background(), u.Repository, u.Refspec)
		return err
	},
}

// lakectl branch revert lakefs://myrepo@master --commit commitId --path path --object path-to-object
var branchRevertCmd = &cobra.Command{
	Use:   "revert [branch uri] [flags]",
	Short: "revert changes to specified commit, or revert uncommitted changes all or by path ",
	Args: ValidationChain(
		HasNArgs(1),
		IsBranchURI(0),
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		clt, err := getClient()
		if err != nil {
			return err
		}
		u := uri.Must(uri.Parse(args[0]))

		commitId, err := cmd.Flags().GetString("commit")
		if err != nil {
			return err
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}
		isObject, err := cmd.Flags().GetBool("object")
		if err != nil {
			return err
		}
		isCommit := len(commitId) > 0

		isPath := len(path) > 0

		if isCommit && isPath {
			return xerrors.Errorf("can't revert by path and commit, please choose only one!")
		}

		var revert models.Revert
		var confirmationMsg string
		if isCommit {
			confirmationMsg = fmt.Sprintf("Are you sure you want to revert all changes to commit: %s ?", commitId)
			revert = models.Revert{
				Commit: commitId,
				Type:   models.RevertTypeCOMMIT,
			}
		} else if isPath {
			var modelType string
			if isObject {
				modelType = models.RevertTypeOBJECT
				confirmationMsg = fmt.Sprintf("Are you sure you want to revert all changes for object: %s to last commit?", path)
			} else {
				confirmationMsg = fmt.Sprintf("Are you sure you want to revert all changes from path: %s to last commit?", path)
				modelType = models.RevertTypePATH
			}
			revert = models.Revert{
				Path: path,
				Type: modelType,
			}
		} else {
			confirmationMsg = "are you sure you want all uncommitted changes?"
			revert = models.Revert{
				Type: models.RevertTypeRESET,
			}
		}

		confirmation, err := confirm(confirmationMsg)
		if err != nil || !confirmation {
			fmt.Println("Revert Aborted")
			return nil
		}
		err = clt.RevertBranch(context.Background(), u.Repository, u.Refspec, &revert)
		return err
	},
}

var branchShowCmd = &cobra.Command{
	Use:   "show [branch uri]",
	Short: "show branch metadata",
	Args: ValidationChain(
		HasNArgs(1),
		IsBranchURI(0),
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getClient()
		if err != nil {
			return err
		}
		u := uri.Must(uri.Parse(args[0]))
		resp, err := client.GetBranch(context.Background(), u.Repository, u.Refspec)
		if err != nil {
			return err
		}
		fmt.Printf("%s\t%s\n", *resp.ID, *resp.CommitID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.AddCommand(branchCreateCmd)
	branchCmd.AddCommand(branchDeleteCmd)
	branchCmd.AddCommand(branchListCmd)
	branchCmd.AddCommand(branchShowCmd)
	branchCmd.AddCommand(branchRevertCmd)

	branchListCmd.Flags().Int("amount", -1, "how many results to return, or-1 for all results (used for pagination)")
	branchListCmd.Flags().String("after", "", "show results after this value (used for pagination)")

	branchCreateCmd.Flags().StringP("source", "s", "", "source branch uri")
	_ = branchCreateCmd.MarkFlagRequired("source")

	branchDeleteCmd.Flags().BoolP("sure", "y", false, "do not ask for confirmation")

	branchRevertCmd.Flags().StringP("commit", "c", "", "commit ID to revert branch to ")
	branchRevertCmd.Flags().StringP("path", "p", "", "path to revert it's changes")
	branchRevertCmd.Flags().BoolP("object", "o", false, "path to revert it's changes")
}
