package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex 1.4实现int64转[]byte
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if err != nil {
		log.Panicf("int transact to []byte failed! %v\n", err)
	}
	return buffer.Bytes()
}
