package handler

import (
	"net/http"
	"strconv"
	"time"

	testApp "githhub.com/VSBrilyakov/test-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createSubscribe(c *gin.Context) {
	var input testApp.Subscription
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateSubscription(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getSubscribe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	sub, err := h.services.GetSubscription(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sub.GetJSON())
}

func (h *Handler) updateSubscribe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input testApp.UpdSubscription
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateSubscription(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "ok"})
}

func (h *Handler) deleteSubscribe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.DeleteSubscription(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "ok"})
}

func (h *Handler) getAllSubscribes(c *gin.Context) {
	subs, err := h.services.GetAllSubscriptions()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, subs)
}

func (h *Handler) getSubsSum(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	serviceName := c.Query("service_name")
	if serviceName == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid service name")
		return
	}

	dateFrom, err := time.Parse("01-2006", c.Query("date_from"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid date_from param")
		return
	}

	dateTo, err := time.Parse("01-2006", c.Query("date_to"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid date_to param")
		return
	}

	if dateFrom.After(dateTo) {
		newErrorResponse(c, http.StatusBadRequest, "invalid date_from/date_to param")
		return
	}

	sum, err := h.services.GetSubsSumByUserID(userId, serviceName, dateFrom, dateTo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
