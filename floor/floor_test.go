package floor

import (
	"reflect"
	"testing"
)

func TestUn(t *testing.T) {
	x := [][]int{
		[]int{1, 1, 3, 4},
		[]int{1, 1, 4, 3},
		[]int{0, 0, 2, 2},
		[]int{0, 0, 2, 2}}
	y := readFloorFromFile("../floor-files/exemple")
	if !reflect.DeepEqual(x, y) {
		t.Fail()
	}
}

func TestDeux(t *testing.T) {
	x := [][]int{
		[]int{0, 0, 1, 0, 2, 2, 2, 2},
		[]int{0, 0, 1, 0, 2, 2, 2, 2},
		[]int{0, 0, 1, 0, 2, 2, 2, 2},
		[]int{0, 0, 1, 1, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 1, 1, 0, 0},
		[]int{2, 2, 0, 0, 0, 1, 0, 0},
		[]int{2, 2, 2, 0, 0, 1, 0, 0},
		[]int{2, 2, 2, 0, 0, 1, 0, 0}}
	y := readFloorFromFile("../floor-files/beaupasbeau")
	if !reflect.DeepEqual(x, y) {
		t.Fail()
	}
}

func TestTrois(t *testing.T) {
	/*x := [][]int{
	[]int{0, 0, 1, 1, 2, 2},
	[]int{0, 0, 1, 1, 0, 0},
	[]int{3, 3, 3, 3, 3, 3, 3}}*/
	y := readFloorFromFile("../floor-files/fail")
	if y != nil {
		t.Fail()
	}
}

func TestVide(t *testing.T) {
	y := readFloorFromFile("../floor-files/vide")
	if y != nil {
		t.Fail()
	}
}

// test pour updateFromFileFloor
func copyFloor(facopier Floor, n int) (f, cempty [][]int) {
	matrix := facopier.fullContent
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	cempty = make([][]int, n)
	for elt := range cempty {
		for i := 0; i < n; i++ {
			cempty[elt] = append(cempty[elt], 0)
		}
	}
	return duplicate, cempty
}

func TestUpFF(t *testing.T) {
	f1 := Floor{
		content: [][]int{
			[]int{1, 1, 3, 4},
			[]int{1, 1, 4, 3},
			[]int{0, 0, 2, 2},
			[]int{0, 0, 2, 2}},

		fullContent: [][]int{
			[]int{1, 1, 3, 4},
			[]int{1, 1, 4, 3},
			[]int{0, 0, 2, 2},
			[]int{0, 0, 2, 2}}}

	var f2 Floor
	f2.fullContent, f2.content = copyFloor(f1, 4)
	f2.updateFromFileFloor(0, 0)
	if !reflect.DeepEqual(f1, f2) {
		t.Fail()
	}
}

func TestUpFF2(t *testing.T) {
	f1 := Floor{
		content: [][]int{
			[]int{4, -1, -1, -1},
			[]int{3, -1, -1, -1},
			[]int{2, -1, -1, -1},
			[]int{2, -1, -1, -1}},

		fullContent: [][]int{
			[]int{1, 1, 3, 4},
			[]int{1, 1, 4, 3},
			[]int{0, 0, 2, 2},
			[]int{0, 0, 2, 2}}}

	var f2 Floor
	f2.fullContent, f2.content = copyFloor(f1, 4)
	f2.updateFromFileFloor(3, 0)
	if !reflect.DeepEqual(f1, f2) {
		t.Fail()
	}
}

func TestUpFF3(t *testing.T) {
	f1 := Floor{
		content: [][]int{
			[]int{0, 0, 2, 2},
			[]int{-1, -1, -1, -1},
			[]int{-1, -1, -1, -1},
			[]int{-1, -1, -1, -1}},

		fullContent: [][]int{
			[]int{1, 1, 3, 4},
			[]int{1, 1, 4, 3},
			[]int{0, 0, 2, 2},
			[]int{0, 0, 2, 2}}}

	var f2 Floor
	f2.fullContent, f2.content = copyFloor(f1, 4)
	f2.updateFromFileFloor(0, 3)
	if !reflect.DeepEqual(f1, f2) {
		t.Fail()
	}
}

func TestUpFF4(t *testing.T) {
	f1 := Floor{
		content: [][]int{
			[]int{2, -1, -1, -1},
			[]int{-1, -1, -1, -1},
			[]int{-1, -1, -1, -1},
			[]int{-1, -1, -1, -1}},

		fullContent: [][]int{
			[]int{1, 1, 3, 4},
			[]int{1, 1, 4, 3},
			[]int{0, 0, 2, 2},
			[]int{0, 0, 2, 2}}}

	var f2 Floor
	f2.fullContent, f2.content = copyFloor(f1, 4)
	f2.updateFromFileFloor(3, 3)
	if !reflect.DeepEqual(f1, f2) {
		t.Fail()
	}
}

func TestUpFF5(t *testing.T) {
	f1 := Floor{
		content: [][]int{
			[]int{1, 1, 3, -1},
			[]int{1, 1, 4, -1},
			[]int{0, 0, 2, -1},
			[]int{-1, -1, -1, -1}},

		fullContent: [][]int{
			[]int{1, 1, 3},
			[]int{1, 1, 4},
			[]int{0, 0, 2}}}

	var f2 Floor
	f2.fullContent, f2.content = copyFloor(f1, 4)
	f2.updateFromFileFloor(0, 0)
	if !reflect.DeepEqual(f1, f2) {
		t.Fail()
	}
}

func TestUpFF6(t *testing.T) {
	f1 := Floor{
		content: [][]int{
			[]int{-1, -1, -1, -1},
			[]int{-1, 1, 1, 3},
			[]int{-1, 1, 1, 4},
			[]int{-1, 0, 0, 2}},

		fullContent: [][]int{
			[]int{1, 1, 3, 4},
			[]int{1, 1, 4, 3},
			[]int{0, 0, 2, 2},
			[]int{0, 0, 2, 2}}}

	var f2 Floor
	f2.fullContent, f2.content = copyFloor(f1, 4)
	f2.updateFromFileFloor(-1, -1)
	if !reflect.DeepEqual(f1, f2) {
		t.Fail()
	}
}
