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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stormentt/simple-single-share/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run a sss server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindEnv("uploads.api_key", "SSS_UPLOADS_API_KEY")

		err := server.Serve()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("unable to serve")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
