package api

import (
	"net/http"
	"sb30_5/internal/logger"
	"sb30_5/internal/service"
	"strconv"
)

type Api struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	SourceId int    `json:"source_id"`
	TargetId int    `json:"target_id"`
	NewAge   int    `json:"new age"`
}

func ServerUp(host *string) {

	logs := logger.InitLogger()

	mux := http.NewServeMux()
	req := &Api{}
	mux.HandleFunc("/create", req.CreateUser)
	mux.HandleFunc("/make_friends", req.MakeFriends)
	mux.HandleFunc("/user", req.DeleteUser)
	mux.HandleFunc("/friends/", req.GetFriends)
	mux.HandleFunc("/", req.UpdateUser)
	mux.HandleFunc("/get", req.GetAllUsers)

	server := &http.Server{
		Addr:     *host,
		Handler:  mux,
		ErrorLog: logs.LogErr,
	}

	logs.LogInf.Printf("Server is running on %s", *host)
	logs.Screen.Printf("Server is running on %s\n", *host)
	err := server.ListenAndServe()
	logs.LogErr.Fatal(err)
}

func (req *Api) CreateUser(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "POST", w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	id, name, err := service.CreateUser((*service.Service)(req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User " + name + " was created with id" + strconv.Itoa(id)))
}

func (req *Api) MakeFriends(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "POST", w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	source, target, err := service.MakeFriends((*service.Service)(req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(source.Name + " and " + target.Name + " are now friends"))
}

func (req *Api) DeleteUser(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "DELETE", w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = service.DeleteUser((*service.Service)(req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user id" + strconv.Itoa(req.TargetId) +
		" was successfully deleted"))
}

func (req *Api) GetFriends(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "GET", w) {
		return
	}
	reqUserID, err := strconv.Atoi(r.URL.Path[9:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("enter an existing user ID"))
		return
	}
	name, list, err := service.GetFriends(reqUserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\t" + name + "`s friends:\n" + list))
}

func (req *Api) UpdateUser(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "PUT", w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	req.TargetId, err = strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("enter an existing user ID"))
		return
	}
	err = service.UpdateUser((*service.Service)(req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user id" + strconv.Itoa(req.TargetId) +
		" was successfully updated"))
}

func (req *Api) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "GET", w) {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\tAll users:\n" + service.GetAllUsers()))
}

/*
Уровень транспорта (api) (в нашем случае - HTTP) отвечает за обработку сетевых
запросов - преобразование JSON-ов в модельки, валидацию запроса, формирование
ответа на запрос, установку статус кодов ответа и т.д. Функция уровня транспорта
вызывает функцию уровня сервиса.
*/
