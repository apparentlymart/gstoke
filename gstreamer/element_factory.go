package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type ElementFactory struct {
	raw *C.struct__GstElementFactory
}

func FindElementFactory(name string) (*ElementFactory, error) {
	rawName := C.CString(name)
	defer C.free(unsafe.Pointer(rawName))
	raw := C.gst_element_factory_find((*C.gchar)(rawName))
	if raw == nil {
		return nil, fmt.Errorf("no element factory named %s", name)
	}
	return newElementFactory(raw), nil
}

func (ef *ElementFactory) Metadata() map[string]string {
	ret := make(map[string]string)
	keys := C.gst_element_factory_get_metadata_keys(ef.raw)
	q := uintptr(unsafe.Pointer(keys))
	for {
		p := (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		key := C.GoString(*p)
		rawVal := C.gst_element_factory_get_metadata(ef.raw, (*C.gchar)(*p))
		val := C.GoString((*C.char)(rawVal))
		ret[key] = val
		q += unsafe.Sizeof(q)
	}
	C.g_strfreev(keys)
	return ret
}

func (ef *ElementFactory) CreateElement(name string) (*Element, error) {
	rawName := C.CString(name)
	defer C.free(unsafe.Pointer(rawName))
	raw := C.gst_element_factory_create(ef.raw, (*C.gchar)(rawName))
	if raw == nil {
		return nil, fmt.Errorf("failed to create element")
	}
	return newElement(raw), nil
}

func newElementFactory(raw *C.struct__GstElementFactory) *ElementFactory {
	ret := &ElementFactory{
		raw: raw,
	}
	runtime.SetFinalizer(ret, (*ElementFactory).unref)
	return ret
}

func (ef *ElementFactory) unref() {
	raw := C.gpointer(unsafe.Pointer(ef.raw))
	C.gst_object_unref(raw)
}
