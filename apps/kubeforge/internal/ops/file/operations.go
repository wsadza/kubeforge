// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

package file

import (
	"io"
	"io/fs"
	"os"
)

func Open(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err 
	}
	return file, nil
}

func Remove(filePath string) error {
  err := os.RemoveAll(filePath)
  if err != nil {
    return err
  }
  return nil
}

func Create(filePath string) (*os.File, error) {
  file, err := os.Create(filePath)
  if err != nil {
		return nil, err
  }
  return file, nil
}

func Rename(filePathOld, filePathNew string) error {
	err := os.Rename(filePathOld, filePathNew)
	if err != nil {
		return err 
	}
	return nil
}

func Copy(filePathNew, filePathOld string) error {

  source, err := Open(filePathOld)
  if err != nil {
		return err
  }

  target, err := os.Create(filePathNew)
	if err != nil {
		return err
	}
	defer target.Close()

  _, err = io.Copy(target, source)
	if err != nil {
		return err
	}
  return nil
}

func SetMode(filePath string, fileMode fs.FileMode) error {
  return os.Chmod(filePath, fileMode)
}
