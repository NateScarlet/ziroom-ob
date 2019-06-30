package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RoomResponse from ziroom API
type RoomResponse struct {
	Status       string   `json:"status"`
	ErrorCode    int      `json:"error_code"`
	ErrorMessage string   `json:"error_message"`
	Data         RoomData `json:"data"`
}

// RoomData in RoomResponse
type RoomData struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	CityCode   string `json:"city_code"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	NoticeWord string `json:"notice_word,omitempty"`
}

// RoomID for ziroom API
type RoomID struct {
	CityCode string
	ID       string
}

// RoomURL on ziroom pc website
func (r *RoomID) RoomURL() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   Host(r.CityCode),
		Path:   "/z/vr/" + r.ID + ".html",
	}
}

func (r *RoomID) detailURL() url.URL {
	ret := url.URL{
		Scheme: "http",
		Host:   "m.ziroom.com",
		Path:   "/wap/detail/room.json",
	}
	q := ret.Query()
	q.Set("city_code", r.CityCode)
	q.Set("id", r.ID)
	ret.RawQuery = q.Encode()
	return ret
}

//FetchData from ziroom api
func (r *RoomID) FetchData() (ret *RoomData, err error) {
	u := r.detailURL()
	resp, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error during fetch url: %s: %s", u.String(), resp.Status)
		return
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result := &RoomResponse{}
	json.Unmarshal(d, result)
	if result.Status != "success" {
		err = fmt.Errorf("Api result fail: %d %s: %s",
			result.ErrorCode, result.Status, result.ErrorMessage)
		return
	}
	ret = &result.Data
	if ret.ID != r.ID {
		err = fmt.Errorf("Expect result on room id: %s != %s", ret.ID, r.ID)
		return
	}
	if ret.CityCode != r.CityCode {
		err = fmt.Errorf("Expect result on room city code: %s != %s", ret.CityCode, r.CityCode)
		return
	}
	return
}

// FormatStatus to human-readable string.
func (d *RoomData) FormatStatus() string {
	ret := RoomStatus(d.Status)
	if d.NoticeWord != "" {
		ret += fmt.Sprintf("(%s)", d.NoticeWord)
	}
	return ret
}
