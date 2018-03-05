package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type Element struct {
	raw *C.struct__GstElement
}

type State int

//go:generate stringer -type State

const NullState State = C.GST_STATE_NULL
const ReadyState State = C.GST_STATE_READY
const PausedState State = C.GST_STATE_PAUSED
const PlayingState State = C.GST_STATE_PLAYING

type StateChangeResult int

//go:generate stringer -type StateChangeResult

const StateChangeSuccess StateChangeResult = C.GST_STATE_CHANGE_SUCCESS
const StateChangeFailure StateChangeResult = C.GST_STATE_CHANGE_FAILURE
const StateChangeAsync StateChangeResult = C.GST_STATE_CHANGE_ASYNC
const StateChangeNoPreroll StateChangeResult = C.GST_STATE_CHANGE_NO_PREROLL

func NewElement(factoryName, name string) (*Element, error) {
	factory, err := FindElementFactory(factoryName)
	if err != nil {
		return nil, err
	}
	return factory.CreateElement(name)
}

func (e *Element) Name() string {
	raw := gstObject(unsafe.Pointer(e.raw))
	rawName := C.gst_object_get_name(raw)
	defer C.g_free(gpointer(unsafe.Pointer(rawName)))
	return C.GoString((*C.char)(rawName))
}

func (e *Element) SetProperty(name string, value interface{}) {
	rawName := C.CString(name)
	defer C.free(unsafe.Pointer(rawName))

	switch tv := value.(type) {
	case string:
		rawVal := C.CString(tv)
		C.gst_element_set_addr(e.raw, rawName, unsafe.Pointer(rawVal))
		C.free(unsafe.Pointer(rawVal))
	default:
		panic(fmt.Errorf("set property with unsupported type %T", value))
	}
}

func (e *Element) PropertySpecs() PropertySpecs {
	return objectProperties((*C.struct__GObject)(unsafe.Pointer(e.raw)))
}

func (e *Element) SetState(state State) (StateChangeResult, error) {
	result := StateChangeResult(C.gst_element_set_state(e.raw, C.GstState(state)))
	if result == StateChangeFailure {
		return result, fmt.Errorf("state change failed")
	}
	return result, nil
}

func LinkElements(src, dst *Element) error {
	success := C.gst_element_link(src.raw, dst.raw)
	if int(success) != 0 {
		return fmt.Errorf("no link is possbile between %q and %q", src.Name(), dst.Name())
	}
	return nil
}

func (s State) String() string {
	switch s {
	case NullState:
		return "null"
	case ReadyState:
		return "ready"
	case PausedState:
		return "paused"
	case PlayingState:
		return "playing"
	default:
		return "<unknown>"
	}
}

func newElement(raw *C.struct__GstElement) *Element {
	ret := &Element{
		raw: raw,
	}
	runtime.SetFinalizer(ret, (*Element).unref)
	return ret
}

func (e *Element) ref() {
	C.gst_object_ref(gpointer(unsafe.Pointer(e.raw)))
}

func (e *Element) unref() {
	C.gst_object_unref(gpointer(unsafe.Pointer(e.raw)))
}
