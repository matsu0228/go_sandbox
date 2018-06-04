package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// // Execute は、サブコマンドを登録しておく関数
// func Execute() {
// 	rootCmd := NewCmdRoot()
// 	// cmd.SetOutput(os.Stdout)
// 	if err := rootCmd.Execute(); err != nil {
// 		log.Fatal(err)
// 		// cmd.SetOutput(os.Stderr)
// 		// cmd.Println(err)
// 		// os.Exit(1)
// 	}
// }

// NewCmdRoot は、rootコマンドを作成するconstoractor
func NewCmdRoot() *cobra.Command {

	// Commandのオプションについてはここ
	//  https://godoc.org/github.com/spf13/cobra#Command
	rootCmd := &cobra.Command{
		Use:   "cobra-cli",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
    example:
      hoge fuga
    `,
	}

	// configの読み込み
	cobra.OnInitialize(initConfig)

	// グローバル変数: コマンドオプションを変数にセット。全てのサブコマンドで使える
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-cli.yaml)")

	// サブコマンドの登録
	rootCmd.AddCommand(NewCmdShow())
	return rootCmd
}

// initConfig は、「confFile or 環境変数」から読み込み
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	//  } else { // homeディレクトリ配下の「.cobra-cli」を読み込み
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".cobra-cli")
	// }
	viper.AutomaticEnv() // 環境変数から読み込み

	// confFileの読み込み
	if err := viper.ReadInConfig(); err == nil {
		log.Print("Using config file:", viper.ConfigFileUsed())
	}
}
