package filestorage

import (
	"bytes"
	"os"
)

func New(pathToFile string) *FileStorage {
	return &FileStorage{pathToFile: pathToFile}
}

type FileStorage struct {
	pathToFile string
}

func (fs *FileStorage) SaveImage(imageData bytes.Buffer, fileName string) error {

	file, err := os.OpenFile(fs.pathToFile+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := imageData.WriteTo(file); err != nil {
		return err
	}
	return nil
}