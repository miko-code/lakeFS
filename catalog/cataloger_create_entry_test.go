package catalog

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"

	"github.com/treeverse/lakefs/testutil"
)

func TestCataloger_CreateEntry(t *testing.T) {
	ctx := context.Background()
	c := testCataloger(t)
	// test data
	repo := testCatalogerRepo(t, ctx, c, "repo", "master")
	testutil.MustDo(t, "create entry on master for testing",
		c.CreateEntry(ctx, repo, "master", Entry{Path: "/aaa/bbb/ddd", Checksum: "cc", PhysicalAddress: "xx", Size: 1}))
	testutil.MustDo(t, "create branch b1 based on master",
		c.CreateBranch(ctx, repo, "b1", "master"))

	type args struct {
		repository      string
		branch          string
		path            string
		checksum        string
		physicalAddress string
		size            int64
		metadata        Metadata
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "new",
			args: args{
				repository:      repo,
				branch:          "master",
				path:            "/aaa/bbb/ccc",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"k": "v"},
			},
			wantErr: false,
		},
		{
			name: "new on branch",
			args: args{
				repository:      repo,
				branch:          "b1",
				path:            "/aaa/bbb/ccc",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"k": "v"},
			},
			wantErr: false,
		},
		{
			name: "existing",
			args: args{
				repository:      repo,
				branch:          "master",
				path:            "/aaa/bbb/ddd",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"kk": "vv"},
			},
			wantErr: false,
		},
		{
			name: "unknown repo",
			args: args{
				repository:      "norepo",
				branch:          "master",
				path:            "/aaa/bbb/ccc",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"k": "v"},
			},
			wantErr: true,
		},
		{
			name: "unknown branch",
			args: args{
				repository:      repo,
				branch:          "masterX",
				path:            "/aaa/bbb/ccc",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"k": "v"},
			},
			wantErr: true,
		},
		{
			name: "missing repo",
			args: args{
				repository:      "",
				branch:          "master",
				path:            "/aaa/bbb/ccc",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"k": "v"},
			},
			wantErr: true,
		},
		{
			name: "missing branch",
			args: args{
				repository:      repo,
				branch:          "",
				path:            "/aaa/bbb/ccc",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"k": "v"},
			},
			wantErr: true,
		},
		{
			name: "missing path",
			args: args{
				repository:      repo,
				branch:          "master",
				path:            "",
				checksum:        "1234",
				physicalAddress: "5678",
				size:            100,
				metadata:        Metadata{"k": "v"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.CreateEntry(ctx, tt.args.repository, tt.args.branch, Entry{
				Path:            tt.args.path,
				Checksum:        tt.args.checksum,
				PhysicalAddress: tt.args.physicalAddress,
				Size:            tt.args.size,
				Metadata:        tt.args.metadata,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			// in case there is no error - get the entry and compare
			ref := MakeReference(tt.args.branch, UncommittedID)
			ent, err := c.GetEntry(ctx, tt.args.repository, ref, tt.args.path)
			testutil.MustDo(t, "get entry we just created", err)
			if ent.Size != tt.args.size {
				t.Fatalf("entry size %d, expected %d", ent.Size, tt.args.size)
			}
			if ent.PhysicalAddress != tt.args.physicalAddress {
				t.Fatalf("entry physical address %s, expected %s", ent.PhysicalAddress, tt.args.physicalAddress)
			}
			if ent.Checksum != tt.args.checksum {
				t.Fatalf("entry checksum %s, expected %s", ent.Checksum, tt.args.checksum)
			}
			if !reflect.DeepEqual(ent.Metadata, tt.args.metadata) {
				t.Fatalf("entry metadata %+v, expected %+v", ent.Metadata, tt.args.metadata)
			}
		})
	}
}

func randomFilepath(basename string) string {
	var sb strings.Builder
	depth := rand.Intn(10)
	for i := 0; i < depth; i++ {
		level := fmt.Sprintf("dir%d/", rand.Intn(3))
		sb.WriteString(level)
	}
	sb.WriteString(basename)
	return sb.String()
}

func BenchmarkCataloger_CreateEntry(b *testing.B) {
	ctx := context.Background()
	c := testCataloger(b)
	repo := testCatalogerRepo(b, ctx, c, "repo", "master")
	for i := 0; i < b.N; i++ {
		entPath := randomFilepath("test_entry")
		testCatalogerCreateEntry(b, ctx, c, repo, "master", entPath, nil, "")
	}
}
