package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"runtime"
	"unsafe"
)

type Bus struct {
	raw *C.struct__GstBus
}

const Infinite = int64(-1)

// Poll blocks the current goroutine until a matching message is emitted
// or until the timeout is reached. Pass Infinite to disable the timeout.
//
// If the timeout is reached, the return value is nil.
func (b *Bus) Poll(types MessageType, timeoutNS int64) *Message {
	raw := C.gst_bus_poll(b.raw, C.GstMessageType(types), C.GstClockTime(timeoutNS))
	if raw == nil {
		return nil
	}
	return newMessage(raw)
}

func newBus(raw *C.struct__GstBus) *Bus {
	ret := &Bus{
		raw: raw,
	}
	runtime.SetFinalizer(ret, (*Bus).unref)
	return ret
}

func (ef *Bus) unref() {
	raw := C.gpointer(unsafe.Pointer(ef.raw))
	C.gst_object_unref(raw)
}
