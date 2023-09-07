/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// updateCustomerCmd represents the updateCustomer command
var updateCustomerCmd = &cobra.Command{
	Use:   "updateCustomer",
	Short: "update customer name and email by id (id, name, email)",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Fatal("Please enter a customer id, name and email")
		}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		var customerId = args[0]
		var customerName = args[1]
		var customerEmail = args[2]
		sql := "UPDATE customer SET name=?, email=? WHERE id=?"
		results, err := db.ExecContext(context.Background(), sql, customerName, customerEmail, customerId)

		// Get the data
		if err != nil {
			log.Fatalf("Could not update customer: %s", err)
			fmt.Println(err)
		}
		affectedRows, err := results.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated %d rows\n", affectedRows)
	},
}

func init() {
	rootCmd.AddCommand(updateCustomerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCustomerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCustomerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
