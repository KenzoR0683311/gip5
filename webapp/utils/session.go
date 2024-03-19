package utils

import (
  "os"
  "io"
  "fmt"
  "path/filepath"
  "html/template" 
)

func TemplateHelpers() map[string]interface{} {
  funcs := template.FuncMap{
    "Add": func(a, b int) int {
      return a + b
    },
    "Sub": func(a, b int) int {
      return a - b
    },
  }

  return funcs
}

func Upload(file io.Reader, filename string) {
  // Create a file in the specified directory
  f, err := os.OpenFile(filepath.Join("assets", "videos", filename), os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    //http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  defer f.Close()

  // Copy the file to the destination
  _, err = io.Copy(f, file)
  if err != nil {
    //http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func Download(fileName string) {
  // Get the user's home directory
  homeDir, err := os.UserHomeDir()
  if err != nil {
    //return err
  }

  // Construct the path to the Downloads folder
  downloadsDir := filepath.Join(homeDir, "Downloads")

  // Construct the file paths
  srcFilePath := fmt.Sprintf("assets/videos/%s", fileName)
  destFilePath := filepath.Join(downloadsDir, fileName)

  // Open the source file for reading
  srcFile, err := os.Open(srcFilePath)
  if err != nil {}
  
  defer srcFile.Close()

  // Create a new file in the Downloads folder
  destFile, err := os.Create(destFilePath)
  if err != nil {
            //return err
  }
  
  defer destFile.Close()

  // Copy the content from the source file to the destination file
  _, err = io.Copy(destFile, srcFile)
  if err != nil {
            //return err
  }

}
