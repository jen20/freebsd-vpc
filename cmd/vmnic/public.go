package vmnic

import (
	"github.com/pkg/errors"
	"github.com/sean-/vpc/cmd/vmnic/create"
	"github.com/sean-/vpc/cmd/vmnic/destroy"
	"github.com/sean-/vpc/cmd/vmnic/list"
	"github.com/sean-/vpc/internal/command"
	"github.com/spf13/cobra"
)

const _CmdName = "vmnic"

var Cmd = &command.Command{
	Name: _CmdName,

	Cobra: &cobra.Command{
		Use:     _CmdName,
		Aliases: []string{"nic", "if", "iface"},
		Short:   "VM network interface management",
	},

	Setup: func(self *command.Command) error {
		subCommands := command.Commands{
			create.Cmd,
			destroy.Cmd,
			list.Cmd,
		}

		if err := self.Register(subCommands); err != nil {
			return errors.Wrapf(err, "unable to register sub-commands under %s", _CmdName)
		}

		return nil
	},
}
