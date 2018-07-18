package main

import (
  "os"
  "fmt"
  "github.com/tealeg/xlsx"
  "encoding/json"
)

func main() {
  filesname := "O:\\Go\\src\\standard\\drivers\\2_3_2_201_1_1_16.xlsx"
  xlsx, err := xlsx.OpenFile(filesname)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  baseDevice := GetDevice(xlsx)
  if _,ok := baseDevice["modbus"];ok{
    ModbusPro(xlsx)
  }else if _,ok := baseDevice["modbus_tcp"];ok{
    ModbusPro(xlsx)
  }else if _,ok := baseDevice["less"];ok{
    LessPro(xlsx)
  }else if _,ok := baseDevice["snmp"];ok{
    SnmpPro(xlsx)
  }else if _,ok := baseDevice["santak"];ok{
    SantakPro(xlsx)
  }
}

func GetDevice(xlsx *xlsx.File) (map[string]string){
  f := xlsx.Sheets[0] /*获取sheet 设备信息表*/
  deviceMap := map[string]string{
//    "company":f.Cell(1,2).Value, /*厂家*/
//    "model":f.Cell(2,2).Value,  /*设备型号*/
//    "type":f.Cell(3,2).Value, /*分类(ups、空调、电量仪等)*/
//    "board_type":f.Cell(4,2).Value, /*设备7层结构*/
//    "auth_code":f.Cell(5,2).Value,  /*授权码*/
//    "procotol_filename":f.Cell(6,2).Value,  /*协议文档名称*/
//    "procotol_caption":f.Cell(7,2).Value, /*协议内容标题*/
//    "author":f.Cell(8,2).Value, /*创建人*/
//    "create_time":f.Cell(9,2).Value,  /*创建时间*/
//    "procotol_type":f.Cell(10,2).Value, /*协议类型*/
//    "limit_interval":f.Cell(11,2).Value,  /*命令间隔（单位ms）*/
//    "standard_model":f.Cell(12,2).Value,  /*标准化模板*/
//    "command_max_register":f.Cell(13,2).Value,  /*命令最大连续地址数*/
    f.Cell(10,2).Value:f.Cell(13,2).Value,  /*flag 用来代替if in 语句*/
  }
  return deviceMap
}

type REDATA struct {
  TableInfo []map[string]string `json:"table_info"`
  TableHead []map[string]string `json:"table_head"`
  Code map[string]string `json:"code"`
  Msg map[string]string `json:"msg"`
}

func ModbusPro(xlsx *xlsx.File)map[string]interface{}{
//  read table
  readInfo := xlsx.Sheets[1]
  redata := &REDATA{}
  for index,value := range readInfo.Rows{
    reInfo := map[string]string{
      "number": value.Sheet.Cell(index+1,0).Value,
      "spotName": value.Sheet.Cell(index+1,1).Value,
      "standardName": "",
      "status": "",
      "unit": "",
      "mapper": "",
      "rw": "0",
      "precision": "",
      "id": "",
    }
    redata.TableInfo = append(redata.TableInfo,reInfo)
  }
//  write table
  writeinfo := xlsx.Sheets[2]
  for count,data := range writeinfo.Rows{
    reInfo := map[string]string{
      "number": data.Sheet.Cell(count+1,0).Value,
      "spotName": data.Sheet.Cell(count+1,1).Value,
      "standardName": "",
      "status": "",
      "unit": "",
      "mapper": "",
      "rw": "1",
      "precision": "",
      "id": "",
    }
    redata.TableInfo = append(redata.TableInfo,reInfo)
  }
  b,err := json.Marshal(redata)
  if err != nil {
    fmt.Println(err.Error())
  }
  dic := make(map[string]interface{})
  if err := json.Unmarshal(b,&dic); err != nil {fmt.Println(err)}
  return dic
}
func LessPro(xlsx *xlsx.File)  {

}
func SnmpPro(xlsx *xlsx.File)  {

}

func SantakPro(xlsx *xlsx.File)  {

}

