package api

import "strings"

var cityCodeMap = map[string]string{
	"BJ": "110000",
	"SH": "310000",
	"SZ": "440300",
	"HZ": "330100",
	"NJ": "320100",
	"GZ": "440100",
	"CD": "510100",
	"WH": "420100",
	"TJ": "120000",
}

//CityCode returns ziroom citycode from city short name.
func CityCode(name string) string {
	key := strings.ToUpper(name)
	value := cityCodeMap[key]
	if value == "" {
		value = "110000"
	}
	return value
}

//Host for each city
func Host(cityCode string) string {
	for k, v := range cityCodeMap {
		if k == "BJ" {
			k = "www"
		}
		if v == cityCode {
			return strings.ToLower(k) + ".ziroom.com"
		}
	}
	return "www.ziroom.com"
}
