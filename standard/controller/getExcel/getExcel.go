package getExcel

import (
  "os"
  "standard/pkg/pathJoin"
  "fmt"
  "github.com/tealeg/xlsx"
  "encoding/json"
)

func GetXlsxInfo(filename string) map[string]interface{}{
  dir, _ := os.Getwd()
  drivers := pathJoin.Join(dir, "drivers")
  excelPath := pathJoin.Join(drivers, filename)
  _xlsx, err := xlsx.OpenFile(excelPath)
  if err != nil {
    var dic = make(map[string]interface{})
    dic["code"]="1"
    dic["msg"]= err
    return dic
  }
  res := Pro(_xlsx)
//  baseDevice := GetDevice(xlsx)
//  if _,ok := baseDevice["modbus"];ok{
//    return ModbusPro(xlsx)
//  }else if _,ok := baseDevice["modbus_tcp"];ok{
//    return ModbusPro(xlsx)
//  }else if _,ok := baseDevice["less"];ok{
//    LessPro(xlsx)
//  }else if _,ok := baseDevice["snmp"];ok{
//    SnmpPro(xlsx)
//  }else if _,ok := baseDevice["santak"];ok{
//    SantakPro(xlsx)
//  }
  return res
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
    f.Cell(10,2).Value:f.Cell(10,2).Value,  /*flag 用来代替if in 语句*/
  }
  return deviceMap
}

type REDATA struct {
  TableInfo []map[string]string `json:"table_info"`
  TableHead []map[string]string `json:"table_head"`
  Code map[string]string `json:"code"`
  Msg map[string]string `json:"msg"`
}


func Pro(xlsx *xlsx.File)map[string]interface{}{
  //  read table
  readInfo := xlsx.Sheets[1]
  redata := &REDATA{}
  for index,value := range readInfo.Rows{
    if index != 0{
      if value.Sheet.Cell(index,1).Value == ""{continue}
      reInfo := map[string]string{
        "number": value.Sheet.Cell(index,0).Value,
        "spotName": value.Sheet.Cell(index,1).Value,
        "standardName": "",
        "status": "",
        "unit": "",
        "rw":"0",
        "mapper": "",
        "precision": "",
        "id": "",
      }
      redata.TableInfo = append(redata.TableInfo,reInfo)
    }
  }
  //  write table
  writeinfo := xlsx.Sheets[2]
  for count,data := range writeinfo.Rows{
    if count != 0 {
      if data.Sheet.Cell(count,1).Value == ""{continue}
      reInfo := map[string]string{
        "number": data.Sheet.Cell(count,0).Value,
        "spotName": data.Sheet.Cell(count,1).Value,
        "standardName": "",
        "status": "0",
        "unit": "",
        "rw":"1",
        "mapper": "",
        "precision": "",
        "id": "",
      }
      redata.TableInfo = append(redata.TableInfo,reInfo)
    }
  }
//  table head
  head := []map[string]string{
    {
      "key":"number",
      "value":"序号",
      "eidtable":"0",
    },{
      "key":"spotName",
      "value":"测点名称",
      "eidtable":"0",
    },{
      "key":"standardName",
      "value":"标准化测点名称",
      "eidtable":"0",
    },{
      "key":"unit",
      "value":"单位",
      "eidtable":"0",
    },{
      "key":"mapper",
      "value":"状态描述",
      "eidtable":"0",
    },
  }
  redata.TableHead = head
  b,err := json.Marshal(redata) /*struct 转字符串*/
  if err != nil {
    fmt.Println(err.Error())
  }
  dic := make(map[string]interface{}) /*字符串转map*/
  if err := json.Unmarshal(b,&dic); err != nil {fmt.Println(err)}
  dic["code"] = "0"
  dic["msg"] = "success"
  return dic
}

func ModbusPro(xlsx *xlsx.File)  {

}

func LessPro(xlsx *xlsx.File)  {

}
func SnmpPro(xlsx *xlsx.File)  {

}

func SantakPro(xlsx *xlsx.File)  {

}