package repo

import (
  "testProject/models"
  "database/sql"
  "testProject/db"
)

func SetlongestWatchTime(data models.WatchTime, userId int) {
  var watchTime int
  err := db.DbCon.QueryRow("SELECT watch_time FROM analytics_tbl WHERE user_id = ? AND video_id = ?", userId, data.VideoID).Scan(&watchTime)
 
  if(watchTime < int(data.Duration)) {
     _, err = db.DbCon.Exec("UPDATE analytics_tbl SET watch_time = ?, start_time = ?, end_time = ? WHERE user_id = ? AND video_id = ?", data.Duration, data.StartTime, data.EndTime, userId, data.VideoID)
  }

  _ = err
}

func UpdateClickCounter(userId int, id string) {
  var clicked int
  
  err := db.DbCon.QueryRow("SELECT clicked FROM analytics_tbl WHERE user_id = ? AND video_id = ?", userId, id).Scan(&clicked)
  if err == sql.ErrNoRows {
    // Combination not found, insert a new row
    _, err = db.DbCon.Exec("INSERT INTO analytics_tbl (user_id, video_id, clicked, downloaded, watch_time, start_time, end_time) VALUES (?, ?, 1, false, 0, 0, 0)", userId, id)
    if err != nil {
      panic(err.Error())
    }
  } else if err != nil {
    panic(err.Error())
  } else {
    // Combination found, update the clicked column
     _, err = db.DbCon.Exec("UPDATE analytics_tbl SET clicked = clicked + 1 WHERE user_id = ? AND video_id = ?", userId, id)
    if err != nil {
      panic(err.Error())
    }
  }
}

func UpdateDownloadStatus(checkboxes []string, userId int) {
  for _, id := range checkboxes {
    // Check if the combination of user_id and video_id exists in the analytics table
    var downloaded bool
    
    err := db.DbCon.QueryRow("SELECT downloaded FROM analytics_tbl WHERE user_id = ? AND video_id = ?", userId, id).Scan(&downloaded)
    if err == sql.ErrNoRows {
      // Combination not found, insert a new row
      _, err = db.DbCon.Exec("INSERT INTO analytics_tbl (user_id, video_id, clicked, downloaded, watch_time, start_time, end_time) VALUES (?, ?, 1, false, 0, 0, 0)", userId, id)
      if err != nil {
        panic(err.Error())
      }
    } else if err != nil {
      panic(err.Error())
    } else {
      if(downloaded == false) {
        _, err = db.DbCon.Exec("UPDATE analytics_tbl SET downloaded = true WHERE user_id = ? AND video_id = ?", userId, id)
        if err != nil {
          panic(err.Error())
        }
      }
    } 
  }
}

func GetVideoAnalytics(id string, page int, pageSize int) map[string]interface{} {
  results, _ := db.DbCon.Query("SELECT a.id, u.Name, a.video_id, a.clicked, a.downloaded, a.watch_time, a.start_time, a.end_time, (SELECT AVG(clicked) FROM analytics_tbl WHERE video_id = a.video_id) AS avg_clicks, (SELECT SUM(downloaded) FROM analytics_tbl WHERE video_id = a.video_id) AS total_downloads,(SELECT AVG(watch_time) FROM analytics_tbl WHERE video_id = a.video_id) AS avg_watch_time FROM analytics_tbl a INNER JOIN user_tbl u ON a.user_id = u.id WHERE a.video_id = ? LIMIT ? OFFSET ?", id, pageSize, (page * pageSize))
  var analytics []models.VideoAnalytics
  var avgClicks float64
  var totalDownloads int
  var avgWatchTime float64

  // Iterate through the results and populate the videos slice
  for results.Next() {
    var an models.VideoAnalytics
    if err := results.Scan(&an.Id, &an.UserName, &an.VideoId, &an.Clicked, &an.Downloaded, &an.WatchTime, &an.StartTime, &an.EndTime, &avgClicks, &totalDownloads, &avgWatchTime); err != nil {
        panic(err.Error())
      }
      analytics = append(analytics, an)
  }

  // Create a map with the videos slice and execute the template
  //videoMap := map[string][]models.VideoAnalytics{"Analytics": analytics}
  videoMap := map[string]interface{}{
    "Analytics": analytics,
    "Ida": id,
    "Page": page,
    "PageSize": pageSize,
    "AvgClicks": avgClicks,
    "TotalDownloads": totalDownloads,
    "AvgWatchTime": avgWatchTime,
  }

  return videoMap
}
