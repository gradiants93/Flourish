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

// getOrderByCustomerCmd represents the getOrderByCustomer command
var getOrderByCustomerCmd = &cobra.Command{
	Use:   "getOrderByCustomer",
	Short: "gets all orders by customer id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//if len(args) < 1 {
		//	log.Fatal("Please enter a customer id to search for")
		//} else if len(args) > 1 {
		//	log.Fatal("Please enter a single id")
		//}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		customerId, _ := cmd.Flags().GetInt("customerId")

		results, err := db.Query("SELECT * FROM `order` WHERE customer_id = ?", customerId)

		fmt.Printf("Trying to get all orders for customer ID: %d\n", customerId)

		// Get the data
		if err != nil {
			fmt.Println(err)
		}
		if !results.Next() {

			fmt.Printf("No orders found for customer %d\n", customerId)
		}
		for results.Next() {

			var order Order
			err := results.Scan(&order.id, &order.customer_id, &order.date, &order.qty_ordered, &order.total_price)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", order)
		}
	},
}

func init() {
	rootCmd.AddCommand(getOrderByCustomerCmd)
	getOrderByCustomerCmd.PersistentFlags().Int("customerId", 0, "Customer id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getOrderByCustomerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getOrderByCustomerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
