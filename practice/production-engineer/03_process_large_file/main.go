package main

// go run 03_process_large_file.go 127.0.0.1 2222

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// sshStuff()
	streamFile()
}

func streamFile() {
	file, err := os.Open("03_process_large_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	chunkSize := 13
	chunk := make([]byte, chunkSize)
	for {
		if _, err := file.Read(chunk); err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				return
			}

			fmt.Println("ERROR:", err)
			return
		}

		// TODO: Add custom logic here during stream process
		fmt.Println(string(chunk))
	}
}

func sshStuff() {
	if len(os.Args) < 3 {
		fmt.Println("ERROR: not enough args provided. ${1} == host ${2} == port")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	fmt.Printf("connecting to: %s %s\n", host, port)

	pemBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/.ssh/id_rsa_joynbio", os.Getenv("HOME")))
	if err != nil {
		log.Fatal(err)
	}

	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		log.Fatal(err)
	}

	config := &ssh.ClientConfig{
		User: "vagrant",
		Auth: []ssh.AuthMethod{
			ssh.Password("vagrant"),
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			// use OpenSSH's known_hosts file if you care about host validation
			return nil
		},
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	return
}
