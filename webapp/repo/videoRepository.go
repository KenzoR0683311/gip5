package repo

import (
  _ "database/sql"
  "testProject/models"
  "testProject/utils"
  "testProject/db"
  "fmt"
  "strings"
)

func GetVideos(page int, pageSize int) ([]models.Video, error) {
    var videos []models.Video

    rows, err := db.DbCon.Query("SELECT * FROM video_tbl LIMIT ? OFFSET ?", pageSize, (page * pageSize))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var video models.Video
        if err := rows.Scan(&video.Id, &video.Title, &video.Content); err != nil {
            return nil, err
        }
        videos = append(videos, video)
    }
    return videos, nil
}


func AddVideo(title, content string) (string, error) {
  var id string 
  _, err := db.DbCon.Exec("INSERT INTO video_tbl (title, content) VALUES (?, ?)", title, content)
  err = db.DbCon.QueryRow("SELECT id FROM video_tbl ORDER BY id DESC LIMIT 1").Scan(&id)
  return id, err
}


func GetVideo(videoID string) (models.Video, error) {
    var video models.Video
    err := db.DbCon.QueryRow("SELECT * FROM video_tbl WHERE id = ?", videoID).Scan(&video.Id, &video.Title, &video.Content)
    
    if err != nil {
        return models.Video{}, err
    }

    return video, nil
}


func UpdateVideo(videoID, title string) error {
    _, err := db.DbCon.Exec("UPDATE video_tbl SET title=? WHERE id=?", title, videoID)
    return err
}

func RemoveVideo(id string) error {
  _, err := db.DbCon.Exec("DELETE FROM analytics_tbl WHERE video_id = ?", id)
  _, err = db.DbCon.Exec("DELETE FROM video_tbl WHERE id = ?", id)
 
  return err
}


func UpdateVideoPage(id string) map[string]interface{} {
  var video models.Video
  row := db.DbCon.QueryRow("SELECT * FROM video_tbl WHERE id = ?", id)
  
  err := row.Scan(&video.Id, &video.Title, &video.Content)
  if err != nil {}
  
  varmap := map[string]interface{}{
    "Id": video.Id,
    "Title": video.Title,
  }

  return varmap
}

func DownloadVideo(checkboxes []string) {
  query := fmt.Sprintf("SELECT content FROM video_tbl WHERE id IN (%s)", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(checkboxes)), ","), "[]"))
  rows, err := db.DbCon.Query(query)
  _ = err

  for rows.Next() {
    var content string
      
    // Scan the content (filename) from the row
    if err := rows.Scan(&content); err != nil {}

    utils.Download(content)
  }

}

