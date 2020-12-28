package cmd

import (
	"consulapi/consul"
	"consulapi/env"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push cmd",
	Run: func(cmd *cobra.Command, args []string) {
		push()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	pushCmd.Flags().StringVarP(&env.ConsulHost, "consul server", "s", "127.0.0.1:8500", "consul address")
	pushCmd.Flags().StringVarP(&env.ConsulClient, "consul client", "c", "", "consul client address")
}

func push() {
	// 來源 - 目的
	consul.CloneKV(env.ConsulHost, env.ConsulClient)
}
