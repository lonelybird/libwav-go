# 基于golang的wav文件操作库
示例:
```go
/*
    author : WUID
    qq : 3310975439
    email : 3310975439@qq.com
    github : https://github.com/WUID
*/

package main

import (
	"fmt"
	"github.com/wuid/libwav-go"
)

func main() {
	fmt.Println("我是WUID，一个广door学生")

	one := WAV.NewWav()
	two := WAV.NewWav()

	one.Load("1.wav")	//使用之前要加载文件
	two.Load("2.wav")

	one.Splice(two)		//把tow合并在one的后面

	one.Save("3.wav")	//保存为3.wav
}

```
