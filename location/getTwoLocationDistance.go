package location

import "math"

//获取两个经纬度之间的距离
func GetTwoLocationDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const EarthRadius = 6378.137 // 地球半径，单位为 km

	radLat1 := rad(lat1)
	radLat2 := rad(lat2)
	radLng1 := rad(lng1)
	radLng2 := rad(lng2)

	a := radLat1 - radLat2
	b := radLng1 - radLng2

	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2))) * EarthRadius

	return s
}

func rad(d float64) float64 {
	// 将角度转换为弧度
	return d * math.Pi / 180.0
}
