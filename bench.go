package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/niubaoshu/gotiny"
	"github.com/vmihailenco/msgpack"
)

type AI interface {
	GetName() string
}

type A struct {
	Name  string
	Body  []byte
	Value float64
}

func (a A) GetName() string {
	return a.Name
}

func init() {
	msgpack.RegisterExt(20, A{})
	gob.Register(A{})
	gotiny.Register(A{})
}

func createStruct() A {
	return A{
		Name:  "blah",
		Body:  []byte{1, 2, 3, 4, 5},
		Value: 1.666,
	}
}

func createMap(max int) map[int64]float64 {
	m := make(map[int64]float64)
	for i := 0; i < max; i++ {
		m[int64(i)] = float64(i)
	}
	return m
}

func createSliceMap(max int) []map[int64]float64 {
	list := make([]map[int64]float64, max)
	for i := 0; i < max; i++ {
		list[i] = createMap(max)
	}
	return list
}

func encodeGob(v interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func encodeMsgpack(v interface{}) []byte {
	buf, err := msgpack.Marshal(v)
	if err != nil {
		panic(err)
	}

	return buf
}

func encodeGobStruct(enc *gob.Encoder, a AI) {
	err := enc.Encode(&a)
	if err != nil {
		panic(err)
	}
}

func encodeMsgpackStruct(enc *msgpack.Encoder, a A) {
	err := enc.Encode(a)
	if err != nil {
		panic(err)
	}
}

func encodeGotiny(v interface{}) []byte {
	buf := gotiny.Marshal(v)

	return buf
}

func encodeGotinyStruct(a AI) []byte {
	buf := gotiny.Marshal(&a)

	return buf
}

func decodeGob(b []byte, result interface{}) {
	buf := bytes.NewBuffer(b)
	enc := gob.NewDecoder(buf)

	err := enc.Decode(result)
	if err != nil {
		panic(err)
	}
}

func decodeGotinyStruct(b []byte) AI {
	var result AI
	gotiny.Unmarshal(b, &result)

	return result
}

func decodeMsgpackStruct(b []byte) AI {
	var result AI
	err := msgpack.Unmarshal(b, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func decodeGobStruct(b []byte) AI {
	var result AI
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}
func encodeJSON(v interface{}) []byte {
	byt, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return byt
}

func decodeJSON(b []byte, result interface{}) {
	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
	}
}
