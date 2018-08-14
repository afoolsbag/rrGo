package rrgo

import "testing"

func TestIfElse(t *testing.T) {
	// (1)
	if true {
		// ...
	} else {
		// ...
	}

	// (2)
	if lmt := 3; lmt < 7 {
		// ...
	}
}

func TestSwitchCaseDefault(t *testing.T) {
	// (1)
	const val = 0

	switch val {
	case 0:
		fallthrough
	case 1:
		// ...
	case 2:
		// ...
	default:
		// ...
	}

	// (2)
	switch {
	case true:
		// ...
	default:
		// ...
	}

	// (3)
	switch rslt := 0; {
	case rslt < 0:
		// ...
	case rslt < 10:
		// ...
	default:
		// ...
	}
}

func TestFor(t *testing.T) {
	// (1)
	for i := 0; i <= 5; i++ {
		// ...
	}

	// (2)
	cond := true
	for cond {
		// ...
		cond = false
	}

	// (3)
	for {
		// ...
		break
	}

	// (4)
	str := "We Are the World"
	for ix, val := range str {
		ix++
		val++
	}
	for val := range str {
		val++
	}
}
