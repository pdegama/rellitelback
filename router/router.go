package router

import (
	"github.com/gorilla/mux"
	"linka/handlers"
	"net/http"
)

func HTTPRouter(rh *mux.Router) {

	rh.HandleFunc("/", handlers.Repo.MainAuth)
	rh.HandleFunc("/newuser", handlers.Repo.NewUser)
	rh.HandleFunc("/authuser", handlers.Repo.AuthUser)
	rh.HandleFunc("/link", handlers.Repo.AllLink)
	rh.HandleFunc("/link/new", handlers.Repo.NewLink)
	rh.HandleFunc("/link/edit", handlers.Repo.EditLink)
	rh.HandleFunc("/link/info", handlers.Repo.LinkInfo)
	rh.HandleFunc("/link/delete", handlers.Repo.DeleteLink)
	rh.HandleFunc("/page", handlers.Repo.PageInfo)
	rh.HandleFunc("/page/update", handlers.Repo.UpdatePage)
	rh.HandleFunc("/profile", handlers.Repo.UserProfile)
	rh.HandleFunc("/profile/verify", handlers.Repo.VerifyProfile)
	rh.HandleFunc("/profile/update", handlers.Repo.UpdateProfile)
	rh.HandleFunc("/profile/password", handlers.Repo.NewPassword)
	rh.HandleFunc("/profile/photo", handlers.Repo.ProfilePhoto)
	rh.HandleFunc("/verify", handlers.Repo.Verify)
	rh.HandleFunc("/forgot", handlers.Repo.ForgotPass)
	rh.HandleFunc("/reset", handlers.Repo.ResetPass)
	rh.HandleFunc("/tools/c", handlers.Repo.ContantChecker)
	rh.HandleFunc("/logout", handlers.Repo.UserLogOut)
	
	rh.NotFoundHandler = http.HandlerFunc(handlers.Repo.Error404)

}
