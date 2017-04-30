package morfos_test

import (
	"testing"
	"unsafe"

	"github.com/katcipis/morfos"
)

func TestStructsSameSize(t *testing.T) {
	type original struct {
		x int
		y int
	}
	type notoriginal struct {
		z int
		w int
	}

	orig := original{x: 100, y: 200}
	_, ok := interface{}(orig).(notoriginal)
	if ok {
		t.Fatal("casting should be invalid")
	}

	morphed := morfos.Morph(orig, notoriginal{})
	morphedNotOriginal, ok := morphed.(notoriginal)

	if !ok {
		t.Fatal("casting should be valid now")
	}

	if orig.x != morphedNotOriginal.z {
		t.Fatalf("expected x[%d] == z[%d]", orig.x, morphedNotOriginal.z)
	}

	if orig.y != morphedNotOriginal.w {
		t.Fatalf("expected y[%d] == w[%d]", orig.y, morphedNotOriginal.w)
	}
}

func TestMutatingString(t *testing.T) {

	type stringStruct struct {
		str unsafe.Pointer
		len int
	}

	var rawstr [5]byte
	rawstr[0] = 'h'
	rawstr[1] = 'e'
	rawstr[2] = 'l'
	rawstr[3] = 'l'
	rawstr[4] = 'o'

	hi := stringStruct{
		str: unsafe.Pointer(&rawstr),
		len: len(rawstr),
	}

	somestr := ""

	morphed := morfos.Morph(hi, somestr)
	mutableStr := morphed.(string)

	if mutableStr != "hello" {
		t.Fatalf("expected hello, got: %s", mutableStr)
	}

	rawstr[0] = 'h'
	rawstr[1] = 'a'
	rawstr[2] = 'c'
	rawstr[3] = 'k'
	rawstr[4] = 'd'

	if mutableStr != "hackd" {
		t.Fatalf("expected hackd, got: %s", mutableStr)
	}
}
