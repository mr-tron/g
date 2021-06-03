package g

import "testing"

func TestMaxInt(t *testing.T) {
	var r int
	var err error
	_, err = MaxIntInSlice([]int{})
	if err == nil {
		t.Fatal("should be error")
	}
	_, err = MaxIntInSlice(nil)
	if err == nil {
		t.Fatal("should be error")
	}
	r, err = MaxIntInSlice([]int{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MaxIntInSlice([]int{3, 2, 1})
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MaxIntInSlice([]int{1, -2, 3})
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MaxIntInSlice([]int{-1, -2, -3})
	if err != nil {
		t.Fatal(err)
	}
	if r != -1 {
		t.Fatalf("bad result: %v", r)
	}
}

func TestMaxIntWithDefault(t *testing.T) {
	var r int
	r = MaxIntWithDefault(3, []int{}...)
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault(3)
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault(3, 1, 2)
	if r != 2 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault(3, 4, 5)
	if r != 5 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault(3, -4, -5)
	if r != -4 {
		t.Fatalf("bad result: %v", r)
	}
}
