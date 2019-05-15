package topology

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lhwloc
// #include <stdlib.h>
// #include <stdio.h>
// #include <hwloc.h>
import "C"

type Topology struct {
	HwlocObject
	hwloc_topology C.hwloc_topology_t
}

func NewTopology() (*Topology, error) {
	var topology C.hwloc_topology_t = &C.struct_hwloc_topology{}
	C.hwloc_topology_init(&topology) // initialization

	return &Topology{
		hwloc_topology: topology,
	}, nil
}

func (t *Topology) Load() error {
	C.hwloc_topology_set_icache_types_filter(t.hwloc_topology, C.HWLOC_TYPE_FILTER_KEEP_ALL)
	C.hwloc_topology_load(t.hwloc_topology) // actual detection
	t.HwlocObject.Depth = int(C.hwloc_topology_get_depth(t.hwloc_topology))
	return nil
}

// Check Run internal checks on a topology structure
// The program aborts if an inconsistency is detected in the given topology.
// This routine is only useful to developers.
// The input topology should have been previously loaded with Load().
func (t *Topology) Check() error {
	C.hwloc_topology_check(t.hwloc_topology)
	return nil
}

// GetDepth Object levels, depths and types
/** \defgroup hwlocality_levels Object levels, depths and types
 * @{
 *
 * Be sure to see the figure in \ref termsanddefs that shows a
 * complete topology tree, including depths, child/sibling/cousin
 * relationships, and an example of an asymmetric topology where one
 * package has fewer caches than its peers.
 *
 * \brief Get the depth of the hierarchical tree of objects.
 *
 * This is the depth of ::HWLOC_OBJ_PU objects plus one.
 *
 * \note NUMA nodes, I/O and Misc objects are ignored when computing
 * the depth of the tree (they are placed on special levels).
 */
func (t *Topology) GetDepth() (int, error) {
	depth := C.hwloc_topology_get_depth(t.hwloc_topology)
	return int(depth), nil
}

// GetTypeDepth Retruns the depth of objects of type
/** \brief Returns the depth of objects of type \p type.
 *
 * If no object of this type is present on the underlying architecture, or if
 * the OS doesn't provide this kind of information, the function returns
 * ::HWLOC_TYPE_DEPTH_UNKNOWN.
 *
 * If type is absent but a similar type is acceptable, see also
 * hwloc_get_type_or_below_depth() and hwloc_get_type_or_above_depth().
 *
 * If ::HWLOC_OBJ_GROUP is given, the function may return ::HWLOC_TYPE_DEPTH_MULTIPLE
 * if multiple levels of Groups exist.
 *
 * If a NUMA node, I/O or Misc object type is given, the function returns a virtual
 * value because these objects are stored in special levels that are not CPU-related.
 * This virtual depth may be passed to other hwloc functions such as
 * hwloc_get_obj_by_depth() but it should not be considered as an actual
 * depth by the application. In particular, it should not be compared with
 * any other object depth or with the entire topology depth.
 * \sa hwloc_get_memory_parents_depth().
 *
 * \sa hwloc_type_sscanf_as_depth() for returning the depth of objects
 * whose type is given as a string.
 */
func (t *Topology) GetTypeDepth(ht HwlocObjType) (int, error) {
	depth := C.hwloc_get_type_depth(t.hwloc_topology, C.hwloc_obj_type_t(ht))
	return int(depth), nil
}

// GetMemoryParentsDepth Return the depth of parents where memory objects are attached.
/*
 * Memory objects have virtual negative depths because they are not part of
 * the main CPU-side hierarchy of objects. This depth should not be compared
 * with other level depths.
 *
 * If all Memory objects are attached to Normal parents at the same depth,
 * this parent depth may be compared to other as usual, for instance
 * for knowing whether NUMA nodes is attached above or below Packages.
 *
 * \return The depth of Normal parents of all memory children
 * if all these parents have the same depth. For instance the depth of
 * the Package level if all NUMA nodes are attached to Package objects.
 *
 * \return ::HWLOC_TYPE_DEPTH_MULTIPLE if Normal parents of all
 * memory children do not have the same depth. For instance if some
 * NUMA nodes are attached to Packages while others are attached to
 * Groups.
 */
