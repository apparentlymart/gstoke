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
