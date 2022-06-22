package main

import (
	"github.com/alexflint/go-arg"
	"github.com/nosarthur/gom/monitor"
)

type StartCmd struct {
}

type QuitCmd struct {
}

type ShowCmd struct {
	Verbose bool   `arg:"-v"`
	Type    string `arg:"-t"`
}

type ConfigCmd struct {
}

type RunCmd struct {
	cmds string
}

type args struct {
	Start  *StartCmd  `arg:"subcommand:start" help:"Start/reset monitoring"`
	Quit   *QuitCmd   `arg:"subcommand:quit" help:"Quit monitoring"`
	Show   *ShowCmd   `arg:"subcommand:show" help:"Show jobs"`
	Config *ConfigCmd `arg:"subcommand:config" help:"Show configurations"`
	Run    *RunCmd    `arg:"subcommand:run" help:"Run python script"`
}

func (args) Version() string {
	return "gom 0.0.1"
}

func main() {
	var args args
	arg.MustParse(&args)

	switch {
	case args.Start != nil:
		if monitor.IsRunning() {
			monitor.Reset()
		} else {
			monitor.Spinup()
		}
	case args.Config != nil:
		println("show config ")
	case args.Quit != nil:
		println("not implemented; just `killall gom`")
	case args.Show != nil:
		fallthrough
	default:
		monitor.Connect()
	}
}
