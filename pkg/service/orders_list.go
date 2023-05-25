package service

import (
	lavka "github.com/EltIsma/YandexLavka"
	"github.com/EltIsma/YandexLavka/pkg/repository"
)

type OrdersListService struct {
	repo repository.OrderList
}


func NewOrdersListService(repo repository.OrderList) *OrdersListService{
	return &OrdersListService{repo: repo}
}

func (s *OrdersListService) Create(list lavka.Orders) (int, error) {
	return s.repo.Create(list)
}

func(s *OrdersListService) GetAll(limit int, offset int)([]lavka.Orders, error){
	return s.repo.GetAll(limit, offset)
}
func(s *OrdersListService) GetById(orderId int) (lavka.Orders, error){
	return s.repo.GetById(orderId)
}

func (s *OrdersListService) Update(ordComp lavka.OrdersComplete) (int, error) {
	return s.repo.Update(ordComp)
}