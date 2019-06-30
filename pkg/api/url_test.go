package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRoomURL_bj_pc(t *testing.T) {
	room, err := ParseRoomURLString("http://www.ziroom.com/z/vr/61819181.html")
	assert.NoError(t, err)
	assert.Equal(t, "110000", room.CityCode)
	assert.Equal(t, "61819181", room.ID)
}
func TestParseRoomURL_sh_pc(t *testing.T) {
	room, err := ParseRoomURLString("http://sh.ziroom.com/z/vr/61306999.html")
	assert.NoError(t, err)
	assert.Equal(t, "310000", room.CityCode)
	assert.Equal(t, "61306999", room.ID)
}

func TestParseRoomURL_bj_wap(t *testing.T) {
	room, err := ParseRoomURLString("http://m.ziroom.com/BJ/room?id=61819181")
	assert.NoError(t, err)
	assert.Equal(t, "110000", room.CityCode)
	assert.Equal(t, "61819181", room.ID)
}

func TestParseRoomURL_bj_pc_https(t *testing.T) {
	room, err := ParseRoomURLString("https://www.ziroom.com/z/vr/61819181.html")
	assert.NoError(t, err)
	assert.Equal(t, "110000", room.CityCode)
	assert.Equal(t, "61819181", room.ID)
}

func TestParseRoomURL_bj_wap_https(t *testing.T) {
	room, err := ParseRoomURLString("https://m.ziroom.com/BJ/room?id=61819181")
	assert.NoError(t, err)
	assert.Equal(t, "110000", room.CityCode)
	assert.Equal(t, "61819181", room.ID)
}
