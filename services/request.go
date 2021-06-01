package services

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestWithMD5Response(address string) (md5Resp string, err error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 Safari/537.36")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	respData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("Something wrong with request data, response status is " + fmt.Sprintf("%d", res.StatusCode))
	}

	md5Resp = fmt.Sprintf("%x", md5.Sum(respData))
	return
}
