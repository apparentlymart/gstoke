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

// These are initialized in function Init
var gstTypeState C.GType
var gstTypeURIType C.GType
var gstTypeEvent C.GType
var gstTypeTask C.GType
var gstTypeElement C.GType
var gstTypeFormat C.GType
var gstTypeFourCC C.GType
var gstTypeIntRange C.GType
var gstTypeDoubleRange C.GType
var gstTypeTagList C.GType

type UnsupportedType uint64

func (t UnsupportedType) String() string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}
