package avatar

//Encode information from any personal data
type infoEncoder interface {
	EncodeInformation(userInformation string) (encodedInformation []byte, err error)
}

//Transform encode information to avatar
type imageGenerator interface {
	BuildAndSaveImage(encodedInformation []byte) error
}

//Build funcionalities between interfaces
type Service struct {
	encoder   infoEncoder
	generator imageGenerator
}

type Information struct {
	//Here go all the information that you want to encode could be eMail string
}

func (s *Service) GenerateAndSaveAvatar(information Information) error {
	//Here will be all  logic
	return nil
}
