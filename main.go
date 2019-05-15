package main

import (
	"fmt"
	"log"

	"github.com/carmark/gohwloc/topology"
)

func main() {
	t, err := topology.NewTopology()
	if err != nil {
		log.Fatal(err)
	}
	err = t.Load()
	if err != nil {
		log.Fatal(err)
	}
	defer t.Destroy()
	nbcores, err := t.GetNbobjsByType(topology.HwlocObjCore)
	if err != nil {
		log.Fatal(err)
	}
	obj, err := t.GetRootObj()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cores: %d\nDepth: %d\n%#v\n", nbcores, t.Depth, obj)
	res, _ := obj.GetInfo("")
	fmt.Printf("Info: %v\n", res)
}
