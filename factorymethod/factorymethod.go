package factorymethod

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"strings"
)

type CodecFactory interface {
	NewCodec(string) CodecInterface
}

type BaseCodeFactory struct {
}

func (*BaseCodeFactory) NewCodec(codec string) CodecInterface {
	codec = strings.ToLower(codec)
	switch codec {
	case "json":
		return &JsonCodec{}
	case "gob":
		return &GobCodec{}
	default:
		return nil
	}
}

type CodecInterface interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

type JsonCodec struct{}

func (*JsonCodec) Encode(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

func (*JsonCodec) Decode(data []byte, out interface{}) error {
	return json.Unmarshal(data, out)
}

type GobCodec struct{}

func (*GobCodec) Encode(input interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(input)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (*GobCodec) Decode(data []byte, out interface{}) error {
	var buffer bytes.Buffer
	_, err := buffer.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buffer)
	return dec.Decode(out)
}
