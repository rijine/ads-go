package counters

type Service interface {
	Count(of string) int64
}

type service struct{}

func NewCounterService() Service {
	return &service{}
}

func (s *service) Count(of string) int64 {
	return 1
}
