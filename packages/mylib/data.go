package mylib

type data struct{
	x int 
	y int 
}

func NewData() *data{
	return &data{1, 2}
}

func (d *data) test(){
	println(d.x, d.y)
}

// 类型别名
type DataTmp = data 
func(d *DataTmp) SetY(y int){
	d.y = y
 }

 func(d *DataTmp) Test(){
	d.test()
 }