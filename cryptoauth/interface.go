package cryptoauth

import (
	ma "github.com/jbenet/go-multiaddr"
)

type Conn interface {
	io.Reader
	io.Writer

	LocalAddr() ma.Multiaddr
	RemoteAddr() ma.Multiaddr
}

type Session interface {
}

func NewSessionManager(private)

type SessionManager interface {
}

type CryptoAuthPacket interface {
}
