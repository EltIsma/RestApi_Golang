package lavka

import "time"

type Couriers struct {
	Id            int      `json:"courier_id" `
	Type          string   `json:"courier_type"  binding:"required"`
	Region        []int32  `json:"regions" binding:"required"`
	Working_Hours []string `json:"working_hours" binding:"required"`
	Earning       int      `json:"earnings,omitempty"`
	Ratings       int      `json:"rating,omitempty"`
}

const price = 1000

func (c *Couriers) Earnings(count_completed_orders int) int {
	var earnings int
	if c.Type == "FOOT" {
		earnings = 2 * (price * count_completed_orders)
	} else if c.Type == "BIKE" {
		earnings = 3 * (price * count_completed_orders)
	} else {
		earnings = 4 * (price * count_completed_orders)
	}

	return earnings
}

func (c *Couriers) Rating(count_completed_orders int, start_date string, end_date string) int {
	var coeff int
	if c.Type == "FOOT" {
		coeff = 2
	} else if c.Type == "BIKE" {
		coeff = 3
	} else {
		coeff = 4
	}
	startT, err := time.Parse("2006-01-02", start_date)
	if err != nil {
		return 0
	}
	endT, err := time.Parse("2006-01-02", end_date)
	if err != nil {
		return 0
	}
	duration := endT.Sub(startT)
	hours := int(duration.Hours())
	rating := (count_completed_orders / hours) * coeff
	return rating

}


func (c *Couriers) MaxWeight() int {
	var weight int
	if c.Type == "FOOT" {
		weight = 10
	} else if c.Type == "BIKE" {
		weight = 20
	} else {
		weight = 40
	}
	return weight
}

func (c *Couriers) MaxCapacity() int {
	var count int
	if c.Type == "FOOT" {
		count = 2
	} else if c.Type == "BIKE" {
		count = 4
	} else {
		count = 7
	}
	return count
}

func (c *Couriers) MaxCountRegion() int {
	var count int
	if c.Type == "FOOT" {
		count = 1
	} else if c.Type == "BIKE" {
		count = 2
	} else {
		count = 3
	}
	return count
}

/*func (c *Couriers) orderAccep(ordersRegion int) bool {
	if !contains(c.Region, ordersRegion)
}*/