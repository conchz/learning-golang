package main

import (
	"log"

	"github.com/valyala/fasthttp"
	"github.com/json-iterator/go"
)

type IpDetailsInfo struct {
	Country   string        `json:"country"`
	CountryId string        `json:"country_id"`
	Area      string        `json:"area"`
	AreaId    string        `json:"area_id"`
	Region    string        `json:"region"`
	RegionId  string        `json:"region_id"`
	City      string        `json:"city"`
	CityId    string        `json:"city_id"`
	County    string        `json:"county"`
	CountyId  string        `json:"county_id"`
	Isp       string        `json:"isp"`
	IspId     string        `json:"isp_id"`
	Ip        string        `json:"ip"`
}

func main() {
	c := &fasthttp.Client{}

	statusCode, body, err := c.Get(nil, "http://ip.taobao.com/service/getIpInfo.php?ip=61.139.2.69")

	if err != nil {
		log.Fatalf("%s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}

	useResponseBody(body)
}
func useResponseBody(body []byte) {
	//respStr := string(body[:])
	//log.Printf("Response string: %s\n", respStr)

	var ipInfo *IpDetailsInfo
	iter := jsoniter.ParseBytes(body)
	err := jsoniter.UnmarshalFromString(iter.ReadAny().Get("data").ToString(), &ipInfo)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(ipInfo)

	var f interface{}

	_err := jsoniter.Unmarshal(body, &f)
	if _err != nil {
		log.Fatal(_err)
	}
	log.Println(f)
}
