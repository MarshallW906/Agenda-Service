// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"logger"
	"service"
	"utils"

	"github.com/spf13/cobra"
)

// findUserCmd represents the findUser command
var findUserCmd = &cobra.Command{
	Use:   "findUser",
	Short: "Find a user",
	Long: `Find a user
	- 用户删除
	- args: username string
	- notes: 要求已登录
	`,
	Run: func(cmd *cobra.Command, args []string) {
		username := utils.GetNonEmptyString(cmd, "username")

		service.FindUser(username)
		logger.Info("FindUser called with username:[%+v]", username)
	},
}

func init() {
	RootCmd.AddCommand(findUserCmd)

	findUserCmd.Flags().StringP("username", "u", "", "specify username")
}
