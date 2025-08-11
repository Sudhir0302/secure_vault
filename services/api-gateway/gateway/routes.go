package gateway

import (
	"net/http/httputil"
	"net/url"
)

func ConfigGateway(tar string) *httputil.ReverseProxy {
	url, _ := url.Parse(tar)
	return httputil.NewSingleHostReverseProxy(url)
}
