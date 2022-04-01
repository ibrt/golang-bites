package gzipz

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/ibrt/golang-bites/internal"
)

// MustCompress implements in-memory GZIP compression.
func MustCompress(buf []byte) []byte {
	cBuf := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(cBuf, gzip.BestCompression)
	internal.MaybePanic(err)
	defer func() {
		internal.MaybePanic(w.Close())
	}()

	_, err = w.Write(buf)
	internal.MaybePanic(err)
	internal.MaybePanic(w.Close())

	return cBuf.Bytes()
}

// MustDecompress implements in-memory GZIP decompression.
func MustDecompress(buf []byte) []byte {
	r, err := gzip.NewReader(bytes.NewReader(buf))
	internal.MaybePanic(err)
	defer func() {
		internal.MaybePanic(r.Close())
	}()

	dBuf, err := ioutil.ReadAll(r)
	internal.MaybePanic(err)
	return dBuf
}
