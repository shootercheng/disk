# 文件扫描、文件清除命令行工具

如何运行程序，提供一下3种方式
1. 使用源码运行
2. go install 安装
3. 直接下载可执行文件

## 使用源码运行
1. 下载源码
2. 准备测试文件体验,测试文件结构如下
```txt
E:\CODE\GO\TOOLSET\ROOT
│  root_1.txt
│  root_2.txt
│  
├─dir1
│  │  1.txt
│  │  
│  └─dir1-1
│          1-1.txt
│          
├─dir2
│      2.txt
│      
└─dir3
        3.txt
```
3. 运行扫描程序，在项目根目录下面生成 scan.txt 文件
```bash
go run main.go -r=E:\CODE\GO\TOOLSET\ROOT -t=5
```
命令行参数解析，-t=5 超过5字节的文件会记录到 scan.txt中，下一步清除任务的时候会读取该文件
```bash
请输入提示的命令行参数, 无默认值(default)的为必须输入参数
  -m string
        执行方法:scan或者clean (default "scan")
  -o string
        扫描输出文件 (default "scan.txt")
  -r string
        扫描根路径
  -t int
        文件大小阈值 (default 1073741824)
```
检查生成的扫描文件内容如下
```txt
[文件]:E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt,7
[文件]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt,11
[文件夹]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1,11
[文件夹]:E:\CODE\GO\TOOLSET\ROOT\dir1,18
[文件]:E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt,8
[文件夹]:E:\CODE\GO\TOOLSET\ROOT\dir3,8
[文件]:E:\CODE\GO\TOOLSET\ROOT\root_1.txt,13
[文件]:E:\CODE\GO\TOOLSET\ROOT\root_2.txt,11
[文件夹]:E:\CODE\GO\TOOLSET\ROOT,51
```

4. 运行清除文件, 在命令行中交互式清除文件
```bash
E:\code\go\toolset\disk>go run main.go -m=clean
开始执行文件清理任务,扫描结果文件路径:scan.txt
确认删除文件:E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt?Y或者N:y
删除文件:E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt,成功
确认删除文件:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt?Y或者N:y
删除文件:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt,成功
确认删除文件:E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt?Y或者N:y
删除文件:E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt,成功
确认删除文件:E:\CODE\GO\TOOLSET\ROOT\root_1.txt?Y或者N:y
删除文件:E:\CODE\GO\TOOLSET\ROOT\root_1.txt,成功
确认删除文件:E:\CODE\GO\TOOLSET\ROOT\root_2.txt?Y或者N:y
删除文件:E:\CODE\GO\TOOLSET\ROOT\root_2.txt,成功
确认删除文件夹:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1?Y或者N:y
删除文件夹:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1,成功
确认删除文件夹:E:\CODE\GO\TOOLSET\ROOT\dir1?Y或者N:y
删除文件夹:E:\CODE\GO\TOOLSET\ROOT\dir1,成功
确认删除文件夹:E:\CODE\GO\TOOLSET\ROOT\dir3?Y或者N:y
删除文件夹:E:\CODE\GO\TOOLSET\ROOT\dir3,成功
确认删除文件夹:E:\CODE\GO\TOOLSET\ROOT?Y或者N:n
```
检查清除文件的结果，查看 scan.txt 内容如下, 删除成功的文件会标记为 {deleted}
```txt
{deleted}[文件]:E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt,7
{deleted}[文件]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt,11
{deleted}[文件夹]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1,11
{deleted}[文件夹]:E:\CODE\GO\TOOLSET\ROOT\dir1,18
{deleted}[文件]:E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt,8
{deleted}[文件夹]:E:\CODE\GO\TOOLSET\ROOT\dir3,8
{deleted}[文件]:E:\CODE\GO\TOOLSET\ROOT\root_1.txt,13
{deleted}[文件]:E:\CODE\GO\TOOLSET\ROOT\root_2.txt,11
[文件夹]:E:\CODE\GO\TOOLSET\ROOT,51
```

## go install 安装
```bash
go install github.com/shootercheng/disk@latest
```
1. 检查是否安装成功，执行disk命令
```bash
E:\code\go\toolset\disk>disk
请输入提示的命令行参数, 无默认值(default)的为必须输入参数
  -m string
        执行方法:scan或者clean (default "scan")
  -o string
        扫描输出文件 (default "scan.txt")
  -r string
        扫描根路径
  -t int
        文件大小阈值 (default 1073741824)
```
2. 运行扫描程序, -t 指定单位为字节
```
disk -r=E:\CODE\GO\TOOLSET\ROOT -t=5
```
3. 运行清除文件
```
disk -m=clean
```

## 下载已发布的可执行文件运行
1. windows amd64
[https://github.com/shootercheng/disk/releases/download/v1.0.0/disk.exe](https://github.com/shootercheng/disk/releases/download/v1.0.0/disk.exe)
2. linux amd64 
[https://github.com/shootercheng/disk/releases/download/v1.0.0/disk](https://github.com/shootercheng/disk/releases/download/v1.0.0/disk)