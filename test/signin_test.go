package test

import (
	"testing"
	"time"

	"github.com/xgzlucario/structx/app"
)

func BenchmarkSignIn(b *testing.B) {
	s := app.NewSignIn()
	now := app.ParseDateInt(time.Now())

	for i := 0; i < b.N; i++ {
		s.Sign(app.UserID(i), now)
	}
}
