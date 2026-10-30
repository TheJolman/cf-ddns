package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var ipinfo_api_token = os.Getenv("IPINFO_API_TOKEN")

func help() {
	fmt.Printf("%s\nUpdate dynamic public IPs\n", os.Args[0])
}

func get_ip(client http.Client, version int) (string, error) {
	if version != 4 && version != 6 {
		return "", fmt.Errorf("version must be 4 or 6")
	}

	url, err := url.Parse(fmt.Sprintf("https://v%d.api.ipinfo.io/lite/me", version))
	if err != nil {
		return "", err
	}

	fmt.Printf("%s\n", url)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer"+ipinfo_api_token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error: got status %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	ip := string(body)

	return ip, nil
}

func update_dns_record(version int) {
	client := http.Client{}

	ip, err := get_ip(client, version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	fmt.Println(ip)
}

func main() {
	var version int

	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			help()
		case "-4":
			version = 4
		case "-6":
			version = 6
		default:
			fmt.Fprintf(os.Stderr, "Error: unrecognized argument '%s'\n", arg)
			help()
		}
	}

	update_dns_record(version)
}
