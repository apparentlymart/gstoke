package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

type Message struct {
	ty  MessageType
	raw *C.struct__GstMessage
}

func (m *Message) Type() MessageType {
	return m.ty
}

// ParseError parses the message as an error, or panics if the message type
// is not ErrorMessage.
func (m *Message) ParseError() error {
	if m.ty != ErrorMessage {
		panic(fmt.Sprintf("ParseError called on %s message", m.ty))
	}
	var rawErr *C.GError
	C.gst_message_parse_error(m.raw, &rawErr, nil)
	err := errors.New(C.GoString((*C.char)(rawErr.message)))
	C.g_error_free(rawErr)
	return err
}

// SourceElement returns the element that generated the message, or panics
// if the message type is not ElementMessage.
func (m *Message) SourceElement() *Element {
	if m.ty != ElementMessage {
		panic(fmt.Sprintf("SourceElement called on %s message", m.ty))
	}
	src := newElement((*C.struct__GstElement)(unsafe.Pointer(m.raw.src)))
	src.ref() // so the element will outlive the message
	return src
}

// Content returns the dynamic structure of the message as a map, along
// with the GStreamer type name for that structure.
//
// Not all messages have content. If the receiver does not have content
// then the result is "", nil.
func (m *Message) Content() (string, map[string]interface{}) {
	st := C.gst_message_get_structure(m.raw)
	if st == nil {
		return "", nil
	}
	rawName := (*C.char)(C.gst_structure_get_name(st))
	name := C.GoString(rawName)

	vals := make(map[string]interface{})
	fieldCount := int(C.gst_structure_n_fields(st))
	for i := 0; i < fieldCount; i++ {
		rawFieldName := C.gst_structure_nth_field_name(st, C.guint(i))
		fieldType := C.gst_structure_get_field_type(st, rawFieldName)
		fieldName := C.GoString((*C.char)(rawFieldName))

		switch fieldType {
		case C.G_TYPE_BOOLEAN:
			var val C.gboolean
			C.gst_structure_get_boolean(st, rawFieldName, &val)
			vals[fieldName] = !(val == 0)
		case C.G_TYPE_INT, C.G_TYPE_INT64: // We're assuming a 64-bit system here, because the GStreamer API doesn't make anything else easy
			var val C.gint
			C.gst_structure_get_int(st, rawFieldName, &val)
			vals[fieldName] = int(val)
		case C.G_TYPE_UINT, C.G_TYPE_UINT64: // We're assuming a 64-bit system here, because the GStreamer API doesn't make anything else easy
			var val C.guint
			C.gst_structure_get_uint(st, rawFieldName, &val)
			vals[fieldName] = uint(val)
		case C.G_TYPE_DOUBLE:
			var val C.gdouble
			C.gst_structure_get_double(st, rawFieldName, &val)
			vals[fieldName] = float64(val)
		case C.G_TYPE_STRING:
			rawVal := C.gst_structure_get_string(st, rawFieldName)
			vals[fieldName] = C.GoString((*C.char)(rawVal))
		case gstTypeState:
			var val C.gint
			C.gst_structure_get_enum(st, rawFieldName, fieldType, &val)
			vals[fieldName] = State(val)
		case gstTypeElement:
			val := C.gst_structure_get_value(st, rawFieldName)
			ptr := C.g_value_peek_pointer(val)
			raw := (*C.struct__GstElement)(unsafe.Pointer(ptr))
			elem := newElement(raw)
			elem.ref()
			vals[fieldName] = elem
		default:
			vals[fieldName] = UnsupportedType(fieldType)
		}
	}

	return name, vals
}

type MessageType int

//go:generate stringer -type MessageType

const ErrorMessage MessageType = C.GST_MESSAGE_ERROR
const WarningMessage MessageType = C.GST_MESSAGE_WARNING
const InfoMessage MessageType = C.GST_MESSAGE_INFO
const TagMessage MessageType = C.GST_MESSAGE_TAG
const BufferingMessage MessageType = C.GST_MESSAGE_BUFFERING
const ElementMessage MessageType = C.GST_MESSAGE_ELEMENT
const ProgressMessage MessageType = C.GST_MESSAGE_PROGRESS
const AnyMessages MessageType = C.GST_MESSAGE_ANY

func (t MessageType) String() string {
	rawName := C.gst_message_type_get_name(C.GstMessageType(t))
	return C.GoString((*C.char)(rawName))
}

func newMessage(raw *C.struct__GstMessage) *Message {
	ret := &Message{
		ty:  MessageType(raw._type),
		raw: raw,
	}
	runtime.SetFinalizer(ret, (*Message).unref)
	return ret
}

func (m *Message) unref() {
	raw := C.gpointer(unsafe.Pointer(m.raw))
	C.gst_object_unref(raw)
}
