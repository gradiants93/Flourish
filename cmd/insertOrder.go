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

// insertOrderCmd represents the insertOrder command
var insertOrderCmd = &cobra.Command{
	Use:   "insertOrder",
	Short: "insert a new order (customerID, quantity, total price)",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Fatal("Please enter a customer id, quantity and total price")
		}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		var customer_Id = args[0]
		var qty_Ordered = args[1]
		var total_Price = args[2]
		sql := "INSERT INTO `order`(customer_id, date, qty_ordered, total_price) VALUES (?, NOW(), ?, ?)"
		results, err := db.ExecContext(context.Background(), sql, customer_Id, qty_Ordered, total_Price)

		// Get the data
		if err != nil {
			log.Fatalf("Could not insert order: %s", err)
			fmt.Println(err)
		}
		id, err := results.LastInsertId()
		if err != nil {
			log.Fatalf("impossible to retrieve last inserted id: %s", err)
		}
		fmt.Println("inserted order id: %d", id)
	},
}

func init() {
	rootCmd.AddCommand(insertOrderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertOrderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertOrderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
