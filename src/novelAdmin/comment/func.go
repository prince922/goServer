package comment

import (
	"crypto/md5"
	"fmt"
)

func Str2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}
