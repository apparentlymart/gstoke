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

	gstTypeState = C.gst_state_get_type()
	gstTypeURIType = C.gst_uri_type_get_type()
	gstTypeEvent = C.gst_event_get_type()
	gstTypeTask = C.gst_task_get_type()
	gstTypeElement = C.gst_element_get_type()
	gstTypeFormat = C.gst_format_get_type()
	gstTypeIntRange = C.gst_int_range_get_type()
	gstTypeDoubleRange = C.gst_double_range_get_type()
	gstTypeTagList = C.gst_tag_list_get_type()

	// This one isn't getting picked up by CGo for some reason
	//gstTypeFourCC = C.gst_fourcc_get_type()
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
