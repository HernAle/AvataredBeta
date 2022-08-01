package avatar

//infoEncoder encode and normalize information from any personal data
type infoEncoder interface {
	EncodeInformation(userInformation string) (encodedInformation []byte, err error)
}

//imageGenerator create a single avatar with encoded information
type imageGenerator interface {
	BuildAndSaveImage(encodedInformation []byte) error
}

//Service has both dependencies
type Service struct {
	encoder   infoEncoder
	generator imageGenerator
}

func NewService(e infoEncoder, g imageGenerator) *Service {
	return &Service{
		encoder:   e,
		generator: g,
	}
}

func (s *Service) GenerateAndSaveAvatar(email string) error {
	//Here will be all  logic
	encodedEmail, err := s.encoder.EncodeInformation(email)
	if err != nil {
		panic(err)
	}
	genErr := s.generator.BuildAndSaveImage(encodedEmail)
	if genErr != nil {
		panic(genErr)
	}
	return nil
}
