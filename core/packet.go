package core

import (
	"bytes"
	"encoding/binary"
)

type Packet struct {
	buffer *bytes.Buffer
}

func NewPacket(buffer *bytes.Buffer) *Packet {
	if buffer == nil {
		buffer = new(bytes.Buffer)
	}
	return &Packet{buffer}
}

func (p *Packet) Write(v interface{}) {
	binary.Write(p.buffer, binary.LittleEndian, v)
}

func (p *Packet) ReadByte() (v byte) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadRune() (v rune) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadInt() (v int) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadInt8() (v int8) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadInt16() (v int16) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadInt32() (v int32) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadInt64() (v int64) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadUint() (v uint) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadUint8() (v uint8) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadUint16() (v uint16) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadUint32() (v uint32) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}

func (p *Packet) ReadUint64() (v uint64) {
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return
}
