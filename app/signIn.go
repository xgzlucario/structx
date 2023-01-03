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

// SignIn: Daily SignIn Data Structure
type SignIn struct {
	dateLogs *structx.SyncMap[DateID, *structx.BitMap]
	userLogs *structx.SyncMap[UserID, *structx.BitMap]
}

var (
	// ZeroTime: Make sure the signIn date is greater than ZeroTime
	ZeroTime, _ = time.Parse("2006-01-02", "2023-01-01")
)

// ParseDateInt: Return days to ZeroTime
func ParseDateInt(date time.Time) DateID {
	return DateID(date.Sub(ZeroTime).Hours() / 24)
}

// NewSignIn
func NewSignIn() *SignIn {
	return &SignIn{
		dateLogs: structx.NewSyncMapStringer[DateID, *structx.BitMap](),
		userLogs: structx.NewSyncMapStringer[UserID, *structx.BitMap](),
	}
}

// Sign
func (s *SignIn) Sign(userID UserID, dateID DateID) {
	// userLog
	bm, ok := s.userLogs.Get(userID)
	if !ok {
		bm = structx.NewBitMap()
		s.userLogs.Set(userID, bm)
	}
	bm.Add(uint(dateID))

	// dateLog
	bm, ok = s.dateLogs.Get(dateID)
	if !ok {
		bm = structx.NewBitMap()
		s.dateLogs.Set(dateID, bm)
	}
	bm.Add(uint(userID))
}

// UserCount: Get the number of days users have signed in
func (s *SignIn) UserCount(id UserID) int {
	bm, ok := s.userLogs.Get(id)
	if !ok {
		return -1
	}
	return bm.Len()
}

// UserDetails: Get user sign-in date slices
func (s *SignIn) UserDetails(id UserID) []uint {
	bm, ok := s.userLogs.Get(id)
	if !ok {
		return nil
	}
	return bm.ToSlice()
}

// UserGetMax: Get the user's most recent sign-in date
func (s *SignIn) UserGetMax(id UserID) int {
	bm, ok := s.userLogs.Get(id)
	if !ok {
		return -1
	}
	return bm.GetMax()
}

// DateCount: Get the total number of sign-in for the day
func (s *SignIn) DateCount(id DateID) int {
	bm, ok := s.dateLogs.Get(id)
	if !ok {
		return -1
	}
	return bm.Len()
}
