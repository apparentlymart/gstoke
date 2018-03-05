package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"unsafe"
)

func gpointer(raw unsafe.Pointer) C.gpointer {
	return C.gpointer(raw)
}

func gstObject(raw unsafe.Pointer) *C.struct__GstObject {
	return (*C.struct__GstObject)(raw)
}

type ValueType C.GType

func (t ValueType) String() string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}

func (t ValueType) raw() C.GType {
	return C.GType(t)
}

// The type variables are initialized in function Init

const TypeBoolean ValueType = ValueType(C.G_TYPE_BOOLEAN)
const TypeInt ValueType = ValueType(C.G_TYPE_INT)
const TypeInt64 ValueType = ValueType(C.G_TYPE_INT64)
const TypeUInt ValueType = ValueType(C.G_TYPE_UINT)
const TypeUInt64 ValueType = ValueType(C.G_TYPE_UINT64)
const TypeDouble ValueType = ValueType(C.G_TYPE_DOUBLE)
const TypeString ValueType = ValueType(C.G_TYPE_STRING)

var TypeState ValueType
var TypeURIType ValueType
var TypeEvent ValueType
var TypeTask ValueType
var TypeElement ValueType
var TypeFormat ValueType
var TypeIntRange ValueType
var TypeDoubleRange ValueType
var TypeTagList ValueType

//var TypeFourCC ValueType

type UnsupportedType uint64

func (t UnsupportedType) String() string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}
