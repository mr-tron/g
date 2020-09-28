package g

import (
	"testing"
	"time"
)

func TestInSlice(t *testing.T) {
	if !IntInSlice(1, []int{1, 2, 3}) {
		t.Fatal("should be found")
	}
	if !IntInSlice(1, []int{2, 1, 3}) {
		t.Fatal("should be found")
	}
	if IntInSlice(4, []int{1, 2, 3}) {
		t.Fatal("should be not found")
	}
	if IntInSlice(1, []int{}) {
		t.Fatal("should be not found")
	}

	if Float32InSlice(0.3, []float32{0.2, 0.1}) {
		t.Fatal("should be not found")
	}
	if !Float32InSlice(0.3, []float32{0.2, 0.3, 0.4}) {
		t.Fatal("should be found")
	}

	if !StringInSlice("", []string{"", "123", "321"}) {
		t.Fatal("should be found")
	}
	if !StringInSlice("123", []string{"", "123", "321"}) {
		t.Fatal("should be found")
	}
	if StringInSlice("", []string{}) {
		t.Fatal("should be not found")
	}
	if StringInSlice("123", []string{}) {
		t.Fatal("should be not found")
	}

	if Complex64InSlice(0.3+0.4i, []complex64{1 + 2i, 0.3 + 0i}) {
		t.Fatal("should be not found")
	}
	if !Complex64InSlice(0.3+0.4i, []complex64{1 + 2i, 0.3 + 0.4i}) {
		t.Fatal("should be found")
	}
	if TimeInSlice(time.Now(), []time.Time{}) {
		t.Fatal("should be not found")
	}
	if !TimeInSlice(time.Unix(1232141414, 32423423), []time.Time{time.Now(), time.Unix(1232141414, 32423423), time.Unix(243254365, 32452)}) {
		t.Fatal("should be found")
	}
}
