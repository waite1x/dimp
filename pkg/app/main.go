package app

import "github.com/spf13/cobra"

type AppContext struct {
	cmdBuidlers []func() *cobra.Command
}

func (a *AppContext) AddCmd(cmdCreator func() *cobra.Command) {
	a.cmdBuidlers = append(a.cmdBuidlers, cmdCreator)
}

func (a *AppContext) Build() (*App, error) {
	rootCmd := &cobra.Command{
		Use:   "dexp",
		Short: "dexp",
		Long:  "dexp",
	}

	for _, cmdBuilder := range a.cmdBuidlers {
		cmd := cmdBuilder()
		rootCmd.AddCommand(cmd)
	}
	return NewApp(rootCmd), nil
}
