package cmd

import (
	"fmt"

	"github.com/WhiCu/mangazeya/internal/core/network"
	tui "github.com/WhiCu/mangazeya/internal/tui/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ioCmd represents the io command
var ioCmd = &cobra.Command{
	Use:   "io",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Version: "0.0.1",

	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ios, err := network.Networks()
		if err != nil {
			return err
		}

		if viper.GetBool("tui") {
			t := tui.NewProgram(ios)
			if _, err := t.Run(); err != nil {
				return err
			}
			return nil
		}

		if name := viper.GetString("name"); name != "" {
			io, err := ios.Network(name)
			if err != nil {
				return err
			}

			if viper.GetBool("json") {
				json, err := io.JSON()
				if err != nil {
					return err
				}
				fmt.Println(string(json))
				return nil
			}

			if viper.GetBool("cooljson") {
				json, err := io.CoolJSON()
				if err != nil {
					return err
				}
				fmt.Println(string(json))
				return nil
			}

			fmt.Println(io)
			return nil
		}

		if viper.GetBool("json") {
			json, err := ios.JSON()
			if err != nil {
				return err
			}
			fmt.Println(string(json))
			return nil
		}

		if viper.GetBool("cooljson") {
			json, err := ios.CoolJSON()
			if err != nil {
				return err
			}
			fmt.Println(string(json))
			return nil
		}

		fmt.Println(ios)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ioCmd)

	ioCmd.Flags().StringP("name", "n", "", "name")
	err := viper.BindPFlag("name", ioCmd.Flags().Lookup("name"))
	if err != nil {
		panic(err)
	}

	ioCmd.Flags().BoolP("json", "j", false, "JSON")
	err = viper.BindPFlag("json", ioCmd.Flags().Lookup("json"))
	if err != nil {
		panic(err)
	}

	ioCmd.Flags().BoolP("cooljson", "c", false, "cooljson")
	err = viper.BindPFlag("cooljson", ioCmd.Flags().Lookup("cooljson"))
	if err != nil {
		panic(err)
	}

	ioCmd.Flags().BoolP("tui", "t", false, "tui")
	err = viper.BindPFlag("tui", ioCmd.Flags().Lookup("tui"))
	if err != nil {
		panic(err)
	}
}
