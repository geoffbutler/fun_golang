package main

import "testing"


func TestReturnTrue(t *testing.T) {
  if !ReturnTrue() {
    t.Fatalf("ReturnTrue result should be true")
  }
}

func TestReturnFalse(t *testing.T) {
  if ReturnFalse() {
    t.Fatalf("ReturnFalse result should be false")
  }
}
