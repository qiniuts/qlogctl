package main

import (
	"os"

	"github.com/qiniuts/qlogctl/api"
	"github.com/qiniuts/qlogctl/cmd"
	"github.com/qiniuts/qlogctl/log"
	"github.com/urfave/cli/v2"
)

func main() {
	qlogctl := &cli.App{
		Name:      "qlogctl",
		Usage:     "query logs from logdb",
		UsageText: " command [command options] [arguments...] ",
		Version:   "0.0.8",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "debug",
			},
			&cli.StringFlag{
				Name:  "ak",
				Usage: "设置 ak ，即 AccessKey ；优先级高于配置文件内容",
			},
			&cli.StringFlag{
				Name:  "sk",
				Usage: "设置 sk ，即 SecretKey ；优先级高于配置文件内容",
			},
			&cli.StringFlag{
				Name:  "repo",
				Usage: "设置 repo，即 logdb 的名称 ；优先级高于配置文件内容",
			},
			&cli.StringFlag{
				Name:  "endpoint",
				Usage: "设置 endpoint ；优先级高于配置文件内容",
			},
		},
		Commands: []*cli.Command{
			cmd.QueryByReqid, cmd.Query,
			cmd.ListRepo, cmd.SetRepo,
			cmd.LoginAccount, cmd.ShowAccounts,
			cmd.SwitchAccount, cmd.DelLoginAccount,
			cmd.QuerySample, cmd.SetRange, cmd.ClearLoginInfo,
			cmd.LogdbConf,
		},
		Before: func(c *cli.Context) error {
			api.SetDebug(c.Bool("debug"))
			if ak := c.String("ak"); len(ak) > 10 {
				api.CurrentAK(ak)
			}
			if sk := c.String("sk"); len(sk) > 10 {
				api.CurrentSK(sk)
			}
			if repo := c.String("repo"); len(repo) > 0 {
				api.CurrentRepo(repo)
			}
			if endpoint := c.String("endpoint"); len(endpoint) > 10 {
				api.CurrentEndpoint(endpoint)
			}
			return nil
		},
	}

	err := qlogctl.Run(os.Args)
	if err != nil {
		log.Errorln(err)
		os.Exit(1)
	}
}
