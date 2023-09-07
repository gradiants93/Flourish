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

// updateOrderCmd represents the updateOrder command
var updateOrderCmd = &cobra.Command{
	Use:   "updateOrder",
	Short: "update order by id (order id, customer id, date, quantity and total price)",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 5 {
			log.Fatal("Please enter an order id, customer id, date, quantity and total price")
		}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		var order_Id = args[0]
		var customer_Id = args[1]
		var date = args[2]
		var qty_Ordered = args[3]
		var total_Price = args[4]

		resultsCustomer, err := db.Query("SELECT * FROM customer WHERE id = ?", customer_Id)

		// Get the data
		if err != nil {
			fmt.Println(err)
		}
		if resultsCustomer.Next() {
			sql := "UPDATE `order` SET customer_id = ?, date = ?, qty_ordered = ?, total_price = ? WHERE id = ?"
			results, err := db.ExecContext(context.Background(), sql, customer_Id, date, qty_Ordered, total_Price, order_Id)

			// Get the data
			if err != nil {
				log.Fatalf("Could not update order: %s", err)
				fmt.Println(err)
			}
			affectedRows, err := results.RowsAffected()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Updated %d rows\n", affectedRows)
		} else {

			fmt.Printf("Customer %s does not exist", customer_Id)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateOrderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateOrderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateOrderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
