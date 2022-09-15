package requests

import (
	//"bytes"

	"bytes"
	"encoding/base64"
	"encoding/json"
	//"fmt"

	//"io"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/onsi/gomega"
)

//This is for bearer tokens if needed
func GetToken() string {
	access_key := "XXXXXXXXXXXXXXXXX"
	secret_key:= "XXXXXXXXXXXXXXXXX"
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://XXXXXXXXXXXXX.XXX/auth/token", nil)
	req.Close = true
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(access_key, secret_key))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	j := make(map[string]interface{})
	json.Unmarshal(body, &j)
	return j["access_token"].(string)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func GetResult(data string, endpoint string, method string, token string) (*http.Response, []byte) {
	client := &http.Client{}
	var request *http.Request
	var resp *http.Response
	var err error
	var body []byte
	url := "https://api.agify.io/" + endpoint
	if data != "" {
		var b = []byte(data)
		request, err = http.NewRequest(method, url, bytes.NewBuffer(b))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	Expect(err).ShouldNot(HaveOccurred())
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Add("Authorization", "Bearer "+token)
	resp, err = client.Do(request)
	body, err = ioutil.ReadAll(resp.Body)
	Expect(err).ShouldNot(HaveOccurred())
	defer resp.Body.Close()
	return resp, body
}
