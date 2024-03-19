package services

import (
    "testProject/repo"
    "testProject/models"
    "golang.org/x/crypto/bcrypt"
)

func RegisterUser(name, firstName, email, password string) error {
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Register user in the database
    err = repo.RegisterUser(name, firstName, email, string(hashedPassword))
    return err
}


func LoginUser(email, password string) ( models.USER, error) {
  user, err := repo.GetUserByEmail(email)
  if err != nil {
    return user, err
  }

  // Compare hashed password
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if err != nil {
    return user, err
  }
	  
  return user, nil
}

