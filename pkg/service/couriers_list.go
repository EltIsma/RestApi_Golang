package service

import (
	lavka "github.com/EltIsma/YandexLavka"
	"github.com/EltIsma/YandexLavka/pkg/repository"
)

type CouriersListService struct {
	repo repository.CouriersList
}


func NewCouriersListService(repo repository.CouriersList) *CouriersListService{
	return &CouriersListService{repo: repo}
}

func (s *CouriersListService) Create(list lavka.Couriers) (int, error) {
	return s.repo.Create(list)
}

func(s *CouriersListService) GetAll(limit int, offset int)([]lavka.Couriers, error){
	return s.repo.GetAll(limit, offset)
}
func(s *CouriersListService) GetById(courierId int) (lavka.Couriers, error){
	return s.repo.GetById(courierId)
}

func(s *CouriersListService) GetCouriersSalaryRating(courier_id int, start_date string, end_date string)(lavka.Couriers, error){
	metacourier, countOrdersCompleted, err := s.repo.GetCouriersSalaryRating(courier_id, start_date, end_date)
   if err != nil{
      return lavka.Couriers{},  err
   }
	earning := metacourier.Earnings(countOrdersCompleted)
	rating := metacourier.Rating(countOrdersCompleted, start_date, end_date)
	metacourier.Earning=earning
	metacourier.Ratings=rating

	return metacourier, nil
}