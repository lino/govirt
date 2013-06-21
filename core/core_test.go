package core

import "testing"

func TestConnection(t *testing.T) {
	_, err := Connect("test:///default")
	if err != nil {
		t.Error(err)
	}
}

func TestFailingConnection(t *testing.T) {
	_, err := Connect("foo42:///bar")
	if err == nil {
		t.Error("Can connect to impossible URL")
	}
}
