package stacks

import (
	"testing"
)

func TestMinMaxStack(t *testing.T) {
	mms := New()

	_, err := mms.Pop()

	if err == nil {
		t.Fatal()
	}

	_, err = mms.Peek()

	if err == nil {
		t.Fatal()
	}
	
	lastPush := 1
	mms.Push(lastPush)

	val, err := mms.Peek()
	if err != nil {
		t.Fatal()
	}

	if val != lastPush {
		t.Fatal()
	}

	min, err := mms.GetMin()
	if err != nil {
		t.Fatal()
	}

	if min != lastPush {
		t.Fatal()
	}

	max, err := mms.GetMax()
	if err != nil {
		t.Fatal()
	}

	if max != lastPush {
		t.Fatal()
	}

	nextPush := 2
	mms.Push(nextPush)

	val, err = mms.Peek()

	if err != nil {
		t.Fatal()
	}

	if val != nextPush{
		t.Fatal()
	}

	min, err = mms.GetMin()
	if err != nil {
		t.Fatal()
	}

	if min != lastPush {
		t.Fatal()
	}

	max, err = mms.GetMax()
	if err != nil {
		t.Fatal()
	}

	if max != nextPush {
		t.Fatal()
	}

	expectedLen := 2
	if mms.length() != expectedLen {
		t.Fatal()
	}

	val, err = mms.Pop()

	if val != nextPush {
		t.Fatal()
	}

	expectedLen = 1
	if mms.length() != expectedLen {
		t.Fatal()
	}

	min, err = mms.GetMin()
	if err != nil {
		t.Fatal()
	}

	if min != lastPush {
		t.Fatal()
	}

	max, err = mms.GetMax()
	if err != nil {
		t.Fatal()
	}

	if max != lastPush {
		t.Fatal()
	}
} 
