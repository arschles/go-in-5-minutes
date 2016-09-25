package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

// handles the stream of new incoming channels from a server connection. nomenclature is a bit confusing here because we talk about an SSH channel and a Go channel that receives SSH channels. we'll distinguish by saying "Go channel" and "SSH channel"
func handleNewSSHChannels(keyID string, sshChanCh <-chan ssh.NewChannel) {
	// use a range to accept the new SSH channels that come over the Go channel
	for sshCh := range sshChanCh {
		// make sure we're dealing with a session
		if sshCh.ChannelType() != "session" {
			sshCh.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		sshCh, reqCh, err := sshCh.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleSSHChannel(sshCh, reqCh)

	}
}

func handleSSHChannel(sshCh ssh.Channel, reqsCh <-chan *ssh.Request) {
	defer sshCh.Close()
	for req := range reqsCh {
		payload := string(req.Payload)
		str := fmt.Sprintf("received payload %s", payload)
		log.Println(str)
		sshCh.Write([]byte(str))
	}
}
