/*
Copyright © 2023 Glen Baker <iepathos@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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

var Branch string

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout <branch>",
	Short: "Git checkout the specified branch in every repository.",
	Long: `Execute git checkout <branch> in every repository
	in the given grit yaml.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Branch = args[0]
		// log.Println("checkout", Branch)
		config_path := viper.ConfigFileUsed()
		paths, err := conf.ParseYml(config_path)
		if err != nil {
			log.Fatalf("%v", err)
		}

		var wg sync.WaitGroup
		wg.Add(len(paths))
		errCh := make(chan error, 10)

		for localPath := range paths {
			go git.CheckoutBranch(localPath, Branch, &wg, errCh)
		}

		go func() {
			wg.Wait()
			close(errCh)
		}()

		for err := range errCh {
			log.Fatalf("%v", err)
		}

		log.Println("checkout complete")

	},
}

func init() {
	// checkoutCmd.Flags().StringVarP(&Branch, "branch", "b", "", "Branch or tag to checkout")
	// checkoutCmd.MarkFlagRequired("branch")
	rootCmd.AddCommand(checkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
