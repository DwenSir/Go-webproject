package introduceModel

import (
  "io/ioutil"
  "os"
  "standard/pkg/pathJoin"
  "fmt"
  "encoding/json"
  "strconv"
)

type StandardTable struct {
  TableInfo []map[string]interface{} `json:"table_info"`
  TableHeader []string  `json:"table_header"`
  ModelInfo []string  `json:"model_info"`
}

type ModelTable struct {
  ModelInfo []map[string]string `json:"table_info"`
  ModelHeader []map[string]string  `json:"table_head"`
}
//返回map数据
func GetModelInfo(filesName string) map[string]interface{}{
  retMap := make(map[string]interface{})
  pwd,_ := os.Getwd()
  modelPath := pathJoin.Join(pwd,"model")
  fileStr := GetIdInfo(modelPath,filesName)
  if fileStr == "nil"{
    retMap["code"] = "1"
    retMap["msg"] = "can't find file"
    return retMap
  }
  jsonPath := pathJoin.Join(modelPath,fileStr+".json")
  i,h := GetData(jsonPath)
  retMap["code"] = "0"
  retMap["msg"] = "get data success"
  retMap["table_head"] = h
  retMap["table_info"] = i
  return retMap
}

//读取索引表信息 用来查找标准化模版表
func GetIdInfo(modelPath string,fileName string) string{
  standardTable := &StandardTable{}
  StandardIndexTable := pathJoin.Join(modelPath,"StadardIndexTable.json")
  b,err := ioutil.ReadFile(StandardIndexTable)
  if err != nil {fmt.Println(err)}
  if err := json.Unmarshal(b, standardTable); err != nil {fmt.Println(err)}
  dict := standardTable.TableInfo
  for i:= 0;i<len(standardTable.TableInfo);i++{
    if MoName, ok := dict[i]["MoName"].(string); ok {
      if ids, ok := dict[i]["ID"]; ok{
        if fileName == MoName{
          str := strconv.FormatFloat(ids.(float64), 'f', 1, 32)
          return str
        }
      }
    }
  }
  return "None"
}

//读取标准化模版信息
func GetData(modelPath string) ([]map[string]string,[]map[string]string){
  tempMap := &ModelTable{}

  f,err := ioutil.ReadFile(modelPath)
  if err != nil{fmt.Println(err)}
  if err := json.Unmarshal(f, tempMap); err != nil {fmt.Println(err)}
  i := tempMap.ModelInfo
  h := tempMap.ModelHeader
  return i,h
}