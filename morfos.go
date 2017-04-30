package morfos

import "unsafe"

// tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	runtime/type.go
type tflag uint8

type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype

type rtype struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32   // hash of type; avoids computation in hash tables
	tflag      tflag    // extra type information flags
	align      uint8    // alignment of variable with this type
	fieldAlign uint8    // alignment of struct field with this type
	kind       uint8    // enumeration for C
	alg        *typeAlg // algorithm table
	gcdata     *byte    // garbage collection data
	str        nameOff  // string form
	ptrToThis  typeOff  // type for pointer to this type, may be zero
}

type eface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func getEface(i interface{}) eface {
	return *(*eface)(unsafe.Pointer(&i))
}

// Morph will coerce the given value to the type stored on desiredtype
// without copying or changing any data on value. The result will
// be a merge of the data stored on value with the type stored on
// desiredtype, basically a frankstein :-).
//
// The result value should be castable to the type of desiredtype.
func Morph(value interface{}, desiredtype interface{}) interface{} {
	return nil
}
