package main

import (
	"os"

	// "github.com/ovrclk/akash/cmd/akash/cmd"

	// authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	// bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ovrclk/akash/app"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
	// "github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
    Use:   "akash",
    SilenceUsage: true,
		Short:        "Akash Blockchain Application",
		Long:         "Akash CLI Utility.\n\nAkash is a peer-to-peer marketplace for computing resources and \na deployment platform for heavily distributed applications. \nFind out more at https://akash.network",
    PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
      encodingConfig := app.MakeEncodingConfig()
      initClientCtx := client.Context{}.
        WithJSONMarshaler(encodingConfig.Marshaler).
        WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
        WithTxConfig(encodingConfig.TxConfig).
        WithLegacyAmino(encodingConfig.Amino).
        WithInput(os.Stdin).
        WithBroadcastMode(flags.BroadcastBlock).
        WithHomeDir(app.DefaultHome)
        
      return client.SetCmdClientContextHandler(initClientCtx, cmd)
    },
}



func init() {


    rootCmd.AddCommand(txCmd())
}

func main() {
    if err := rootCmd.Execute(); err != nil {
      os.Exit(1)
    }

  }