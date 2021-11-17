# multipipelines

`multipipes` is a simple Go module to create pipelines to process data.

## 流式任务处理模式

Pipline is a simple Go module to create pipelines to process data.

**适合场景：**

数据需经过多个阶段处理，每个阶段耗时不同，应使用不同的并发处理

**两个核心对象：**

    Node：一个任务组（func），
    Pipeline：多个不同任务组成的管道

**使用方法：**

    1. 定义Nodes,即，数据处理过程中需要用到的方法
    2. 定义Pipline
    3. 执行Setup方法，Pipline两端注册两个 indata，outData（可选）,都是Node类型
    4. 执行Start方法

详细使用方法参考`pipline_test.go`


![image](http://github.com/altairlee/awesomeGo/blob/master/images/pipline.jpg?raw=true)




