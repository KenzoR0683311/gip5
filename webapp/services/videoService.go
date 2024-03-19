package services

import (
    "testProject/models"
    "testProject/repo"
)

func GetVideos(page int, pageSize int) []models.Video {
    videos, _ := repo.GetVideos(page, pageSize)
    return videos
}

func AddVideo(title, content string) (string, error) {
    id, err := repo.AddVideo(title, content)
    return id, err
}

func GetVideo(videoID string) (models.Video, error) {
    video, err := repo.GetVideo(videoID)
    return video, err
}


func UpdateVideo(videoID, title string) error {
    err := repo.UpdateVideo(videoID, title)
    return err
}


func RemoveVideo(videoID string) error {
    err := repo.RemoveVideo(videoID)
    return err
}

func UpdateVideoPage(id string) map[string]interface{}  {
  return repo.UpdateVideoPage(id)
}

func DownloadVideo(checkboxes []string) {
  repo.DownloadVideo(checkboxes)
}
