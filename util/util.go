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

// DoStuff does stuff.
func DoStuff(dirpath string) {
	files, err := ioutil.ReadDir(dirpath)
	check(err)

	for _, f := range files {
		if f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			continue
		}

		filename := filepath.Join(dirpath, f.Name())
		extDir := GetExtension(filename)
		_ = os.Mkdir(filepath.Join(dirpath, extDir), os.ModePerm)

		Move(
			filename,
			filepath.Join(filepath.Dir(filename), extDir, filepath.Base(filename)),
		)
	}
}

// Move moves a file from old location to new location.
func Move(oldName, newName string) {
	if _, err := os.Stat(newName); !os.IsNotExist(err) {
		Move(oldName, NewName(newName))
		return
	}

	err := os.Rename(oldName, newName)
	check(err)
	fmt.Printf("%s -> %s\n", oldName, newName) // TODO: if verbose
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
	pat, err := regexp.Compile(`(.+) \((\d+)\)\.(.+)`)
	check(err)
	res := pat.FindAllStringSubmatch(filename, -1)

	if len(res) == 0 {
		li := strings.LastIndex(filename, ".")
		return fmt.Sprintf("%s (2)%s", filename[:li], filename[li:])
	}

	base := res[0][1]
	attempt, err := strconv.Atoi(res[0][2])
	check(err)
	ext := res[0][3]
	return fmt.Sprintf("%s (%d).%s", base, attempt+1, ext)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
