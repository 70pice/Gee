# fmt

## Print系列

Print系列会将内容输出到Console中去，这个系列主要由三个函数

- Print：直接输出给定的内容，占位符不会起到作用
- Printf：可以输出给定的占位符，占位符前面需要跟进百分号%
  - %s：字符串占位符
  - %q：字符串占位符，可以把双引号也打出来
  - %p：打印地址，指针占位符
  - %T：打印类型
  - %b：二进制占位符
  - 
  
- Println：同Print，但是后面增加换行符



## Fprint系列

Fprint系列的函数，会将内容输出到一个io.Writer接口类型的变量中

- Fprint：同上
- Fprintf：同上
- Fprintln：同上

```
func main() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出(控制台)写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "70pice is nb"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中(标准输出)写如信息：%s", name)

}
```

## Sprint系列

Sprint系列函数会把传入的数据合并成一个字符串

