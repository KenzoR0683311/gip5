package handlers

import (
    "net/http"
    "testProject/models"
    "html/template" 
    "testProject/services"
    "time"
    "testProject/utils"
    "testProject/db"
    "fmt"
    "github.com/google/uuid"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
  tmpl := template.New("index.html")
  tmpl.Funcs(utils.TemplateHelpers())
  _, err := tmpl.ParseFiles("web/index.html")
  if err != nil {
    panic(err)
  }

  videos := services.GetVideos(0, 10)
  videoMap := map[string]interface{}{
    "Videos": videos,
    "Page": 0,
    "PageSize": 10,
  }

  tmpl.Execute(w, videoMap)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
  // Parse form data
  name := r.FormValue("name")
  firstName := r.FormValue("firstName")
  email := r.FormValue("email")
  password := r.FormValue("password")
  
  // Register user
  err := services.RegisterUser(name, firstName, email, password)
  if err != nil {
    // Handle error (e.g., send error response)
    return
  }

  http.Redirect(w, r, "/login/", http.StatusSeeOther)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("email")
  password := r.FormValue("password")
  
  user, err := services.LoginUser(email, password)
  if err != nil {
    htmlStr := fmt.Sprintf("<p class='text-red-600'>Email or password is incorrect.</p>")
    tmpl, _ := template.New("q").Parse(htmlStr)
    tmpl.Execute(w, nil)
    return
  }

  sessionToken := uuid.New().String()
	expiresAt := time.Now().Add(12000 * time.Second)
  session := models.Session{
        UserId:   user.Id,
        UserRole: user.Role,
  }
  db.SaveSession(sessionToken, session, expiresAt)
  
  /*
	utils.Sessions[sessionToken] = models.Session{
		User: user,
		Expiry:   expiresAt,
	}*/
  
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
    Path: "/",
	})
  

  w.Header().Set("HX-Redirect", "/")
}


func LogoutUser(w http.ResponseWriter, r *http.Request) {
  sessionCookie, err := r.Cookie("session_token")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  db.DeleteSession(sessionCookie.Value)

  expiredCookie := &http.Cookie{
    Name:    "session_token",
    Value:   "",
    Expires: time.Now().AddDate(0, 0, -1), // Set to expire in the past
    MaxAge:  -1,
    Path:    "/",
  }
  
  http.SetCookie(w, expiredCookie)
  http.Redirect(w, r, "/login", http.StatusSeeOther)
  return
}


func RegisterPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/register.html"))
    tmpl.Execute(w, nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/login.html"))
    tmpl.Execute(w, nil)
}


func AdminPage(w http.ResponseWriter, r *http.Request) {
  tmpl := template.New("admin.html")
  tmpl.Funcs(utils.TemplateHelpers())
  _, err := tmpl.ParseFiles("web/admin.html")
  if err != nil {
    panic(err)
  }

  videos := services.GetVideos(0, 10)

  videoMap := map[string]interface{}{
    "Videos": videos,
    "Page": 0,
    "PageSize": 10,
  }

  tmpl.Execute(w, videoMap)
}

