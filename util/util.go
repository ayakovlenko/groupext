package util

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// DoStuff does stuff.
func DoStuff() {
	fmt.Println("grouping files…")
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
