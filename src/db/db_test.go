package db

import "testing"

func TestDbConnect(t *testing.T) {
	result, err := dbConnect()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 1 {
		t.Fatal("failed test")
	}
}
