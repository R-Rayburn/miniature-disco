package main

import "testing"

func TestNewDeck(t *testing.T) {
    d:= newDeck(false)

    if len(d) != 52 {
        t.Errorf("expected deck length of 52, but got %v", len(d))
    }
}