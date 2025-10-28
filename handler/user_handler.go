package handler

import (
	"ensweb_crud_demo/db"
	"ensweb_crud_demo/model"
	"github.com/EnsurityTechnologies/ensweb"
	"net/http"
)

func GetUsers(req *ensweb.Request) *ensweb.Result {
	var users []model.User
	db.DB.Find(&users)
	return req.Server.RenderJSON(req, users, http.StatusOK)
}

func AddUser(req *ensweb.Request) *ensweb.Result {
	var user model.User
	if err := req.Server.ParseJSON(req, &user); err != nil {
		return req.Server.RenderJSONError(req, http.StatusBadRequest, "Invalid JSON", err.Error())
	}
	db.DB.Create(&user)
	return req.Server.RenderJSONSuccessResponse(req, "User added successfully!", true)
}

func UpdateUser(req *ensweb.Request) *ensweb.Result {
	var user model.User
	if err := req.Server.ParseJSON(req, &user); err != nil {
		return req.Server.RenderJSONError(req, http.StatusBadRequest, "Invalid JSON", err.Error())
	}

	var existing model.User
	if err := db.DB.First(&existing, "id = ?", user.ID).Error; err != nil {
		return req.Server.RenderJSONError(req, http.StatusNotFound, "User not found", err.Error())
	}

	existing.Name = user.Name
	existing.Email = user.Email
	db.DB.Save(&existing)

	return req.Server.RenderJSONSuccessResponse(req, "User updated successfully!", true)
}

func DeleteUser(req *ensweb.Request) *ensweb.Result {
	id := req.Server.GetQuerry(req, "id")
	if id == "" {
		return req.Server.RenderJSONError(req, http.StatusBadRequest, "Missing user ID", "")
	}
	db.DB.Delete(&model.User{}, "id = ?", id)
	return req.Server.RenderJSONSuccessResponse(req, "User deleted successfully!", true)
}
