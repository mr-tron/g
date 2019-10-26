package g

import "testing"

func TestMinInt(t *testing.T) {
	var r int
	var err error
	_, err = MinInt([]int{})
	if err == nil {
		t.Fatal("should be error")
	}
	_, err = MinInt(nil)
	if err == nil {
		t.Fatal("should be error")
	}
	r, err = MinInt([]int{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	if r != 1 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MinInt([]int{3, 2, 1})
	if err != nil {
		t.Fatal(err)
	}
	if r != 1 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MinInt([]int{3, -2, 1})
	if err != nil {
		t.Fatal(err)
	}
	if r != -2 {
		t.Fatalf("bad result: %v", r)
	}
	r, err = MinInt([]int{-1, -2, -3})
	if err != nil {
		t.Fatal(err)
	}
	if r != -3 {
		t.Fatalf("bad result: %v", r)
	}
}

func TestMinIntWithDefault(t *testing.T) {
	var r int
	r = MinIntWithDefault([]int{}, 3)
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r = MinIntWithDefault(nil, 3)
	if r != 3 {
		t.Fatalf("bad result: %v", r)
	}
	r = MinIntWithDefault([]int{1, 2}, 3)
	if r != 1 {
		t.Fatalf("bad result: %v", r)
	}
	r = MinIntWithDefault([]int{4, 5}, 3)
	if r != 4 {
		t.Fatalf("bad result: %v", r)
	}
	r = MinIntWithDefault([]int{-4, -5}, 3)
	if r != -5 {
		t.Fatalf("bad result: %v", r)
	}
}
