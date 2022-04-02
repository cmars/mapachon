package main

import (
	"context"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-tika/tika"
	_ "github.com/mattn/go-sqlite3"

	"github.com/cmars/mapachon/ent"
)

var (
	tikaURL = flag.String("tika", "http://localhost:9998", "Apache Tika API")
	out     = flag.String("out", "test.db", "output database file")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	var root string
	var err error
	if len(flag.Args()) > 0 {
		root, err = filepath.Abs(flag.Arg(0))
	} else {
		root, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("extracting from", root)
	cl := tika.NewClient(&http.Client{}, *tikaURL)
	db, err := ent.Open("sqlite3", *out+"?_fk=1") //"file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	err = fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
		path = filepath.Join(root, path)
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			return nil
		}
		log.Println(path)
		err = func() error {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			metas, err := cl.MetaRecursive(ctx, f)
			if err != nil {
				return err
			}
			for _, meta := range metas {
				for i, content := range meta["X-TIKA:content"] {
					if content == "" {
						continue
					}
					artifactID, err := db.Artifact.Create().
						SetFilePath(path).
						SetArchivePath(getMeta(meta, i, "X-TIKA:embedded_resource_path")).
						SetFileType(getMeta(meta, i, "Content-Type")).
						SetContent(content).
						OnConflict().
						UpdateNewValues().
						ID(ctx)
					if err != nil {
						return err
					}
					for k, v := range meta {
						if k == "X-TIKA:content" {
							// Do not store the actual content as metadata
							continue
						}
						err := db.Metadata.Create().
							SetArtifactID(artifactID).
							SetKey(k).
							SetValue(v[i]).
							OnConflict().
							UpdateNewValues().
							Exec(ctx)
						if err != nil {
							return err
						}
					}
				}
			}
			return nil
		}()
		if err != nil {
			log.Println("warning:", err)
		}
		log.Println()
		// TODO: read the file and archive contents / metadata
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func getMeta(meta map[string][]string, i int, keys ...string) string {
	for _, key := range keys {
		if len(meta[key]) > i {
			return meta[key][i]
		}
	}
	return ""
}

func keys[T any](m map[string]T) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
