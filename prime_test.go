package main

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	actual := split(86981717)
	expected := []int{86981717}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad factors: %v", actual)
	}
	actual = split(86981716)
	expected = []int{2, 2, 21745429}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad factors: %v", actual)
	}
}
