package main

import (
	"github.com/HernAle/AvataredBeta/pkg/avatar"
	"github.com/HernAle/AvataredBeta/pkg/avatar/encoder"
	"github.com/HernAle/AvataredBeta/pkg/avatar/images"
)

func main() {
	enc := encoder.NewMD5Encoder()
	imgGen := images.NewAvatar()

	avt := avatar.NewService(enc, imgGen)
	avt.GenerateAndSaveAvatar("lara@gmail.com")
}
