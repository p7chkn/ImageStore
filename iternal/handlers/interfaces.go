package handlers

import "bytes"

type RepositoryInterfaceHttp interface {
	GetImage(imageName string) ([]byte, error)
}

type RepositoryInterfaceGrpc interface {
	SaveImage(imageData bytes.Buffer, title string) error
	DeleteImage(fileName string) error
}
