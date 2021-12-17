package puzzlelib

import "fmt"

func Same(ai, bi []byte) bool {
	a := map[byte]int{}

	lai := len(ai)
	lbi := len(bi)

	if lai != lbi {
		return false
	}
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

func IntersectS(groupA, groupB string) (output string) {
	return string(Intersect([]byte(groupA), []byte(groupB)))
}

func FilterEmptyS(input []string) (output []string) {
	for _, v := range input {
		if v != "" {
			output = append(output, v)
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

func SplitInTwo(input []byte, delim byte) (outputOne, outputTwo []byte) {
	for k := range input {
		if input[k] == delim {
			outputOne = input[:k]
			outputTwo = input[k+1:]
			break
		}
	}
	return outputOne, outputTwo
}

func Trim(input []byte) []byte {
	var s, e int
	var c byte
	for s = 0; s < len(input); s++ {
		c = input[s]
		if c != ' ' && c != '\n' {
			break
		}
	}
	for e = len(input) - 1; e >= 0; e-- {
		c = input[e]
		if c != ' ' && c != '\n' {
			break
		}
	}
	return input[s : e+1]
}

type ByteStack struct {
	stack []byte
}

func (s *ByteStack) Push(b byte) {
	s.stack = append(s.stack, b)
}

func (s *ByteStack) Peek() (byte, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("error, can't do Peek(), the stack is empty")
	}
	return s.stack[len(s.stack)-1:][0], nil
}

func (s *ByteStack) Pop() (b byte, err error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("error, can't do Pop(), the stack is empty")
	}
	b, _ = s.Peek()
	s.stack = s.stack[:len(s.stack)-1]
	return b, nil
}

func (s *ByteStack) Stack() []byte {
	return s.stack
}
