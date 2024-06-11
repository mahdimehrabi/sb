package handlers

import (
	"errors"
	"m1-article-service/application/http/requests"
	"m1-article-service/domain/service/address_user"
	"net/http"
)
import "github.com/gin-gonic/gin"

type AddressUser struct {
	addrService address_user.Service
}

func NewAddressUser(addrService address_user.Service) *AddressUser {
	return &AddressUser{addrService: addrService}
}

func (h AddressUser) CreateUser(c *gin.Context) {
	user := &requests.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.addrService.Create(user.Addresses.ToEntities(), user.ToEntity())
	if err != nil {
		if errors.Is(err, address_user.ErrServiceUnavailable) {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Service Unavailable"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{})
	return
}
