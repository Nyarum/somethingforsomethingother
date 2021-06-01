package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"sync"

	"somethingforsomethingother/services"
)

func validURL(str string) (string, bool) {
	u, err := url.Parse(str)
	if err != nil {
		return "", false
	}

	if u.RequestURI() == "" {
		return "", false
	}

	if u.Scheme == "" {
		u.Scheme = "http"
	}

	return u.String(), true
}

func main() {
	numParallel := flag.Int("parallel", 10, "How maximum executing parallel requests (Default: 10)")
	flag.Parse()

	addresses := []string{}
	for _, val := range os.Args {
		formatURL, ok := validURL(val)
		if !ok {
			continue
		}

		addresses = append(addresses, formatURL)
	}

	if len(addresses) == 0 {
		fmt.Println("You didn't set any address to execute it")
	}

	var (
		wg          = sync.WaitGroup{}
		controlWork = make(chan struct{}, *numParallel)
	)
	for _, addr := range addresses {
		addrCopy := addr

		wg.Add(1)
		go func() {
			defer wg.Done()
			controlWork <- struct{}{}

			md5Resp, err := services.RequestWithMD5Response(addrCopy)
			if err != nil {
				// We don't need to log any errors by the task
				//log.Println("Can't execute, error -" + err.Error())
			} else {
				fmt.Printf("%s %s\n", addrCopy, md5Resp)
			}

			<-controlWork
		}()
	}
	wg.Wait()
}
