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

func TestSplitInTwo(t *testing.T) {
	input := []byte("abc\ndef\nghi")

	var partA, partB []byte
	partA, partB = SplitInTwo(input, '\n')

	if !Same(partA, []byte("abc")) || !Same(partB, []byte("def\nghi")) {
		t.Errorf("error: SplitInTwo([]byte(\"abc\\ndef\\nghi\"), '\\n') should return []byte(\"abc\"), []byte(\"def\nghi\"), got %s, %s", string(partA), string(partB))
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
