package handler

import (
	"net/http"
	"strconv"

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

	subscription, err := h.services.GetSubscription(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, subscription.GetJSON())
}

func (h *Handler) updateSubscribe(c *gin.Context) {

}

func (h *Handler) deleteSubscribe(c *gin.Context) {

}

func (h *Handler) getAllSubscribes(c *gin.Context) {

}
