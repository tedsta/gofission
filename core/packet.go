package core

import (
	"bytes"
	"encoding/gob"
	"io"
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

func (p *OutPacket) Append(p2 *OutPacket) {
	p2.WriteTo(p.buffer)
}

func (p *OutPacket) Write(v ...interface{}) {
	for _, w := range v {
		p.encoder.Encode(w)
	}
}

func (p *OutPacket) String() string {
	return p.buffer.String()
}

func (p *OutPacket) Clear() {
	p.buffer.Reset()
}

func (p *OutPacket) WriteTo(w io.Writer) {
	p.buffer.WriteTo(w)
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

func (p *InPacket) Read(v ...interface{}) {
	for _, w := range v {
		p.decoder.Decode(w)
	}
}

func (p *InPacket) String() string {
	return p.buffer.String()
}

func (p *InPacket) Clear() {
	p.buffer.Reset()
}
