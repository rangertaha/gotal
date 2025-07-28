package indicators

type Opts func(*Option) error

type Option struct {
	Period int
}

func NewOption(options ...Opts) (*Option, error) {
	o := &Option{
		Period: 10,
	}

	// Apply config options
	for _, opt := range options {
		err := opt(o)
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

func Period(period int) func(*Option) (err error) {
	return func(o *Option) (err error) {
		o.Period = period
		return nil
	}
}
