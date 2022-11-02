package toolkit

import "testing"

// test the RandomString function
func TestRandomString(t *testing.T) {
	var testTools Tool

	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("RandomString length should be 10")
	}

	s = testTools.RandomString(0)
	if len(s) != 0 {
		t.Error("RandomString length should be 0")
	}
}
