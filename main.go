package main

import (
	"fmt"
	"os"

	cmd "github.com/qiniuts/qlogctl/cmd"
	cli "gopkg.in/urfave/cli.v2"
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
		},
		Commands: []*cli.Command{
			cmd.QueryByReqid, cmd.Query,
			cmd.ListRepo, cmd.SetRepo,
			cmd.LoginAccount, cmd.ShowAccounts,
			cmd.SwitchAccount, cmd.DelLoginAccount,
			cmd.QuerySample, cmd.SetRange, cmd.ClearLoginInfo,
			cmd.LogdbConf,
		},
	}

	err := qlogctl.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
