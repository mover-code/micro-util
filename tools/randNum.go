package tools

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 生成 6位随机数字验证码
func GetRandomNumber() string {
	rand.Seed(time.Now().UnixNano() - time.Now().Unix()) // rand seed
	randNums := strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
	return randNums
}

/**
len: 长度
max: 个数
随机生成长度为len的max个纯数字随机数
*/
func RandNum(length int, max int) []string {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	//rand.Seed(42)

	digitNumber := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0",
	}
	// 用make创建map

	set := New()
	for set.Len() < max {
		ranNumber := ""
		for j := 1; j < length; j++ {
			ranNumber += digitNumber[rand.Intn(len(digitNumber))]
		}
		if !set.Has(ranNumber) {
			set.Add(ranNumber)
		}
	}

	return set.List()

}

/**
构造set类型
*/
type Set struct {
	m map[string]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[string]bool{},
	}
}

func (s *Set) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	var list []string
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
