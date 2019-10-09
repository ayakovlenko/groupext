package util

import (
	"io/ioutil"
	"testing"
)

func TestGetExtension(t *testing.T) {

	t.Run("normal filename", func(t *testing.T) {
		assertString(t,
			GetExtension("whatever/index.js"),
			"js",
		)
	})

	t.Run("filename with dots", func(t *testing.T) {
		assertString(t,
			GetExtension("whatever/main.test.js"),
			"js",
		)
	})

	t.Run("hidden file with extension", func(t *testing.T) {
		assertString(t,
			GetExtension("whatever/.prettierrc.js"),
			"",
		)
	})

	t.Run("hidden file without extension", func(t *testing.T) {
		assertString(t,
			GetExtension("whatever/.gitignore"),
			"",
		)
	})
}

func TestNewName(t *testing.T) {

	t.Run("first name collision", func(t *testing.T) {
		assertString(t,
			NewName("filename.ext"),
			"filename copy.ext",
		)
	})

	t.Run("second name collision", func(t *testing.T) {
		assertString(t,
			NewName("filename copy.ext"),
			"filename copy 2.ext",
		)
	})

	t.Run("third name collision", func(t *testing.T) {
		assertString(t,
			NewName("filename copy 2.ext"),
			"filename copy 3.ext",
		)
	})
}

func TestFileExists(t *testing.T) {

	t.Run("file does not exist", func(t *testing.T) {
		have := fileExists("3fd7kht7xz8gvshr.txt")
		want := false

		if have != want {
			t.Fatalf("expected file to not exist")
		}
	})

	t.Run("file exists", func(t *testing.T) {
		tmpfile, _ := ioutil.TempFile("", "example.*.txt")

		have := fileExists(tmpfile.Name())
		want := true

		if have != want {
			t.Fatalf("expected file to not exist")
		}
	})
}

func assertString(t *testing.T, have, want string) {
	t.Helper()

	if have != want {
		t.Fatalf("have: %s; want: %s", have, want)
	}
}
