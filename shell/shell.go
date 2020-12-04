/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package shell

import (
	"github.com/desertbit/grumble"
)

// App is the global shell app.
var App *grumble.App

// TODO(jiashuo) The version name my need rename
var NameWithVersion = "Pegasus-AdminCli-1.0.0"

// AddCommand registers the command to the global shell app.
func AddCommand(cmd *grumble.Command) {
	App.AddCommand(cmd)
}

// TODO(jiashuo1) some command with table no support verify the table exists
func init() {
	App = grumble.New(&grumble.Config{
		Name:        NameWithVersion,
		Description: "Pegasus administration command line tool",
		Flags: func(f *grumble.Flags) {
			f.String("m", "meta", "127.0.0.1:34601,127.0.0.1:34602", "a list of MetaServer IP:Port addresses")
		},
		HistoryFile: ".admin-cli-history",
	})
}
