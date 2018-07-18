package pathJoin

import "runtime"

func Join(paths string,fileName string) string{
  system := runtime.GOOS
  if system == "windows"{
    paths = paths + "\\" + fileName
  }else if system == "linux"{
    paths = paths +"/" + fileName
  }
  return  paths
}