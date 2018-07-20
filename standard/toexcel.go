package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "encoding/json"
  "github.com/dexyk/stringosim"
)

type STANDARD struct {
  CODE string `json:"code"`
  MSG string `json:"msg"`
  TABLEHEAD []map[string]string `json:"table_head"`
  TABLEINFO []map[string]string `json:"table_info"`
}

func GetDictInfo()  {

}

func GetModelInfo()  {

}

func ReadFiles()  {

}

func TODict(dic map[string]map[string]string)  {
  jsIndent,_ := json.MarshalIndent(&dic, "", "\t")
  ioutil.WriteFile("test.json",jsIndent,os.ModeAppend)
}

func BackInfo(exc *STANDARD,mod *STANDARD)(map[string]interface{},map[string]interface{}){
  //  struct 转 map
  excelData, err := json.Marshal(exc)
  if err != nil{fmt.Println(err)}
  modelData, err := json.Marshal(mod)
  if err != nil{fmt.Println(err)}
  resExcel := make(map[string]interface{})
  resModel := make(map[string]interface{})
//  string to map
  if err := json.Unmarshal(excelData,&resExcel);err != nil {fmt.Println(err)}
  if err := json.Unmarshal(modelData,&resModel);err != nil {fmt.Println(err)}
  return resExcel,resModel
}

func Binding() (map[string]interface{},map[string]interface{}){
  pwd, _ := os.Getwd()
  ec := pwd + "\\excel.json"
  mod := pwd + "\\2.1.json"
  dict := pwd + "\\dict.json"
  excel := &STANDARD{} /*绑定的Excel数据*/
  model := &STANDARD{} /*标准模版数据*/
  dictionary := make(map[string]map[string]string) /*字典数据*/
  //excel file
  b, err := ioutil.ReadFile(ec)
  if err != nil {fmt.Println(err.Error())}
  if e := json.Unmarshal(b, &excel); e != nil {fmt.Println(e)}
  //  model file
  fp, err := ioutil.ReadFile(mod)
  if err != nil {fmt.Println(err)}
  if mo := json.Unmarshal(fp, &model); mo != nil {fmt.Println(mo)}
  //  dictionary info
  dicts, err := ioutil.ReadFile(dict)
  if err != nil {fmt.Println(err)}
  if mo := json.Unmarshal(dicts, &dictionary); mo != nil {fmt.Println(mo)}
//  实现快速绑定
  for i, value := range excel.TABLEINFO{
    excelName := value["spotName"]
    //  先查看字典中有没有
    if succ, ok := dictionary[excelName]; ok{
      excel.TABLEINFO[i] = map[string]string{
        "spotName" : excelName,
        "id" : succ["id"],
        "mapper" : succ["mapper"],
        "number" : value["number"],
        "precision" : succ["precision"],
        "standardName" : succ["spotName"],
        "status" : "3",
        "unit" : succ["unit"],
        "rw" : value["rw"],
      }
    }else{
      /*字符串相识度匹配 查找模版*/
      for n, info := range model.TABLEINFO{
        resVal := stringosim.Jaro([]rune(excelName),[]rune(info["spotName"]))*100
        if resVal > 90{ /*匹配率需要写成可配置的*/
          excel.TABLEINFO[i] = map[string]string{
            "spotName" : excelName,
            "id" : info["id"],
            "mapper" : info["mapper"],
            "number" : value["number"],
            "precision" : info["precision"],
            "standardName" : info["spotName"],
            "status" : "2",
            "unit" : info["unit"],
            "rw" : value["rw"],
          }
          model.TABLEINFO[n]["status"] = "1" /*改变其状态 0=未绑定 1=已绑定*/
          dictionary[excelName] = info /*写回到字典*/
        }
      }
    }
  }
  TODict(dictionary) /*写入字典*/
  resE, resM := BackInfo(excel,model) /*return web data*/
  return resE, resM
}

func main()  {
}