package bucket

import "errors"

// ErrNotFound represents the error returned when an object is not found in the storage bucket.
// This error is used to indicate that the requested object does not exist in the bucket.
var ErrNotFound = errors.New("bucket: not found")

// ErrFailedToUpload represents the error returned when an object upload operation fails.
// This error is used to indicate that the object could not be uploaded to the storage bucket.
var ErrFailedToUpload = errors.New("bucket: failed to upload")

// ErrFailedToDownload represents the error returned when an object download operation fails.
// This error is used to indicate that the object could not be downloaded from the storage bucket.
var ErrFailedToDownload = errors.New("bucket: failed to download")

// ErrFailedToStats represents the error returned when an object stats operation fails.
// This error is used to indicate that the metadata of the object could not be retrieved from the storage bucket.
var ErrFailedToStats = errors.New("bucket: failed to get stats")
