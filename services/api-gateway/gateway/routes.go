package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Gateway struct {
	*httputil.ReverseProxy
}

func ConfigGateway(tar string) *httputil.ReverseProxy {
	url, _ := url.Parse(tar)
	return httputil.NewSingleHostReverseProxy(url)
}

// auth
func AuthRoute(auth *httputil.ReverseProxy) {
	AuthGateway := &Gateway{auth}
	http.HandleFunc("/auth/", AuthGateway.AuthHandler)
}

// share
func ShareRoute(share *httputil.ReverseProxy) {
	ShareGateway := &Gateway{share}
	http.HandleFunc("/share/", ShareGateway.ShareHandler)
}

// store
func StoreRoute(store *httputil.ReverseProxy) {
	StoreGateway := &Gateway{store}
	http.HandleFunc("/store/", StoreGateway.StoreHandler)
}
