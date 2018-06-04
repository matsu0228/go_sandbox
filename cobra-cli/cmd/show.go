package cmd

import (
	"github.com/spf13/cobra"
)

// NewCmdShow は、rootコマンドを作成するconstoractor
func NewCmdShow() *cobra.Command {
	type Options struct {
		Optint int    `validate:"min=0,max=10"`
		Optstr string `validate:"required,alphanum"`
	}
	var (
		o = &Options{}
	)

	showCmd := &cobra.Command{
		Use:   "show",
		Short: "A brief description of your command",
		// command実行前に呼ばれる
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateShowOpts(*o)
		},
		// command実行内容
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("show called: optint: %d, optstr: %s \n", o.Optint, o.Optstr)
		},
		// for tests
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	showCmd.Flags().IntVarP(&o.Optint, "int", "i", 0, "int option")
	showCmd.Flags().StringVarP(&o.Optstr, "str", "s", "", "string option")

	return showCmd
}
