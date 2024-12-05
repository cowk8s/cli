// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ory/cli/cmd/cloudx"
	"github.com/ory/cli/cmd/cloudx/client"
	"github.com/ory/cli/cmd/cloudx/proxy"
	"github.com/ory/kratos/cmd/jsonnet"
	"github.com/ory/x/cmdx"
)

var commandTemplatingOnce sync.Once

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "ory",
		Short: "The Ory CLI",
	}

	c.AddCommand(devCommands...)
	commandTemplatingOnce.Do(func() {
		cmdx.EnableUsageTemplating(c)
	})

	return c
}

