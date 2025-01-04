package DataStructures

import (
	"bytes"
	"encoding/gob"
)

/*
This package is for serializing data strucutres/ either to json or to disk
*/

func (node *bNode[T]) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(node)
	if err != nil {
		panic(err) // Handle serialization error
	}
	return buffer.Bytes()
}

func Deserialize[T Ordered](data []byte) *bNode[T] {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	var node bNode[T]
	err := decoder.Decode(&node)
	if err != nil {
		panic(err) // Handle deserialization error
	}
	return &node
}
