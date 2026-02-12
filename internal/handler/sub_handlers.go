package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	testApp "githhub.com/VSBrilyakov/test-app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	invalidBodyError         = "invalid request body"
	invalidIdError           = "invalid id parameter"
	invalidUserIdError       = "invalid user_id parameter"
	invalidServiceNameError  = "invalid service_name parameter"
	invalidDateFromError     = "invalid date_from parameter"
	invalidDateToError       = "invalid date_to parameter"
	dateFromAfterDateToError = "date_from cannot be after date_to"
)

func (h *Handler) createSubscribe(c *gin.Context) {
	logrus.Debug(fmt.Sprintf("incoming: %s", c.Request.URL.String()))

	var input testApp.Subscription
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidBodyError)
		return
	}

	id, err := h.services.CreateSubscription(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	newSuccessResponse(c, gin.H{"id": id})
}

func (h *Handler) getSubscribe(c *gin.Context) {
	logrus.Debug(fmt.Sprintf("incoming: %s", c.Request.URL.String()))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidIdError)
		return
	}

	sub, err := h.services.GetSubscription(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, sub.GetJSON())
}

func (h *Handler) updateSubscribe(c *gin.Context) {
	logrus.Debug(fmt.Sprintf("incoming: %s", c.Request.URL.String()))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidIdError)
		return
	}

	var input testApp.UpdSubscription
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidBodyError)
		return
	}

	if err := h.services.UpdateSubscription(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, gin.H{"success": "ok"})
}

func (h *Handler) deleteSubscribe(c *gin.Context) {
	logrus.Debug(fmt.Sprintf("incoming: %s", c.Request.URL.String()))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidIdError)
		return
	}

	if err := h.services.DeleteSubscription(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, gin.H{"success": "ok"})
}

func (h *Handler) getAllSubscribes(c *gin.Context) {
	logrus.Debug(fmt.Sprintf("incoming: %s", c.Request.URL.String()))

	subs, err := h.services.GetAllSubscriptions()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, subs)
}

func (h *Handler) getSubsSum(c *gin.Context) {
	logrus.Debug(fmt.Sprintf("incoming: %s", c.Request.URL.String()))

	userId := c.Query("user_id")
	if userId == "" {
		newErrorResponse(c, http.StatusBadRequest, invalidUserIdError)
		return
	}

	serviceName := c.Query("service_name")
	if serviceName == "" {
		newErrorResponse(c, http.StatusBadRequest, invalidServiceNameError)
		return
	}

	dateFrom, err := time.Parse("01-2006", c.Query("date_from"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidDateFromError)
		return
	}

	dateTo, err := time.Parse("01-2006", c.Query("date_to"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidDateToError)
		return
	}

	if dateFrom.After(dateTo) {
		newErrorResponse(c, http.StatusBadRequest, dateFromAfterDateToError)
		return
	}

	sum, err := h.services.GetSubsSumByUserID(userId, serviceName, dateFrom, dateTo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, gin.H{"sum": sum})
}
