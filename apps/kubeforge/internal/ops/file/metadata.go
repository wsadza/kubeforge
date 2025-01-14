// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

package file

import (
	"io/fs"
	"os"
)

func Metadata(filePath string) (fs.FileInfo, error) {
  file, err := os.Stat(filePath)
  if err != nil {
		return nil, err
  }
  return file, nil
}
