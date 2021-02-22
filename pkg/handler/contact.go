package handler

import (
	"net/http"
	"strconv"

	"github.com/Mirobidjon/contact-list"
	"github.com/gin-gonic/gin"
)


type allContact struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}


func (h *Handler) createContact(c *gin.Context){
	
	userID, err := getUserID(c)

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
		"id":id,
	})
}

func (h *Handler) getContact(c *gin.Context){
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	mycontact, err := h.service.Contacts.GetByID(userID, id)
	output := contact.DefaultContact{
		Name: mycontact.Name,
		Phone: mycontact.Phone,
	}
	c.JSON(http.StatusOK, output)
}

func (h *Handler) getAllContact(c *gin.Context){
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

func (h *Handler) updateContact(c *gin.Context){
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil{
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
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":"ok",
	})
}

func (h *Handler) deleteContact(c *gin.Context){
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.Contacts.Delete(userID, id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
