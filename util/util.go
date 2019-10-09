package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const verbose = true // TODO: command-line flag

// DoStuff does stuff.
func DoStuff(dirpath string) {
	files, err := ioutil.ReadDir(dirpath)
	check(err)

	for _, f := range files {
		if f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			continue
		}

		filename := filepath.Join(dirpath, f.Name())
		extDir := filepath.Join(dirpath, GetExtension(filename))

		if !fileExists(extDir) {
			err = os.Mkdir(extDir, os.ModePerm)
			check(err)

			if verbose {
				fmt.Printf("mkdir %q\n", extDir)
			}
		}

		Move(
			filename,
			filepath.Join(extDir, filepath.Base(filename)),
		)
	}
}

// Move moves a file from old location to new location.
func Move(oldName, newName string) {
	if fileExists(newName) {
		Move(oldName, NewName(newName))
		return
	}

	err := os.Rename(oldName, newName)
	check(err)

	if verbose {
		fmt.Printf("mv %q %q\n", oldName, newName)
	}
}

// GetExtension returns extension of a file.
func GetExtension(filename string) string {
	abs, err := filepath.Abs(filename)
	check(err)
	base := filepath.Base(abs)

	if strings.HasPrefix(base, ".") {
		return ""
	}

	ext := filepath.Ext(base)

	if len(ext) > 0 {
		ext = ext[1:]
	}

	return ext
}

// NewName generates a new name for a colliding name.
func NewName(filename string) string {
	pat, err := regexp.Compile(`(.+) copy( (\d+))?\.(.+)`)
	check(err)
	res := pat.FindAllStringSubmatch(filename, -1)

	if len(res) == 0 {
		li := strings.LastIndex(filename, ".")
		return fmt.Sprintf("%s copy%s", filename[:li], filename[li:])
	}

	base := res[0][1]
	ext := res[0][4]
	attempt := res[0][3]
	if attempt == "" {
		return fmt.Sprintf("%s copy 2.%s", base, ext)
	}
	nextAttempt, err := strconv.Atoi(attempt)
	check(err)
	return fmt.Sprintf("%s copy %d.%s", base, nextAttempt+1, ext)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
