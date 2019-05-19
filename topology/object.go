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

// AddInfo Add the given info name and value pair to the given object.
/*
 * The info is appended to the existing info array even if another key
 * with the same name already exists.
 *
 * The input strings are copied before being added in the object infos.
 *
 * \return \c 0 on success, \c -1 on error.
 *
 * \note This function may be used to enforce object colors in the lstopo
 * graphical output by using "lstopoStyle" as a name and "Background=#rrggbb"
 * as a value. See CUSTOM COLORS in the lstopo(1) manpage for details.
 *
 * \note If \p value contains some non-printable characters, they will
 * be dropped when exporting to XML, see hwloc_topology_export_xml() in hwloc/export.h.
 */
func (o *HwlocObject) AddInfo(name, value string) error {
	cname := C.CString(name)
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(cvalue))
	ret := C.hwloc_obj_add_info(o.hwloc_obj_t(), cname, cvalue)
	_ = ret
	return nil
}
