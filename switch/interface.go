package packetswitch

import (
	ma "github.com/jbenet/go-multiaddr"
)

type Conn interface {
	io.Reader
	io.Writer

	LocalAddr() ma.Multiaddr
	RemoteAddr() ma.Multiaddr
}

type SwitchController interface {
	AddTransport(maddr ma.Multiaddr) error

	// Opens a peering connection to the given remote,
	// wraps it in a cryptoauth.Session, then adds it to the Switch.
	// Then sends a switch-ping, which gets the cryptoauth handshake going.
	// Returns once the switch-ping has returned or timed out.
	Connect(rmaddr ma.Multiaddr) error

	// Adds an existing connection.
	// Doesn't wrap in cryptoauth.Session, but sends switch-ping
	// This is used to add a "local" connection,
	// which isn't backed by a network connection, but by local code.
	AddLocalConn(conn *Conn) error

	// Kills the cryptoauth session,
	// and removes the connection from the Switch.
	Disconnect(rmaddr ma.Multiaddr) error

	Conns() ([]*Conn, error)
}

type Switch interface {
	io.Writer

	AddConn(slot int, conn *Conn)
	RemoveConn(rmaddr ma.Multiaddr)
}

//                     1               2               3
//     0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//  0 |                                                               |
//    +                         Switch Label                          +
//  4 |                                                               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//  8 |   Congest   |S| V |labelShift |            Penalty            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// 12 |                                                               |
//    +                 Data Packet or Control Packet                 +
//    |
//
type SwitchPacket interface {
	Label() uint64
	Congest() uint8
	SuppressErrors() uint8
	Version() uint8
	LabelShift() uint8
	Penalty() uint16
	DataPacket() DataPacket
	ControlPacket() ControlPacket
}

// If a received packet is for LocalConn, it isn't a CryptoPacket!

//                     1               2               3
//     0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//  0 |                             Handle                            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//  4 |            Checksum           |            Type               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//  8 |                             Magic                             |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// 12 |                            Version                            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// 16 |                                                               |
//    +                        Control Payload                        +
//    |
//
type ControlPacket interface {
	Handle() uint32
	Checksum() uint16
	Type() uint16
	Magic() uint32
	Version() uint32
	Payload() []byte
}

//                    1               2               3
//    0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// 0 |  ver  | unusd |     unused    |         Content Type          |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// 1 |                                                               |
//   +                         Data Payload                          +
//   |
//
type DataPacket interface {
	Version() uint16
	ContentType() uint16
	Payload() []byte
}
