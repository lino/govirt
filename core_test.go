package govirt

import "testing"

func TestConnection(t *testing.T) {
	conn, err := Connect("test:///default")
	if err != nil {
		t.Error(err)
	}
	conn.Disconnect()
}

func TestFailingConnection(t *testing.T) {
	conn, err := Connect("foo42:///bar")
	if err == nil {
		t.Error("Can connect to impossible URL")
	}
	conn.Disconnect()
}
