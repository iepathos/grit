/*
Copyright © 2023 Glen Baker <iepathos@gmail.com>
*/
package cmd

import (
	"log"

	conf "grit/config"
	git "grit/git"

	"sync"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config_path := conf.GetDefaultYml()
		paths := conf.ParseYml(config_path)

		var wg sync.WaitGroup
		wg.Add(len(paths))

		for localPath, remotePath := range paths {
			go git.PullRepository(localPath, remotePath, &wg)
		}

		wg.Wait()
		log.Println("pull complete")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
