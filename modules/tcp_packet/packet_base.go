package tcp_packet

import (
	"bytes"
	"encoding/gob"
)

type PacketType uint16

const (
	PacketTypePressed PacketType = iota
)

type PacketBase struct {
	PacketType
}

func SerializePacket(data any) ([]byte, error) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)
	gob.Register(data)
	err := enc.Encode(&data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DeserializePacket(data []byte) (any, error) {
	buf := bytes.NewBuffer(data)

	dec := gob.NewDecoder(buf)

	gob.Register(&PacketPressed{})

	var decodedData any
	err := dec.Decode(&decodedData)
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}
