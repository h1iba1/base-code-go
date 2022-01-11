目录遍历两种方法：

## 1. 递归
```go
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
```

## 2. 栈
```go
func main()  {
	path :="/Users/h11ba1/Desktop/go/codeing"
	// 数组字符串
	files:=[]string{}
	mystack :=StackArray.NewStack()
	// 入栈
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
```