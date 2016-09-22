package main

import (
	"net"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// call getConfig first!
func listen(config *ssh.ServerConfig, port int, errCh chan<- error) error {
	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	for {
		// Once a ServerConfig has been configured, connections can be accepted.
		rawConn, err := listener.Accept()
		if err != nil {
			errCh <- err
			continue
		}

		// Before use, a handshake must be performed on the incoming net.Conn.
		serverConn, sshChanCh, reqCh, err := ssh.NewServerConn(rawConn, config)
		if err != nil {
			errCh <- err
			continue
		}

		// The incoming Request channel must be serviced.
		go ssh.DiscardRequests(reqCh)
		go handleServerConn(sConn.Permissions.Extensions["key-id"], chans)
	}
}
