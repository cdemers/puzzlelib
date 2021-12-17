package puzzlelib

import "testing"

func TestSame(t *testing.T) {
	if !Same([]byte{'a', 'b'}, []byte{'a', 'b'}) {
		t.Errorf("error: Same([]byte{'a','b'}, []byte{'a', 'b'}) should return true")
	}
	if !Same([]byte{}, []byte{}) {
		t.Errorf("error: Same([]byte{}, []byte{}) should return true")
	}
}

func TestNotSame(t *testing.T) {
	if Same([]byte{'b'}, []byte{'a', 'b'}) {
		t.Errorf("error: Same([]byte{'b'}, []byte{'a', 'b'}) should return false")
	}
	if Same([]byte{'a', 'b'}, []byte{'b'}) {
		t.Errorf("error: Same([]byte{'a', 'b'}, []byte{'b'}) should return false")
	}
	if Same([]byte{}, []byte{'a'}) {
		t.Errorf("error: Same([]byte{}, []byte{'a'}) should return false")
	}
	if Same([]byte{'a'}, []byte{}) {
		t.Errorf("error: Same([]byte{'a'}, []byte{}) should return false")
	}
}

func TestSubtract(t *testing.T) {
	var ensableA []byte = []byte{'a', 'b', 'c'}
	var ensableB []byte = []byte{'b'}

	var output []byte

	output = Subtract(ensableA, ensableB)

	if !Same(output, []byte{'a', 'c'}) {
		t.Errorf("expected output to be {'a', 'c'}, got %#v", output)
	}

	output = Subtract(ensableB, ensableA)
	if !Same(output, []byte{}) {
		t.Errorf("expected output to be {}, got %#v", output)
	}

	output = Subtract([]byte{}, []byte{})
	if !Same(output, []byte{}) {
		t.Errorf("expected output to be {}, got %#v", output)
	}
}

func TestSplit(t *testing.T) {
	var output [][]byte

	output = Split([]byte("abc\ndef"), '\n')
	if len(output) != 2 {
		t.Errorf("expected lenght 2, got %d", len(output))
	}
	if !Same(output[0], []byte("abc")) {
		t.Errorf("expected output[0] to be []byte(\"abc\"), got %s", string(output[0]))
	}
	if !Same(output[1], []byte("def")) {
		t.Errorf("expected output[1] to be []byte(\"def\"), got %s", string(output[1]))
	}

	output = Split([]byte("a\ndef\n"), '\n')
	if len(output) != 2 {
		t.Errorf("expected lenght 2, got %d", len(output))
	}
	if !Same(output[1], []byte("def")) {
		t.Errorf("expected output[1] to be []byte(\"def\"), got %s", string(output[1]))
	}
}

func TestIntersect(t *testing.T) {
	var intersect []byte

	intersect = Intersect([]byte{'a', 'b', 'c'}, []byte{'b', 'c', 'd'})
	if !Same(intersect, []byte{'b', 'c'}) {
		t.Errorf("error: Intersect([]byte{'a', 'b', 'c'}, []byte{'b', 'c', 'd'}) should return []byte{'b', 'c'}")
	}

	intersect = Intersect([]byte{'a', 'b'}, []byte{'c', 'd'})
	if !Same(intersect, []byte{}) {
		t.Errorf("error: Intersect([]byte{'a', 'b'}, []byte{'c', 'd'}) should return []byte{}")
	}
}

func TestIntersectS(t *testing.T) {
	var intersect string

	intersect = IntersectS("abc", "bcd")
	if !Same([]byte(intersect), []byte{'b', 'c'}) {
		t.Errorf("error: IntersectS(\"abc\",\"bcd\") should return \"bc\"")
	}

	intersect = IntersectS("ab12", "bc23")
	if !Same([]byte(intersect), []byte{'b', '2'}) {
		t.Errorf("error: IntersectS(\"ab12\",\"bc23\") should return \"b2\"")
	}

	intersect = IntersectS("ab", "cd")
	if !Same([]byte(intersect), []byte{}) {
		t.Errorf("error: Intersect([]byte{'a', 'b'}, []byte{'c', 'd'}) should return []byte{}")
	}
}

