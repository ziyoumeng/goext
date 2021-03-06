// Copyright 2016 ~ 2018 AlexStocks(https://github.com/AlexStocks).
// All rights reserved.  Use of this source code is
// governed by Apache License 2.0.

// Package gxfilepath encapsulates some path.filepath functions
// ref: https://github.com/rogpeppe/go-internal/blob/master/dirhash/hash.go
package gxfilepath

import (
	"os"
	"path/filepath"
	"strings"
)

// Get the fiels in @dir, and then put @prefix as all filenames's prefix
func DirFiles(dir, prefix string) ([]string, error) {
	var files []string
	dir = filepath.Clean(dir)
	err := filepath.Walk(dir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel := file
		if dir != "." {
			rel = file[len(dir)+1:]
		}
		f := filepath.Join(prefix, rel)
		files = append(files, filepath.ToSlash(f))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func Mkdirf(dir string) error {
	// the owner perm must be 7
	if err := os.MkdirAll(dir, 0766); err != nil &&
		!strings.Contains(err.Error(), "file exists") {
		return err
	}

	return nil
}

func Rmdirf(dir string) error {
	// the owner perm must be 7
	if err := os.RemoveAll(dir); err != nil {
		return err
	}

	return nil
}
