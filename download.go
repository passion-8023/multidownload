package multidownload

import (
	"fmt"
)

// File represents a download target
type File struct {
	URL  string
	Path string
}

// Download downloads all given files.
func Download(files []File) error {
	for _, file := range files {
		fmt.Printf("Downloading %s -> %s\n", file.URL, file.Path)
		// 这里只是演示，真实下载逻辑另写
	}
	return nil
}
