package morfos

import "unsafe"

type eface struct {
	Type unsafe.Pointer
	Word unsafe.Pointer
}

func geteface(i *interface{}) *eface {
	return (*eface)(unsafe.Pointer(i))
}

// Morph will coerce the given value to the type stored on desiredtype
// without copying or changing any data on value. The result will
// be a merge of the data stored on value with the type stored on
// desiredtype, basically a frankstein :-).
//
// The result value should be castable to the type of desiredtype.
func Morph(value interface{}, desiredtype interface{}) interface{} {
	valueeface := geteface(&value)
	typeeface := geteface(&desiredtype)
	valueeface.Type = typeeface.Type
	return value
}
