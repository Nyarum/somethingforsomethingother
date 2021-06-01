package services_test

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/http/httptest"
	"somethingforsomethingother/services"
	"testing"
)

func Test_RequestSuccess(t *testing.T) {
	body := "demo response"
	md5Body := fmt.Sprintf("%x", md5.Sum([]byte(body)))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(body))
	}))
	defer ts.Close()

	md5Resp, err := services.RequestWithMD5Response(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if md5Body != md5Resp {
		t.Errorf("md5 of response isn't right, got - %s and expected - %s", md5Resp, md5Body)
	}
}
