/*
Copyright © 2023 Har Preet Singh <singhhp1069@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	urlPath string
)

func init() {
	createTableCmd.Flags().StringVarP(&urlPath, "url-path", "u", "", "URL path")
	if err := createTableCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
	TableCmd.AddCommand(createTableCmd)
	TableCmd.AddCommand(updateTableCmd)
	TableCmd.AddCommand(deleteTableCmd)
	TableCmd.AddCommand(readTableCmd)
	TableCmd.AddCommand(importTableCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tablesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// tablesCmd represents the tables command
var TableCmd = &cobra.Command{
	Use:   "table",
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

var createTableCmd = &cobra.Command{
	Use:     "create-table",
	Short:   "",
	Example: exampleQuery("create-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createTable called")
	},
}

var updateTableCmd = &cobra.Command{
	Use:     "update-table",
	Short:   "",
	Example: exampleQuery("update-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateTableCmd called")
	},
}

var readTableCmd = &cobra.Command{
	Use:     "read-table",
	Short:   "",
	Example: exampleQuery("read-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readTableCmd called")
	},
}

var deleteTableCmd = &cobra.Command{
	Use:     "delete-table",
	Short:   "",
	Example: exampleQuery("delete-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readTableCmd called")
	},
}

var importTableCmd = &cobra.Command{
	Use:     "import-table",
	Short:   "",
	Example: exampleQuery("import-table"),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readTableCmd called")
	},
}

// Relationship
// Raw SQL builder (future version)
func exampleQuery(cmd ...string) string {
	return fmt.Sprint("Example query ", cmd)
}
