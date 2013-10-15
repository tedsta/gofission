package core

import (
	//"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	out := NewOutPacket(nil)
	out.Write("hello")
	out.Write(byte(42))
	out.Write(rune(42))
	out.Write(int(42))
	out.Write(int8(42))
	out.Write(int16(42))
	out.Write(int32(42))
	out.Write(int64(42))
	out.Write(uint(42))
	out.Write(uint8(42))
	out.Write(uint16(42))
	out.Write(uint32(42))
	out.Write(uint64(42))

	out2 := NewOutPacket(nil)
	out2.Write(int(84))
	out.Append(out2)

	var s string
	var b byte
	var r rune
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	var app int // Append tester

	in := NewInPacket(out.buffer)
	in.Read(&s)
	in.Read(&b)
	in.Read(&r)
	in.Read(&i)
	in.Read(&i8)
	in.Read(&i16)
	in.Read(&i32)
	in.Read(&i64)
	in.Read(&u)
	in.Read(&u8)
	in.Read(&u16)
	in.Read(&u32)
	in.Read(&u64)
	in.Read(&app)

	if s != "hello" {
		t.Fail()
	}
	if b != 42 {
		t.Fail()
	}
	if r != 42 {
		t.Fail()
	}
	if i != 42 {
		t.Fail()
	}
	if i8 != 42 {
		t.Fail()
	}
	if i16 != 42 {
		t.Fail()
	}
	if i32 != 42 {
		t.Fail()
	}
	if i64 != 42 {
		t.Fail()
	}
	if u != 42 {
		t.Fail()
	}
	if u8 != 42 {
		t.Fail()
	}
	if u16 != 42 {
		t.Fail()
	}
	if u32 != 42 {
		t.Fail()
	}
	if u64 != 42 {
		t.Fail()
	}
	if app != 84 {
		t.Fail()
	}
}
