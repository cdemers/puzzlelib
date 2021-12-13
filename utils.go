package puzzlelib

func Same(ai, bi []byte) bool {
	a := map[byte]int{}

	lai := len(ai)
	lbi := len(bi)

	if lai == 0 && lbi == 0 {
		return true
	}
	if lai == 0 && lbi != 0 {
		return false
	}
	if lai != 0 && lbi == 0 {
		return false
	}

	for _, v := range ai {
		a[v]++
	}
	for _, v := range bi {
		a[v]++
	}
	for k := range a {
		if a[k] != 2 {
			return false
		}
	}
	return true
}

func Subtract(minuend, subtrahend []byte) (o []byte) {
	for _, b := range minuend {
		keep := true
		for _, a := range subtrahend {
			if a == b {
				keep = false
				break
			}
		}
		if keep {
			o = append(o, b)
		}
	}
	return o
}

func Intersect(groupA, groupB []byte) (output []byte) {
	for _, a := range groupA {
		keep := false
		for _, b := range groupB {
			if b == a {
				keep = true
				break
			}
		}
		if keep {
			output = append(output, a)
		}
	}
	return output

}

func Split(input []byte, delim byte) (output [][]byte) {
	var p, k int
	for k = range input {
		if input[k] == delim {
			output = append(output, input[p:k])
			p = k + 1
		}
	}
	if k > p {
		output = append(output, input[p:k+1])
	}
	return output
}
