package repository

import (
	"fmt"
	"net/http"

	lavka "github.com/EltIsma/YandexLavka"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	
	_ "github.com/lib/pq"
)

type OrdersListPostgres struct {
	db *sqlx.DB
}

func NewOrdersListPostgres(db *sqlx.DB) *OrdersListPostgres {
	return &OrdersListPostgres{db: db}
}

func (r *OrdersListPostgres) Create(list lavka.Orders) (int, error) {
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (weight, cost, region, delivery_hours) VALUES($1, $2, $3, $4) RETURNING id", ordersTable)
	row := r.db.QueryRow(createListQuery, list.Weight, list.Cost, list.Region, list.DeliveryHours)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrdersListPostgres) GetAll(limit int, offset int) ([]lavka.Orders, error) {
	var ordersList []lavka.Orders

	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", ordersTable, limit, offset)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order lavka.Orders
		//var region []uint8
		//var workingHours []uint8
		err := rows.Scan(&order.Id, &order.Weight, &order.Cost, &order.Region, &order.DeliveryHours)
		if err != nil {
			return nil, err
		}
		//courier.Region = make([]int, len(region))
		//courier.Working_Hours = make([]string, len(workingHours))
		/*for i, r := range region {
		      courier.Region[i] = int(r)
		  }
		  for i, w := range workingHours {
		      courier.Working_Hours[i] = string(w)
		  }*/
		ordersList = append(ordersList, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	//rr := r.db.Select(&couriersList, query)
	return ordersList, nil
}

func (r *OrdersListPostgres) GetById(orderId int) (lavka.Orders, error) {
	var oList lavka.Orders
	
	query := fmt.Sprintf("SELECT *  FROM %s WHERE id = %d ", ordersTable, orderId)

	row, err := r.db.Query(query)
	if err != nil {
		return lavka.Orders{}, err
	}
	defer row.Close()
	for row.Next() {
		var order lavka.Orders
		err := row.Scan(&order.Id, &order.Weight, &order.Cost, &order.Region, &order.DeliveryHours,  &order.CompleteOrders)
		if err != nil {
			return lavka.Orders{}, err
		}
		oList = order
	}
	if err := row.Err(); err != nil {
		return lavka.Orders{}, err
	}
	//err := r.db.Get(&cList, query, courierId)
	return oList, err
}

func (r *OrdersListPostgres) Update(ordComp lavka.OrdersComplete) (int, error) {

	var exist bool

	err := r.db.QueryRow("select exists(select couriers_id from completeOrderdate where couriers_id=$1 AND $2= ANY(orders_id))", ordComp.CurierId, ordComp.OrderId).Scan(&exist)
	if err != nil {
		return 0, err
	}
	/*tx, err :=  r.db.Begin()

		if err != nil {
	    return 0, err
		}*/
	if exist {
		updateListQuery := fmt.Sprintf("UPDATE %s SET complete_time = $1 WHERE id= $2 AND complete_time is NULL", ordersTable)
		_, err := r.db.Exec(updateListQuery, ordComp.CompleteTime, ordComp.OrderId)
		if err != nil {

			return 0, err
		}
	} else {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "HTTP 400 Bad Request")
	}
	/*createCompleteQuery := fmt.Sprintf("INSERT INTO %s (couriers_id, orders_id, complete_time) VALUES($1, $2, $3)",completeOrdersTable )
	_, err = tx.Exec(createCompleteQuery, ordComp.CurierId, ordComp.OrderId, ordComp.CompleteTime )

	if err!= nil{
		tx.Rollback()
		return 0, err
	}*/

	return ordComp.OrderId, nil
}
