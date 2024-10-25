package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateMatch() error {
	return nil
}

func (s *Service) Place(i int) error {
	return nil
}
