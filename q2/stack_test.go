package main

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	s := &Stack{limit: 3}
	s.Push(1)
	if r := s.Pop(); r != 1 {
		t.Fatalf("want 1, got %v", r)
	}
	if r := s.Pop(); r != nil {
		t.Fatalf("want nil, got %v", r)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	if r := s.Pop(); r != 4 {
		t.Fatalf("want 4, got %v", r)
	}
	if r := s.Pop(); r != 3 {
		t.Fatalf("want 3, got %v", r)
	}
	if r := s.Pop(); r != 2 {
		t.Fatalf("want 2, got %v", r)
	}
	if r := s.Pop(); r != nil {
		t.Fatalf("want nil, got %v", r)
	}
}
