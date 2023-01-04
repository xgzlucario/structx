package app

import (
	"errors"
	"strconv"
	"sync"
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
	// Use Add() instead of AddDate()
	return ZeroTime.Add(time.Duration(id) * DateDuration)
}

var (
	// ZeroTime: Make sure the sign date is greater than ZeroTime
	ZeroTime, _ = time.Parse("2006-01-02", "2023-01-01")

	// Sign date duration
	DateDuration = time.Hour * 24
)

// SignIn: Threadsafe Sign-In Data Structure
type SignIn struct {
	mu      sync.RWMutex
	dateMap *structx.SyncMap[dateID, *structx.BitMap]
	userMap *structx.SyncMap[userID, *structx.BitMap]
}

// NewSignIn
func NewSignIn() *SignIn {
	return &SignIn{
		dateMap: structx.NewSyncMapStringer[dateID, *structx.BitMap](),
		userMap: structx.NewSyncMapStringer[userID, *structx.BitMap](),
	}
}

// Insert: Insert a sign-in record
func (s *SignIn) Insert(userId uint, date time.Time) error {
	userID, dateID := userID(userId), s.parseDateID(date)
	s.mu.Lock()
	defer s.mu.Unlock()

	// userLog
	bm, ok := s.userMap.Get(userID)
	if !ok {
		bm = structx.NewBitMap()
		s.userMap.Set(userID, bm)
	}
	// check if signed in
	if ok = bm.Add(uint(dateID)); !ok {
		return errors.New("sign-in record already exist")
	}

	// dateLog
	bm, ok = s.dateMap.Get(dateID)
	if !ok {
		bm = structx.NewBitMap()
		s.dateMap.Set(dateID, bm)
	}
	// check if signed in
	if ok = bm.Add(userId); !ok {
		return errors.New("sign-in record already exist")
	}

	return nil
}

// UserCount: Get the number of days users have signed in
// 用户签到总天数
func (s *SignIn) UserCount(userId uint) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	bm, ok := s.userMap.Get(userID(userId))
	if !ok {
		return -1
	}

	return bm.Len()
}

// UserDates: Get user sign-in date slices
// 用户签到日期列表
func (s *SignIn) UserDates(userId uint) []time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()

	bm, ok := s.userMap.Get(userID(userId))
	if !ok {
		return nil
	}

	// parse timeSlice
	times := make([]time.Time, 0, bm.Len())
	for _, id := range bm.ToSlice() {
		times = append(times, dateID(id).ToDate())
	}

	return times
}

// UserRecentDate: Get the user's most recent sign-in date
// 用户最近签到日期
func (s *SignIn) UserRecentDate(userId uint) time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()

	bm, ok := s.userMap.Get(userID(userId))
	if !ok {
		return time.Time{}
	}

	return dateID(bm.Max()).ToDate()
}

// UserContinuousCount: Count the number of consecutive days the user has signed in
// 用户连续签到统计
func (s *SignIn) UserContinuousCount(userId uint) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	bm, ok := s.userMap.Get(userID(userId))
	if !ok {
		return -1
	}
	return bm.MaxBitCount()
}

// DateCount: Get the total number of sign-in for the day
// 当日签到总量统计
func (s *SignIn) DateCount(date time.Time) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	id := s.parseDateID(date)
	bm, ok := s.dateMap.Get(id)
	if !ok {
		return -1
	}

	return bm.Len()
}

// parseDateID: Return days to ZeroTime
func (s *SignIn) parseDateID(date time.Time) dateID {
	return dateID(date.Sub(ZeroTime) / DateDuration)
}
