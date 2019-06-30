package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomIDDetailURL(t *testing.T) {
	r := RoomID{CityCode: "110000", ID: "61819181"}
	u := r.detailURL()
	assert.Equal(t, "http://m.ziroom.com/wap/detail/room.json?city_code=110000&id=61819181", u.String())
}

func TestFetchRoomData(t *testing.T) {
	r := RoomID{CityCode: "110000", ID: "62337045"}
	d, err := r.FetchData()
	assert.NoError(t, err)
	assert.Equal(t, r.CityCode, d.CityCode)
	assert.Equal(t, r.ID, d.ID)
	assert.Equal(t, "空气检测中", d.NoticeWord)
	assert.Equal(t, "dzz", d.Status)
}

func TestFetchRoomDataWhenNoNoticeWord(t *testing.T) {
	r := RoomID{CityCode: "110000", ID: "61819181"}
	d, err := r.FetchData()
	assert.NoError(t, err)
	assert.Equal(t, r.CityCode, d.CityCode)
	assert.Equal(t, r.ID, d.ID)
	assert.Equal(t, "", d.NoticeWord)
	assert.Equal(t, "dzz", d.Status)
}

func TestRoomURL(t *testing.T) {
	r := RoomID{CityCode: "110000", ID: "61819181"}
	u := r.RoomURL()
	assert.Equal(t, "https://www.ziroom.com/z/vr/61819181.html", u.String())
}
