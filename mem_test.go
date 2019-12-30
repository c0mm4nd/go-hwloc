package hwloc

import "testing"

func TestHwlocSetMemBind(t *testing.T) {
	topology, _ := NewTopology()
	topology.Load()
	nodeSet := HwlocGetNUMANodeObjByOSIndex(topology, 0)
	HwlocSetMemBind(topology, nodeSet, HwlocMemBindBind, HwlocMemBindThread|HwlocMemBindByNodeSet)
}
