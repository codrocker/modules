package modules

import (
	"time"
	"math/rand"
)

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterBytesLower = "0123456789abcdefghijklmnopqrstuvwxyz"
const letterBytesNum = "0123456789"
const letterBytesLowercase = "abcdefghijklmnopqrstuvwxyz"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandStringBytesLower(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytesLower[rand.Intn(len(letterBytesLower))]
	}
	return string(b)
}

func RandStringBytesNum(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytesNum[rand.Intn(len(letterBytesNum))]
	}
	return string(b)
}

func RandStringBytesLowercase(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytesLowercase[rand.Intn(len(letterBytesNum))]
	}
	return string(b)
}



//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}