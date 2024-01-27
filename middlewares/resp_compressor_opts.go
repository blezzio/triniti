package middlewares

import (
	"compress/flate"
	"compress/gzip"
	"io"
)

type RespCompressorOpt func(*RespCompressor)

func WithDeflate(level int) RespCompressorOpt {
	return func(respcom *RespCompressor) {
		respcom.wconstr = func(w io.Writer) (io.WriteCloser, string, error) {
			fw, err := flate.NewWriter(w, level)
			return fw, "flate", err
		}
	}
}

func WithGZip(level int) RespCompressorOpt {
	return func(respcom *RespCompressor) {
		respcom.wconstr = func(w io.Writer) (io.WriteCloser, string, error) {
			gzw, err := gzip.NewWriterLevel(w, level)
			return gzw, "gzip", err
		}
	}
}
