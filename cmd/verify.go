/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"prtg_safenet_hsm/safenet"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify slot is available",
	Long: `verify slot is available and return slot count
also offers serial number checking
`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		dir, err := flags.GetString("dir")
		if err != nil {
			log.Fatal("could not get directory %v", err)
		}
		exe, err := flags.GetString("exe")
		if err != nil {
			log.Fatal("could not get exe %v", err)
		}
		serial, err := flags.GetString("serial")
		if err != nil {
			log.Fatal("could not get serial %v", err)
		}
		v := safenet.NewVtl(dir, exe)
		err = v.Verify(serial)
		if err != nil {
			log.Fatal("error verifying connection", err)
		}
	},
}

func init() {
	vtlCmd.AddCommand(verifyCmd)
	verifyCmd.Flags().StringP("serial", "S", "", "serial number to check")
}
