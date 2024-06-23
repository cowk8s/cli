// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package cloudx

import (
	"github.com/spf13/cobra"

	"github.com/ory/cli/cmd/cloudx/project"
)

func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get",
		Short: "Get a resource",
	}

	cmd.AddCommand(

	)
}
