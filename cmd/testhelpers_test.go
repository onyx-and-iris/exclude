package cmd

import "bytes"

// seekBuffer wraps bytes.Buffer to satisfy the readWriteTruncater interface,
// allowing it to be used in place of *os.File in tests.
type seekBuffer struct {
	bytes.Buffer
}

func (s *seekBuffer) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (s *seekBuffer) Truncate(size int64) error {
	s.Buffer.Reset()
	return nil
}
