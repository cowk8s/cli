// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/ory/x/flagx"
)

const projectFlag = "project"

func RegisterProjectFlag(f *flag.FlagSet) {
	f.String(projectFlag, "", "The project to use, either project ID or a (partial) slug.")
}

func ProjectOrDefault(cmd *cobra.Command, h *CommandHelper) (string, error) {
	if flag := h
}
