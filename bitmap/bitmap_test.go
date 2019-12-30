package bitmap

import (
	"log"
	"reflect"
	"testing"
)

// get gets a value from a bitmap, handling
// all the error checking so it's not repeated
// a million times.
func get(b BitMap, i uint64) bool {
	val, err := b.IsSet(i)
	if err != nil {
		log.Fatal("error getting from bitmap", err)
	}
	return val
}

// TestCreate proves we can create a
// bitmap and its size is as set.
func TestCreate(t *testing.T) {
	for _, size := range []uint64{1, 13, 27, 66} {
		b := New(size)
		if b.Size() != size {
			t.Error("size doesn't match")
		}
	}
}

func TestSet(t *testing.T) {
	b := New(10)
	b.Set(2)
	values, err := b.Values()
	if err != nil {
		t.Error(err)
	}
	if !slicesEqual(values, []uint64{2}) {
		t.Error("values do not match")
	}
}

func slicesEqual(from []uint64, to []uint64) bool {
	return reflect.DeepEqual(from, to)
}

func TestIsSet0(t *testing.T) {
	b := New(50)
	b.Set(2)
	if !get(b, 2) {
		t.Error("expected true, got false")
	}
	if get(b, 3) {
		t.Error("expected false, got true")
	}
	b.Set(3)
	if !get(b, 3) {
		t.Error("expected true, got false")
	}
	// Larger number (to prove indexing past the first
	// byte works).
	if get(b, 42) {
		t.Error("expected false, got true")
	}
	b.Set(42)
	if !get(b, 42) {
		t.Error("expected true, got false")
	}
}

// Setting a value twice should not unset it.
func TestSetTwice(t *testing.T) {
	b := New(10)
	b.Set(2)
	if !get(b, 2) {
		t.Error("expected it to be set")
	}
	b.Set(2)
	if !get(b, 2) {
		t.Error("expected it to still be set")
	}
}

// Unset a value.
func TestSetUnset(t *testing.T) {
	b := New(10)
	err := b.Set(2)
	if err != nil {
		t.Error(err)
	}
	if !get(b, 2) {
		t.Error("expected it to be set")
	}
	err = b.Unset(2)
	if err != nil {
		t.Error(err)
	}
	if get(b, 2) {
		t.Error("expected it to be unset")
	}
	b.Unset(2)
	if get(b, 2) {
		t.Error("expected it to still be unset")
	}
}

// TestValues tests the retrieval of a slice of
// values from a BitMap.
func TestValues(t *testing.T) {
	b := New(42)
	b.Set(2)
	b.Set(3)
	b.Set(13)
	b.Set(42)
	values, err := b.Values()
	if err != nil {
		t.Error(err)
	}
	if !slicesEqual(values, []uint64{2, 3, 13, 42}) {
		t.Error("didn't receive the expected values")
	}
}

// TestSetOverflow tests dealing with out-of-range issues
// in the Set method.
func TestSetOverflow(t *testing.T) {
	b := New(42)
	err := b.Set(52)
	if err != ErrOutOfRange {
		t.Error("out of range, there should be an error")
	}
	err = b.Set(0)
	if err != ErrOutOfRange {
		t.Error("out of range, there should be an error")
	}
}

// TestIsSet tests dealing with out-of-range issues
// in the IsSet method.
func TestIsSet(t *testing.T) {
	b := New(42)
	_, err := b.IsSet(52)
	if err != ErrOutOfRange {
		t.Error("out of range, there should be an error")
	}
	err = b.Set(0)
	if err != ErrOutOfRange {
		t.Error("out of range, there should be an error")
	}
}
