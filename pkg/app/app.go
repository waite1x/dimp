package app

import "github.com/spf13/cobra"

type App struct {
	RootCmd *cobra.Command
}

func NewApp(rootCmd *cobra.Command) *App {
	return &App{
		RootCmd: rootCmd,
	}
}

func (a *App) Run() error {
	a.RootCmd.Execute()
	return nil
}
