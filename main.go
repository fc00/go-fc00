package main

import (
	"log"
	"net"

	"github.com/fc00/go-fc00/cryptoauth"
)

func main() {
	lstn, err := listen("127.0.0.1:65432")
	log.Printf("listening on %s\n", lstn.LocalAddr())

	pk := "acf0d4098ef24f946c68e6e8c85470246fcfb290b501d4c06f0b24b7c9159582"
	cia := cryptoauth.NewCIA(pk)

	for {
		pkt := make([]byte)
		rsize, raddr, err := lstn.ReadFromUDP(pkt)
		if err != nil {
			log.Printf("read error: %s", err)
			continue
		}

		wsize, err := cia.Write(pkt)
		if err != nil {
			log.Printf("CIA.Write error: %s", err)
			continue
		}
		if rsize != wsize {
			log.Printf("CIA.Write error: read %d bytes, but wrote %d bytes", rsize, wsize)
		}
	}
}

// 2 goroutines per peer connection: 1 read, 1 write

func listen(address string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:65432")
	if err != nil {
		return err
	}

	lstn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return err
	}
	return lstn
}
