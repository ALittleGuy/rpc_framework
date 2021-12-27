package gserver

import (
	"bytes"
	"encoding/binary"
	"io"
	"rpc/internal/giface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) UnPack(reader io.Reader) (giface.IMessage, error) {
	msg := &Message{}

	// read magic num and validate magic num
	if err := binary.Read(reader, binary.LittleEndian, &msg.MagicNum); err != nil || msg.MagicNum != magicNum {
		return nil, err
	}

	// read messsage len
	if err := binary.Read(reader, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}

	// read header len
	if err := binary.Read(reader, binary.LittleEndian, &msg.HeaderLen); err != nil {
		return nil, err
	}

	// read header
	msg.Header = make([]byte, msg.HeaderLen)
	if err := binary.Read(reader, binary.LittleEndian, &msg.Header); err != nil {
		return nil, err
	}

	// read body
	msg.Body = make([]byte, msg.Len-msg.HeaderLen-4)
	if err := binary.Read(reader, binary.LittleEndian, &msg.Body); err != nil {
		return nil, err
	}
	return msg, nil
}

func (d *DataPack) Pack(message giface.IMessage) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	if err := binary.Write(buff, binary.LittleEndian, message.GetLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buff, binary.LittleEndian, message.GetHeaderLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buff, binary.LittleEndian, message.GetHeader()); err != nil {
		return nil, err
	}
	if err := binary.Write(buff, binary.LittleEndian, message.GetBody()); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
