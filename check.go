package main

import (
	"bytes"
	"encoding/json"
	"errors"
	proto "github.com/tiankui626/mgtvPmpProto"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Request(data string) (*proto.Response, error) {
	if verbose {
		log.Println("request body:", data)
	}
	body := bytes.NewBuffer([]byte(data))
	resp, err := http.Post(apiUrl, "application/json", body)

	if verbose {
		log.Printf("response header:%s\n", resp.Header)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if verbose {
		log.Printf("response content :%s\n", result)
	}
	if err != nil {
		return nil, err
	}
	t := proto.Response{}
	err = json.Unmarshal(result, &t)
	return &t, err
}

func check(data string, resp *proto.Response) error {
	request, err := getObject(data)
	if err != nil {
		log.Fatalf("request data err:%s", err)
	}
	if resp == nil {
		return errors.New("response is nil")
	}
	if resp.Version != 3 {
		return errors.New("response version not equal 3")
	}
	if resp.Bid != request.Bid {
		return errors.New("bid not equal request")
	}

	if resp.ErrCode == 204 {
		return errors.New("response code 204, has no bid, please bid the request")
	}
	if resp.ErrCode != 200 {
		return errors.New("response error code != 200")
	}

	if len(resp.Ads) == 0 {
		return errors.New("response ADS is zero")
	}
	ads := resp.Ads[0]
	if ads.Price == 0 || ads.Price < request.Imp[0].MinCpmPrice {
		return errors.New("ads price is 0 or less than request price")
	}

	if ads.AdUrl == "" {
		return errors.New("ads AdUrl is nil ")
	}
	if !strings.HasPrefix(ads.AdUrl, "http://mp4.res.hunantv.com/") {
		return errors.New("ad url not start with http://mp4.res.hunantv/")
	}
	if ads.ClickThroughUrl == "" {
		return errors.New("has no click url")
	}
	if len(ads.IUrl) == 0 {
		return errors.New("has no monitor url")
	}
	if ads.IUrl[0].Event != 0 {
		return errors.New("i url event now just support 0 (start play event)")
	}
	if ads.IUrl[0].Url == "" {
		return errors.New("i url url is nil")
	}
	if ads.Title == "" {
		return errors.New("need set Title ")
	}
	if ads.Ctype != request.Imp[0].Ctype[0] {
		return errors.New("ads ctype not equal requested")
	}
	if ads.Duration == 0 {
		return errors.New("duration is 0")
	}
	return nil
}

func getObject(data string) (*proto.Request, error) {
	var pmpRequest proto.Request
	err := json.Unmarshal([]byte(data), &pmpRequest)
	return &pmpRequest, err
}
