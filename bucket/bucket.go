package bucket

import (
	"context"
	"io"
)

// Bucket represents a storage bucket interface.
// It defines methods for interacting with a storage system,
// such as uploading, downloading, and retrieving metadata of objects.
type Bucket interface {
	// PutObject uploads an object to the storage bucket.
	// It accepts a context, the name of the object, a reader for the object data,
	// and the length of the object data. It returns an error if the operation fails.
	PutObject(ctx context.Context, objectName string, reader io.Reader, readerLen int64) error

	// GetObject downloads an object from the storage bucket.
	// It accepts a context and the name of the object. It returns a reader for the object data
	// and any error encountered during the operation.
	GetObject(ctx context.Context, objectName string) (io.ReadCloser, error)

	// Stats retrieves the metadata of an object in the storage bucket.
	// It accepts a context and the name of the object. It returns the object's metadata
	// and any error encountered during the operation.
	Stats(ctx context.Context, objectName string) (*Stats, error)
}
