package util

import "testing"

func TestGetExtension(t *testing.T) {

	assertExt := func(t *testing.T, filename, want string) {
		t.Helper()

		have := GetExtension(filename)

		if have != want {
			t.Fatalf("have: %s; want: %s", have, want)
		}
	}

	t.Run("normal filename", func(t *testing.T) {
		assertExt(t, "whatever/index.js", "js")
	})

	t.Run("filename with dots", func(t *testing.T) {
		assertExt(t, "whatever/main.test.js", "js")
	})

	t.Run("hidden file with extension", func(t *testing.T) {
		assertExt(t, "whatever/.prettierrc.js", "")
	})

	t.Run("hidden file without extension", func(t *testing.T) {
		assertExt(t, "whatever/.gitignore", "")
	})
}
