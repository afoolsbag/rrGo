package rrgo

import (
	"testing"
	"math/rand"
)

// TestIfBranch if 分支。
func TestIfBranch(t *testing.T) {

	if rand.Intn(2) < 2 {
		// pass
	} else {
		t.Fail()
	}

	if rnd := rand.Intn(2); rnd < 2 {
		// pass
	} else {
		t.Fail()
	}
}

// TestSwitchBranch switch 分支。
func TestSwitchBranch(t *testing.T) {

	switch rand.Intn(2) {
	case 0, 1:
		// pass
	default:
		t.Fail()
	}

	switch {
	case rand.Intn(2) < 2:
		// pass
	default:
		t.Fail()
	}

	switch rnd := rand.Intn(2); {
	case rnd < 1:
		fallthrough
	case rnd < 2:
		// pass
	default:
		t.Fail()
	}

}

// TestForLoop for 循环。
func TestForLoop(t *testing.T) {

	for {
		// ...
		break
	}

	cond := true
	for cond {
		// ...
		cond = false
	}

	for i := 0; i <= 5; i++ {
		// ...
	}

	const str = "We Are the World"
	for ix, val := range str {
		ix++
		val++
	}
	for val := range str {
		val++
	}
}
