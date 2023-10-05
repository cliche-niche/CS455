package blob

import (
	"os"
	"testing"
)

func TestBlob(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := "./test_data"
	os.MkdirAll(tempDir, os.ModePerm)
	defer os.RemoveAll(tempDir)

	// Initialize a Blob instance for testing
	blob := &Blob{}
	err := blob.InitBlob("test_blob.txt", "txt", tempDir)
	if err != nil {
		t.Fatalf("InitBlob failed: %v", err)
	}
	defer blob.CloseBlob()

	// Test EditBlob and SaveBlob methods
	content := "Hello, world!"
	err = blob.EditBlob(content)
	if err != nil {
		t.Fatalf("EditBlob failed: %v", err)
	}

	err = blob.SaveBlob()
	if err != nil {
		t.Fatalf("SaveBlob failed: %v", err)
	}

	// Read the saved content from the file
	savedContent, err := os.ReadFile(tempDir + "/test_blob.txt")
	if err != nil {
		t.Fatalf("Failed to read saved content: %v", err)
	}

	if string(savedContent) != content {
		t.Fatalf("Saved content does not match expected content")
	}

	// Test DeleteBlob method, free ptr then remove file
	err = blob.DeleteBlob()
	if err != nil {
		t.Fatalf("DeleteBlob failed: %v", err)
	}

	// Check if the file has been deleted
	_, err = os.Stat(tempDir + "/test_blob.txt")
	if !os.IsNotExist(err) {
		t.Fatalf("File should have been deleted, but still exists")
	}
}
