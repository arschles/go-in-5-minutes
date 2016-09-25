package main

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ServerConfig{
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			// this return effectively accepts all users. generally you'd want to look up the public key in a database to authorize and authenticate the user
			return nil, nil
		},
	}

	errCh := make(chan error)
	log.Printf("listening on port 8080")
	listen(config, 8080, errCh)
	for err := range errCh {
		log.Println(err)
	}
}
