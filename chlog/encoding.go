package chlog

import (
	"encoding/gob"
	"io"
	"log"
)

func interfaceEncode(enc *gob.Encoder, p interface{}) error {
	return enc.Encode(&p)
}

// decodeAll decodes the next interface value from the stream and returns it.
func decodeAll(dec *gob.Decoder) chan interface{} {
	decodedChan := make(chan interface{})

	go func() {
		for {
			var p interface{}
			err := dec.Decode(&p)
			if err != nil {
				if err == io.EOF {
					close(decodedChan)
					return
				}
				log.Fatal("decode:", err)
			}
			decodedChan <- p
		}
	}()

	return decodedChan
}

func he(err error) {
	if err != nil {
		panic(err)
	}
}
