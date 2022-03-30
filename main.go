package main

import (
	"context"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-tika/tika"
)

func main() {
	var root string
	var err error
	if len(os.Args) > 1 {
		root, err = filepath.Abs(os.Args[1])
	} else {
		root, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	cl := tika.NewClient(&http.Client{}, "http://localhost:9998")
	ctx := context.Background()
	fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
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
			meta, err := cl.MetaRecursive(ctx, f)
			if err != nil {
				return err
			}
			for _, m := range meta {
				log.Println("metadata:", strings.Join(keys(m), ","))
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
}

func keys[T any](m map[string]T) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
