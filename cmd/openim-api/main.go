// Copyright © 2023 OpenIM. All rights reserved.
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

package main

import (
	_ "net/http/pprof"
	"time"

	"github.com/openimsdk/open-im-server/v3/pkg/common/cmd"
	util "github.com/openimsdk/open-im-server/v3/pkg/util/genutil"
)

func panicAfterOneMinute() {
	go func() {
		<-time.After(1 * time.Minute)
		panic("Panic after one minute!")
	}()
}

func main() {
	panicAfterOneMinute()
	apiCmd := cmd.NewApiCmd(cmd.ApiServer)
	apiCmd.AddPortFlag()
	apiCmd.AddPrometheusPortFlag()
	if err := apiCmd.Execute(); err != nil {
		util.ExitWithError(err)
	}
}
