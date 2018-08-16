package rrgo

func niladic() error {
	// ...
	return nil
}

func sum(args ...int) (sum int) {
	sum = 0
	for arg := range args {
		sum += arg
	}
	return
}

func deviceVoltage(dev string) (vte float32, err error) {
	vte = 0
	err = nil
	// ...
	return
}

func typecheck(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
		case float32:
		case string:
		case bool:
		default:
		}
	}
}
