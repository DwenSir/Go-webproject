package getModel

import (
  "os"
  "standard/pkg/pathJoin"
  "io/ioutil"
  "fmt"
  "encoding/json"
)

//获取版本号
func VersionCfg() string {
  dir,_ := os.Getwd()
  cfgDir := pathJoin.Join(dir,"cfg")
  filePaths := pathJoin.Join(cfgDir,"version.ini")
  b, err := ioutil.ReadFile(filePaths)
  if err != nil {fmt.Println(err)}
  var dict map[string]interface{}
  if err := json.Unmarshal([]byte(b), &dict); err != nil {fmt.Println(err)}
  if str, ok := dict["version"].(string); ok {return str}
  return "None"
}

//获取模板选项信息
func ModelInfo() map[string]interface{} {
  dir,_ := os.Getwd()
  modelFile := pathJoin.Join(dir,"model")
  files,_ := ioutil.ReadDir(modelFile)
  var dict = make(map[string]interface{})
  for _, file := range files {
    if file.IsDir() {
      continue
    } else {
      if file.Name() == "StadardIndexTable.json"{
        StadardIndexTable := pathJoin.Join(modelFile,file.Name()) /*读取设备索引表*/
        b, err := ioutil.ReadFile(StadardIndexTable)
        if err != nil {fmt.Println(err)}
        if err := json.Unmarshal(b, &dict); err != nil {fmt.Println(err)}
        return dict
      }
    }
  }
  return dict
}