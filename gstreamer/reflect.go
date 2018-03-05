package gstreamer

// #cgo pkg-config: gstreamer-1.0
// #include "gstbind.h"
import "C"
import (
	"unsafe"
)

type PropertySpec struct {
	Name         string
	Blurb        string
	Type         ValueType
	Readable     bool
	Writable     bool
	Controllable bool
}

type PropertySpecs map[string]PropertySpec

func objectProperties(raw *C.struct__GObject) PropertySpecs {
	class := C.g_object_get_class(raw)
	var rawCount C.guint
	rawSpecsStart := C.g_object_class_list_properties(class, &rawCount)
	count := uint(rawCount)
	if count == 0 {
		return nil
	}
	defer C.g_free(C.gpointer(rawSpecsStart))

	sl := struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(unsafe.Pointer(rawSpecsStart)), int(count), int(count)}
	rawSpecs := *((*[]*C.GParamSpec)(unsafe.Pointer(&sl)))

	ret := make(PropertySpecs)
	for _, rawSpec := range rawSpecs {
		rawName := C.g_param_spec_get_name(rawSpec)
		name := C.GoString((*C.char)(rawName))
		rawBlurb := C.g_param_spec_get_blurb(rawSpec)
		blurb := C.GoString((*C.char)(rawBlurb))

		spec := PropertySpec{
			Name:  name,
			Blurb: blurb,
			Type:  ValueType(rawSpec.value_type),
		}

		if (rawSpec.flags & C.G_PARAM_READABLE) != 0 {
			spec.Readable = true
		}
		if (rawSpec.flags & C.G_PARAM_WRITABLE) != 0 {
			spec.Writable = true
		}
		if (rawSpec.flags & C.GST_PARAM_CONTROLLABLE) != 0 {
			spec.Controllable = true
		}

		ret[name] = spec
	}
	return ret
}
