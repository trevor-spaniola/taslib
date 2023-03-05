package apiRequests

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

/*
GetApiRequestUserPass sends a GET request to the passed in API endpoint request, using
Username and Password Basic Authentication. The API JSON response is returned in a slice of bytes.
*/
func GetApiRequestUserPass(apiReq, apiUser, apiPass string) []byte {

	req, err := http.NewRequest("GET", apiReq, nil)
	if err != nil {
		log.Fatalln("[!] Error with GET request for:", apiReq, "ERROR OUTPUT:", err)
	}

	// Username and Password Basic Authentication
	req.SetBasicAuth(apiUser, apiPass)
	req.Header.Set("Accept", "application/json")

	// Submit GET request to API
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("[!] Error with GET request response for:", apiReq, "ERROR OUTPUT:", err)
	}
	defer resp.Body.Close()

	// Converts HTTP response into readable format
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("[!] Error with reading response for:", apiReq, "ERROR OUTPUT:", err)
	}

	return respBody
}

/*
GetApiRequestToken sends a GET request to the passed in API endpoint request,
using a Token for Basic Authentication. The API JSON response is returned in a slice of bytes.
*/
func GetApiRequestToken(apiReq, apiToken string) []byte {

	req, err := http.NewRequest("GET", apiReq, nil)
	if err != nil {
		log.Fatalln("[!] Error with GET request for:", apiReq, "ERROR OUTPUT:", err)
	}

	// Authorization Token Basic Authentication
	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Set("Accept", "application/json")

	// Submit GET request to API
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("[!] Error with GET request response for:", apiReq, "ERROR OUTPUT:", err)
	}
	defer resp.Body.Close()

	// Converts HTTP response into readable format
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("[!] Error with reading response for:", apiReq, "ERROR OUTPUT:", err)
	}

	return respBody
}

/*
GetApiRequestUserPassInsecureSSL sends a GET request to the passed in API endpoint request,
using Username and Password Basic Authentication, and ignore insecure SSL.
The API JSON response is returned in a slice of bytes.
*/
func GetApiRequestUserPassInsecureSSL(apiReq, apiUser, apiPass string) []byte {

	// For ignoring self-signed SSL cert
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: customTransport}

	req, err := http.NewRequest("GET", apiReq, nil)
	if err != nil {
		log.Fatalln("[!] Error with GET request for:", apiReq, "ERROR OUTPUT:", err)
	}

	// Username and Password Basic Authentication
	req.SetBasicAuth(apiUser, apiPass)
	req.Header.Set("Accept", "application/json")

	// Submit GET request to API
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("[!] Error with GET request response for:", apiReq, "ERROR OUTPUT:", err)
	}
	defer resp.Body.Close()

	// Converts HTTP response into readable format
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("[!] Error with reading response for:", apiReq, "ERROR OUTPUT:", err)
	}

	return respBody
}

/*
GetApiRequestTokenInsecureSSL sends a GET request to the passed in API endpoint request,
using a Token for Basic Authentication, and ignore insecure SSL.
The API JSON response is returned in a slice of bytes.
*/
func GetApiRequestTokenInsecureSSL(apiReq, apiToken string) []byte {

	// For ignoring self-signed SSL cert
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: customTransport}

	req, err := http.NewRequest("GET", apiReq, nil)
	if err != nil {
		log.Fatalln("[!] Error with GET request for:", apiReq, "ERROR OUTPUT:", err)
	}

	// Authorization Token Basic Authentication
	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Set("Accept", "application/json")

	// Submit GET request to API
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("[!] Error with GET request response for:", apiReq, "ERROR OUTPUT:", err)
	}
	defer resp.Body.Close()

	// Converts HTTP response into readable format
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("[!] Error with reading response for:", apiReq, "ERROR OUTPUT:", err)
	}

	return respBody
}
