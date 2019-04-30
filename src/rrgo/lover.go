package rrgo

import (
	"time"
)

const (
	nickname        = "zl"
	birthdayISO8601 = "1992-05-25T00:00:00+08:00"
)

var (
	birthday time.Time
)

func init() {
	birthday, _ = time.Parse(time.RFC3339, birthdayISO8601)
}

func NextBirthdayAnniversary() time.Time {
	nba, now := birthday, time.Now() // next birthday anniversary
	for nba.Before(now) {
		nba = nba.AddDate(1, 0, 0)
	}
	return nba
}

func NextBirthdayAnniversaryCountdown() time.Duration {
	return NextBirthdayAnniversary().Sub(time.Now())
}
