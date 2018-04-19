package tgrpc

import (
	"fmt"
	"time"

	"github.com/toukii/goutils"
)

type Invoke struct {
	Method   string   `toml:"method"`
	Headers  []string `toml:"headers"`
	Data     string   `toml:"data"`
	N        int      `toml:"n"`
	Interval *Ms      `toml:"interval"`
	Resp     *Resp    `toml:"resp"`
}

type Ms struct {
	time.Duration
}

func (d *Ms) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(goutils.ToString(text))
	return err
}

func (d *Ms) MarshalText() ([]byte, error) {
	return goutils.ToByte(fmt.Sprintf("%dms", int64(d.Nanoseconds()/1e6))), nil
}