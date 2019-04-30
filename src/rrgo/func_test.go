package rrgo

// Basic

func niladic() {
	return
}

// Call

func callByValue(val int) {
	val++
}

func callByReference(ref *int) {
	*ref++
}

// Return

func returnError() error {
	return nil
}

func returnMulti() (int, int) {
	return 0, 0
}

func returnNamed() (r1 int, r2 int) {
	r1, r2 = 0, 0
	return
}

// More

type options struct {
	opt1 bool
	opt2 int
	opt3 string
}

func optionsParameters(p1 int, p2 int, opt options) {

}

func callOptionsParameters(){
	optionsParameters(1,2, options{})
}

func variable(args ...int) (sum int) {
	sum = 0
	for arg := range args {
		sum += arg
	}
	return
}

func interfaceVariable(args ...interface{}) {
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
