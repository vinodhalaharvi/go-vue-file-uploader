package main

type FileMetadata struct {
	Name         string `json:"name" binding:"required"`
	AbsolutePath string `json:"absolutePath" binding:"required"`
}

type FileRepository interface {
	Save(file FileMetadata) error
	GetAll() ([]FileMetadata, error)
	Delete(filename string) error
}
