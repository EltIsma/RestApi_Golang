package repository

import (
	lavka "github.com/EltIsma/YandexLavka"
	"github.com/jmoiron/sqlx"
)

type CouriersList interface {
	Create(list lavka.Couriers) (int, error)
	GetAll(limit int, offset int)([]lavka.Couriers, error)
	GetById(courierId int) (lavka.Couriers, error)
	GetCouriersSalaryRating(courier_id int, start_date string, end_date string)(lavka.Couriers, int, error)
}

type OrderList interface {
	Create(list lavka.Orders) (int, error)
	GetAll(limit int, offset int)([]lavka.Orders, error)
	GetById(orderId int) (lavka.Orders, error)
	Update(ordComp lavka.OrdersComplete) (int, error)
}

type Repository struct {
	CouriersList
	OrderList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CouriersList: NewCouriersListPostgres(db),
		OrderList: NewOrdersListPostgres(db),
	}
}