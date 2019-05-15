package hwloc

import "testing"

func TestNewTopology(t *testing.T) {
	topo, _ := NewTopology()
	topo.Load()
	topo.Destroy()
}