func TestSplitInTwo(t *testing.T) {
	input := []byte("abc\ndef\nghi")

	var partA, partB []byte
	partA, partB = SplitInTwo(input, '\n')

	if !Same(partA, []byte("abc")) || !Same(partB, []byte("def\nghi")) {
		t.Errorf("error: SplitInTwo([]byte(\"abc\\ndef\\nghi\"), '\\n') should return []byte(\"abc\"), []byte(\"def\nghi\"), got %s, %s", string(partA), string(partB))
	}
}

func TestFilterEmptyS(t *testing.T) {
	input := []string{"", "a", "", "b", ""}

	output := FilterEmptyS(input)
	if output[0] != "a" || output[1] != "b" || len(output) != 2 {
		t.Errorf("error: FilterEmptyS([]string{\"\", \"a\", \"\", \"b\", \"\"}) should return []string{\"a\",\"b\"}, got %#v", output)
	}
}

func TestTrim(t *testing.T) {
	output := Trim([]byte(" \n abc \n "))
	if !Same(output, []byte("abc")) {
		t.Errorf("error: Trim([]byte(\" \\n abc \\n \")) should return []byte(\"abc\"), got %s", string(output))
	}

	output = Trim([]byte("abc"))
	if !Same(output, []byte("abc")) {
		t.Errorf("error: Trim([]byte(\"abc\")) should return []byte(\"abc\"), got %s", string(output))
	}
}

func TestByteStack_Push(t *testing.T) {
	s := ByteStack{}
	s.Push('a')
	if len(s.stack) != 1 {
		t.Errorf("expected stack length to be `1`, got %v", len(s.stack))
	}
	s.Push('b')
	if len(s.stack) != 2 {
		t.Errorf("expected stack length to be `2`, got %v", len(s.stack))
	}
	s.Push('c')

	if !Same(s.stack, []byte("abc")) {
		t.Errorf("expected stack content to be `abc`, got %v", s.stack)
	}
}

func TestByteStack_Peek(t *testing.T) {
	s := ByteStack{}
	s.Push('a')
	p, err := s.Peek()
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
	if p != 'a' {
		t.Errorf("expected stack peek to return `a`, got %v", p)
	}
	if len(s.stack) != 1 {
		t.Errorf("expected stack length to be `1`, got %v", len(s.stack))
	}
	s.Push('b')
	p, err = s.Peek()
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
	if p != 'b' {
		t.Errorf("expected stack peek to return `b`, got %v", p)
	}
	if len(s.stack) != 2 {
		t.Errorf("expected stack length to be `2`, got %v", len(s.stack))
	}
}

func TestByteStack_Stack(t *testing.T) {
	s := ByteStack{}
	s.Push('a')
	s.Push('b')

	if !Same(s.Stack(), []byte("ab")) {
		t.Errorf("expected the stack to be `ab`, got %v", s.Stack())
	}
}

func TestByteStack_Pop(t *testing.T) {
	s := ByteStack{}
	s.Push('a')
	if len(s.stack) != 1 {
		t.Errorf("expected stack length to be `1`, got %v", len(s.stack))
	}
	p, err := s.Pop()
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
	if len(s.stack) != 0 {
		t.Errorf("expected stack length to be `0`, got %v", len(s.stack))
	}
	if p != 'a' {
		t.Errorf("expected Pop() to return `a`, got %v", p)
	}
	s.Push('a')
	s.Push('b')
	if len(s.stack) != 2 {
		t.Errorf("expected stack length to be `2`, got %v", len(s.stack))
	}
	p, err = s.Pop()
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
	if len(s.stack) != 1 {
		t.Errorf("expected stack length to be `1`, got %v", len(s.stack))
	}
	if p != 'b' {
		t.Errorf("expected Pop() to return `b`, got %v", p)
	}
	p, err = s.Pop()
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
	if len(s.stack) != 0 {
		t.Errorf("expected stack length to be `0`, got %v", len(s.stack))
	}
	p, err = s.Pop()
	if err == nil {
		t.Errorf("expected error while atempting Pop() on an empty stack, got no error")
	}
	if len(s.stack) != 0 {
		t.Errorf("expected stack length to be `0`, got %v", len(s.stack))
	}
}
