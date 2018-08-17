// +build !containers_image_include_torrent

package torrent

import (
	"context"
	"errors"
	"io"

	"github.com/containers/image/types"
	"github.com/containers/image/docker/reference"
)

type TorrentClient struct {
}

func MakeTorrentClient(sys *types.SystemContext, debug bool) (*TorrentClient, error) {
	return nil, errors.New("BitTorrent not supported")
}

func (t *TorrentClient) GetBlobTorrent(ctx context.Context, info types.BlobInfo, registry string, ref reference.Named, trackers []string) (io.ReadCloser, int64, error) {
	return nil, -1, errors.New("BitTorrent not supported")
}

func (t *TorrentClient) Close() {
}

func (t *TorrentClient) Seed(ctx context.Context, srcCtx *types.SystemContext, ref types.ImageReference, refSrc types.ImageReference) error {
	return errors.New("BitTorrent not supported")
}

func (t *TorrentClient) WriteStatus(w io.Writer) {
}
