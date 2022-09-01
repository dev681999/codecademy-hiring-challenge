package file

import (
	"io"
	"os"
)

type Service interface {
	Create(path string, src io.Reader) error
	Delete(path string) error
}

func NewDiskService() Service {
	return &diskService{}
}

type diskService struct {
}

func (s *diskService) Create(path string, src io.Reader) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, src)
	return err
}

func (s *diskService) Delete(path string) error {
	return os.Remove(path)
}
