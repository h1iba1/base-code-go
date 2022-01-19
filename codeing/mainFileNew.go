package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)
//
// --
// ----

var tmplevel int

// GetALLX 递归的方式遍历文件目录
func GetALLX(path string,files[] string, level int) ([]string,error) {
	//tmplevel =level

	fmt.Println("level",tmplevel)

	levelstr :=""
	if level ==1{
		levelstr ="+"
	}else{
		for ;level>=1;level--{
			levelstr+="|--"
		}
		levelstr+="+"
	}

	read,err:=ioutil.ReadDir(path)  //读取文件夹
	if err != nil {
		return files,errors.New("文件夹不可读取")
	}

	for _,fi:=range read{		//循环每个文件或者文件夹
		if fi.IsDir(){			//判断是否是文件夹
			fulldir :=path+"/"+fi.Name()	//构造新路径
			files=append(files,levelstr+fulldir)		//追加路径
			tmplevel=level+1
			files,_=GetALLX(fulldir,files,level+1)	//文件夹递归处理
		}else{
			fulldir :=path+"/"+fi.Name()	//构造新的路径
			files=append(files,levelstr+fulldir)		//追加路径
		}
	}
	return files,nil
}

func mainx()  {
	path :="/Users/h11ba1/Desktop/go/codeing"
	files:=[]string{}	//数组字符串

	tmplevel=1
	files,_=GetALLX(path,files,1)  //抓取所有文件

	for i:=0;i<len(files);i++{
		fmt.Println(files[i])
	}
}













