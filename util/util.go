package util

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

// DoStuff does stuff.
func DoStuff() {
	fmt.Println("grouping files…")
}

// GetExtension returns extension of a file.
func GetExtension(filename string) string {
	abs, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}
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
