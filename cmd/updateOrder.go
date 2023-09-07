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
		//if len(args) != 5 {
		//	log.Fatal("Please enter an order id, customer id, date, quantity and total price")
		//}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		orderId, _ := cmd.Flags().GetInt("orderId")
		customerId, _ := cmd.Flags().GetInt("customerId")
		date, _ := cmd.Flags().GetString("date")
		quantity, _ := cmd.Flags().GetInt("quantity")
		totalPrice, _ := cmd.Flags().GetInt("totalPrice")

		resultsCustomer, err := db.Query("SELECT * FROM customer WHERE id = ?", customerId)

		// Get the data
		if err != nil {
			fmt.Println(err)
		}
		if resultsCustomer.Next() {
			sql := "UPDATE `order` SET customer_id = ?, date = ?, qty_ordered = ?, total_price = ? WHERE id = ?"
			results, err := db.ExecContext(context.Background(), sql, customerId, date, quantity, totalPrice, orderId)

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

			fmt.Printf("Customer %d does not exist", customerId)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateOrderCmd)
	updateOrderCmd.PersistentFlags().Int("orderId", 0, "Order id")
	updateOrderCmd.PersistentFlags().Int("customerId", 0, "Customer id")
	updateOrderCmd.PersistentFlags().String("date", "2000-01-01", "Date - YYYY-MM-DD")
	updateOrderCmd.PersistentFlags().Int("quantity", 0, "Quantity")
	updateOrderCmd.PersistentFlags().Int("totalPrice", 0, "Total Price")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateOrderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateOrderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
