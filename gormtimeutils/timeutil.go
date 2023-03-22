package gormtimeutils

import (
	"database/sql"
	"time"
)

// str to sqlnulltime
func StrToSqlNullTime(str string) sql.NullTime {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: theTime, Valid: true}
}

func SqlNullTimeToStr(sqlNullTime sql.NullTime) string {
	if sqlNullTime.Valid {
		return sqlNullTime.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}
func TransformTimestrToTimestamp(timestr string) int64 {
	//2006-01-02 15:04:05 是golang的时间模板，据说是golang语言的诞生时间，2006-01-02 15:04:05类似于我们熟悉的YYYY-MM-dd HH:mm:ss
	result, err := time.ParseInLocation("2006-01-02T15:04:05+0800", timestr, time.Local)
	//如果错误则退出
	if err != nil {
		return -1
	}
	//转为13位时间戳 乘以1000是因为原来是秒单位，乘以之后则是13位毫秒时间戳单位
	return result.Unix() * 1000
}
