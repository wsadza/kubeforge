// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

package yaml

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
	"kubeforge/internal/ops/file"
)

func Unmarshal(target string, data interface{}, expandData bool) error {
  fileData := []byte(target)
  info, _ := file.Metadata(target)
  if info != nil {
    fileReaded, err := file.Read(target)
    if err != nil {
      return err
    }
    fileData, _ = fileReaded.([]byte)
  }
  if expandData {
    fileData = []byte(os.ExpandEnv(string(fileData)))
  }
  return yaml.Unmarshal(fileData, data)
}

func Marshal(data interface{}, yamlData *interface{}) (error) {
  var err  error
  *yamlData, err = yaml.Marshal(data)
  
  if err != nil {
    return fmt.Errorf("err: %v\n", err)
  }
  return nil 
}
