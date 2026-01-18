# 常见问题

## 执行 scan 任务问题记录
|问题|解决方式|
|:--|:--|
|读取文件夹 F:\ISO\Linux 失败:open F:\ISO\Linux: Access is denied.|切换成管理员方式运行|
|Linux系统运行 disk 命令 显示 $ ./disk
-bash: ./disk: Permission denied|chmod +x ./disk|

## 执行 clean 任务问题记录

|问题|解决方式|
|:--|:--|
|删除文件:F:\BaiduNetdiskDownload\2.zip,失败:unlinkat F:\BaiduNetdiskDownload\2.zip: Access is denied.|切换成管理员方式执行|