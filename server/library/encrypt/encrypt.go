// Package encrypt

package encrypt

import (
	"crypto/md5"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gutil"
	"hash/fnv"
)

// Md5ToString 生成md5
func Md5ToString(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Md5 生成md5
func Md5(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func Hash32(b []byte) uint32 {
	h := fnv.New32a()
	h.Write(b)
	return h.Sum32()
}

func EncodePassword(password string, salt string) string {
	return gmd5.MustEncryptString(password + "@" + salt)
}

func CheckPassword(inputPassword, password, salt string) bool {
	gutil.Dump("####", EncodePassword(inputPassword, salt))
	if EncodePassword(inputPassword, salt) == password {
		return true
	}
	return false
}
