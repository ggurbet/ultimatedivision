// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package mock

import "context"

// Mock mocked struct for remote file storage call.
type Mock struct{}

// Upload mock for upload method.
func (mock *Mock) Upload(ctx context.Context, bucket, objectName string, data []byte) error {
	return nil
}

// Download mock for download method.
func (mock *Mock) Download(ctx context.Context, bucket, objectName string, buffer []byte) ([]byte, error) {
	return nil, nil
}
