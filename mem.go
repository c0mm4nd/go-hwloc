//go:build !windows
// +build !windows

package hwloc

//#cgo LDFLAGS: -lhwloc
//#cgo LDFLAGS: -static -static-libgcc
// #include <stdint.h>
// #include <hwloc.h>
import "C"

// TODO
/*
HWLOC_DECLSPEC int hwloc_set_membind(hwloc_topology_t topology, hwloc_const_bitmap_t set, hwloc_membind_policy_t policy, int flags);

HWLOC_DECLSPEC int hwloc_get_membind(hwloc_topology_t topology, hwloc_bitmap_t set, hwloc_membind_policy_t * policy, int flags);

HWLOC_DECLSPEC int hwloc_set_proc_membind(hwloc_topology_t topology, hwloc_pid_t pid, hwloc_const_bitmap_t set, hwloc_membind_policy_t policy, int flags);

HWLOC_DECLSPEC int hwloc_get_proc_membind(hwloc_topology_t topology, hwloc_pid_t pid, hwloc_bitmap_t set, hwloc_membind_policy_t * policy, int flags);

HWLOC_DECLSPEC int hwloc_set_area_membind(hwloc_topology_t topology, const void *addr, size_t len, hwloc_const_bitmap_t set, hwloc_membind_policy_t policy, int flags);

HWLOC_DECLSPEC int hwloc_get_area_membind(hwloc_topology_t topology, const void *addr, size_t len, hwloc_bitmap_t set, hwloc_membind_policy_t * policy, int flags);

HWLOC_DECLSPEC int hwloc_get_area_memlocation(hwloc_topology_t topology, const void *addr, size_t len, hwloc_bitmap_t set, int flags);

HWLOC_DECLSPEC void *hwloc_alloc(hwloc_topology_t topology, size_t len);

HWLOC_DECLSPEC void *hwloc_alloc_membind(hwloc_topology_t topology, size_t len, hwloc_const_bitmap_t set, hwloc_membind_policy_t policy, int flags) __hwloc_attribute_malloc;

static __hwloc_inline void *
hwloc_alloc_membind_policy(hwloc_topology_t topology, size_t len, hwloc_const_bitmap_t set, hwloc_membind_policy_t policy, int flags) __hwloc_attribute_malloc;

HWLOC_DECLSPEC int hwloc_free(hwloc_topology_t topology, void *addr, size_t len);
*/

func (t *Topology) HwlocGetNUMANodeObjByOSIndex(affinity uint32) *HwlocNodeSet {
	node := C.hwloc_get_numanode_obj_by_os_index(t.hwloc_topology, C.uint32_t(affinity))
	return &HwlocNodeSet{hwloc_nodeset_t: node.nodeset}
}

func (t *Topology) HwlocSetMemBind(set *HwlocNodeSet, policy HwlocMemBindPolicy, flags HwlocMemBindFlag) {
	C.hwloc_set_membind(t.hwloc_topology, set.hwloc_nodeset_t, policy.CType(), flags.CType())
}
