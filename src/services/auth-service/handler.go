package authservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lampronmuehleisenvt571-dev/tm-money-exchange-wallet/src/shared/logger"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
			"details": err.Error(),
		})
		return
	}

	resp, err := h.service.Register(req)
	if err != nil {
		logger.Error("registration failed", "error", err.Error())
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("user registered", "user_id", resp.ID)
	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	resp, err := h.service.Login(req)
	if err != nil {
		logger.Error("login failed", "email", req.Email, "error", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	logger.Info("user logged in", "email", req.Email)
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	resp, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if err := h.service.Logout(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "logout failed",
		})
		return
	}

	logger.Info("user logged out", "user_id", userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}
