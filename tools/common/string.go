package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	kAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Layout    = "20060102150405"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//随机前缀
func RandString32ByPrefix(prefix string) string {

	now := time.Now()
	nanoStr := fmt.Sprintf("%s%d", now.Format(Layout), now.Nanosecond()/1000)
	if len(nanoStr) >= 32 {
		return prefix + nanoStr[0:32]
	}
	b := make([]byte, 32-len(nanoStr))
	for i := range b {
		b[i] = kAlphaNum[rand.Int63()%int64(len(kAlphaNum))]
	}
	return prefix + nanoStr + string(b)
}

func RandString32() string {

	now := time.Now()
	nanoStr := fmt.Sprintf("%s%d", now.Format(Layout), now.Nanosecond()/1000)
	if len(nanoStr) >= 32 {
		return nanoStr[0:32]
	}
	b := make([]byte, 32-len(nanoStr))
	for i := range b {
		b[i] = kAlphaNum[rand.Int63()%int64(len(kAlphaNum))]
	}
	return nanoStr + string(b)
}

func RandStringWithRange(n int) string {
	res := rand.Intn(n)
	return strconv.Itoa(res)
}

func RandNumberStr(n int) string {
	b := ""
	for i := 0; i < n; i++ {
		b += strconv.Itoa(rand.Int() % 10)
	}
	return string(b)
}

//src 源串指针，当src为空时，将dft值赋值给它
func FillStrDefault(src *string, dft string) {
	if len(*src) < 1 {
		*src = dft
	}
}

//MD5编码
func StringToMd5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
