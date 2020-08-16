package g

import "testing"

func TestPbool(t *testing.T) {
	b := Pbool(true)
	if b == nil {
		t.Fatal("b is nil")
	}
	if *b != true {
		t.Fatal("invalid value")
	}
}

func TestPstring(t *testing.T) {
	s := Pstring("ololo")
	if s == nil {
		t.Fatal("s is nil")
	}
	if *s != "ololo" {
		t.Fatal("invalid value")
	}
}

func TestPint(t *testing.T) {
	i := Pint(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPint8(t *testing.T) {
	i := Pint8(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPint16(t *testing.T) {
	i := Pint16(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPint32(t *testing.T) {
	i := Pint32(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPint64(t *testing.T) {
	i := Pint64(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPuint(t *testing.T) {
	i := Puint(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPuint8(t *testing.T) {
	i := Puint8(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPuint16(t *testing.T) {
	i := Puint16(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPuint32(t *testing.T) {
	i := Puint32(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPuint64(t *testing.T) {
	i := Puint64(100)
	if i == nil {
		t.Fatal("i is nil")
	}
	if *i != 100 {
		t.Fatal("invalid value")
	}
}

func TestPfloat32(t *testing.T) {
	f := Pfloat32(1.0)
	if f == nil {
		t.Fatal("f is nil")
	}
	if *f != 1.0 {
		t.Fatal("invalid value")
	}
}

func TestPfloat64(t *testing.T) {
	f := Pfloat64(1.0)
	if f == nil {
		t.Fatal("f is nil")
	}
	if *f != 1.0 {
		t.Fatal("invalid value")
	}
}

func TestPcomplex64(t *testing.T) {
	c := Pcomplex64(2.3 - 4.7i)
	if c == nil {
		t.Fatal("c is nil")
	}
	if *c != 2.3-4.7i {
		t.Fatal("invalid value")
	}
}

func TestPcomplex128(t *testing.T) {
	c := Pcomplex128(2.3 - 4.7i)
	if c == nil {
		t.Fatal("c is nil")
	}
	if *c != 2.3-4.7i {
		t.Fatal("invalid value")
	}
}
