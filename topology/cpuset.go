package topology

// #cgo LDFLAGS: -lhwloc
// #include <hwloc.h>
import "C"
import "github.com/carmark/gohwloc/bitmap"

// HwlocCPUSet A CPU set is a bitmap whose bits are set according to CPU physical OS indexes.
/*
 * It may be consulted and modified with the bitmap API as any
 * ::hwloc_bitmap_t (see hwloc/bitmap.h).
 *
 * Each bit may be converted into a PU object using
 * hwloc_get_pu_obj_by_os_index().
 */
type HwlocCPUSet struct {
	bitmap.BitMap
	hwloc_cpuset_t C.hwloc_bitmap_t
}

// Destroy free the HwlocCPUSet object
func (set HwlocCPUSet) Destroy() {
	if set.hwloc_cpuset_t != nil {
		C.hwloc_bitmap_free(set.hwloc_cpuset_t)
	}
}
