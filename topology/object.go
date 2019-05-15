package topology

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhwloc
// #include <stdlib.h>
// #include <stdio.h>
// #include <hwloc.h>
import "C"

func (o *HwlocObject) GetInfo(name string) (string, error) {
	res := C.hwloc_obj_get_info_by_name(o.hwloc_obj_t(), C.CString(name))
	return C.GoString(res), nil
}

func (o *HwlocObject) hwloc_obj_t() C.hwloc_obj_t {
	var ret C.hwloc_obj_t = &C.struct_hwloc_obj{}
	ret.depth = C.int(o.Depth)
	return ret
}
