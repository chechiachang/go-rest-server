package main

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
)

type User struct {
	Id, Name string
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")

	// mock User
	// u.users[id] = User{"1", "name"}
	 usr := u.users[id]

	testMySQL()
	//usr := getUser(id)

	if len(usr.Id) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: User could not be found.")
		return
	}
	response.WriteEntity(usr)
}

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	usr.Id = strconv.Itoa(len(u.users) + 1) // simple id generation
	u.users[usr.Id] = *usr
	response.WriteHeaderAndEntity(http.StatusCreated, usr)
}

func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	u.users[usr.Id] = *usr
	response.WriteEntity(usr)
}

func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	delete(u.users, id)
}