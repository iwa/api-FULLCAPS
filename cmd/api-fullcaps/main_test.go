package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouteUppercase(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "POST with lowercase",
			method:         "POST",
			body:           "hello world",
			wantStatusCode: http.StatusOK,
			wantBody:       "HELLO WORLD",
		},
		{
			name:           "POST with mixed case",
			method:         "POST",
			body:           "GoLang",
			wantStatusCode: http.StatusOK,
			wantBody:       "GOLANG",
		},
		{
			name:           "POST with empty body",
			method:         "POST",
			body:           "",
			wantStatusCode: http.StatusOK,
			wantBody:       "",
		},
		{
			name:           "GET method",
			method:         "GET",
			body:           "should fail",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Bad request\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/uppercase", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()

			routeUppercase(rr, req)

			if rr.Code != tc.wantStatusCode {
				t.Errorf("expected status %d, got %d", tc.wantStatusCode, rr.Code)
			}
			if rr.Body.String() != tc.wantBody {
				t.Errorf("expected body %q, got %q", tc.wantBody, rr.Body.String())
			}
		})
	}
}

func TestRouteLowercase(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "POST with uppercase",
			method:         "POST",
			body:           "HELLO WORLD",
			wantStatusCode: http.StatusOK,
			wantBody:       "hello world",
		},
		{
			name:           "POST with mixed case",
			method:         "POST",
			body:           "GoLang",
			wantStatusCode: http.StatusOK,
			wantBody:       "golang",
		},
		{
			name:           "POST with empty body",
			method:         "POST",
			body:           "",
			wantStatusCode: http.StatusOK,
			wantBody:       "",
		},
		{
			name:           "GET method",
			method:         "GET",
			body:           "should fail",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Bad request\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/uppercase", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()

			routeLowercase(rr, req)

			if rr.Code != tc.wantStatusCode {
				t.Errorf("expected status %d, got %d", tc.wantStatusCode, rr.Code)
			}
			if rr.Body.String() != tc.wantBody {
				t.Errorf("expected body %q, got %q", tc.wantBody, rr.Body.String())
			}
		})
	}
}
