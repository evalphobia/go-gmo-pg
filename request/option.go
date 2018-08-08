package request

import (
	"time"
)

type Option struct {
	UserAgent string
	Timeout   time.Duration
	Debug     bool
	Retry     bool
}

func (o Option) getUserAgent() string {
	if o.UserAgent != "" {
		return o.UserAgent
	}
	return "go-gmo-pg"
}

func (o Option) getTimeout() time.Duration {
	if o.Timeout > 0 {
		return o.Timeout
	}
	return 30 * time.Second
}
