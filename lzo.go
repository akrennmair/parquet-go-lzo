package lzo

import (
	"bytes"
	"errors"
	"io"

	"github.com/cyberdelia/lzo"
	goparquet "github.com/fraugster/parquet-go"
	"github.com/fraugster/parquet-go/parquet"
)

type LZOBlockCompressor struct{}

func NewLZOBlockCompressor() *LZOBlockCompressor {
	return &LZOBlockCompressor{}
}

func (c *LZOBlockCompressor) CompressBlock(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	w := lzo.NewWriter(buf)
	n, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	if n < len(data) {
		return nil, errors.New("short write")
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *LZOBlockCompressor) DecompressBlock(data []byte) ([]byte, error) {
	r, err := lzo.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

func init() {
	goparquet.RegisterBlockCompressor(parquet.CompressionCodec_LZO, NewLZOBlockCompressor())
}
