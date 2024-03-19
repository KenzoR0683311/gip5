package models


type USER struct {
  Id        int
  Name      string
  FirstName  string
  Email     string
  Password  string
  Role     string
}

type Video struct {
  Id      string
  Title   string
  Content string
}

type Session struct {
  UserId    int
  UserRole  string
}

type WatchTime struct {
  VideoID         int  `json:"videoId"`
	StartTime       float64 `json:"startTime"`
	EndTime         float64 `json:"endTime"`
	Duration        float64 `json:"duration"`
}

type VideoAnalytics struct {
  Id          int
  UserName    string 
  VideoId     int
  Clicked     int
  Downloaded  bool
  WatchTime   int
  StartTime   int
  EndTime     int
}
