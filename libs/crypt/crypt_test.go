package crypt

import (
  "testing"
)

func TestCrypt(t *testing.T) {
  result := Crypt("works")
  if result != "Crypt output works" {
    t.Error("Expected Crypt to append 'works'")
  }
}
