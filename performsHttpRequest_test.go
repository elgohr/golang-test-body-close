package closing_test

import (
	"github.com/elgohr/closing"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBeClosedWhenClosed(t *testing.T) {
	tripper := NewFakeRoundTripper()
	cl := closing.MyClient{
		Client: http.Client{
			Transport: tripper,
		},
	}

	if err := cl.Closing(); err != nil {
		t.Error(err)
	}

	if !tripper.Body.Closed {
		t.Error("Should be closed, but wasn't")
	}
}

func TestShouldBeOpenWhenNotClosed(t *testing.T) {
	tripper := NewFakeRoundTripper()
	cl := closing.MyClient{
		Client: http.Client{
			Transport: tripper,
		},
	}

	if err := cl.NotClosing(); err != nil {
		t.Error(err)
	}

	if tripper.Body.Closed {
		t.Error("Should be open, but wasn't")
	}
}

func NewFakeRoundTripper() *FakeRoundTripper {
	return &FakeRoundTripper{
		Body: &FakeReadCloser{
			ReadCloser: ioutil.NopCloser(strings.NewReader("content")),
		},
	}
}

type FakeRoundTripper struct {
	Body *FakeReadCloser
}

func (r *FakeRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{
		Body: r.Body,
	}, nil
}

type FakeReadCloser struct {
	io.ReadCloser
	Closed bool
}

func (r *FakeReadCloser) Close() error {
	r.Closed = true
	return r.ReadCloser.Close()
}
