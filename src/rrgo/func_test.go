package rrgo

func func1() {
}

func func2(lhv int, rhv int) {
	sum := lhv + rhv
	sum++
}

func func3(args ...int) {
	sum := 0
	for arg := range args {
		sum += arg
	}
}

func func4() (int, int) {
	return 0, 0
}

func func5() (rslt int) {
	rslt = 0
	return
}
