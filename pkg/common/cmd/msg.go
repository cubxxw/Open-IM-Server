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

package cmd

import (
	"context"

	"github.com/openimsdk/open-im-server/v3/internal/rpc/msg"
	"github.com/openimsdk/open-im-server/v3/pkg/common/config"
	"github.com/openimsdk/open-im-server/v3/pkg/common/startrpc"
	"github.com/openimsdk/open-im-server/v3/version"
	"github.com/openimsdk/tools/system/program"
	"github.com/spf13/cobra"
)

type MsgRpcCmd struct {
	*RootCmd
	ctx       context.Context
	configMap map[string]any
	msgConfig *msg.Config
}

func NewMsgRpcCmd() *MsgRpcCmd {
	var msgConfig msg.Config
	ret := &MsgRpcCmd{msgConfig: &msgConfig}
	ret.configMap = map[string]any{
		config.OpenIMRPCMsgCfgFileName:  &msgConfig.RpcConfig,
		config.RedisConfigFileName:      &msgConfig.RedisConfig,
		config.MongodbConfigFileName:    &msgConfig.MongodbConfig,
		config.KafkaConfigFileName:      &msgConfig.KafkaConfig,
		config.ShareFileName:            &msgConfig.Share,
		config.NotificationFileName:     &msgConfig.NotificationConfig,
		config.WebhooksConfigFileName:   &msgConfig.WebhooksConfig,
		config.LocalCacheConfigFileName: &msgConfig.LocalCacheConfig,
		config.DiscoveryConfigFilename:  &msgConfig.Discovery,
	}
	ret.RootCmd = NewRootCmd(program.GetProcessName(), WithConfigMap(ret.configMap))
	ret.ctx = context.WithValue(context.Background(), "version", version.Version)
	ret.Command.RunE = func(cmd *cobra.Command, args []string) error {
		return ret.runE()
	}
	return ret
}

func (a *MsgRpcCmd) Exec() error {
	return a.Execute()
}

func (a *MsgRpcCmd) runE() error {
	return startrpc.Start(a.ctx, &a.msgConfig.Discovery, &a.msgConfig.RpcConfig.Prometheus, a.msgConfig.RpcConfig.RPC.ListenIP,
		a.msgConfig.RpcConfig.RPC.RegisterIP, a.msgConfig.RpcConfig.RPC.AutoSetPorts, a.msgConfig.RpcConfig.RPC.Ports,
		a.Index(), a.msgConfig.Discovery.RpcService.Msg, &a.msgConfig.NotificationConfig, a.msgConfig,
		[]string{
			a.msgConfig.RpcConfig.GetConfigFileName(),
			a.msgConfig.RedisConfig.GetConfigFileName(),
			a.msgConfig.MongodbConfig.GetConfigFileName(),
			a.msgConfig.KafkaConfig.GetConfigFileName(),
			a.msgConfig.NotificationConfig.GetConfigFileName(),
			a.msgConfig.Share.GetConfigFileName(),
			a.msgConfig.WebhooksConfig.GetConfigFileName(),
			a.msgConfig.LocalCacheConfig.GetConfigFileName(),
			a.msgConfig.Discovery.GetConfigFileName(),
		}, nil,
		msg.Start)
}
