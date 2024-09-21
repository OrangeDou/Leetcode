 package main

 type zipList struct {
	zbyte int	//字节长度
	ztail int	//偏移量
	zlen int	//元素个数
	enty []zipListEnty
	zend int
 }

 type zipListEnty struct {
	previousLen int
	encoding string
	content []byte
 }
 ß
 func main () {

 }