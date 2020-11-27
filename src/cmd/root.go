package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version    Version
	vaultSrv   string
	vaultToken string
	verbose    bool
	monoignore []string
	rootCmd    = &cobra.Command{
		Use:   "vaultmv (version)",
		Short: "Move or extend your vault path with ease.",
	}
)

// Execute executes the root command.
func Execute(v Version) error {
	version = v
	return rootCmd.Execute()
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func init() {
	cobra.OnInitialize(initConfig)
	log.SetFormatter(&log.TextFormatter{})
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&vaultSrv, "server", "s", "",  "server address of Hashicorp Vault")
	rootCmd.PersistentFlags().StringVarP(&vaultToken, "token", "t", "",  "server auth token for Hashicorp Vault")
}

func initConfig() {
	/*
		Based on persisten flag we need to define loglevel
	*/
	if verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug log output has been turned on")
	}
	viper.ConfigFileUsed()
}