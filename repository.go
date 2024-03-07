package main

type FileMetadata struct {
	Name         string
	AbsolutePath string
}

type FileRepository interface {
	Save(file FileMetadata) error
	GetAll() ([]FileMetadata, error)
	Delete(filename string) error
}
