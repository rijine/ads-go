package counters

type Service interface {
	Count(of string) (int64, error)
}

var (
	counterRepo = NewCounterRepository()
)

type service struct{}

func NewCounterService() Service {
	return &service{}
}

func (s *service) Count(of string) (int64, error) {
	return counterRepo.GetAndUpdate(of)
}
