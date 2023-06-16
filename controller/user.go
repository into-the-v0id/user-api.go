package controller

import (
	"github.com/gorilla/mux"
	"github.com/into-the-v0id/user-api.go/entity"
	"github.com/into-the-v0id/user-api.go/model"
	"github.com/into-the-v0id/user-api.go/model/payload"
	"github.com/into-the-v0id/user-api.go/repository"
	"github.com/into-the-v0id/user-api.go/service"
	"github.com/into-the-v0id/user-api.go/value_object"
	"github.com/samber/lo"
	"net/http"
	"time"
)

func UserList(writer http.ResponseWriter, request *http.Request) {
	users := repository.User.GetAll()

	publicUsers := lo.Map(users, func(user *entity.User, index int) *model.PublicUser {
		return model.NewPublicUser(user)
	})

	lo.Must0(service.RespondData(writer, publicUsers, request))
}

func UserCreate(writer http.ResponseWriter, request *http.Request) {
	var data payload.CreateUser
	err := service.ParseData(request, &data, writer)
	if err != nil {
		return
	}

	user := entity.NewUser(data.Name, data.Password)
	err = repository.User.Create(user)
	if err != nil {
		http.Error(writer, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	publicUser := model.NewPublicUser(user)
	lo.Must0(service.RespondData(writer, publicUser, request))
}

func UserDetail(writer http.ResponseWriter, request *http.Request) {
	id, err := value_object.NewIdentifierFrom(mux.Vars(request)["id"])
	if err != nil {
		http.Error(writer, "400 Bad Request", http.StatusBadRequest)
		return
	}

	user := repository.User.GetById(id)
	if user == nil {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	}

	publicUser := model.NewPublicUser(user)
	lo.Must0(service.RespondData(writer, publicUser, request))
}

func UserUpdate(writer http.ResponseWriter, request *http.Request) {
	id, err := value_object.NewIdentifierFrom(mux.Vars(request)["id"])
	if err != nil {
		http.Error(writer, "400 Bad Request", http.StatusBadRequest)
		return
	}

	user := repository.User.GetById(id)
	if user == nil {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	}

	var data payload.UpdateUser
	err = service.ParseData(request, &data, writer)
	if err != nil {
		return
	}

	if !data.Id.Equals(user.Id) {
		http.Error(writer, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if !data.DateCreated.Equal(user.DateCreated) || !data.DateUpdated.Equal(user.DateUpdated) {
		http.Error(writer, "409 Conflict", http.StatusConflict)
		return
	}

	if data.Name != user.Name {
		user.Name = data.Name
		user.DateUpdated = time.Now().UTC()
	}

	err = repository.User.Update(user)
	if err != nil {
		http.Error(writer, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	publicUser := model.NewPublicUser(user)
	lo.Must0(service.RespondData(writer, publicUser, request))
}

func UserDelete(writer http.ResponseWriter, request *http.Request) {
	id, err := value_object.NewIdentifierFrom(mux.Vars(request)["id"])
	if err != nil {
		http.Error(writer, "400 Bad Request", http.StatusBadRequest)
		return
	}

	user := repository.User.GetById(id)
	if user == nil {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	}

	err = repository.User.Delete(user)
	if err != nil {
		http.Error(writer, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(204)
}