func (t *Topology) GetMemoryParentsDepth() (int, error) {
	depth := C.hwloc_get_memory_parents_depth(t.hwloc_topology)
	return int(depth), nil
}

// GetTypeOrBelowDepth Returns the depth of objects of type or below
/*
 * If no object of this type is present on the underlying architecture, the
 * function returns the depth of the first "present" object typically found
 * inside \p type.
 *
 * This function is only meaningful for normal object types.
 * If a memory, I/O or Misc object type is given, the corresponding virtual
 * depth is always returned (see hwloc_get_type_depth()).
 *
 * May return ::HWLOC_TYPE_DEPTH_MULTIPLE for ::HWLOC_OBJ_GROUP just like
 * hwloc_get_type_depth().
 */
func (t *Topology) GetTypeOrBelowDepth(ht HwlocObjType) (int, error) {
	depth := C.hwloc_get_type_or_below_depth(t.hwloc_topology, C.hwloc_obj_type_t(ht))
	return int(depth), nil
}

// GetTypeOrAboveDepth Returns the depth of objects of type or above
/*
 * If no object of this type is present on the underlying architecture, the
 * function returns the depth of the first "present" object typically
 * containing \p type.
 *
 * This function is only meaningful for normal object types.
 * If a memory, I/O or Misc object type is given, the corresponding virtual
 * depth is always returned (see hwloc_get_type_depth()).
 *
 * May return ::HWLOC_TYPE_DEPTH_MULTIPLE for ::HWLOC_OBJ_GROUP just like
 * hwloc_get_type_depth().
 */
func (t *Topology) GetTypeOrAboveDepth(ht HwlocObjType) (int, error) {
	depth := C.hwloc_get_type_or_above_depth(t.hwloc_topology, C.hwloc_obj_type_t(ht))
	return int(depth), nil
}

// GetDepthType Returns the type of objects at depth
// depth should between 0 and hwloc_topology_get_depth()-1.
// return (hwloc_obj_type_t)-1 if depth \p depth does not exist.
func (t *Topology) GetDepthType(depth int) (HwlocObjType, error) {
	hw := C.hwloc_get_depth_type(t.hwloc_topology, C.int(depth))
	return HwlocObjType(hw), nil
}

// GetNbobjsByDepth Returns the width of level at depth.
func (t *Topology) GetNbobjsByDepth(depth int) (uint, error) {
	w := C.hwloc_get_nbobjs_by_depth(t.hwloc_topology, C.int(depth))
	return uint(w), nil
}

// GetNbobjsByType Returns the width of level type
// If no object for that type exists, 0 is returned.
// If there are several levels with objects of that type, -1 is returned.
func (t *Topology) GetNbobjsByType(ht HwlocObjType) (int, error) {
	nbcores := C.hwloc_get_nbobjs_by_type(t.hwloc_topology, C.hwloc_obj_type_t(ht))
	return int(nbcores), nil
}

// GetRootObj Returns the top-object of the topology-tree.
// Its type is ::HWLOC_OBJ_MACHINE.
func (t *Topology) GetRootObj() (*HwlocObject, error) {
	obj := C.hwloc_get_root_obj(t.hwloc_topology)
	ret := &HwlocObject{
		Type:         HwlocObjType(obj._type),
		SubType:      C.GoString(obj.subtype),
		OSIndex:      uint(obj.os_index),
		Name:         C.GoString(obj.name),
		TotalMemory:  uint64(obj.total_memory),
		Depth:        int(obj.depth),
		LogicalIndex: uint(obj.logical_index),
	}
	return ret, nil
}

func (t *Topology) Destroy() {
	C.hwloc_topology_destroy(t.hwloc_topology)
}
