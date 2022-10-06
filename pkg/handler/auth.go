package handler

import (
	"net/http"

	"github.com/Mirobidjon/contact-list"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @ID sign-up
// @Router /auth/sign-up [POST]
// @Summary Sign up
// @Description this is sign up
// @Tags auth
// @Accept json
// @Produce json
// @Param user body contact.User true "user"
// @Success 200 {object} contact.ID "desc"
// @Failure 500 {object} errorMessage "desc"
func (h *Handler) signUp(c *gin.Context) {
	var user contact.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signinInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @ID login
// @Router /auth/sign-in [POST]
// @Summary Login
// @Description this is login
// @Tags auth
// @Accept json
// @Produce json
// @Param profession body signinInput true "profession"
// @Success 200 {object} contact.Token "desc"
func (h *Handler) signIn(c *gin.Context) {
	var user signinInput

	if err := c.BindJSON(&user); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(user.Username, user.Password)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
