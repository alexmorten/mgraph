package chlog

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	//ErrNoChangeLog will be returned when the change log does not exist (yet)
	ErrNoChangeLog = errors.New("no changelog file found")
)

//ReadAllChanges in the change log. ErrNoChangeLog will be returned when the log doesn't exist
func ReadAllChanges(c Config) (chan *Change, error) {
	if _, err := os.Stat(c.logFilesPath()); err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoChangeLog
		}
		return nil, err
	}
	files, err := getSortedFileList(c.logFilesPath())
	if err != nil {
		return nil, err
	}

	changeChan := make(chan *Change)
	go func() {
		for _, fileInfo := range files {
			f, err := os.Open(c.logFilePath(fileInfo.name))
			if os.IsNotExist(err) {
				log.Fatal("file disappeared for some reason: ", fileInfo.name, "with error: ", err)
			}

			var decoder *gob.Decoder
			decoder = gob.NewDecoder(f)
			decodeChan := decodeAll(decoder)

			for decoded := range decodeChan {
				if change, ok := decoded.(*Change); ok {
					changeChan <- change
				} else {
					panic(fmt.Sprintf("unexpected value found when decoding into changes: %s", decoded))
				}
			}
			err = f.Close()
			if err != nil {
				panic(err)
			}
		}
		close(changeChan)
	}()
	return changeChan, nil
}

func getSortedFileList(dataPath string) (fileInfoSlice, error) {
	infoList := fileInfoSlice{}
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		ts, err := strconv.ParseInt(f.Name(), 10, 64)
		if err != nil {
			continue
		}
		info := fileInfo{}
		info.name = f.Name()
		info.ts = ts

		infoList = append(infoList, info)
	}

	sort.Sort(infoList)
	return infoList, nil
}

type fileInfo struct {
	name string
	ts   int64
}

type fileInfoSlice []fileInfo

func (s fileInfoSlice) Len() int      { return len(s) }
func (s fileInfoSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s fileInfoSlice) Less(i, j int) bool {
	return s[i].ts < s[j].ts
}
