package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func getClient(knowHosts []byte, ks ssh.Signer, usr string) (*ssh.Client, error) {
	for {

		_, _, hostKey, _, rest, err := ssh.ParseKnownHosts(knowHosts)
		if err != nil {
			return nil, err
		}

		config := &ssh.ClientConfig{
			User: usr,
			Auth: []ssh.AuthMethod{
				// Use the PublicKeys method for remote authentication.
				ssh.PublicKeys(ks),
			},
			HostKeyCallback: ssh.FixedHostKey(hostKey),
			//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		// Connect to the remote server and perform the SSH handshake.
		client, err := ssh.Dial("tcp", "18.236.67.86:22", config)
		if err == nil {
			return client, nil
		}
		knowHosts = rest

	}
}

func getWindClient(usr string, pswrd string ) (*ssh.Client, error){
	config := &ssh.ClientConfig{
		User: usr,
		Auth: []ssh.AuthMethod{
			// Use the Password method for remote authentication.
			ssh.Password(pswrd),
		},
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", "34.212.87.182:22", config)
	if err == nil {
		return client, nil
	}else{
		return nil, err
	}

}

func main() {
	fmt.Println("Hello World")

	//Get know hosts file
	knowHosts, err := os.ReadFile("C:/Users/esunb/.ssh/known_hosts")
	if err != nil {
		log.Fatalf("unable to read know hosts: %v", err)
	}

	//Get Private Key
	key, err := os.ReadFile("C:/Users/esunb/esundblad")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	//Establish SSH connection 
	conn, err := getClient(knowHosts, signer, "esundblad")		//linux machines
	//conn, err = getWindClient("administrator", "3xc0g1t@t3")  //Windows machines
	
	
	if err != nil {
		log.Fatalf("Failed to get client: %v", err)
	}
	defer conn.Close()

	//Creat sftp client
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//Print out files in current dir!
	dirs, err := client.ReadDir(".")
	if err != nil{
		log.Fatalf("Failed to read dir: %v", err)
	}

	for _, dir := range dirs{
		fmt.Printf("Name: %v Size %v\n", dir.Name(), dir.Size())
	}
	
}
