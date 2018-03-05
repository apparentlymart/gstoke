
#include <stdlib.h>
#include <gst/gst.h>

inline void gst_element_set_addr(GstElement* elem, char* name, void* addr) {
    g_object_set(G_OBJECT(elem), name, addr, NULL);
}

inline GObjectClass* g_object_get_class(GObject *obj) {
    return G_OBJECT_GET_CLASS(obj);
}

inline GType g_type_from_class(GObjectClass *cla) {
    return G_TYPE_FROM_CLASS(cla);
}
