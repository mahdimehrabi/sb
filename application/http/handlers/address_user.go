package handlers

import (
	"errors"
	"m1-article-service/application/http/requests"
	userRepo "m1-article-service/domain/repository/user"
	"m1-article-service/domain/service/address_user"
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"

type AddressUser struct {
	addrService *address_user.Service
}

func NewAddressUser(addrService *address_user.Service) *AddressUser {
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
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h AddressUser) DetailUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.addrService.Detail(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, user)
}
