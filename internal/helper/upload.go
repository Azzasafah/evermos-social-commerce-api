package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveUploadedFile(file *multipart.FileHeader, destDir string) (string, error) {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", err
	}

	cleanFileName := filepath.Base(file.Filename)
	cleanFileName = strings.ReplaceAll(cleanFileName, " ", "-")
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), cleanFileName)
	targetPath := filepath.Join(destDir, filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(targetPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, src); err != nil {
		return "", err
	}

	return filename, nil
}
