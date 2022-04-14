package main

import (
	"context"

	"github.com/pPrecel/gardener-agent/internal/command"
	"github.com/pPrecel/gardener-agent/internal/command/serve"
	"github.com/pPrecel/gardener-agent/internal/command/state"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	o := &command.Options{
		Context: context.Background(),
		Logger:  logrus.New(),
	}

	cmd := &cobra.Command{
		Use: "gardenagent",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if o.Verbose == true {
				o.Logger.SetLevel(logrus.DebugLevel)
			}
		},
	}

	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", false, "Displays details of actions triggered by the command.")

	cmd.AddCommand(
		serve.NewCmd(serve.NewOptions(o)),
		state.NewCmd(state.NewOptions(o)),
	)

	err := cmd.Execute()
	if err != nil {
		o.Logger.Fatal(err)
	}
}
