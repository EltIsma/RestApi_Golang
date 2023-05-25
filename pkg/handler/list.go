package handler

import (
	"net/http"
	"strconv"

	lavka "github.com/EltIsma/YandexLavka"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createOrders(c echo.Context) error {
	var input lavka.Orders
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}

	id, err := h.services.OrderList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return c.String(http.StatusOK, "HTTP 200 OK")
}

type getAllOrdersResponse struct {
	Data []lavka.Orders `json:"data"`
}

func (h *Handler) getAllOrders(c echo.Context) error {
	var limit, offset int
	l := c.QueryParam("limit")
	q := c.QueryParam("offset")
	if len(l)>0 {
		limit, _ = strconv.Atoi(l)
	} else {limit =1}

	if len(q)>0 {
		offset, _ = strconv.Atoi(q)
	} else {offset =0}


	orders_lists, err := h.services.OrderList.GetAll(limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}
	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders_lists,
	})
	return c.String(http.StatusOK, "HTTP 200 OK")
}

func (h *Handler) getOrderById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return c.String(http.StatusBadRequest, "invalid id param")
	}
	orders_list, err := h.services.OrderList.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}
	c.JSON(http.StatusOK, orders_list)
	return c.String(http.StatusOK, "HTTP 200 OK")
}

func (h *Handler) ordersComplete(c echo.Context) error {
    var input lavka.OrdersComplete
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}
	id, err := h.services.OrderList.Update(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return c.String(http.StatusOK, "HTTP 200 OK")
}


/*func (h *Handler) getAssign(c echo.Context) error {
	var input lavka.OrdersAssign
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}

	id, err := h.services.OrderList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return c.String(http.StatusInternalServerError, "Not good")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return c.String(http.StatusOK, "HTTP 200 OK")
}*/