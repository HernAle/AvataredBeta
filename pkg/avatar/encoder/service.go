package encoder

import (
	"crypto/md5"
)

type MD5Encoder struct{}

func (*MD5Encoder) EncodeInformation(userInformation string) (encodedInformation []byte, err error) {
	avatarMD5 := md5.Sum([]byte(userInformation))
	return avatarMD5[:], nil
}

func NewMD5Encoder() *MD5Encoder {
	return &MD5Encoder{}
}
