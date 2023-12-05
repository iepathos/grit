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
	"github.com/spf13/viper"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Git pull in every repository in the specified config.  Clone if repository isn't found locally.",
	Long: `Grit will execute git pull in all of the repositories specified in the grit config.
	If the repository isn't found in the config path locally it will clone the repository
	from the remote path in the config.`,
	Run: func(cmd *cobra.Command, args []string) {
		// viper.SetConfigFile(
		config_path := viper.ConfigFileUsed()
		paths, err := conf.ParseYml(config_path)
		if err != nil {
			log.Fatalf("%v", err)
		}

		var wg sync.WaitGroup
		wg.Add(len(paths))
		errCh := make(chan error, 10)

		for localPath, remotePath := range paths {
			go git.PullRepository(localPath, remotePath, &wg, errCh)
		}

		go func() {
			wg.Wait()
			close(errCh)
		}()

		for err := range errCh {
			log.Fatalf("%v", err)
		}

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
