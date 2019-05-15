package topology

// brief Type of topology object.
// Do not rely on the ordering or completeness of the values as new ones
// may be defined in the future!  If you need to compare types, use
// hwloc_compare_types() instead.
type HwlocObjType int

const (
	// HwlocObjMachine Machine.
	// A set of processors and memory with cache coherency.
	// This type is always used for the root object of a topology, and never used anywhere else.
	// Hence its parent is always NULL.
	HwlocObjMachine HwlocObjType = iota

	// HwlocObjPackage Physical package.
	// The physical package that usually gets inserted
	// into a socket on the motherboard.
	// A processor package usually contains multiple cores.
	HwlocObjPackage
	// HwlocObjCore Core
	// A computation unit (may be shared by several logical processors).
	HwlocObjCore
	// HwlocObjPU Processing Unit, or (Logical) Processor.
	// An execution unit (may share a core with some
	// other logical processors, e.g. in the case of an SMT core).
	// This is the smallest object representing CPU resources,
	// it cannot have any child except Misc objects.
	// Objects of this kind are always reported and can
	// thus be used as fallback when others are not.
	HwlocObjPU
	// HwlocObjL1Cache Level 1 Data (or Unified) Cache.
	HwlocObjL1Cache
	// HwlocObjL2Cache Level 2 Data (or Unified) Cache.
	HwlocObjL2Cache
	// HwlocObjL3Cache Level 3 Data (or Unified) Cache.
	HwlocObjL3Cache
	// HwlocObjL4Cache Level 4 Data (or Unified) Cache.
	HwlocObjL4Cache
	// HwlocObjL5Cache Level 5 Data (or Unified) Cache.
	HwlocObjL5Cache

	// HwlocObjL1ICache  Level 1 instruction Cache (filtered out by default).
	HwlocObjL1ICache
	// HwlocObjL2ICache Level 2 instruction Cache (filtered out by default).
	HwlocObjL2ICache
	// HwlocObjL3ICache Level 3 instruction Cache (filtered out by default).
	HwlocObjL3ICache

	// HwlocObjGroup Group objects.
	// Objects which do not fit in the above but are
	// detected by hwloc and are useful to take into
	// account for affinity. For instance, some operating systems
	// expose their arbitrary processors aggregation this
	// way.  And hwloc may insert such objects to group
	// NUMA nodes according to their distances.
	// See also \ref faq_groups.
	// These objects are removed when they do not bring
	// any structure (see ::HWLOC_TYPE_FILTER_KEEP_STRUCTURE).
	HwlocObjGroup

	// HwlocObjNumaNode NUMA node.
	// An object that contains memory that is directly
	// and byte-accessible to the host processors.
	// It is usually close to some cores (the corresponding objects
	// are descendants of the NUMA node object in the hwloc tree).
	// This is the smallest object representing Memory resources,
	// it cannot have any child except Misc objects.
	// However it may have Memory-side cache parents.
	// There is always at least one such object in the topology
	// even if the machine is not NUMA.
	// Memory objects are not listed in the main children list,
	// but rather in the dedicated Memory children list.
	// NUMA nodes have a special depth ::HWLOC_TYPE_DEPTH_NUMANODE
	// instead of a normal depth just like other objects in the main tree.
	HwlocObjNumaNode

	// HwlocObjBridge Bridge (filtered out by default).
	// Any bridge that connects the host or an I/O bus,
	// to another I/O bus.
	// They are not added to the topology unless I/O discovery
	// is enabled with hwloc_topology_set_flags().
	// I/O objects are not listed in the main children list,
	// but rather in the dedicated io children list.
	// I/O objects have NULL CPU and node sets.
	HwlocObjBridge
	// HwlocObjPCIDevice PCI device (filtered out by default).
	// They are not added to the topology unless I/O discovery
	// is enabled with hwloc_topology_set_flags().
	// I/O objects are not listed in the main children list,
	// but rather in the dedicated io children list.
	// I/O objects have NULL CPU and node sets.
	HwlocObjPCIDevice
	// HwlocObjOSDevice Operating system device (filtered out by default).
	// They are not added to the topology unless I/O discovery
	// is enabled with hwloc_topology_set_flags().
	// I/O objects are not listed in the main children list,
	// but rather in the dedicated io children list.
	// I/O objects have NULL CPU and node sets.
	HwlocObjOSDevice

	// HwlocObjMisc Miscellaneous objects (filtered out by default).
	// Objects without particular meaning, that can e.g. be
	// added by the application for its own use, or by hwloc
	// for miscellaneous objects such as MemoryModule (DIMMs).
	// These objects are not listed in the main children list,
	// but rather in the dedicated misc children list.
	// Misc objects may only have Misc objects as children,
	// and those are in the dedicated misc children list as well.
	// Misc objects have NULL CPU and node sets.
	HwlocObjMisc

	// HwlocObjMemCache Memory-side cache (filtered out by default).
	// A cache in front of a specific NUMA node.
	// This object always has at least one NUMA node as a memory child.
	// Memory objects are not listed in the main children list,
	// but rather in the dedicated Memory children list.
	// Memory-side cache have a special depth ::HWLOC_TYPE_DEPTH_MEMCACHE
	// instead of a normal depth just like other objects in the
	// main tree.
	HwlocObjMemCache
)

