package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sb30_5/internal/logger"
	repo "sb30_5/internal/repository/database"
	"sb30_5/internal/service"
	"strconv"
	"time"
)

type Api struct {
	list map[string]interface{}
	db   *repo.Storage
}

type Command struct {
	exit chan bool
}

func ServerUp(host *string) {

	db := repo.NewDatabase()
	req := &Api{
		list: nil,
		db:   db,
	}
	com := &Command{exit: make(chan bool)}

	mux := http.NewServeMux()
	mux.HandleFunc("/create", req.createUser)
	mux.HandleFunc("/make_friends", req.makeFriends)
	mux.HandleFunc("/user", req.deleteUser)
	mux.HandleFunc("/friends/", req.getFriends)
	mux.HandleFunc("/", req.updateUser)
	mux.HandleFunc("/get", req.getAllUsers)
	mux.HandleFunc("/getAR", req.getAllUsersAutoRefresh)
	mux.HandleFunc("/exit", com.killServ)
	mux.HandleFunc("/test", req.test)

	logs := logger.InitLogger()
	server := &http.Server{
		Addr:     *host,
		Handler:  mux,
		ErrorLog: log.New(logs.Log2Way, "SERVER: ", 0),
	}
	logs.Log2Way.Info().Msgf("server is running on %s\n", *host)

	go closeServ(server, com.exit, logs)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logs.LogFile.Fatal().Msg(fmt.Sprint(err))
	}
	logs.Log2Way.Info().Msgf("server close on %s\n", *host)
}

func closeServ(s *http.Server, exit chan bool, logs *logger.Logs) {

	<-exit
	if err := s.Shutdown(context.Background()); err != nil {
		logs.Log2Way.Error().Msgf("on shutdown server: %v", err)
	}
}

func (com *Command) killServ(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("shutdown server on %s\n", r.Host)))
	time.Sleep(time.Millisecond * 200)
	com.exit <- true
}

func (req *Api) test(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte("TEST"))
}

func (req *Api) createUser(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "POST", w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	id, name, err := service.CreateUser(req.list, req.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user " + name + " was created with id" + strconv.Itoa(id)))
}

func (req *Api) makeFriends(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "POST", w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	source, target, err := service.MakeFriends(req.list, req.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(source.Name + " and " + target.Name + " are now friends"))
}

func (req *Api) deleteUser(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, http.MethodDelete, w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = service.DeleteUser(req.list, req.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user id" + fmt.Sprint(req.list["target_id"]) +
		" was successfully deleted"))
}

func (req *Api) getFriends(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "GET", w) {
		return
	}
	reqUserID, err := strconv.Atoi(r.URL.Path[9:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("enter an existing user ID"))
		return
	}

	name, list, err := service.GetFriends(reqUserID, req.db)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\t" + name + "`s friends:\n" + list))
}

func (req *Api) updateUser(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "PUT", w) {
		return
	}
	err := readRequest(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	req.list["target_id"], err = strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("enter an existing user ID"))
		return
	}
	err = service.UpdateUser(req.list, req.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user id" + fmt.Sprint(req.list["target_id"]) +
		" was successfully updated"))
}

func (req *Api) getAllUsers(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "GET", w) {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\tAll users:\n" + service.GetAllUsers(req.db)))
}

func (req *Api) getAllUsersAutoRefresh(w http.ResponseWriter, r *http.Request) {

	if notSupportedMethod(r.Method, "GET", w) {
		return
	}
	w.Header().Set("Refresh", "0.05; url=/getAR")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\tAll users:\n" + service.GetAllUsers(req.db)))
}
