package actions

import (
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/processors"
	"github.com/pkg/errors"
)

type addLocale struct {
	timezone string
}

func init() {
	processors.RegisterPlugin("add_locale", newAddLocale)
}

func newAddLocale(c common.Config) (processors.Processor, error) {
	config := struct {
	}{}

	err := c.Unpack(&config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack add_locale config")
	}

	zone, _ := time.Now().In(time.Local).Zone()

	l := addLocale{timezone: zone}

	return l, nil
}

func (l addLocale) Run(event common.MapStr) (common.MapStr, error) {
	event.Put("beat.timezone", l.timezone)

	return event, nil
}

func (l addLocale) String() string {
	return "add_locale=" + l.timezone
}
