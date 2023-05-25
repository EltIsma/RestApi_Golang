package handler

import (
	"net/http"
	"strconv"

	lavka "github.com/EltIsma/YandexLavka"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createCouriers(c echo.Context) error {
	var input lavka.Couriers
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}

	id, err := h.services.CouriersList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return c.String(http.StatusOK, "HTTP 200 OK")
}

type getAllCouriersResponse struct {
	Data []lavka.Couriers `json:"data"`
}

func (h *Handler) getAllCouriers(c echo.Context) error {
	var limit, offset int
	l := c.QueryParam("limit")
	q := c.QueryParam("offset")
	if len(l) > 0 {
		limit, _ = strconv.Atoi(l)
	} else {
		limit = 1
	}

	if len(q) > 0 {
		offset, _ = strconv.Atoi(q)
	} else {
		offset = 0
	}

	couriers_lists, err := h.services.CouriersList.GetAll(limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}
	c.JSON(http.StatusOK, getAllCouriersResponse{
		Data: couriers_lists,
	})
	return c.String(http.StatusOK, "HTTP 200 OK")
}

func (h *Handler) getCourierById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return c.String(http.StatusBadRequest, "invalid id param")
	}
	couriers_list, err := h.services.CouriersList.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}
	c.JSON(http.StatusOK, couriers_list)
	return c.String(http.StatusOK, "HTTP 200 OK")
}

func (h *Handler) getCourierRatingSalary(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return c.String(http.StatusBadRequest, "invalid id param")
	}
	start := c.QueryParam("start_date")
	end := c.QueryParam("end_date")
	var courierRating lavka.Couriers

	courierRating, err = h.services.CouriersList.GetCouriersSalaryRating(id, start, end)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}


    
	c.JSON(http.StatusOK, courierRating)
	return c.String(http.StatusOK, "HTTP 200 OK")
}

