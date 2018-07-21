package routing

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRouter_getHandler(t *testing.T) {
	t.Parallel()
	cases := []struct {
		path    string
		method  HttpMethod
		handler func(http.ResponseWriter, *http.Request)
	}{
		{
			path:    "/simple",
			method:  HttpMethodGet,
			handler: func(w http.ResponseWriter, r *http.Request) {},
		},
		{
			path:    "/two/slashes",
			method:  HttpMethodGet,
			handler: func(w http.ResponseWriter, r *http.Request) {},
		},
		{
			path:    "no_head_slash",
			method:  HttpMethodGet,
			handler: func(w http.ResponseWriter, r *http.Request) {},
		},
		{
			path:    "complex/no_head_slash",
			method:  HttpMethodGet,
			handler: func(w http.ResponseWriter, r *http.Request) {},
		},
	}
	for i, tc := range cases {
		i := i
		tc := tc
		t.Run(fmt.Sprintf("TestRouter_getHandler cases #%d", i), func(t *testing.T) {
			r := NewRouter()
			h := r.getHandler(tc.path, tc.method)
			if h != nil {
				t.Fatalf("router has been just initialized, did not expect to get a handler")
			}
			r.HandleFunc(tc.path, tc.method, tc.handler)
			h = r.getHandler(tc.path, tc.method)
			if h == nil {
				t.Fatalf("path %s should have been handled, but getHandler returned nil", tc.path)
			}
		})
	}
	noHandlerCases := []struct {
		path    string
		method  HttpMethod
		handler func(http.ResponseWriter, *http.Request)
	}{
		{
			path:    "/simple",
			method:  HttpMethodGet,
			handler: nil,
		},
		{
			path:    "/two/slashes",
			method:  HttpMethodGet,
			handler: nil,
		},
		{
			path:    "no_head_slash",
			method:  HttpMethodGet,
			handler: nil,
		},
		{
			path:    "complex/no_head_slash",
			method:  HttpMethodGet,
			handler: nil,
		},
	}

	for i, tc := range noHandlerCases {
		i := i
		tc := tc
		t.Run(fmt.Sprintf("TestRouter_getHandler noHandlerCases #%d", i), func(t *testing.T) {
			r := NewRouter()
			h := r.getHandler(tc.path, tc.method)
			if h != nil {
				t.Fatalf("router has been just initialized, did not expect to get a handler")
			}
			r.HandleFunc(tc.path, tc.method, tc.handler)
			h = r.getHandler(tc.path, tc.method)
			if h == nil {
				t.Fatalf("path %s should be assoicated with a nil handler, but getHandler returned an http.Handler", tc.path)
			}
		})
	}
}
