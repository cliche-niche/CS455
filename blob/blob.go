package blob

import (
	"os"
	"time"
)

// Blob type
type Blob struct {
	name      string
	blobtype  string
	location  string
	contents  string
	fptr      *os.File
	updatedAt time.Time
}

// Store interface
type Store interface {
	InitBlob(name string, blobtype string, location string) error
	CloseBlob() error
	EditBlob(updatedContent string) error
	SaveBlob() error
	DeleteBlob() error
}

// InitBlob opens a file for editing, creates a new file if it doesn't exist
func (b *Blob) InitBlob(name string, blobtype string, location string) error {
	b.name = name
	b.blobtype = blobtype
	b.location = location

	var err error
	if b.fptr, err = os.OpenFile(b.location+"/"+b.name, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		return err
	}

	b.updatedAt = time.Now()

	// read contents of file into b.contents
	b.fptr.Read([]byte(b.contents))

	return nil
}

// CloseBlob closes the file pointer
func (b *Blob) CloseBlob() error {
	return b.fptr.Close()
}

// EditBlob updates the contents of the blob
func (b *Blob) EditBlob(updatedContent string) error {
	b.contents = updatedContent
	b.updatedAt = time.Now()

	// b.fptr.Write([]byte(b.contents)) // maybe do this only on save
	return nil
}

// SaveBlob saves the contents of the blob to the file
func (b *Blob) SaveBlob() error {
	b.updatedAt = time.Now()
	b.fptr.Write([]byte(b.contents))
	return nil
}

// DeleteBlob closes the file pointer and deletes the file
func (b *Blob) DeleteBlob() error {
	if err := b.fptr.Close(); err != nil {
		return err
	}

	return os.Remove(b.location + "/" + b.name)
}
