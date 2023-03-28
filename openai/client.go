package openai

import (
	"crypto/tls"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func NewClient() *http.Client {
	client := &http.Client{}
	return client
}

func NewHttpProxyClient(proxyConf string) (*http.Client, error) {
	proxyUrl, err := url.Parse(proxyConf)
	if err != nil {
		log.Errorf("%+v", err)
		return nil, err
	}
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	return client, nil
}
