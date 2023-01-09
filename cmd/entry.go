/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	EntryCmd.AddCommand(createEntryCmd)
	EntryCmd.AddCommand(updateEntryCmd)
	EntryCmd.AddCommand(readEntryCmd)
	EntryCmd.AddCommand(deleteEntryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// entryCmd represents the entry command
var EntryCmd = &cobra.Command{
	Use:   "entry",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createEntryCmd = &cobra.Command{
	Use:     "create-entry",
	Short:   "",
	Example: exampleQuery("create-entry"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createEntryCmd called")
	},
}

var updateEntryCmd = &cobra.Command{
	Use:     "update-entry",
	Short:   "",
	Example: exampleQuery("update-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateEntryCmd called")
	},
}

var readEntryCmd = &cobra.Command{
	Use:     "read-entry",
	Short:   "",
	Example: exampleQuery("read-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readEntryCmd called")
	},
}

var deleteEntryCmd = &cobra.Command{
	Use:     "delete-entry",
	Short:   "",
	Example: exampleQuery("delete-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteEntryCmd called")
	},
}
