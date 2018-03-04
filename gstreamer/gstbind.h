
#include <stdlib.h>
#include <gst/gst.h>

inline void gst_element_set_addr(GstElement* elem, char* name, void* addr) {
    g_object_set (G_OBJECT(elem), name, addr, NULL);
}
