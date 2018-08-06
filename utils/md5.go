package utils

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func Md5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}
