package bucket

import "time"

// Stats represents the metadata of an object in the storage bucket.
// It contains information about the total number of objects, the size of the object,
// the content type of the object, and the last modified time of the object.
type Stats struct {
	// Total represents the total number of objects in the storage bucket.
	Total int64 `json:"total" yaml:"total"`
	// Size represents the size of the object in the storage bucket.
	Size int64 `json:"size" yaml:"size"`
	// ContentType represents the content type of the object in the storage bucket.
	ContentType string `json:"contentType" yaml:"contentType"`
	// LastModified represents the last modified time of the object in the storage bucket.
	LastModified time.Time `json:"lastModified" yaml:"lastModified"`
}