type HwlocObjCacheType int

const (
	// HwlocObjCacheUnified Unified cache.
	HwlocObjCacheUnified HwlocObjCacheType = iota
	// HwlocObjCacheData Data cache.
	HwlocObjCacheData
	// HwlocObjCacheInstruction Instruction cache (filtered out by default).
	HwlocObjCacheInstruction
)

type HwlocObjBridgeType int

const (
	// HwlocObjBridgeHost Host-side of a bridge, only possible upstream.
	HwlocObjBridgeHost HwlocObjBridgeType = iota
	// HwlocObjBridgePCI PCI-side of a bridge.
	HwlocObjBridgePCI
)

type HwlocObjOSDevType int

const (
	// HwlocObjOSDevBlock Operating system block device, or non-volatile memory device.
	// For instance "sda" or "dax2.0" on Linux.
	HwlocObjOSDevBlock HwlocObjOSDevType = iota
	// HwlocObjOSDevGPU Operating system GPU device.
	// For instance ":0.0" for a GL display, "card0" for a Linux DRM device.
	HwlocObjOSDevGPU
	// HwlocObjOSDevNetwork Operating system network device.
	// For instance the "eth0" interface on Linux.
	HwlocObjOSDevNetwork
	// HwlocObjOSDevOpenFabrics Operating system openfabrics device.
	// For instance the "mlx4_0" InfiniBand HCA, or "hfi1_0" Omni-Path interface on Linux.
	HwlocObjOSDevOpenFabrics
	// HwlocObjOSDevDMA Operating system dma engine device.
	// For instance the "dma0chan0" DMA channel on Linux.
	HwlocObjOSDevDMA
	// HwlocObjOSDevCoproc Operating system co-processor device.
	// For instance "mic0" for a Xeon Phi (MIC) on Linux, "opencl0d0" for a OpenCL device, "cuda0" for a CUDA device.
	HwlocObjOSDevCoproc
)

// HwlocNumaNodeAttr NUMA node-specific Object Attributes
type HwlocNumaNodeAttr struct {
	LocalMemory     uint64
	PageTypesLength uint
}

// HwlocCacheAttr Cache-specific Object Attributes
type HwlocCacheAttr struct {
	Size          uint64
	Depth         uint
	LineSize      uint
	Associativity int
	Type          HwlocObjCacheType
}

// HwlocGroupAttr Group-specific Object Attribute
type HwlocGroupAttr struct {
	// Depth Depth of group object, It may change if intermediate Group objects are added.
	Depth uint
	// Kind Internally-used kind of group.
	Kind uint
	// SubKind Internally-used subkind to distinguish different levels of groups with same kind.
	SubKind uint
}

