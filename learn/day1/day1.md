# day1

如果使用go去监听一个ip地址的端口，通常使用

```go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

使用ListenAndServe可以监听一个端口，addr为端口号

后面的参数有两个使用方法

## 自主配置Handler方法

Handler是一个接口

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

实现了这个接口的结构体都可以被当做参数传入

```
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (HelloHandler *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprint(w, "hello 70pice")
	} else {
		fmt.Fprintf(w, "hello 70pice!")
	}
}
func main() {
	hello := new(HelloHandler)
	http.ListenAndServe(":9999", hello)
}
```

## 使用方法绑定

```
package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "hello")
	}
}

func hello(w http.ResponseWriter, r http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprint(w, "index")
	}
}
func main() {
	http.ListenAndServe(":9999", nil)
}

```



## 两者的区别和异同

待补充



## type

- 别名

  - type name string。在后续name可以被当做string使用比如

    ```
    type name string
    var qwe name = "ewq"
    ```

- 定义结构体 type Person struct

- 定义结构 type Phone interface

- 定义函数类型 type handler func(str string)，详解一下，这个相当于把函数当做变量来使用。规定了函数的输入和输出，相当于函数就是对象了。



## 关于指针

一直在写java，没想到go当中有对于指针的使用。

*表示这个是一个指针类型，传入的时候需要传入地址，&表示取地址



## make

在golang中，make是Go语言用来分配内存空间的内置函数，默认有三个参数

```
make(Type,len,cap)
```

- Type：数据类型，只能是slice、map、channel这三种数据类型
- len：长度
- cap：预留空间