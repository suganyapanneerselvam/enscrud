package app

import (
	"ensweb_crud_demo/model"
	"net/http"

	"github.com/EnsurityTechnologies/ensweb"
)

func (a *App) GetUsers(req *ensweb.Request) *ensweb.Result {
	var users []model.User
	a.db.Find(&users)
	return a.RenderJSON(req, users, http.StatusOK)
}

func (a *App) AddUser(req *ensweb.Request) *ensweb.Result {
	var user model.User
	if err := a.ParseJSON(req, &user); err != nil {
		return a.RenderJSONError(req, http.StatusBadRequest, "Invalid JSON", err.Error())
	}
	a.db.Create(&user)
	return a.RenderJSONSuccessResponse(req, "User added successfully!", true)
}

func (a *App) UpdateUser(req *ensweb.Request) *ensweb.Result {
	var user model.User
	if err := a.ParseJSON(req, &user); err != nil {
		return a.RenderJSONError(req, http.StatusBadRequest, "Invalid JSON", err.Error())
	}

	var existing model.User
	if err := a.db.First(&existing, "id = ?", user.ID).Error; err != nil {
		return a.RenderJSONError(req, http.StatusNotFound, "User not found", err.Error())
	}

	existing.Name = user.Name
	existing.Email = user.Email
	a.db.Save(&existing)

	return a.RenderJSONSuccessResponse(req, "User updated successfully!", true)
}

func (a *App) DeleteUser(req *ensweb.Request) *ensweb.Result {
	id := a.GetQuerry(req, "id")
	if id == "" {
		return a.RenderJSONError(req, http.StatusBadRequest, "Missing user ID", "")
	}
	a.db.Delete(&model.User{}, "id = ?", id)
	return a.RenderJSONSuccessResponse(req, "User deleted successfully!", true)
}
