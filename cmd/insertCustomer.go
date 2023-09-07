/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// insertCustomerCmd represents the insertCustomer command
var insertCustomerCmd = &cobra.Command{
	Use:   "insertCustomer",
	Short: "Insert a new customer (name and email)",
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

		customerName, _ := cmd.Flags().GetString("name")
		customerEmail, _ := cmd.Flags().GetString("email")

		if customerName == "" && customerEmail == "" {
			log.Fatal("Please enter a customer name and email")
		} else if customerName == "" {
			log.Fatal("Please enter a customer name")
		} else if customerEmail == "" {
			log.Fatal("Please enter a customer email")

		}

		sql := "INSERT INTO customer(name, email) VALUES (?, ?)"
		results, err := db.ExecContext(context.Background(), sql, customerName, customerEmail)

		// Get the data
		if err != nil {
			log.Fatalf("Could not insert customer: %s", err)
			fmt.Println(err)
		}
		id, err := results.LastInsertId()
		if err != nil {
			log.Fatalf("impossible to retrieve last inserted id: %s", err)
		}
		fmt.Println("inserted customer id: %d", id)
	},
}

func init() {
	rootCmd.AddCommand(insertCustomerCmd)
	insertCustomerCmd.PersistentFlags().String("name", "", "Customer Name")
	insertCustomerCmd.PersistentFlags().String("email", "", "Customer Email")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCustomerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCustomerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
