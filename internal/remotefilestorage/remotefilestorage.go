// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package remotefilestorage

import (
	"context"
)

// RemoteFileStorage interface to call s3 with minio.
type RemoteFileStorage interface {
	Upload(ctx context.Context, bucket, objectName string, data []byte) error
	Download(ctx context.Context, bucket, objectName string, buffer []byte) ([]byte, error)
}
