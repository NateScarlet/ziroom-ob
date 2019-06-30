package api

// RoomStatus returns formatted status from ziroom status code.
func RoomStatus(code string) string {
	switch code {
	case "dzz":
		return "待租中"
	case "zxpzz":
		return "装修配置中"
	case "tzpzz":
		return "退租配置中"
	case "ycz":
		return "已出租"
	case "yxd":
		return "已预订"
	case "sfz":
		return "收房中"
	case "zzz":
		return "转租中"
	case "dtz":
		return "待退中"
	case "pzzkyd":
		return "配置中可预订"
	default:
		return "未知"
	}
}
