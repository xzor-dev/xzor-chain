package main

import (
	"net"
	"os"
	"testing"

	"github.com/xzor-dev/xzor/internal/xzor/action"
	"github.com/xzor-dev/xzor/internal/xzor/network"
)

func TestMessenger(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("%v", err)
	}
	instanceA, err := newInstance(dir + "/testdata/instanceA")
	if err != nil {
		t.Fatalf("%v", err)
	}
	connA1, connA2 := net.Pipe()
	instanceA.node.AddListener(&network.MockListener{
		Connections: []net.Conn{connA1},
	})

	_, err = instanceA.Execute(&action.Action{
		Module:    "messenger",
		Command:   "create-board",
		Arguments: []interface{}{"messages"},
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	instanceB, err := newInstance(dir + "/testdata/instanceB")
	if err != nil {
		t.Fatalf("%v", err)
	}
	instanceB.node.AddConnection(connA2)
}
