package api

import (
	"net/url"
	"strings"
)

//ParseRoomURLString get RoomID from url string.
func ParseRoomURLString(s string) (*RoomID, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	return ParseRoomURL(u)
}

//ParseRoomURL get RoomID from url
func ParseRoomURL(u *url.URL) (ret *RoomID, err error) {
	ret = &RoomID{}
	if u.Host == "m.ziroom.com" {
		ret.CityCode = CityCode(strings.Split(u.Path, "/")[0])
		ret.ID = u.Query().Get("id")
	} else {
		ret.CityCode = CityCode(strings.Split(u.Host, ".")[0])
		parts := strings.Split(u.Path, "/")
		lastPart := parts[len(parts)-1]
		ret.ID = strings.Split(lastPart, ".")[0]
	}
	return
}
