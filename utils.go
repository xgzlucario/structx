package structx

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"unicode"
)

// InSlice
func InSlice[T Value](val T, arr []T) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

// Expression
func Expression[T Value](isTrue bool, yes T, no T) T {
	if isTrue {
		return yes
	} else {
		return no
	}
}

// MD5Sum
func Md5Sum(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	val := hex.EncodeToString(m.Sum(nil))
	return val
}

// Check str is Chinese
func IsChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// Mean
func Mean[T Number](arr []T) float64 {
	var sum T
	for _, i := range arr {
		sum += i
	}
	return float64(sum) / float64(len(arr))
}

// Sum
func Sum[T Number](arr []T) T {
	var sum T
	for _, i := range arr {
		sum += i
	}
	return sum
}

// Go Job for every duration
func GoJob(f func(), duration time.Duration, delay ...time.Duration) {
	go func() {
		for _, dl := range delay {
			time.Sleep(dl)
		}
		for {
			f()
			time.Sleep(duration)
		}
	}()
}
