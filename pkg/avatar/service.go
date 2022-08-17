package avatar

import (
	"errors"
	"fmt"

	"github.com/HernAle/AvataredBeta/pkg/avatar/encoder"
	"github.com/HernAle/AvataredBeta/pkg/avatar/images"
)

//dataEncoder encode and normalize information from any personal data
type dataEncoder interface {
	EncodeInformation(userInformation string) (encodedInformation []byte, err error)
}

//avatarGenerator create a single avatar with encoded information
type avatarGenerator interface {
	BuildAndSaveImage(encodedInformation []byte) error
}

//Service list both dependencies
type Service struct {
	encoder   dataEncoder
	generator avatarGenerator
}

//GenerateAndSaveAvatar joins both dependencies to give full functionality
func (service *Service) GenerateAndSaveAvatar(userData string) error {
	encodedData, err := service.encoder.EncodeInformation(userData)
	if err != nil {
		return errors.New(fmt.Sprintf("an error ocurred trying to encode the data"))
	}
	genErr := service.generator.BuildAndSaveImage(encodedData)
	if genErr != nil {
		return errors.New(fmt.Sprintf("an error ocurred trying to build and save image"))
	}
	return nil
}

func NewService() *Service {
	return &Service{
		encoder:   encoder.NewMD5Encoder(),
		generator: images.NewAvatar(),
	}
}
