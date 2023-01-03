package app

import (
	"fmt"
	"time"

	"github.com/xgzlucario/structx"
)

type UserID uint
type DateID uint

func (id DateID) String() string {
	return fmt.Sprintf("%d", id)
}

func (id UserID) String() string {
	return fmt.Sprintf("%d", id)
}

// SignIn 每日签到数据结构
type SignIn struct {
	dateLogs *structx.SyncMap[DateID, *structx.BitMap]
	userLogs *structx.SyncMap[UserID, *structx.BitMap]
}

var zeroTime = time.Time{}

// ParseDateInt
func ParseDateInt(date time.Time) DateID {
	// 距离 zeroTime 所差天数
	days := date.Sub(zeroTime).Hours() / 24
	return DateID(days)
}

// NewSignIn
func NewSignIn() *SignIn {
	return &SignIn{
		dateLogs: structx.NewSyncMapStringer[DateID, *structx.BitMap](),
		userLogs: structx.NewSyncMapStringer[UserID, *structx.BitMap](),
	}
}

// Sign 签到
func (s *SignIn) Sign(userID UserID, dateID DateID) {
	// userLog
	bm, ok := s.userLogs.Get(userID)
	if !ok {
		bm = structx.NewBitMap()
		s.userLogs.Set(userID, bm)
	}
	bm.Add(uint(userID))

	// dateLog
	bm, ok = s.dateLogs.Get(dateID)
	if !ok {
		bm = structx.NewBitMap()
		s.dateLogs.Set(dateID, bm)
	}
	bm.Add(uint(dateID))
}

// UserCount 获取用户签到天数
func (s *SignIn) UserCount(id UserID) int {
	bm, ok := s.userLogs.Get(id)
	if !ok {
		return -1
	}
	return bm.Len()
}

// UserDetails 获取用户签到日期详情
func (s *SignIn) UserDetails(id UserID) []uint {
	bm, ok := s.userLogs.Get(id)
	if !ok {
		return nil
	}
	return bm.ToSlice()
}

// UserGetMax 获取用户最近签到日期
func (s *SignIn) UserGetMax(id UserID) int {
	bm, ok := s.userLogs.Get(id)
	if !ok {
		return -1
	}
	return bm.GetMax()
}

// DateCount 获取当日签到总量
func (s *SignIn) DateCount(id DateID) int {
	bm, ok := s.dateLogs.Get(id)
	if !ok {
		return -1
	}
	return bm.Len()
}
