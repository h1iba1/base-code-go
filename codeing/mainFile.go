package main

import (
	"codeing/Queue"
	"codeing/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

// GetALL 递归的方式遍历文件目录
func GetALL(path string,files[] string) ([]string,error) {
	read,err:=ioutil.ReadDir(path)  //读取文件夹
	if err != nil {
		return files,errors.New("文件夹不可读取")
	}

	for _,fi:=range read{		//循环每个文件或者文件夹
		if fi.IsDir(){			//判断是否是文件夹
			fulldir :=path+"/"+fi.Name()	//构造新路径
			files=append(files,fulldir)		//追加路径
			files,_=GetALL(fulldir,files)	//文件夹递归处理
		}else{
			fulldir :=path+"/"+fi.Name()	//构造新的路径
			files=append(files,fulldir)		//追加路径
		}
	}
	return files,nil
}

func mainRecursion()  {
	path :="/Users/h11ba1/Desktop/go/codeing"
	files:=[]string{}	//数组字符串

	files,_=GetALL(path,files)  //抓取所有文件

	for i:=0;i<len(files);i++{
		fmt.Println(files[i])
	}
}

func mainddd()  {
	path :="/Users/h11ba1/Desktop/go/codeing/test"
	// 数组字符串
	files:=[]string{}
	mystack :=StackArray.NewStack()
	// 路径入栈
	mystack.Push(path)

	for !mystack.IsEmpty(){
		// 路径出栈
		path=mystack.Pop().(string)
		//加入列表
		files=append(files,path)
		//读取文件夹下面所有的路径
		read,err:=ioutil.ReadDir(path)

		if err != nil {fmt.Println("文件读取错误")}

		for _,fi:=range read{
			if fi.IsDir(){
				//构造新的路径
				fulldir:=path+"/"+fi.Name()
				// 是路径就入栈
				mystack.Push(fulldir)
			}else{
				//构造新的路径
				fulldir:=path+"/"+fi.Name()
				files=append(files,fulldir)
			}
		}
	}

	for i:=0;i<len(files);i++{
		fmt.Println(files[i])
	}
}

func main()  {
	path :="/Users/h11ba1/Desktop/go/codeing/test"
	// 数组字符串
	files:=[] string{}

	myqueue :=Queue.NewQueue()

	// 路径入队列
	myqueue.EnQueue(path)

	// 循环直到队列为空才退出
	for ;;{
		//不断从队列取出数据
		path:=myqueue.Dequeue()

		files=append(files,path.(string))
		// 队列为空 退出
		if path==nil{
			break
		}
		// 输出根路径
		//fmt.Println("get",path)
		read ,_:=ioutil.ReadDir(path.(string))

		// 循环路径下 文件夹和文件
		for _,fi:=range read{
			if fi.IsDir(){
				//构造新的路径
				fulldir:=path.(string)+"/"+fi.Name()
				// 将获取的路径也进行输出
				//fmt.Println("dir",fulldir)
				// 文件夹就入栈
				myqueue.EnQueue(fulldir)
			}else{
				//构造新的路径
				fulldir:=path.(string)+"/"+fi.Name()
				// 不是文件夹就添加到files列表
				files=append(files,fulldir)
				//fmt.Println("File",fulldir)
			}
		}
	}

for i:=0;i<len(files);i++{
		fmt.Println(files[i])
	}
}
















