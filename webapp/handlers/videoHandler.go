package handlers

import (
    "strconv"
    "net/http"
    "html/template" 
    _ "testProject/models"
    "testProject/services"
    "testProject/utils"
    "fmt"
    "strings"
    "testProject/db"
)

func GetVideos(w http.ResponseWriter, r *http.Request) {  
  referer := r.Header.Get("Referer")
  parts := strings.Split(referer, "/")
  lastPart := parts[len(parts)-2]
  fileName := "videoList.html"

  if( lastPart == "admin" ) {
    fileName = "adminVideoList.html"
  }

  tmpl := template.New(fileName)
  tmpl.Funcs(utils.TemplateHelpers())

  _, err := tmpl.ParseFiles("web/" + fileName)
  if err != nil {
    panic(err)
  }

  page, _ := strconv.Atoi(r.PathValue("page"))
  pageSize, _ := strconv.Atoi(r.PathValue("pageSize"))
  videos := services.GetVideos(page, pageSize)
  videoMap := map[string]interface{}{
    "Videos": videos,
    "Page": page,
    "PageSize": pageSize,
  }

  tmpl.Execute(w, videoMap)
}


func AddVideo(w http.ResponseWriter, r *http.Request) {
  err := r.ParseMultipartForm(10 << 30) // 10 gB max file size
  if err != nil {
    return
  }

  title := r.FormValue("title")
  file, handler, err := r.FormFile("videoFile")
  if err != nil {
    return
  }
  
  defer file.Close()

  content := handler.Filename
  utils.Upload(file, content)
  id, _ := services.AddVideo(title, content)
  _ = id
  //tmpl := template.Must(template.ParseFiles("web/admin.html"))
  //tmpl.ExecuteTemplate(w, "Video-list-element", models.Video{Id: id, Title: title, Content: content})
  w.Header().Set("HX-Location", "/admin/")
}

func GetVideo(w http.ResponseWriter, r *http.Request) {
    c, _ := r.Cookie("session_token")
    sessionToken := c.Value

    sess, _ := db.GetSession(sessionToken) 
    //sess, _ := utils.Sessions[sessionToken]
    userId := sess.UserId

    id := r.PathValue("id")

    // Get video details from service
    video, err := services.GetVideo(id)
    if err != nil {
        return
    }

    
    services.UpdateClickCounter(userId, id)
    htmlStr := fmt.Sprintf("<video id='vid-player' src='assets/videos/%s' controls onplay='handlePlay(this)' onpause='handlePause(this, %s)' class='w-full rounded-lg shadow-lg'>Your browser does not support the video tag.</video>", video.Content, video.Id)
    tmpl, _ := template.New("j").Parse(htmlStr)
    tmpl.Execute(w, nil)
}

func UpdateVideo(w http.ResponseWriter, r *http.Request) {
  err := r.ParseMultipartForm(10 << 30) // 10 gB max file size
  if err != nil {
    return
  }

  // Retrieve form values
  videoID := r.FormValue("id")
  title := r.FormValue("title")

  // Update video in database
  err = services.UpdateVideo(videoID, title)
  if err != nil {
      // Handle error
      return
  }

  w.Header().Set("HX-Location", "/admin/")
}


func RemoveVideo(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    
    err := services.RemoveVideo(id)
    if err != nil {
      return
    }

    w.Header().Set("HX-Location", "/admin/")
}

func DownloadVideos(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    c, _ := r.Cookie("session_token")
    sessionToken := c.Value

    sess, _ := db.GetSession(sessionToken) 
    userId := sess.UserId
    _= userId
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

    checkboxes := r.Form["checkbox[]"]
    services.DownloadVideo(checkboxes)
    services.UpdateDownloadStatus(checkboxes, userId)
}

func UpdateVideoPage(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  varmap := services.UpdateVideoPage(id)
  tmpl, _ := template.ParseFiles("web/videoForm.html")
  tmpl.Execute(w, varmap)
}

