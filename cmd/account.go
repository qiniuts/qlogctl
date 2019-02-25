package cmd

import (
	"errors"

	"github.com/qiniuts/qlogctl/api"
	"gopkg.in/urfave/cli.v2"
)

var (
	LogdbConf = &cli.Command{
		Name:  "logdb",
		Usage: "设置 logdb 配置参数",
		Subcommands: []*cli.Command{
			&cli.Command{
				Name:  "endpoint",
				Usage: "设置对端地址， 默认： https://jjh-insight.qiniuapi.com",
				Action: func(c *cli.Context) error {
					if c.Args().Len() == 1 {
						endpoint := c.Args().Get(0)
						api.Endpoint(endpoint)
						return nil
					}
					return errors.New("参数错误： endpoint <endpoint> ")
				},
			},
		},
	}

	LoginAccount = &cli.Command{
		Name:      "account",
		Usage:     "设置后续查询时需要的 ak sk 以及方便使用的别名",
		ArgsUsage: "<ak> <sk> <name>",
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 3 {
				api.Account(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
				return nil
			}
			return errors.New("参数错误： <ak> <sk> <name> ")
		},
	}

	ShowAccounts = &cli.Command{
		Name:  "accounts",
		Usage: "查看已设置账号列表",
		Action: func(c *cli.Context) error {
			api.UserList()
			return nil
		},
	}

	SwitchAccount = &cli.Command{
		Name:      "switch",
		Usage:     "通过别名切换不同的账号",
		ArgsUsage: "<name>",
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 1 {
				api.Switch(c.Args().Get(0))
				return nil
			}
			return errors.New("参数错误： <name> ")
		},
	}

	DelLoginAccount = &cli.Command{
		Name:      "deluser",
		Usage:     "通过别名删除账号信息",
		ArgsUsage: "<name>",
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 1 {
				api.Deluser(c.Args().Get(0))
				return nil
			}
			return errors.New("参数错误： <name> ")
		},
	}

	ClearLoginInfo = &cli.Command{
		Name:  "clear",
		Usage: "清理保存在临时文件中的信息",
		Action: func(c *cli.Context) error {
			api.Clear()
			return nil
		},
	}
)
