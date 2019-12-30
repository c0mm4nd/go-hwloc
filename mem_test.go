package hwloc

import "testing"

func TestHwlocSetMemBind(t *testing.T) {
	topology, _ := NewTopology()
	nodeSet := HwlocGetNUMANodeObjByOSIndex(topology, 0)
	HwlocSetMemBind(topology, nodeSet, HwlocMemBindBind, HwlocMemBindThread|HwlocMemBindByNodeSet)
}
