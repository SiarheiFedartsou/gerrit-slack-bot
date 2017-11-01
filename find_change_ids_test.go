package main

import "testing"

func TestFindsChangeIdInUsualURL(t *testing.T) {
	message := "Hey, my commit is here https://test.url.com/42"

	ids := FindChangeIDs("https://test.url.com", message)

	if len(ids) != 1 {
		t.Errorf("got length == %d, want 1", len(ids))
	}

	if ids[0] != "42" {
		t.Errorf("got ID == %s, want 42", ids[0])
	}
}

func TestFindsChangeIdInURLWithCPrefix(t *testing.T) {
	message := "Hey, my commit is here https://test.url.com/#/c/42"

	ids := FindChangeIDs("https://test.url.com", message)

	if len(ids) != 1 {
		t.Errorf("got length == %d, want 1", len(ids))
	}

	if ids[0] != "42" {
		t.Errorf("got ID == %s, want 42", ids[0])
	}
}

func TestFindsChangeIdInURLWithRevisionNumberPostfix(t *testing.T) {
	message := "Hey, my commit is here https://test.url.com/#/c/42/1"

	ids := FindChangeIDs("https://test.url.com", message)

	if len(ids) != 1 {
		t.Errorf("got length == %d, want 1", len(ids))
	}

	if ids[0] != "42" {
		t.Errorf("got ID == %s, want 42", ids[0])
	}
}

func TestFindsAllIDs(t *testing.T) {
	message := "Hey, my commit is here https://test.url.com/#/c/42/1 and here https://test.url.com/43"

	ids := FindChangeIDs("https://test.url.com", message)

	if len(ids) != 2 {
		t.Errorf("got length == %d, want 2", len(ids))
	}

	if ids[0] != "42" {
		t.Errorf("got ID == %s, want 42", ids[0])
	}

	if ids[1] != "43" {
		t.Errorf("got ID == %s, want 43", ids[1])
	}
}
