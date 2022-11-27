// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package servertesting

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/internal/logger/zaplog"
	"ultimatedivision/pkg/fileutils"
)

// Run will create and run ultimatedivision representation, initialize testdb, run tests and close peer after.
func Run(t *testing.T, test func(ctx context.Context, t *testing.T, peer *ultimatedivision.Peer)) {
	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		log := zaplog.NewLog()

		var group errgroup.Group

		defaultConfigDir := fileutils.ApplicationDir(filepath.Join("ultimatedivision", "servertesting"))

		config, err := readConfig(defaultConfigDir)
		assert.NoError(t, err)

		peer, err := ultimatedivision.New(log, config, db)
		assert.NoError(t, err)
		defer func() {
			err = errs.Combine(peer.Close())
		}()

		group.Go(func() error {
			err = peer.Run(ctx)
			require.NoError(t, err)
			return nil
		})

		test(ctx, t, peer)
	})
}

// readConfig reads config from default config dir.
func readConfig(defaultConfigDir string) (config ultimatedivision.Config, err error) {
	configBytes, err := os.ReadFile(path.Join(defaultConfigDir, "config.json"))
	if err != nil {
		return ultimatedivision.Config{}, err
	}

	return config, json.Unmarshal(configBytes, &config)
}
