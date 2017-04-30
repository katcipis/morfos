package morfos_test

import (
	"testing"

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
