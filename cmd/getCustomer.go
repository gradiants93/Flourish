/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "getCustomer",
	Short: "Gets customer info via id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please enter a customer id to search for")
		} else if len(args) > 1 {
			log.Fatal("Please enter a single id")
		}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		customerID, err := strconv.Atoi(args[0])

		results, err := db.Query("SELECT * FROM customer WHERE id = ?", customerID)

		fmt.Printf("Trying to get customer ID: %d\n", customerID)

		// Get the data
		if err != nil {
			fmt.Println(err)
		}
		if results.Next() {

			var customer Customer
			err := results.Scan(&customer.id, &customer.name, &customer.email)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", customer)
		} else {

			fmt.Println("No customer found")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
