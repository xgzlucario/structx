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
		for j := 0; j < 10; j++ {
			s.Sign(app.UserID(i), now+app.DateID(j))
		}
	}
}
