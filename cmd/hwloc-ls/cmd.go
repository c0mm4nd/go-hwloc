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
	"github.com/spf13/cobra"
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
	return nil
}
