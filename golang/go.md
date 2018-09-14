## 《GO语言实践》

### 第六章 并发
- 1.5^ Go语言的运行时默认会为每个可用的物理处理器分配一个逻辑处理器
- 并发（concurrency）不是并行（parallelism）。并行是让不同的代码片段同时在不同的物理处理器上执行。并行的关键是同时做很多事情，而并发是指同时管理很多事情，这些事情可能只做了一半就被暂停去做别的事情了。在很多情况下，并发的效果比并行好，因为操作系统和硬件的总资源一般很少，但能支持系统同时做很多事情。这种“使用较少的资源做更多的事情”的哲学，也是指导 Go 语言设计的哲学。
- 当通道关闭后，goroutine 依旧可以从通道接收数据，但是不能再向通道里发送数据。能够从已经关闭的通道接收数据这一点非常重要，因为这允许通道关闭后依旧能取出其中缓冲的全部值，而不会有数据丢失。
- 

### 第七章 并发模式
- select case 针对通道会被阻塞，如果有default则流程变为非阻塞方式。

### 第八章 单元测试
``` golang
server := mockServer()
defer server.Close()
```
- defer这个代码设计非常爽，避免代码分支写很多的server.Close()语句
- 如果包使用"_test"这种方式命名,不与文件夹名称保持一致，测试代码只能访问包里公开的标识符。即便测试代码文件和被测试的代码放在同一个文件夹中，也只能访问公开的标识符。
- 基准测试函数必须以 Benchmark 开头，接受一个指向 testing.B 类型的指针作为唯一参数。


### Think
``` golang
jsonByte,err := json.Marshal(&u)
if(err != nil){		
    json.NewEncoder(rw).Encode(&u)
    fmt.Println("io.Write")
}else{
    rw.Write(jsonByte);
    fmt.Println("custom C#");
}
```
json.NewEncoder(rw).Encode(&u) 和 rw.Write(jsonByte) 作用相同，后者是常用编码思维习惯的代码，前者为go风格代码,主要体现为io.Write接口的设计和运用。

``` golang
r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
fmt.Println(r.Replace("This is <b>HTML</b>!"))
```
函数式风格代码，先定义业务规则，再传入需要处理的数据

## 《Go语言标准库践》
https://books.studygolang.com/The-Golang-Standard-Library-by-Example/


