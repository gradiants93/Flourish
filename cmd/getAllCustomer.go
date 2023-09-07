/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

type Customer struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	email string `json:"email"`
}

// getAllCustomerCmd represents the getAllCustomer command
var getAllCustomerCmd = &cobra.Command{
	Use:   "getAllCustomer",
	Short: "Gets all customer info",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		results, err := db.Query("SELECT * FROM customer")

		fmt.Println("Getting list of customers...")

		// Get the data
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {

			var customer Customer
			err := results.Scan(&customer.id, &customer.name, &customer.email)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", customer)
		}
	},
}

func init() {
	rootCmd.AddCommand(getAllCustomerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAllCustomerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAllCustomerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
