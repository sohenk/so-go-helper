package gormtimeutils

import "testing"

//go test -v
func TestTTT(t *testing.T) {
	t.Log("statm,", TransformTimestrToTimestamp("2021-01-30T07:37:29.000+0000"))
}
