package handlers

import (
  "strconv"
  _ "testProject/db"
  "html/template" 
  _ "database/sql"
  "testProject/models"
  "testProject/utils"
  "net/http"
  "encoding/json"
  "testProject/services"
  "testProject/db"

)


func GetVideoAnalytics(w http.ResponseWriter, r *http.Request) {
  funcMap := utils.TemplateHelpers()
  tmpl := template.Must(template.New("analytics.html").Funcs(funcMap).ParseFiles("web/analytics.html"))

  id := r.PathValue("id")
  page, _ := strconv.Atoi(r.PathValue("page"))
  pageSize, _ := strconv.Atoi(r.PathValue("pageSize"))
  
  videoMap := services.GetVideoAnalytics(id, page, pageSize)
  tmpl.Execute(w, videoMap)

}


func SetlongestWatchTime(w http.ResponseWriter, r *http.Request) {
  c, _ := r.Cookie("session_token")
  // We then get the session from our session map

  sess, _ := db.GetSession(c.Value) 
  userId := sess.UserId


  var data models.WatchTime
  err := json.NewDecoder(r.Body).Decode(&data)

  if err != nil {
  }
  
  services.SetlongestWatchTime(data, userId)
}


