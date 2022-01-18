// helloworld.go
package main

import (
"fmt"
"./myMath"
)

func main(){
    fmt.Println("Hello World!")
    fmt.Println(mathClass.Add(1,1))
    fmt.Println(mathClass.Sub(1,1))

    // %d 表示整型数字， $s 表示字符串
    var stockcode = 123
    var enddate = "2022-02-17"
    var url = "Code=%d&enddate=%s"
    var target_url = fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)

    var a string  = "Runoob"
    fmt.Println(a)
    //var b,c int = 1,2
    //fmt.Println(b,c)

    //var d int
   // fmt.Println(d)

    var e bool
    fmt.Println(e)

    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)

    var d = true
    fmt.Println(d)

    //var intVal int 
    //intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

    //intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句

    g := "Runoob" // var f string = "Runoob"

    fmt.Println(g)
}