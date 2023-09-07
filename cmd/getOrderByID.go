/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// getOrderByIDCmd represents the getOrderByID command
var getOrderByIDCmd = &cobra.Command{
	Use:   "getOrderByID",
	Short: "get single order by id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please enter an order id to search for")
		} else if len(args) > 1 {
			log.Fatal("Please enter a single id")
		}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		orderID, err := strconv.Atoi(args[0])

		results, err := db.Query("SELECT * FROM `order` WHERE id = ?", orderID)

		fmt.Printf("Trying to get order ID: %d\n", orderID)

		// Get the data
		if err != nil {
			fmt.Println(err)
		}
		if results.Next() {

			var order Order
			err := results.Scan(&order.id, &order.customer_id, &order.date, &order.qty_ordered, &order.total_price)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", order)
		} else {

			fmt.Println("No order found")
		}
	},
}

func init() {
	rootCmd.AddCommand(getOrderByIDCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getOrderByIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getOrderByIDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
