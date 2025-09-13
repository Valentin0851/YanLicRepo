package service

import test "other_files/commonlibsLesson/pkg/api/test/api"

type Service struct {
	test.OrderServiceServer
}

func New() *Service {
	return &Service{}
}
