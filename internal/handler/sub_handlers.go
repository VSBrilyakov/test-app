package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	testApp "github.com/VSBrilyakov/test-app"
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

// @Summary CreateSubscription
// @Tags subscriptions
// @Description Add a subscription information into database
// @ID create-subscription
// @Accept  json
// @Produce  json
// @Param input body test_app.Subscription true "Subscription main info"
// @Success 200 {object} successNewSubMsg
// @Failure 500 {object} errorMsg
// @Router /api/v1/subscribe [post]
func (h *Handler) createSubscription(c *gin.Context) {
	logrus.Debug(fmt.Sprintf("incoming: %s", c.Request.URL.String()))

	var input testApp.Subscription
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidBodyError)
		return
	}

	id, err := h.services.CreateSubscription(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, successNewSubMsg{Id: id})
}

// @Summary GetSubscription
// @Tags subscriptions
// @Description Get a subscription information from database by subscription id
// @ID get-subscription
// @Produce  json
// @Param id path int true "Subscription ID"
// @Success 200 {object} test_app.SubscriptionJSON
// @Failure 400 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Router /api/v1/subscribe/{id} [get]
func (h *Handler) getSubscription(c *gin.Context) {
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

// @Summary UpdSubscription
// @Tags subscriptions
// @Description Update subscription information
// @ID upd-subscription
// @Accept  json
// @Param input body test_app.UpdSubscription true "Subscription updated info"
// @Param id path int true "Subscription ID"
// @Success 200 {object} emptyResponse
// @Failure 400 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Router /api/v1/subscribe/{id} [put]
func (h *Handler) updateSubscription(c *gin.Context) {
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

	newSuccessResponse(c, emptyResponse{})
}

// @Summary DeleteSubscription
// @Tags subscriptions
// @Description Remove subscription information from database by subscription id
// @ID del-subscription
// @Param id path int true "Subscription ID"
// @Success 200 {object} emptyResponse
// @Failure 400 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Router /api/v1/subscribe/{id} [delete]
func (h *Handler) deleteSubscription(c *gin.Context) {
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

	newSuccessResponse(c, emptyResponse{})
}

func (h *Handler) getAllSubscriptions(c *gin.Context) {
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
