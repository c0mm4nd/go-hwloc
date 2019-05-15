package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhwloc
// #include <stdlib.h>
// #include <stdio.h>
// #include <hwloc.h>
import "C"
import (
	"fmt"
)

func main() {
	var topology C.hwloc_topology_t = &C.struct_hwloc_topology{}

	C.hwloc_topology_init(&topology) // initialization
	C.hwloc_topology_set_icache_types_filter(topology, C.HWLOC_TYPE_FILTER_KEEP_ALL)
	C.hwloc_topology_load(topology) // actual detection
	defer C.hwloc_topology_destroy(topology)

	nbcores := C.hwloc_get_nbobjs_by_type(topology, C.HWLOC_OBJ_CORE)
	fmt.Printf("%d cores\n", nbcores)

	//root := C.hwloc_get_root_obj(topology)
	topodepth := C.hwloc_topology_get_depth(topology)
	fmt.Printf("%d\n", topodepth)

	for depth := C.int(0); depth < topodepth; depth++ {
		t := C.hwloc_get_depth_type(topology, depth)
		var types string
		nbobjs := C.hwloc_get_nbobjs_by_depth(topology, depth)
		if nbobjs != 0 {
			if depth < 0 {
				fmt.Printf("Special depth %d:", depth)
			} else {
				fmt.Printf("%*sdepth %d:", depth, "", depth)
			}

			if depth < 0 {
				// use plain type, we don't want OSdev subtype since it may differ for other objects in the level
				types = C.GoString(C.hwloc_obj_type_string(t))
			} else {
				// use verbose type name, those are identical for all objects on normal levels
				var _types = make([]C.char, 64)
				C.hwloc_obj_type_snprintf(&_types[0], 64, C.hwloc_get_obj_by_depth(topology, depth, 0), 1)
				types = C.GoString(&_types[0])
			}

			fmt.Printf("%*s%d %s (type #%d)\n", 10, "", nbobjs, types, t)
		}
	}
}
