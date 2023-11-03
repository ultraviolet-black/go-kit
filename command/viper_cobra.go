package command

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func PostInitCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		presetRequiredFlags(cmd)
		if cmd.HasSubCommands() {
			PostInitCommands(cmd.Commands())
		}
	}
}

func presetRequiredFlags(cmd *cobra.Command) {
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().VisitAll(func(f *pflag.Flag) {

		name := strings.ReplaceAll(f.Name, "-", "_")

		if viper.IsSet(name) {

			val := viper.GetString(name)

			if val != "" {
				cmd.Flags().Set(f.Name, val)
			}

		}

	})
}
