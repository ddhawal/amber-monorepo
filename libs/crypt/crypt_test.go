package crypt

import (
  "testing"
)

func TestCrypt(t *testing.T) {
  result := Crypt("works")
  if result != "Crypt ahe mi works" {
    t.Error("Expected Crypt to append 'works'")
  }
}
