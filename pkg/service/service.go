package service

import (
	lavka "github.com/EltIsma/YandexLavka"
	"github.com/EltIsma/YandexLavka/pkg/repository"
)

type CouriersList interface {
	Create(list lavka.Couriers) (int, error)
	GetAll(limit int, offset int)([]lavka.Couriers, error)
	GetById(courierId int) (lavka.Couriers, error)
	GetCouriersSalaryRating(courier_id int, start_date string, end_date string)(lavka.Couriers, error)
}

type OrderList interface {
	Create(list lavka.Orders) (int, error)
	GetAll(limit int, offset int)([]lavka.Orders, error)
	GetById(orderId int) (lavka.Orders, error)
	Update(ordComp lavka.OrdersComplete) (int, error)
}

type Service struct {
	CouriersList
	OrderList
}

func NewSrvice(repos *repository.Repository) *Service {
	return &Service{
		CouriersList: NewCouriersListService(repos.CouriersList),
		OrderList: NewOrdersListService(repos.OrderList),
	}
}
