package test

import (
	"testing"
	"time"

	"github.com/xgzlucario/structx/app"
)

const signDays = 1000

func getSignIn() *app.SignIn {
	s := app.NewSignIn()
	now := time.Now()
	for i := time.Duration(0); i < signDays; i++ {
		now = now.Add(time.Hour * 24)
		s.Insert(1, now)
	}
	return s
}

func BenchmarkSignIn(b *testing.B) {
	s := app.NewSignIn()
	now := time.Now()
	for i := 0; i < b.N; i++ {
		s.Insert(uint(i), now)
	}
}

func BenchmarkSignIn2(b *testing.B) {
	s := app.NewSignIn()
	now := time.Now()
	for i := 0; i < b.N; i++ {
		now = now.Add(time.Hour * 24)
		s.Insert(1, now)
	}
}

func BenchmarkUserCount(b *testing.B) {
	s := getSignIn()
	for i := 0; i < b.N; i++ {
		s.UserCount(1)
	}
}

func BenchmarkUserRecentDate(b *testing.B) {
	s := getSignIn()
	for i := 0; i < b.N; i++ {
		s.UserRecentDate(1)
	}
}

func BenchmarkUserDates(b *testing.B) {
	s := getSignIn()
	for i := 0; i < b.N; i++ {
		s.UserDates(1)
	}
}

func BenchmarkUserContinuousCount(b *testing.B) {
	s := getSignIn()
	for i := 0; i < b.N; i++ {
		s.UserContinuousCount(1)
	}
}
