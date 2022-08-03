package main

import (
	"github.com/HernAle/AvataredBeta/pkg/avatar"
)

func main() {
	avt := avatar.NewService()
	avt.GenerateAndSaveAvatar("anybody@email.com")
}
