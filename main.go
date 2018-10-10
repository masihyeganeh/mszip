package mszip

import (
	"bytes"
	"compress/flate"
	"errors"
	"io"
)

type MsZip struct {
	dictionary []byte
}

func New() MsZip {
	return MsZip{
		[]byte(""),
	}
}

func (m MsZip) Decompress(input io.ReadSeeker, decompressedSize int) ([]byte, error) {
	output := make([]byte, decompressedSize)

	magic := make([]byte, 2)
	input.Read(magic)

	if !bytes.Equal(magic, []byte("CK")) { // CK = Chris Kirmse, official Microsoft purloiner
		return nil, errors.New("file is corrupted")
	}

	// last 32k of decompressed data of previous block is used as dictionary
	dictOffset := len(output) - (1 << 15)
	if dictOffset < 0 {
		dictOffset = 0
	}

	decompressor := flate.NewReaderDict(input, m.dictionary)
	decompressor.Read(output)
	decompressor.Close()

	m.dictionary = output[dictOffset:]

	return output, nil
}
