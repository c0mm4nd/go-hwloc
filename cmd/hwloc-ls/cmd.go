/*
Copyright 2019 All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/carmark/gohwloc/topology"
)

type lstopoOptions struct {
	PS      bool
	Version bool
	Verbose bool
}

func newCommand() *cobra.Command {
	opts := lstopoOptions{}
	var cmd = &cobra.Command{
		Use:     "hwloc-ls",
		Aliases: []string{"lstopo"},
		Short:   "Show the topology of the system",
		Long:    `Show the topology of the system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Do Stuff Here
			return lstopo(opts)
		},
	}
	flags := cmd.Flags()
	flags.BoolVar(&opts.PS, "ps", false, "Display processes within the hierarchy")
	flags.BoolVar(&opts.Version, "version", false, "Report version and exit")
	flags.BoolVarP(&opts.Verbose, "verbose", "v", false, "Include additional details")
	return cmd
}

func lstopo(opts lstopoOptions) error {
	t, _ := topology.NewTopology()
	t.Load()
	defer t.Destroy()
	n, _ := t.GetNbobjsByType(topology.HwlocObjOSDevice)
	fmt.Printf("t.GetNbobjsByType: %d\n", n)
	for i := 0; i < n; i++ {
		obj, _ := t.GetObjByType(topology.HwlocObjOSDevice, uint(i))
		fmt.Printf("%s:\n", obj.Name)
		// obj->subtype also contains CUDA or OpenCL since v2.0
		s, _ := obj.GetInfo("Backend")
		//fmt.Printf("Info: %s\n", s)
		if s == "CUDA" {
			// TODO This is a CUDA device
		} else if s == "OpenCL" {
			// This is an OpenCL device
			platformid, err := strconv.Atoi(obj.Name[len("opencl") : len("opencl")+1])
			if err != nil {
				return err
			}
			fmt.Printf("OpenCL platform %d\n", platformid)
			devid, err := strconv.Atoi(obj.Name[len("opencl0d"):])
			if err != nil {
				return err
			}
			fmt.Printf("OpenCL device %d\n", devid)
			s, _ = obj.GetInfo("GPUModel")
			if s != "" {
				fmt.Printf("Model: %s\n", s)
			}
			s, _ = obj.GetInfo("OpenCLGlobalMemorySize")
			if s != "" {
				fmt.Printf("Memory: %s\n", s)
			}
		}
		for {
			if obj != nil {
				if obj.CPUSet == nil {
					obj = obj.Parent
				} else if iz, _ := obj.CPUSet.IsZero(); iz {
					obj = obj.Parent
				} else {
					break
				}
			} else {
				break
			}
		}
		/* Find out cpuset this is connected to */
		if obj != nil {
			fmt.Printf("Location: %v P#%d\n", obj.Type.String(), obj.OSIndex)
			fmt.Printf("Cpuset: %s\n", obj.CPUSet.String())
		}
		fmt.Printf("\n")
	}
	return nil
}
