package inter

type Option func(*Interface)

func (l *Interface) With(options ...Option) {
	for _, opt := range options {
		opt(l)
	}
}

func (l InterfaceList) With(options ...Option) InterfaceList {
	list := make(InterfaceList, len(l))
	for key, value := range l {
		value.With(options...)
		list[key] = value
	}
	return list
}

func WithoutMTU() Option {
	return func(i *Interface) {
		i.MTU = 0
	}
}

func WithoutFlags() Option {
	return func(i *Interface) {
		i.Flags = nil
	}
}

func WithoutAddrs() Option {
	return func(i *Interface) {
		i.Addrs = nil
	}
}

func WithoutHardwareAddr() Option {
	return func(i *Interface) {
		i.HardwareAddr = ""
	}
}
