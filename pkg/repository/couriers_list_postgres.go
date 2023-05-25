package repository

import (
	"fmt"

	lavka "github.com/EltIsma/YandexLavka"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type CouriersListPostgres struct {
	db *sqlx.DB
}

func NewCouriersListPostgres(db *sqlx.DB) *CouriersListPostgres {
	return &CouriersListPostgres{db: db}
}

func (r *CouriersListPostgres) Create(list lavka.Couriers) (int, error) {
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (type, region, working_hours) VALUES($1, $2, $3) RETURNING id", couriersTable)
	row := r.db.QueryRow(createListQuery, list.Type, pq.Int32Array(list.Region), pq.StringArray(list.Working_Hours))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CouriersListPostgres) GetAll(limit int, offset int) ([]lavka.Couriers, error) {
	var couriersList []lavka.Couriers

	query := fmt.Sprintf("SELECT id,  type, region,working_hours FROM %s LIMIT %d OFFSET %d", couriersTable, limit, offset)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var courier lavka.Couriers
		err := rows.Scan(&courier.Id, &courier.Type, pq.Array(&courier.Region), pq.Array(&courier.Working_Hours))
		if err != nil {
			return nil, err
		}
		couriersList = append(couriersList, courier)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return couriersList, nil
}

func (r *CouriersListPostgres) GetById(courierId int) (lavka.Couriers, error) {
	var cList lavka.Couriers
	query := fmt.Sprintf("SELECT id, type, region, working_hours  FROM %s WHERE id = %d", couriersTable, courierId)

	row, err := r.db.Query(query)
	if err != nil {
		return lavka.Couriers{}, err
	}
	defer row.Close()
	for row.Next() {
		var courier lavka.Couriers
		err := row.Scan(&courier.Id, &courier.Type, pq.Array(&courier.Region), pq.Array(&courier.Working_Hours))
		if err != nil {
			return lavka.Couriers{}, err
		}
		cList = courier
	}
	if err := row.Err(); err != nil {
		return lavka.Couriers{}, err
	}
	//err := r.db.Get(&cList, query, courierId)
	return cList, err
}

func (r *CouriersListPostgres) GetCouriersSalaryRating(courierId int, start_date string, end_date string) (lavka.Couriers, int, error) {

	var count_completed_orders int
	courierRE, err := r.GetById(courierId)
	if err != nil {
		return lavka.Couriers{}, 0, err
	}

	query := fmt.Sprintf("select count(id) from orders where complete_time is not null and id = any(select unnest(orders_id) from completeOrderdate where couriers_id =%d and deliveryDate between $1 and $2)", courierId)

	err = r.db.QueryRow(query, start_date, end_date).Scan(&count_completed_orders)
	if err != nil {
		return lavka.Couriers{}, 0, err
	}

	if count_completed_orders == 0 {
		return lavka.Couriers{}, 0, nil
	}
	/*var courierType string
	queryFortype := fmt.Sprintf("select type from %s where id = %d", couriersTable, courierId)
	err2 := r.db.QueryRow(queryFortype).Scan(&courierType)
	if err2 != nil {
		return 0,  0, err
	}*/

	return courierRE, count_completed_orders, err
}
