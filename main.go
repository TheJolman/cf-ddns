package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type ipinfo_result struct {
	IP            string `json:"ip"`
	ASN           string `json:"asn"`
	AsName        string `json:"as_name"`
	AsDomain      string `json:"as_domain"`
	CountryCode   string `json:"country_code"`
	Country       string `json:"country"`
	ContinentCode string `json:"continent_code"`
	Continent     string `json:"continent"`
}

func help() {
	fmt.Printf("%s\nUpdate dynamic public IPs\n", os.Args[0])
}

func get_ip(client http.Client, version int) (string, error) {
	if version != 4 && version != 6 {
		return "", fmt.Errorf("version must be 4 or 6")
	}
	ipinfo_api_token := os.Getenv("IPINFO_API_TOKEN")
	if ipinfo_api_token == "" {
		return "", fmt.Errorf("IPINFO_API_TOKEN is empty")
	}

	url, err := url.Parse(fmt.Sprintf("https://v%d.api.ipinfo.io/lite/me?token=%s", version, ipinfo_api_token))
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http status %s", resp.Status)
	}

	var result ipinfo_result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}
	return result.IP, nil
}

func update_dns_record(version int) {
	client := http.Client{}

	ip, err := get_ip(client, version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to check ip: %v", err)
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
