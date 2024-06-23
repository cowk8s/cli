// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display this binary's version, build time, and git hash of this build",
	Run: func(cmd *cobra.Command, args []string) {
		
	}
}