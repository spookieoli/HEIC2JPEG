package utils

import (
	"io"
)

type Utils struct {
	w    io.Writer
	Skip int
}

func New(w io.Writer, exif []byte) (*Utils, error) {
	writer := &Utils{w: w, Skip: 2}

	if _, err := w.Write([]byte{0xff, 0xd8}); err != nil {
		return nil, err
	}

	if exif != nil {
		marker := []byte{0xff, 0xe1, byte(len(exif) >> 8), byte(len(exif) & 0xff)}
		if _, err := w.Write(marker); err != nil {
			return nil, err
		}

		if _, err := w.Write(exif); err != nil {
			return nil, err
		}
	}

	return writer, nil
}

// Write will write the line to the Writer
func (u *Utils) Write(data []byte) (n int, err error) {
	if u.Skip <= 0 {
		return u.w.Write(data)
	}

	if dataLen := len(data); dataLen < u.Skip {
		u.Skip -= dataLen
		return dataLen, nil
	}

	if u.Skip > 0 {
		n, err = u.w.Write(data[u.Skip:])
	}
	n += u.Skip
	u.Skip = 0
	return n, err
}
