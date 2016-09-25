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
		// Once a ServerConfig has been configured, we can start accepting connections from clients
		rawConn, err := listener.Accept()
		if err != nil {
			errCh <- err
			continue
		}

		// Before use, we have to perform the SSH handshake
		serverConn, sshChanCh, reqCh, err := ssh.NewServerConn(rawConn, config)
		if err != nil {
			errCh <- err
			continue
		}

		// we have to service the incoming request channel. discarding it is acceptable in many cases
		go ssh.DiscardRequests(reqCh)
		// now we can handle the stream of new SSH channels
		go handleNewSSHChannels(serverConn.Permissions.Extensions["key-id"], sshChanCh)
	}
}
