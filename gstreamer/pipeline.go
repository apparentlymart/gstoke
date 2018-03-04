package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"runtime"
	"unsafe"
)

type Pipeline struct {
	Bin
}

func NewPipeline(name string) *Pipeline {
	rawName := C.CString(name)
	defer C.free(unsafe.Pointer(rawName))
	raw := C.gst_pipeline_new((*C.gchar)(rawName))
	return newPipeline(raw)
}

func (p *Pipeline) AsBin() *Bin {
	return &p.Bin
}

func (p *Pipeline) Bus() *Bus {
	raw := C.gst_pipeline_get_bus(p.pipelinePtr())
	return newBus(raw)
}

func newPipeline(raw *C.struct__GstElement) *Pipeline {
	ret := &Pipeline{
		Bin{
			Element{
				raw: raw,
			},
		},
	}
	runtime.SetFinalizer(&ret.Bin.Element, (*Element).unref)
	return ret
}

func (b *Bin) pipelinePtr() *C.struct__GstPipeline {
	return (*C.struct__GstPipeline)(unsafe.Pointer(b.raw))
}
