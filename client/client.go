package client

import (
	"os"
	"fmt"
	"net/http"
)

type login struct {
	Ip string `json:"ip"`
}

type Client struct {
}

func (c *Client) Login(serverAddress string) {

	ip := ResolveHostIp()
	fmt.Println("Using host IP: ", ip)

	// Register at Server
	response, err := loginRequest(ip, serverAddress)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if response.StatusCode != 200 {
		fmt.Println("got unexpected status code after login: ", response.StatusCode)
	}
}

/*
	todo implement login request

	Your struct is good to go. Just encode to json and send the encoded value (POST body)
 */
func loginRequest(ip string, serverAddress string) (*http.Response, error) {
	l := login{Ip: ip}
	fmt.Println(l)
	return http.Post(serverAddress+"/login", "application/json", nil)
}

/*
	todo return host ip

	1. Easy - hardcode your ip address
	2. Advanced - implement function to retrieve your ipv4 address and return it (don't use loopback address)
*/
func ResolveHostIp() string {
	return "192.168.188.20"
}
