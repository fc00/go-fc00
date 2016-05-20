package main

type Node interface {
	io.Reader
	io.Writer
}

type Self interface {
	Node
}

type SessionManager interface{}

type Switch interface{}

type PeeringManager interface{}

type Transport interface{}


package main

import (
  "github.com/fc00/go-fc00/cryptoauth"
  "github.com/fc00/go-fc00/node"
  "github.com/fc00/go-fc00/peering"
  "github.com/fc00/go-fc00/switch"
)

func main() {
  pk := "acf0d4098ef24f946c68e6e8c85470246fcfb290b501d4c06f0b24b7c9159582"
  sess := cryptoauth.NewSessionManager(pk)
  slf := node.NewNode(pk, sess)

  ctl := peering.NewPeeringManager(sess)
  tr := peering.NewTransport("/ip4/127.0.0.1/udp/54321")
  ctl.AddTransport(tr)

  sw := swtch.NewSwitch(ctl)
  sw.AddPeer(slf)

  go tr.Listen()
  go func() {
    for {
      log.Debugf("received: %+v", read(slf))
    }
  }()

  write(slf, "asdf.k", "hello")
}

type message struct {
  dest     string
  category int8
  payload  string
}

func read(n *node.Node) message {
  msg := &message{}
  n.Read(msg)
  return msg
}

// n.Write figures out a switch label
func write(n *node.Node, dest string, payload string) {
  msg := &message{dest: dest, category: 42, payload: payload}
  n.Write(msg)
}

// one goroutine per peering connection
// a goroutine can write to any peering connection (threadsafe?)
