package handlers

import "bytes"

type RepositoryInterface interface {
	SaveImage(imageData bytes.Buffer, title string) error
}
