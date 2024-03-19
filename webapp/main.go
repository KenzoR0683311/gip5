package main

import (
  "log"
  "net/http"
  "testProject/handlers"
  "testProject/db"
  "testProject/middleware"
)

func main() {
  db.InitializeDB()
  mux := http.NewServeMux()

  mux.HandleFunc("/register/", handlers.RegisterPage)
  mux.HandleFunc("POST /register/", handlers.RegisterUser)
  mux.HandleFunc("/login/", handlers.LoginPage)
  mux.HandleFunc("POST /login/", handlers.LoginUser)
  mux.HandleFunc("/logout/", handlers.LogoutUser)

  mux.HandleFunc("/admin/", middleware.Auth(handlers.AdminPage, []string{"admin"}))
  mux.HandleFunc("/get_video_analytics/{id}/{page}/{pageSize}", middleware.Auth(handlers.GetVideoAnalytics, []string{"admin"}))
  mux.HandleFunc("POST /add-video/", middleware.Auth(handlers.AddVideo, []string{"admin"}))
  mux.HandleFunc("PUT /update-video/", middleware.Auth(handlers.UpdateVideo, []string{"admin"}))
  mux.HandleFunc("/update-page/{id}", middleware.Auth(handlers.UpdateVideoPage, []string{"admin"}))
  mux.HandleFunc("DELETE /remove-video/{id}", middleware.Auth(handlers.RemoveVideo, []string{"admin"}))
  
  mux.HandleFunc("/", middleware.Auth(handlers.HomePage, []string{"student", "admin"}))
  mux.HandleFunc("/get-videos/{page}/{pageSize}", middleware.Auth(handlers.GetVideos, []string{"student", "admin"}))
  mux.HandleFunc("/get-video/{id}", middleware.Auth(handlers.GetVideo, []string{"student", "admin"}))
  mux.HandleFunc("/download/", middleware.Auth(handlers.DownloadVideos, []string{"student", "admin"}))
  mux.HandleFunc("PATCH /watch-time-analytics/", middleware.Auth(handlers.SetlongestWatchTime, []string{"student", "admin"}))
  

  mux.Handle("/admin/assets/", http.StripPrefix("/admin/assets/", http.FileServer(http.Dir("assets"))))
  mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
  log.Fatalln(http.ListenAndServe(":8082", mux)) 
}

