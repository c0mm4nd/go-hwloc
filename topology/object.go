package topology

// #cgo LDFLAGS: -lhwloc
// #include <hwloc.h>
import "C"
import "unsafe"

// GetInfo Search the given key name in object infos and return the corresponding value.
/*
 * If multiple keys match the given name, only the first one is returned.
 *
 * \return \c NULL if no such key exists.
 */
func (o *HwlocObject) GetInfo(name string) (string, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	res := C.hwloc_obj_get_info_by_name(o.hwloc_obj_t(), cname)
	return C.GoString(res), nil
}

// hwloc_obj_t Return C.hwloc_obj_t for HwlocObject
func (o *HwlocObject) hwloc_obj_t() C.hwloc_obj_t {
	return o.private
}
