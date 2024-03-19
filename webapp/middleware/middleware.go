package middleware

import (
  "net/http"
  _ "testProject/utils"
  "testProject/db"
)

func Auth(next http.HandlerFunc, roles []string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie("session_token")
    
    if err != nil {
		  if err == http.ErrNoCookie {
        http.Redirect(w, r, "/login/", http.StatusSeeOther)
		    return
      }
    
      http.Redirect(w, r, "/login/", http.StatusSeeOther)
	    return
    }
  
    sessionToken := c.Value
    userSession, _ := db.GetSession(sessionToken) 
    
    userRole := userSession.UserRole // Assuming Role is a field in the Session struct
    found := false

    for _, role := range roles {
      if userRole == role {
        found = true
      }
    }

  if(!found) {
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }
    next.ServeHTTP(w, r)
  }
}
