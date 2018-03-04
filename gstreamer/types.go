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

type UnsupportedType uint64

func (t UnsupportedType) String() string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}
