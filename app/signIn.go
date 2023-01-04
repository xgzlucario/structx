package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xgzlucario/structx"
)

type userID uint
type dateID uint

func (id userID) String() string {
	return strconv.Itoa(int(id))
}

func (id dateID) String() string {
	return strconv.Itoa(int(id))
}

func (id dateID) ToDate() time.Time {
	return ZeroTime.Add(time.Duration(id) * oneDay)
}

var (
	// ZeroTime: Make sure the signIn date is greater than ZeroTime
	ZeroTime, _ = time.Parse("2006-01-02", "2023-01-01")
)

const (
	oneDay = time.Hour * 24
)

// SignIn: Daily SignIn Data Structure
type SignIn struct {
	dateLogs *structx.SyncMap[dateID, *structx.BitMap]
	userLogs *structx.SyncMap[userID, *structx.BitMap]
}

// NewSignIn 构造函数
func NewSignIn() *SignIn {
	return &SignIn{
		dateLogs: structx.NewSyncMapStringer[dateID, *structx.BitMap](),
		userLogs: structx.NewSyncMapStringer[userID, *structx.BitMap](),
	}
}

// Sign 签到
func (s *SignIn) Sign(userId uint, date time.Time) error {
	userID, dateID := userID(userId), s.parseDateID(date)

	// userLog
	bm, ok := s.userLogs.Get(userID)
	if !ok {
		bm = structx.NewBitMap()
		s.userLogs.Set(userID, bm)
	}
	// check if sign-in
	if ok = bm.Add(uint(dateID)); ok {
		return fmt.Errorf("user[%v] date[%v] already signed in", userID, dateID)
	}

	// dateLog
	bm, ok = s.dateLogs.Get(dateID)
	if !ok {
		bm = structx.NewBitMap()
		s.dateLogs.Set(dateID, bm)
	}
	// check if sign-in
	if ok = bm.Add(uint(userID)); ok {
		return fmt.Errorf("user[%v] date[%v] already signed in", userID, dateID)
	}

	return nil
}

// UserCount: Get the number of days users have signed in
// 用户签到总天数
func (s *SignIn) UserCount(userId uint) int {
	bm, ok := s.userLogs.Get(userID(userId))
	if !ok {
		return -1
	}

	return bm.Len()
}

// UserDates: Get user sign-in date slices
// 用户签到日期列表
func (s *SignIn) UserDates(userId uint) []time.Time {
	bm, ok := s.userLogs.Get(userID(userId))
	if !ok {
		return nil
	}

	// parse timeSlice
	dateIDs := bm.ToSlice()
	times := make([]time.Time, 0, len(dateIDs))
	for _, id := range dateIDs {
		times = append(times, dateID(id).ToDate())
	}

	return times
}

// UserRecentDate: Get the user's most recent sign-in date
// 用户最近签到日期
func (s *SignIn) UserRecentDate(userId uint) time.Time {
	bm, ok := s.userLogs.Get(userID(userId))
	if !ok {
		return time.Time{}
	}
	id := bm.GetMax()
	
	return dateID(id).ToDate()
}

// DateCount: Get the total number of sign-in for the day
// 当日签到数量统计
func (s *SignIn) DateCount(date time.Time) int {
	id := s.parseDateID(date)
	bm, ok := s.dateLogs.Get(id)
	if !ok {
		return -1
	}

	return bm.Len()
}

// parseDateID: Return days to ZeroTime
func (s *SignIn) parseDateID(date time.Time) dateID {
	return dateID(date.Sub(ZeroTime) / oneDay)
}
