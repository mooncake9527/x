package common

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	ClientIP := c.ClientIP()
	//fmt.Println("ClientIP:", ClientIP)
	RemoteIP := c.RemoteIP()
	//fmt.Println("RemoteIP:", RemoteIP)
	ip := c.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = c.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if RemoteIP != "127.0.0.1" {
		ip = RemoteIP
	}
	if ClientIP != "127.0.0.1" {
		ip = ClientIP
	}
	return ip
}

func init() {
	//var dbPath = "ip2region.xdb"
	//xdb.InitIPInfo(dbPath)
	// 	xdb.LoadHeaderFromFile(dbPath)
	// 	cBuff, err := xdb.LoadContentFromFile(dbPath)
	// 	if err != nil {
	// 		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
	// 		return
	// 	}
	// 	searcher, err = xdb.NewWithBuffer(cBuff)
	// 	if err != nil {
	// 		fmt.Printf("failed to create searcher with content: %s\n", err)
	// 		return
	// 	}
}

// func GetGetLocationByIp2Region(ip string, location *IPLocationData) error {
// 	res, err := xdb.GetipInfo(ip)
// 	if err != nil {
// 		return err
// 	}
// 	// b, _ := json.Marshal(res)
// 	// fmt.Println(string(b))
// 	location.Country = res.Country
// 	location.City = res.City
// 	location.Province = res.Province
// 	location.Isp = res.ISP
// 	//location.CityCode = res.CityId
// 	location.Ip = res.Ip
// 	return nil
// }

// type IPLocation struct {
// 	Continent      string `json:"Continent"`
// 	Country        string `json:"Country"`
// 	CountryEnglish string `json:"CountryEnglish"`
// 	CountryCode    string `json:"CountryCode"`
// 	Province       string `json:"Province"`
// 	ProvinceEn     string `json:"ProvinceEn"` //省: "江苏",
// 	City           string `json:"City"`
// 	CityEn         string `json:"CityEn"` //: "徐州",
// 	District       string `json:"District"`
// 	AreaCode       string `json:"AreaCode"`
// 	ISP            string `json:"ISP"`
// 	Longitude      string `json:"Longitude"`
// 	Latitude       string `json:"Latitude"`
// 	LocalTime      string `json:"LocalTime"`
// 	Elevation      string `json:"Elevation"`
// 	WeatherStation string `json:"WeatherStation"`
// 	ZipCode        string `json:"ZipCode"`
// 	CityCode       string `json:"CityCode"`
// 	Asn            string `json:"Asn"`
// 	Ip             string `json:"Ip"`
// 	Err            string `json:"err"`
// }

// type IPLocationData struct {
// 	AreaCode       string   `json:"area_code"`       //: "320311",
// 	Province       string   `json:"province"`        //省: "江苏",
// 	City           string   `json:"city"`            //: "徐州",
// 	District       string   `json:"district"`        //: "丰县",
// 	CityCode       string   `json:"city_code"`       //: "0516",
// 	Continent      string   `json:"continent"`       //: "亚洲",
// 	Country        string   `json:"country"`         //: "中国",
// 	CountryCode    string   `json:"country_code"`    //: "CN",
// 	CountryEnglish string   `json:"country_english"` //: "",
// 	Elevation      string   `json:"elevation"`       //: "40",
// 	Ip             string   `json:"ip"`              //: "114.234.76.140",
// 	Isp            string   `json:"isp"`             //: "电信",
// 	Latitude       string   `json:"latitude"`        //: "34.227883",
// 	LocalTime      string   `json:"local_time"`      //: "2023-08-02 14:36",
// 	Longitude      string   `json:"longitude"`       //: "117.213995",
// 	MultiStreet    []Street `json:"multi_street"`
// 	Street         string   `json:"street"`          //: "解放路168号",
// 	Version        string   `json:"version"`         //: "V4",
// 	WeatherStation string   `json:"weather_station"` //: "CHXX0437",
// 	ZipCode        string   `json:"zip_code"`        //: "221006"
// }

// type Street struct {
// 	Lng          string `json:"lng"`           //经度: "116.60833",
// 	Lat          string `json:"lat"`           //纬度: "34.701533",
// 	Province     string `json:"province"`      //省: "江苏",
// 	City         string `json:"city"`          //: "徐州",
// 	District     string `json:"district"`      //: "丰县",
// 	Street       string `json:"street"`        //: "解放路168号",
// 	StreetNumber string `json:"street_number"` //: "解放路168号"
// }

// func GetLocationByIp(ip string, location *IPLocationData) error {

// 	//return GetGetLocationByIp2Region(ip, location)

// 	// url := "http://myip.yunlogin.com/?ip=" + ip
// 	url := "http://myip.yunlogin.com/?db=ipdata&ip=" + ip
// 	client := http_util.HTTPClient{}
// 	data, err := client.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	var ipd IPLocation
// 	if err := json.Unmarshal(data, &ipd); err != nil {
// 		return err
// 	}
// 	if ipd.Err != "" {
// 		return errors.New(fmt.Sprintf("获取出错 code:{%s} msg{%s}", ip, ipd.Err))
// 	}
// 	location.Country = ipd.Country
// 	location.City = ipd.City
// 	location.Province = ipd.Province
// 	location.Isp = ipd.ISP
// 	location.Ip = ip
// 	if config.Ext.SystemName == "MT" {
// 		location.Country = ipd.CountryEnglish
// 		location.City = ipd.CityEn
// 		location.Province = ipd.ProvinceEn
// 	}
// 	return nil
// }
