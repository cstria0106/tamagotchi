package util

import "encoding/binary"

func EncodeU16(i uint16) (buffer []byte) {
	buffer = make([]byte, 2)
	binary.LittleEndian.PutUint16(buffer, i)
	return buffer
}

func DecodeU16(buffer []byte) uint16 {
	return binary.LittleEndian.Uint16(buffer)
}

func EncodeU32(i uint32) (buffer []byte) {
	buffer = make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, i)
	return buffer
}

func DecodeU32(buffer []byte) uint32 {
	return binary.LittleEndian.Uint32(buffer)
}

func EncodeI16(i int16) (buffer []byte) {
	buffer = make([]byte, 2)
	binary.LittleEndian.PutUint16(buffer, uint16(i))
	return buffer
}

func DecodeI16(buffer []byte) int16 {
	return int16(binary.LittleEndian.Uint16(buffer))
}

func EncodeI32(i int32) (buffer []byte) {
	buffer = make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, uint32(i))
	return buffer
}

func DecodeI32(buffer []byte) int32 {
	return int32(binary.LittleEndian.Uint32(buffer))
}
