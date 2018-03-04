package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type Bin struct {
	Element
}

func NewBin(name string) *Bin {
	rawName := C.CString(name)
	defer C.free(unsafe.Pointer(rawName))
	raw := C.gst_bin_new((*C.gchar)(rawName))
	return newBin(raw)
}

func (b *Bin) AsElement() *Element {
	return &b.Element
}

func (b *Bin) AddElement(elem *Element) error {
	success := C.gst_bin_add(b.binPtr(), elem.raw)
	if int(success) == 0 {
		return fmt.Errorf("bin does not accept element")
	}
	return nil
}

func newBin(raw *C.struct__GstElement) *Bin {
	ret := &Bin{
		Element{
			raw: raw,
		},
	}
	runtime.SetFinalizer(&ret.Element, (*Element).unref)
	return ret
}

func (b *Bin) binPtr() *C.struct__GstBin {
	return (*C.struct__GstBin)(unsafe.Pointer(b.raw))
}
