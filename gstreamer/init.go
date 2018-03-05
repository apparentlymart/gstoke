package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"

type Version struct {
	Major uint
	Minor uint
	Micro uint
	Nano  uint
}

func Init() {
	C.gst_init(nil, nil)

	TypeState = ValueType(C.gst_state_get_type())
	TypeURIType = ValueType(C.gst_uri_type_get_type())
	TypeEvent = ValueType(C.gst_event_get_type())
	TypeTask = ValueType(C.gst_task_get_type())
	TypeElement = ValueType(C.gst_element_get_type())
	TypeFormat = ValueType(C.gst_format_get_type())
	TypeIntRange = ValueType(C.gst_int_range_get_type())
	TypeDoubleRange = ValueType(C.gst_double_range_get_type())
	TypeTagList = ValueType(C.gst_tag_list_get_type())

	// This one isn't getting picked up by CGo for some reason
	//TypeFourCC = ValueType(C.gst_fourcc_get_type())
}

func DeInit() {
	C.gst_deinit()
}

func GetVersion() Version {
	var major, minor, micro, nano C.guint
	C.gst_version(&major, &minor, &micro, &nano)
	return Version{
		Major: uint(major),
		Minor: uint(minor),
		Micro: uint(micro),
		Nano:  uint(nano),
	}
}
