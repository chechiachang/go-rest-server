package main

import (
	"net/http"
	"fmt"
	"html"
	"log"

	//"github.com/gorilla/mux"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
)

func (u UserResource) Register(container *restful.Container){
	ws := new(restful.WebService)

	ws.Path("/users").
		Doc("Manage Users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Operation("findUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(User{}))

	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
		Doc("update a user").
		Operation("updateUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		ReturnsError(400, "duplicate user-id", nil).
		Reads(User{}))

	ws.Route(ws.POST("").To(u.createUser).
	// docs
		Doc("create a user").
		Operation("createUser").
		Reads(User{})) // from the request

	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
	// docs
		Doc("delete a user").
		Operation("removeUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	container.Add(ws)
}

func main(){
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	//router.HandleFunc("/todos", TodoIndex)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	wsContainer := restful.NewContainer()
	u := UserResource{map[string]User{}}
	u.Register(wsContainer)

	config := swagger.Config{
		WebServices: wsContainer.RegisteredWebServices(),
		WebServicesUrl: "http://localhost:8080",
		ApiPath: "/apidocs.json",

		SwaggerPath: "/apidocs",
		SwaggerFilePath: "",
	}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Todo")
}




