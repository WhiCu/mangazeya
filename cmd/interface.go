/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/WhiCu/mangazeya/internal/core/inter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// interfaceCmdCmd represents the interfaceCmd command
var interfaceCmd = &cobra.Command{
	Use: "interface",
	Aliases: []string{
		"i",
		"in",
		"int",
		"inte",
		"inter",
		"interf",
		"itf",
		"itfc",
	},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Version: "0.0.1",

	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		itf, err := inter.Interfaces()
		if err != nil {
			return err
		}
		for _, i := range itf {
			fmt.Println(i.Addrs)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(interfaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interfaceCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interfaceCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	var err error

	interfaceCmd.Flags().StringP("name", "n", "", "name")
	err = viper.BindPFlag("name", interfaceCmd.Flags().Lookup("name"))
	if err != nil {
		panic(err)
	}

	interfaceCmd.Flags().BoolP("mtu", "m", false, "MTU")
	err = viper.BindPFlag("mtu", interfaceCmd.Flags().Lookup("mtu"))
	if err != nil {
		panic(err)
	}

	interfaceCmd.Flags().BoolP("flags", "f", false, "Flags")
	err = viper.BindPFlag("flags", interfaceCmd.Flags().Lookup("flags"))
	if err != nil {
		panic(err)
	}

	interfaceCmd.Flags().BoolP("addrs", "a", false, "Addrs")
	err = viper.BindPFlag("addrs", interfaceCmd.Flags().Lookup("addrs"))
	if err != nil {
		panic(err)
	}

	interfaceCmd.Flags().BoolP("hardwareAddr", "w", false, "Hardware address")
	err = viper.BindPFlag("hardwareAddr", interfaceCmd.Flags().Lookup("hardwareAddr"))
	if err != nil {
		panic(err)
	}

	interfaceCmd.Flags().BoolP("json", "j", false, "JSON")
	err = viper.BindPFlag("json", interfaceCmd.Flags().Lookup("json"))
	if err != nil {
		panic(err)
	}
}
