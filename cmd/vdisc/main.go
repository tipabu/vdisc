// Copyright © 2019 NVIDIA Corporation
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
	"github.com/alecthomas/kong"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/NVIDIA/vdisc/pkg/vdisc/cli"
)

func main() {
	logAtomic := zap.NewAtomicLevel()
	logCfg := zap.NewProductionConfig()
	logCfg.Level = logAtomic
	logCfg.Encoding = "console"
	logCfg.DisableStacktrace = true
	logger, err := logCfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	undo = zap.RedirectStdLog(logger)
	defer undo()

	maxprocs.Set(maxprocs.Logger(logger.Sugar().Infof))

	cli := vdisc_cli.CLI{}

	ctx := kong.Parse(&cli,
		kong.Name("vdisc"),
		kong.Description("A virtual disc image tool"),
		kong.UsageOnError(),
		vdisc_cli.UUIDTypeMapper(),
		vdisc_cli.SITypeMapper(),
		vdisc_cli.GcThresholdTypeMapper(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{},
	)

	var ll zapcore.Level
	ll.Set(cli.Globals.LogLevel)
	logAtomic.SetLevel(ll)

	if err := ctx.Run(&cli.Globals); err != nil {
		logger.Fatal("run", zap.Error(err))
	}
}
