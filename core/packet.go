package core

import (
	"bytes"
	"encoding/gob"
)

// OutPacket ###################################################################

type OutPacket struct {
	buffer  *bytes.Buffer
	encoder *gob.Encoder
}

func NewOutPacket(buffer *bytes.Buffer) *OutPacket {
	if buffer == nil {
		buffer = new(bytes.Buffer)
	}
	return &OutPacket{buffer, gob.NewEncoder(buffer)}
}

func (p *OutPacket) Write(v interface{}) {
	p.encoder.Encode(v)
}

// InPacket ####################################################################
type InPacket struct {
	buffer  *bytes.Buffer
	decoder *gob.Decoder
}

func NewInPacket(buffer *bytes.Buffer) *InPacket {
	if buffer == nil {
		buffer = new(bytes.Buffer)
	}
	return &InPacket{buffer, gob.NewDecoder(buffer)}
}

func (p *InPacket) Read(v interface{}) {
	p.decoder.Decode(v)
}
