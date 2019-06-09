package chlog

import (
	"encoding/gob"
	"os"
	"path"
	"strconv"
	"time"
)

//Writer writes the changes to disk
type Writer struct {
	f       *os.File
	encoder *gob.Encoder
}

//NewWriter ...
func NewWriter(c Config) (*Writer, error) {
	f, err := obtainWriteFileDescriptor(c)
	if err != nil {
		return nil, err
	}

	encoder := gob.NewEncoder(f)
	return &Writer{f: f, encoder: encoder}, nil
}

//Close the underlying file descriptor
func (w *Writer) Close() error {
	return w.f.Close()
}

//Sync file Buffer to disk
func (w *Writer) Sync() error {
	return w.f.Sync()
}

//WriteChange to log
func (w *Writer) WriteChange(change *Change) error {
	err := interfaceEncode(w.encoder, change)
	if err != nil {
		return err
	}
	return nil
}

func obtainWriteFileDescriptor(c Config) (*os.File, error) {
	err := os.MkdirAll(c.logFilesPath(), 0777)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}

	}

	return os.Create(path.Join(c.logFilesPath(), strconv.Itoa(int(time.Now().Unix()))))
}
