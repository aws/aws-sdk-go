package aws

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var metadataClient = http.Client{
	Timeout: time.Second * 1,
}

func GetMetadata(path string) (contents []byte, err error) {
	url := "http://169.254.169.254/" + path

	res, err := metadataClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("Code %d returned for url %s", res.StatusCode, url)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), nil
}
