package main

import (
	"github.com/HernAle/AvataredBeta/pkg/avatar"
	"github.com/HernAle/AvataredBeta/pkg/avatar/encoder"
)

func main() {
	enc := encoder.NewMD5Encoder()

	avatar.NewService(enc)
}
