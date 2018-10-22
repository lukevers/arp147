package programs

import (
	"engo.io/engo"
	"flag"
	"github.com/lukevers/arp147/clock"
	"github.com/lukevers/arp147/input"
)

type Date struct {
	err error
	out string

	Clock *clock.Clock
	flags *flag.FlagSet
}

func (d *Date) Name() string {
	return "date"
}

func (d *Date) Headless() bool {
	return true
}

func (d *Date) Error() error {
	return d.err
}

func (d *Date) Output() string {
	return d.out
}

func (d *Date) Init() error {
	d.flags = flag.NewFlagSet("date", flag.ContinueOnError)
	d.flags.Bool("help", false, "Help menu")

	return nil
}

func (d *Date) Run(args []string) uint {
	// Parse arguments
	if d.err = d.flags.Parse(args); d.err != nil {
		return 1
	}

	// TODO: use arguments
	return 0
}

func (d *Date) HandleKey(key engo.Key, mods *input.Modifiers) {
	// No need to handle key input since this command doesn't capture/modify
	// the window at all.
}