// HwlocPCIDevAttr PCI Device specific Object Attributes
type HwlocPCIDevAttr struct {
	Domain      uint
	Bus         string
	Dev         string
	Func        string
	ClassID     uint
	VednorID    uint
	DeviceID    uint
	SubVendorID uint
	SubDeviceID uint
	Revision    string
	LinkSpeed   float64 // in GB/s
}

// HwlocBridgeAttr specific Object Attribues
type HwlocBridgeAttr struct {
	UpstreamPCI                 *HwlocPCIDevAttr
	UpstreamType                HwlocObjBridgeType
	DownStreamPCIDomain         uint
	DownStreamPCISecondaryBus   string
	DownStreamPCISubordinateBus string
	DownStreamType              HwlocObjBridgeType
	Depth                       uint
}

// HwlocObjAttr Object type-specific Attributes
type HwlocObjAttr struct {
	NumaNode  *HwlocNumaNodeAttr
	Cache     *HwlocCacheAttr
	Group     *HwlocGroupAttr
	PCIDev    *HwlocPCIDevAttr
	Bridge    *HwlocBridgeAttr
	OSDevType HwlocObjOSDevType
}

// HwlocObject Structure of a topology object
type HwlocObject struct {
	// HwlocObjType Type of object
	Type HwlocObjType
	// Subtype string to better describe the type field
	SubType string
	// OSIndex OS-provided physical index number.
	// It is not guaranteed unique across the entire machine, except for PUs and NUMA nodes.
	// Set to HWLOC_UNKNOWN_INDEX if unknown or irrelevant for this object.
	OSIndex uint
	// Name Object-specific name if any.
	// Mostly used for identifying OS devices and Misc objects where
	// a name string is more useful than numerical indexes.
	Name string
	// TotalMemory Total memory (in bytes) in NUMA nodes below this object.
	TotalMemory uint64
	// Attributes Object type-specific Attributes, may be NULL if no attribute value was found global position.
	Attributes *HwlocObjAttr
	// Depth Vertical index in the hierarchy.
	// For normal objects, this is the depth of the horizontal level
	// that contains this object and its cousins of the same type.
	// If the topology is symmetric, this is equal to the parent depth
	// plus one, and also equal to the number of parent/child links
	// from the root object to here.
	// For special objects (NUMA nodes, I/O and Misc) that are not
	// in the main tree, this is a special negative value that
	// corresponds to their dedicated level,
	// see hwloc_get_type_depth() and ::hwloc_get_type_depth_e.
	// Those special values can be passed to hwloc functions such
	// hwloc_get_nbobjs_by_depth() as usual.
	Depth int
	// LogicalIndex Horizontal index in the whole list of similar objects,
	// hence guaranteed unique across the entire machine.
	// Could be a "cousin_rank" since it's the rank within the "cousin" list below
	// Note that this index may change when restricting the topology
	// or when inserting a group.
	LogicalIndex uint
	// NextCousin Next object of same type and depth
	NextCousin *HwlocObject
	// PrevCousin Previous object of same type and depth
	PrevCousin  *HwlocObject
	Parent      *HwlocObject
	SiblingRank uint
	NextSibling *HwlocObject
	PrevSibling *HwlocObject
	// Arity Number of normal children.
	// Memory, Misc and I/O children are not listed here
	// but rather in their dedicated children list.
	Arity uint

	Children   []*HwlocObject
	FirstChild *HwlocObject
	LastChild  *HwlocObject
	// Set if the subtree of normal objects below this object is symmetric,
	// which means all normal children and their children have identical subtrees.
	// Memory, I/O and Misc children are ignored.
	// If set in the topology root object, lstopo may export the topology as a synthetic string.
	SymmetricSubTree int

	MemoryArity      uint
	MemoryFirstChild *HwlocObject
	IOArity          uint
	IOFirstChild     *HwlocObject
	MiscArity        uint
	MiscFirstChild   *HwlocObject
	CPUSet           []byte
	CompleteCPUSet   []byte
	NodeSet          []byte
	CompleteNodeSet  []byte
	// Infos Array of stringified info type=name.
	Infos map[string]string

	// misc

	// UserData Application-given private data pointer,
	// initialized to \c NULL, use it as you wish.
	UserData []byte
}
