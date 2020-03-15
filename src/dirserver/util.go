package main

import (
	"reflect"
	"unsafe"
)

func unsafeStrToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func unsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type WriteCounter struct {
	Len uint64
}

func (c *WriteCounter) Write(b []byte) (int, error) {
	c.Len += uint64(len(b))
	return len(b), nil
}

type DevZero struct {}

func (_ *DevZero) Read(p []byte) (int, error) {
	return len(p), nil;
}
