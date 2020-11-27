package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	vault_api "github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dataFile string;

var goСmd = &cobra.Command{
	Use:   "go (--source=[path] | --dest=[path] | --clean_source | --data=[csv])",
	Short: "Execute move",
	Args: cobra.MaximumNArgs(0),
	Long: `vaultmv go
============
Execute move of the path based on arguments
Complete documentation available on https://vaultmv.github.io/#go`,
	Run: func(goCmd *cobra.Command, args []string) {
		source, _ := goCmd.Flags().GetString("source")
		destination, _ := goCmd.Flags().GetString("dest")
		client := vaultAuth(vaultSrv, vaultToken)
		dataFile, _ = goCmd.Flags().GetString("data")

		data := make([]map[string]interface{}, 0)
		if (dataFile != "") {
			file, err := os.Open(dataFile)
			if err != nil {
				log.Error(err)
			}
			defer file.Close()
			reader := csv.NewReader(file)
			for {
				record, e := reader.Read()
				if e != nil {
					break
				}
				if (len(record) > 1) {
					data = append(data, map[string]interface{}{
						"source": record[0],
						"dest": record[1],
						"clean_source": record[2],
					})
				}
			}
		} else {
			data = append(data, map[string]interface{}{
				"source": source,
				"dest": destination,
				"clean_source": true,
			})
		}
		moveSecret(client, data)
	},
}

func init() {
	goСmd.Flags().String("source", "", "Source path at the Hashicorp Vault")
	goСmd.Flags().String("dest", "", "Destination path at the Hashicorp Vault")
	goСmd.Flags().Bool("clean_source", true, "Delete source path after move")
	goСmd.Flags().String("data", "", "Path to CSV data file for the bulk action")
	rootCmd.AddCommand(goСmd)
}

func vaultAuth(server string, token string) *vault_api.Client {
	
	// Detect server address configuration
	vaultServer := os.Getenv(vault_api.EnvVaultAddress); 
	
	if (server != "") {
		vaultServer = server
		log.Debugf("vaultmv will work with instance at %v passed as cli argument", vaultServer) 
	} else if (vaultServer != "") {
		log.Debugf("vaultmv will work with instance at %v configured as %v", vaultServer, vault_api.EnvVaultAddress)
	}	else {
		log.Error("VAULT_ADDR was not found at this environment, you can use --server flag")
		os.Exit(89) // Destination address is required
	}
	
	// Create client
	clientConfig := &vault_api.Config{Address: vaultServer}
	client, err := vault_api.NewClient(clientConfig)
	if err != nil {
		log.Error(err)
	}

	// Read Token configuration
	vaultToken := os.Getenv(vault_api.EnvVaultToken); 
	if (token != "") {
		log.Debug("vaultmv will use token as cli argument --token to authorize.")
		client.SetToken(token)
	} else if (vaultToken != "") {
		log.Debugf("vaultmv will use token from %v", vault_api.EnvVaultToken)
	} else {
		log.Error("VAULT_TOKEN was not found at this environment, you can use --token flag to authorize")
		os.Exit(89)
	}
	// fmt.Print(client.Logical().Read("/secrets/foo/bar"))
	return client
}

func moveSecret(client *vault_api.Client, data []map[string]interface{}) {
	// TODO: Current method is loop. Create same using go routines.
	c := client.Logical()
	for _, val := range data {
		secretData, err := c.Read(fmt.Sprintf("%v?version=1", val["source"]))
		if err != nil {
			log.Error(err)
		}
		fmt.Println(secretData.Data)
	}
}