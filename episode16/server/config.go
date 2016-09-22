package main

import (
	"golang.org/x/crypto/ssh"
)

func pkCallback(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	return &ssh.Permissions{}, nil
}

func getConfig(pkBytes []byte) (*ssh.ServerConfig, error) {
	config := &ssh.ServerConfig{
		PublicKeyCallback: pkCallback,
	}

	pk, err := ssh.ParsePrivateKey(pkBytes)
	if err != nil {
		return nil, err
	}
	config.AddHostKey(pk)

	return config, nil
}
