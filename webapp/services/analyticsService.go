package services

import (
    "testProject/models"
    "testProject/repo"
)

func SetlongestWatchTime(data models.WatchTime, userId int) {
  repo.SetlongestWatchTime(data, userId)
}

func GetVideoAnalytics(id string, page int, pageSize int) map[string]interface{} {
  return repo.GetVideoAnalytics(id, page, pageSize)
}

func UpdateClickCounter(userId int, id string) {
  repo.UpdateClickCounter(userId, id)
}

 
func UpdateDownloadStatus(checkboxes []string, userId int) {
  repo.UpdateDownloadStatus(checkboxes, userId)
}
