package gserver

import (
	"bytes"
	"encoding/gob"
	"rpc/internal/giface"
)

type Serializer struct{}

func (s *Serializer) Encode(request *giface.IMessage) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(request)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *Serializer) Decode(data []byte, value interface{}) error {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(value)
	if err != nil {
		return err
	}
	return nil
}
