package main
import "fmt"
import "unsafe"

const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

const (
    aa = iota
    bb = iota
    cc = iota
)

const (
    aaa = iota
    bbb
    ccc
)

const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main(){
	const LENGH int  =10
	const WIDTH int =5
	var area int
	//const a,b,c=1,false,"str"//多重赋值
	area = LENGH*WIDTH
	fmt.Printf("面积为：%d",area)
	println()
	//println(a,b,c)
	//println(a, b, c)

	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
)

fmt.Println(a,b,c,d,e,f,g,h,i)

fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)

}