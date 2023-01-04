package test

import (
	"testing"
	"time"

	"github.com/xgzlucario/structx/app"
)

func BenchmarkSignIn(b *testing.B) {
	s := app.NewSignIn()
	now := time.Now()

	for i := 0; i < b.N; i++ {
		s.Sign(uint(i), now)
	}
}

func BenchmarkGetRecentSignInDate(b *testing.B) {
	s := app.NewSignIn()
	now := time.Now()

	for i := 0; i < b.N; i++ {
		s.Sign(1, now.AddDate(0, 0, -i))
	}
}
