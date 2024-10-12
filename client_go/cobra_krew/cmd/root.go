/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var KubernetesConfigFlags *genericclioptions.ConfigFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8simg",
	Short: "demo for cobra",
	Long:  ``,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.k8simg.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	imageCmd.Flags().BoolP("deployments", "d", false, "show deployments image")
	imageCmd.Flags().BoolP("daemonsets", "e", false, "show daemonsets image")
	imageCmd.Flags().BoolP("statefulsets", "f", false, "show statefulsets image")
	imageCmd.Flags().BoolP("jobs", "o", false, "show jobs image")
	imageCmd.Flags().BoolP("cronjobs", "b", false, "show cronjobs image")
	imageCmd.Flags().BoolP("json", "j", false, "show json format")

	KubernetesConfigFlags = genericclioptions.NewConfigFlags(true)
	KubernetesConfigFlags.AddFlags(rootCmd.PersistentFlags())
}
