# File Scan and File Cleanup Command Line Tool

[中文](README_zh.md)

How to run the program, provides 3 ways:
1. Run using source code
2. Install using go install
3. Download executable file directly

## Run Using Source Code
1. Download the source code
2. Prepare test files for experience, test file structure as follows:
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
3. Run the scan program, generate scan.txt file in the project root directory
```bash
go run main.go -r=E:\CODE\GO\TOOLSET\ROOT -t=5
```
Command line parameter parsing, -t=5 files exceeding 5 bytes will be recorded in scan.txt, which will be read during the cleanup task in the next step
```bash
Please enter the prompted command line parameters, parameters without default value (default) are required
  -m string
        Execution method: scan or clean (default "scan")
  -o string
        Scan output file (default "scan.txt")
  -r string
        Scan root path
  -t int
        File size threshold (default 1073741824)
```
Check the generated scan file content as follows:
```txt
[File]:E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt,7
[File]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt,11
[Folder]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1,11
[Folder]:E:\CODE\GO\TOOLSET\ROOT\dir1,18
[File]:E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt,8
[Folder]:E:\CODE\GO\TOOLSET\ROOT\dir3,8
[File]:E:\CODE\GO\TOOLSET\ROOT\root_1.txt,13
[File]:E:\CODE\GO\TOOLSET\ROOT\root_2.txt,11
[Folder]:E:\CODE\GO\TOOLSET\ROOT,51
```

4. Run file cleanup, interactively clean files in the command line
```bash
E:\code\go\toolset\disk>go run main.go -m=clean
Starting file cleanup task, scan result file path: scan.txt
Confirm delete file: E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt? Y or N: y
Delete file: E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt, successful
Confirm delete file: E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt? Y or N: y
Delete file: E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt, successful
Confirm delete file: E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt? Y or N: y
Delete file: E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt, successful
Confirm delete file: E:\CODE\GO\TOOLSET\ROOT\root_1.txt? Y or N: y
Delete file: E:\CODE\GO\TOOLSET\ROOT\root_1.txt, successful
Confirm delete file: E:\CODE\GO\TOOLSET\ROOT\root_2.txt? Y or N: y
Delete file: E:\CODE\GO\TOOLSET\ROOT\root_2.txt, successful
Confirm delete folder: E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1? Y or N: y
Delete folder: E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1, successful
Confirm delete folder: E:\CODE\GO\TOOLSET\ROOT\dir1? Y or N: y
Delete folder: E:\CODE\GO\TOOLSET\ROOT\dir1, successful
Confirm delete folder: E:\CODE\GO\TOOLSET\ROOT\dir3? Y or N: y
Delete folder: E:\CODE\GO\TOOLSET\ROOT\dir3, successful
Confirm delete folder: E:\CODE\GO\TOOLSET\ROOT? Y or N: n
```
Check the cleanup results, view scan.txt content as follows, successfully deleted files will be marked as {deleted}
```txt
{deleted}[File]:E:\CODE\GO\TOOLSET\ROOT\dir1\1.txt,7
{deleted}[File]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1\1-1.txt,11
{deleted}[Folder]:E:\CODE\GO\TOOLSET\ROOT\dir1\dir1-1,11
{deleted}[Folder]:E:\CODE\GO\TOOLSET\ROOT\dir1,18
{deleted}[File]:E:\CODE\GO\TOOLSET\ROOT\dir3\3.txt,8
{deleted}[Folder]:E:\CODE\GO\TOOLSET\ROOT\dir3,8
{deleted}[File]:E:\CODE\GO\TOOLSET\ROOT\root_1.txt,13
{deleted}[File]:E:\CODE\GO\TOOLSET\ROOT\root_2.txt,11
[Folder]:E:\CODE\GO\TOOLSET\ROOT,51
```

## Install Using go install
```bash
go install github.com/shootercheng/disk@latest
```
1. Check if installation is successful, execute disk command
```bash
E:\code\go\toolset\disk>disk
Please enter the prompted command line parameters, parameters without default value (default) are required
  -m string
        Execution method: scan or clean (default "scan")
  -o string
        Scan output file (default "scan.txt")
  -r string
        Scan root path
  -t int
        File size threshold (default 1073741824)
```
2. Run the scan program, -t specifies file threshold size in bytes
```
disk -r=E:\CODE\GO\TOOLSET\ROOT -t=5
```
3. Run file cleanup
```
disk -m=clean
```

## Download Released Executable File to Run
1. Windows amd64
[https://github.com/shootercheng/disk/releases/download/v1.0.0/disk.exe](https://github.com/shootercheng/disk/releases/download/v1.0.0/disk.exe)
2. Linux amd64
[https://github.com/shootercheng/disk/releases/download/v1.0.0/disk](https://github.com/shootercheng/disk/releases/download/v1.0.0/disk)

```bash
$ wget https://github.com/shootercheng/disk/releases/download/v1.0.0/disk
```
```bash
$ chmod +x ./disk
```
```bash
$ ./disk -h
```

# Run Tests
```bash
go test ./tests/...
