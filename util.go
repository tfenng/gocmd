package main

import "time"

const (
	DEFAULT_DT_FMT = "2006-01-02 15:04:05"
)

var week time.Duration

func Now() time.Time {
	return time.Now()
}

func NowTs() int32 {
	return int32(time.Now().Unix())
}

func NowStr() string {
	t := time.Now()
	s := t.Format(DEFAULT_DT_FMT)
	return s
}

func FormatTime(dt time.Time) string {
	return dt.Format(DEFAULT_DT_FMT)
}

func FormatTs(ts int) string {
	dt := time.Unix(int64(ts), 0)
	return dt.Format(DEFAULT_DT_FMT)
}

func Parse(dt_str string) time.Time {
	t, _ := time.Parse(DEFAULT_DT_FMT, dt_str)
	return t
}
