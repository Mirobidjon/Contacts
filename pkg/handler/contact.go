package handler

import (
	"net/http"
	"strconv"

	"github.com/Mirobidjon/contact-list"
	"github.com/gin-gonic/gin"
)

type allContact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Create Contact godoc
// @ID create-contact
// @Security ApiKeyAuth
// @Router /api/contact [POST]
// @Summary create contact
// @Description Create Contact
// @Tags contact
// @Accept json
// @Produce json
// @Param profession body contact.DefaultContact true "profession"
// @Success 200 {object} contact.DefaultContact "desc"
// @Failure 500 {object} errorMessage "desc"
func (h *Handler) createContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input contact.DefaultContact
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Name == "" || input.Phone == "" {
		newErrorResponce(c, http.StatusBadRequest, "invalid name or phone contact")
		return
	}

	id, err := h.service.Contacts.Create(input.Name, input.Phone, userID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// Get Contact godoc
// @ID get-contact
// @Security ApiKeyAuth
// @Router /api/contact/{id} [get]
// @Summary get contact by id
// @Description Get Contact by id
// @Tags contact
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} contact.DefaultContact "desc"
// @Failure 400 {object} errorMessage "desc"
// @Failure 500 {object} errorMessage "desc"
func (h *Handler) getContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	mycontact, err := h.service.Contacts.GetByID(userID, id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	output := contact.DefaultContact{
		Name:  mycontact.Name,
		Phone: mycontact.Phone,
	}
	c.JSON(http.StatusOK, output)
}

// Get All Contacts godoc
// @ID get-all-contacts
// @Security ApiKeyAuth
// @Router /api/contact [get]
// @Summary get all contacts
// @Description Get All Contacts
// @Tags contact
// @Accept json
// @Produce json
// @Success 200 {object} []allContact "desc"
// @Failure 400 {object} errorMessage "desc"
// @Failure 500 {object} errorMessage "desc"
func (h *Handler) getAllContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	contacts, err := h.service.Contacts.GetAll(userID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	var output []allContact
	for _, value := range contacts {
		output = append(output, allContact{
			value.ID,
			value.Name,
			value.Phone,
		})
	}

	c.JSON(http.StatusOK, output)
}

// Update Contact godoc
// @ID update-contact
// @Security ApiKeyAuth
// @Router /api/contact/{id} [PUT]
// @Summary update contact
// @Description Update Contact
// @Tags contact
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Param input body contact.DefaultContact true "input"
// @Success 200 {object} contact.ID "desc"
// @Failure 400 {object} errorMessage "desc"
// @Failure 500 {object} errorMessage "desc"
func (h *Handler) updateContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input contact.DefaultContact
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Contacts.Update(userID, id, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// Delete Contact godoc
// @ID delete-contact
// @Security ApiKeyAuth
// @Router /api/contact/{id} [DELETE]
// @Summary delete contact
// @Description Delete Contact
// @Tags contact
// @Accept json
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} contact.ID "desc"
// @Failure 400 {object} errorMessage "desc"
// @Failure 500 {object} errorMessage "desc"
func (h *Handler) deleteContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.Contacts.Delete(userID, id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
