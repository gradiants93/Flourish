/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type Order struct {
	id          int    `json:"id"`
	customer_id string `json:"customer_Id"`
	date        string `json:"date"`
	qty_ordered int    `json:"qty_Ordered"`
	total_price int    `json:"total_Price"`
}

// getAllOrderCmd represents the getAllOrder command
var getAllOrderCmd = &cobra.Command{
	Use:   "getAllOrder",
	Short: "Get all orders",
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

		results, err := db.Query("SELECT * FROM `order`")

		fmt.Println("Getting list of orders...")

		// Get the data
		if err != nil {
			fmt.Println(err)
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
	rootCmd.AddCommand(getAllOrderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAllOrderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAllOrderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
