package hwloc

import "testing"

func TestHwlocSetMemBind(t *testing.T) {
	topology, _ := NewTopology()
	topology.Load()
	nodeSet := topology.HwlocGetNUMANodeObjByOSIndex(0)
	topology.HwlocSetMemBind(nodeSet, HwlocMemBindBind, HwlocMemBindThread|HwlocMemBindByNodeSet)
}
