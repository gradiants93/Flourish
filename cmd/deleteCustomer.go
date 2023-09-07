/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCustomerCmd represents the deleteCustomer command
var deleteCustomerCmd = &cobra.Command{
	Use:   "deleteCustomer",
	Short: "delete customer and all related orders",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please enter a customer id to delete for")
		} else if len(args) > 1 {
			log.Fatal("Please enter a single id")
		}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		customerID, err := strconv.Atoi(args[0])
		sqlOrder := "DELETE FROM `order`WHERE customer_id = ?"

		resultsOrder, err := db.ExecContext(context.Background(), sqlOrder, customerID)

		// Get the data
		if err != nil {
			panic(err.Error())
		}

		affectedRowsOrder, err := resultsOrder.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted %d rows from orders\n", affectedRowsOrder)
		sqlCustomer := "DELETE FROM customer WHERE id = ?"
		resultsCustomer, err := db.ExecContext(context.Background(), sqlCustomer, customerID)
		if err != nil {
			panic(err.Error())
		}

		affectedRowsCustomer, err := resultsCustomer.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted %d rows from customers\n", affectedRowsCustomer)
	},
}

func init() {
	rootCmd.AddCommand(deleteCustomerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCustomerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCustomerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
