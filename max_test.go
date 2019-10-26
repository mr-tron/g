package g

import "testing"

func TestMaxInt(t *testing.T) {
	var r int
	var err error
	_, err = MaxInt([]int{})
	if err == nil {
		t.Fatal("should be error")
	}
	_, err = MaxInt(nil)
	if err == nil {
		t.Fatal("should be error")
	}
	r, err = MaxInt([]int{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MaxInt([]int{3, 2, 1})
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MaxInt([]int{1, -2, 3})
	if err != nil {
		t.Fatal(err)
	}
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MaxInt([]int{-1, -2, -3})
	if err != nil {
		t.Fatal(err)
	}
	if r != -1 {
		t.Fatalf("bad result: %v", r)
	}
}

func TestMaxIntWithDefault(t *testing.T) {
	var r int
	r = MaxIntWithDefault([]int{}, 3)
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault(nil, 3)
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault([]int{1, 2}, 3)
	if r != 2 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault([]int{4, 5}, 3)
	if r != 5 {
		t.Fatalf("bad result: %v", r)
	}
	r = MaxIntWithDefault([]int{-4, -5}, 3)
	if r != -4 {
		t.Fatalf("bad result: %v", r)
	}
}
