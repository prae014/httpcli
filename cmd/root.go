/*
Copyright © 2023 github.com/prae014
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpcli",
	Short: "httpcli is a HTTP client command line program",
	Long: `httpcli is a HTTP client command line program developed in GO
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.httpcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//test persistent flag
	//rootCmd.PersistentFlags().StringVar(&head, "header", "", "", "return the specified header")

	//note: StringSliceVarP has a shorthand name, StringSliceP does not have
	rootCmd.PersistentFlags().StringSliceVarP(&head_flags, "header", "H", []string{}, "return specified header")
	rootCmd.PersistentFlags().StringSliceVarP(&query_flags, "query", "q", []string{}, "return specified query")

}
