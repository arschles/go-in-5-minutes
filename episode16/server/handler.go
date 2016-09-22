package main

import (
	"golang.org/x/crypto/ssh"
)

func handleServerConn(keyID string, sshChanCh <-chan ssh.NewChannel) {
	for sshCh := range sshChanCh {
		if sshCh.ChannelType() != "session" {
			sshCh.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		sshCh, reqCh, err := sshCh.Accept()
		if err != nil {
			// handle error
			continue
		}

	}
}
