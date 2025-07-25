package multidownload

import "testing"

func TestDownload(t *testing.T) {
	files := []File{
		{"https://example.com/file1.txt", "/tmp/file1.txt"},
	}

	err := Download(files)
	if err != nil {
		t.Fatal(err)
	}
}
