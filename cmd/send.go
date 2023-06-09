/*
Copyright © 2023 Tanner Storment

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HostString string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "sends a note to a sss server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send called")
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&HostString, "host", "h", "localhost:8080", "Help message for toggle")
}
