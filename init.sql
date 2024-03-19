CREATE DATABASE video_managment_db;

use video_managment_db;

CREATE TABLE IF NOT EXISTS user_tbl (
  id INT AUTO_INCREMENT PRIMARY KEY,
  Name VARCHAR(255),
  FirstName VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  role VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS video_tbl (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255),
  content VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS analytics_tbl (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  video_id INT,
  clicked INT,
  downloaded TINYINT(1),
  watch_time INT,
  start_time INT,
  end_time INT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (video_id) REFERENCES videos(id)
);


