package ArrayList

import (
	"errors"
	"fmt"
)

// List 接口
type List interface {
	Size() int 									//数组大小
	Get(index int)(interface{},error) 			//获取第几个元素
	Set(index int, newval interface{}) error 	//修改数据
	Insert(index int, newval interface{}) error //插入元素
	Append(newval interface{}) error  			//追加
	Clear()  									//清空
	Delete(index int) error 					//删除
	String() string 							//返回字符串

	Iterator() Iterator         //构造初始化接口
}

// ArrayList 数据结构 字符串 整数 实数
type ArrayList struct{
	DataStore[] interface{}  //数组存储
	theSize int 			//数组大小
}

func NewArrayList() *ArrayList {
	list:=new(ArrayList)					//初始化结构体
	list.DataStore=make([]interface{},0,10)	//开辟空间10个
	list.theSize=0
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int)(interface{},error){
	if index<0 || index>list.theSize{
		return nil, errors.New("数组越界")
	}
	return list.DataStore[index],nil

}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index<0 || index>list.theSize{
		return errors.New("数组越界")
	}
	list.DataStore[index]=newval
	return nil
}

func (list *ArrayList) checkisFull()  {
	if list.theSize==cap(list.DataStore){
		//第二个参数指定开辟空间长度，第三个参数指定预留空间
		//当为0时空间cap虽然=list.DataStore*2,但是len=0,copy无法存入数据
		newDataStore :=make([]interface{},list.theSize,list.theSize*2)

		copy(newDataStore,list.DataStore)
		//将添加数据的方式改为append追加也可以
		//for i:=0;i<list.theSize;i++{
		//	newDataStore=append(newDataStore,list.DataStore[i])
		//}
		list.DataStore=newDataStore
	}
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index<0 || index>list.theSize{
		return errors.New("数组越界")
	}
	list.checkisFull()
	//fmt.Println("test:",len(list.DataStore))
	list.DataStore=list.DataStore[:list.theSize+1]  //增加len可获取值
	//fmt.Println("test:",len(list.DataStore))

	//向后移位，移到index为止
	for i:=list.theSize;i>index;i--{
		list.DataStore[i]=list.DataStore[i-1]
	}
	//插入newval
	list.DataStore[index]=newval
	list.theSize++
	return nil
}

func (list *ArrayList) Append(newval interface{}) error {
	list.DataStore=append(list.DataStore,newval)
	list.theSize++
	return nil
}

func (list *ArrayList) Clear(){
	list.DataStore=make([]interface{},0,10)
	list.theSize=0
}

func (list *ArrayList) Delete(index int) error{
	//以index，index+1为分界线 切片DataStore 然后进行组合
	list.DataStore=append(list.DataStore[:index],list.DataStore[index+1:]...)
	list.theSize--
	return nil
}

func (list *ArrayList) String() string{
	return fmt.Sprint(list.DataStore)
}

//func (list *ArrayList) Iterator() Iterator {
//	it :=new(ArrayListIterator.md)
//	it.currentIndex=0
//	it.list=list
//	return it
//}











