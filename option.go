package latency

type option struct {
	format FormatFunc
}

type Option func(o *option)

// WithReportFormat set report format
func WithReportFormat(f FormatFunc) Option {
	return func(o *option) {
		o.format = f
	}
}

var defaultOption = option{
	format: DefaultFormat,
}
