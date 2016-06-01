package utils

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// Utility to help parse http responses. It works with a reader stream and decodes into the given interface
// Param 't' is the destination struct, etc. for the parsed data.
func DecodeResponseBody(respBody io.ReadCloser, t interface{}) error {
	defer closeRespBody(respBody)

	if err := json.NewDecoder(respBody).Decode(t); err == nil {
		//log.Printf("Decoded interface [%v]\n", t)
		return nil
	} else {
		log.Printf("[ERROR] Decode process returned the following err [%s]\n", err.Error())
		return err
	}
}

//Returns an HTTP Client with a connection timeout and an I/O timeout.
func NewTimeoutClient(connectTimeout time.Duration, readWriteTimeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(connectTimeout, readWriteTimeout),
		},
	}
}

//Sets up a net dialer with connection and read/write timeouts used when creating HTTP Clients with timeout capability.
func timeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(netw, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func closeRespBody(respBody io.ReadCloser) {
	if err := respBody.Close(); err != nil {
		log.Printf("[ERROR] Failed to close response Body [%s]\n", err.Error())
	} else {
		//		log.Println("***OK:UnmarshalResponseBody/respBody io.ReadCloser.Close() ***")
	}
}
