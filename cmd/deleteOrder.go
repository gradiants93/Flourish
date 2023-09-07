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

// deleteOrderCmd represents the deleteOrder command
var deleteOrderCmd = &cobra.Command{
	Use:   "deleteOrder",
	Short: "delete order by id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//if len(args) < 1 {
		//	log.Fatal("Please enter an order id to delete")
		//} else if len(args) > 1 {
		//	log.Fatal("Please enter a single id")
		//}
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/flourish")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		orderId, _ := cmd.Flags().GetInt("orderId")
		sql := "DELETE FROM `order`WHERE id = ?"

		results, err := db.ExecContext(context.Background(), sql, orderId)

		// Get the data
		if err != nil {
			panic(err.Error())
		}

		affectedRows, err := results.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted %d rows\n", affectedRows)
	},
}

func init() {
	rootCmd.AddCommand(deleteOrderCmd)
	deleteOrderCmd.PersistentFlags().Int("orderId", 0, "Order id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteOrderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteOrderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
