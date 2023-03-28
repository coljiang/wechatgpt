package test

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"wechatbot/openai"
)

func TestProxy(t *testing.T) {
	client, err := openai.NewHttpProxyClient("http://10.222.96.210:7890")
	if err != nil {
		log.Errorf("%+v", err)
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "https://myip.ipip.net", bytes.NewBufferString(""))
	if err != nil {
		log.Errorf("%+v", err)
		t.Fatal(err)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Errorf("%+v", err)
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("%+v", err)
			t.Fatal(err)
		}
	}(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("%+v", err)
		t.Fatal(err)
	}
	log.Infof("body : %s", body)
}
