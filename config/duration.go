package config

import (
	"time"
)

type Duration struct {
	d time.Duration
}

func MakeDuration(d time.Duration) Duration {
	return Duration{d: d}
}

func ParseDuration(s string) (Duration, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return Duration{}, err
	}

	return MakeDuration(d), nil
}

func (d Duration) Duration() time.Duration {
	return d.d
}

func (d Duration) String() string {
	return d.Duration().String()
}

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(d.d.String()), nil
}

func (d *Duration) UnmarshalText(input []byte) error {
	v, err := time.ParseDuration(string(input))
	if err != nil {
		return err
	}
	*d = MakeDuration(v)
	return nil
}
