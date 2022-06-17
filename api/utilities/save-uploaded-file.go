package utilities

import (
	"mime/multipart"
	"os"
)

func SaveUploadedFile(file *multipart.FileHeader, path string) error {
	// remove existing file if any
	_ = os.Remove(path)

	// create new file
	_, err := os.Create(path)
	if err != nil {
		return err
	}

	return nil
}
