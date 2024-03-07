package main

import (
	"fmt"
	"os"
)

type InMemoryFileRepository struct {
	files []FileMetadata
}

func (repo *InMemoryFileRepository) Delete(filename string) error {
	for i, file := range repo.files {
		if file.Name == filename {
			// Delete the file from the filesystem
			err := os.Remove(file.AbsolutePath)
			if err != nil {
				return fmt.Errorf("unable to delete file %s from filesystem: %v", filename, err)
			}

			// Delete the file from the in-memory database
			repo.files = append(repo.files[:i], repo.files[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("file %s not found", filename)
}

func NewInMemoryFileRepository() FileRepository {
	return &InMemoryFileRepository{
		files: []FileMetadata{},
	}
}

func (repo *InMemoryFileRepository) Save(file FileMetadata) error {
	repo.files = append(repo.files, file)
	return nil // In a real implementation, you might return errors if any
}

func (repo *InMemoryFileRepository) GetAll() ([]FileMetadata, error) {
	return repo.files, nil
}
