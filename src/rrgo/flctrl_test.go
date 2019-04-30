package rrgo

import "testing"

// TestIf
//
// if[ initialization;] condition {
// }[ else {
// }]
func TestIf(t *testing.T) {
	condition := true
	initialization := func() { condition = !condition }

	if condition {
		// ...
	}

	if initialization(); condition {
		// ...
	}
}

// TestSwitch
//
// switch val {
//     case val1, val2: [fallthrough]
//     case val3:
//     default:
// }
//
// switch {
//     case condition1: [fallthrough]
//     case condition2:
//     default:
// }
func TestSwitch(t *testing.T) {
	val := 3

	switch val {
	case 1, 2, 3:
		// ...
	case 4, 5, 6:
		fallthrough
	default:
		// ...
	}

	switch {
	case val < 0:
		// ...
	case 0 < val:
		//...
	default:
		// ...
	}
}

// TestFor
//
// for {
// }
//
// for condition {
// }
//
// for initialization; condition; op {
// }
func TestFor(t *testing.T) {
	for {
		break
	}

	i := 0
	for i < 9 {
		i++
	}

	for i := 0; i < 9; i++ {
		// ...
	}
}
