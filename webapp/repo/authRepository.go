package repo

import (
  "testProject/models"
  _ "database/sql"
  "testProject/db"
)

func RegisterUser(name, firstName, email, hashedPassword string) error {
  _, err := db.DbCon.Exec("INSERT INTO user_tbl (name, firstName, email, password, role) VALUES(?, ?, ?, ?, ?)", name, firstName, email, hashedPassword, "student")
    return err
}

func GetUserByEmail(email string) (models.USER, error) {
  var user models.USER
  rows, err := db.DbCon.Query("SELECT * FROM user_tbl WHERE email = ?", email)
  if err != nil {
    return user, err
  }
  
  defer rows.Close()

  // Check if the user exists
  if !rows.Next() {
    return user, err
  }

  if err := rows.Scan(&user.Id, &user.Name, &user.FirstName, &user.Email ,&user.Password, &user.Role); err != nil {
    return user, err
  }

  return user, nil
}
