package cryptoauth

import (
	"io"
)

type SessionManager struct {
	io.Reader
	io.Writer
}

func NewCIA(privateKey string) *CIA {
	cia := CIA{}
	return &cia
}

func (c *CIA) Write(p []byte) (int, error) {

}
