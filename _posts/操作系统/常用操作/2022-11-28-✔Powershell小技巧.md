# 操作环境变量

https://sysin.org/blog/windows-env/
cmd中使用`%CD%`访问环境变量，powershell使用`$env:windir`
临时设置环境变量`set 变量名=``

## 常用环境变量

```sh
%ALLUSERSPROFILE% 局部 返回所有“用户配置文件”的位置。  
%APPDATA% 局部 返回默认情况下应用程序存储数据的位置。  
%CD% 局部 返回当前目录字符串。  
%CMDCMDLINE% 局部 返回用来启动当前的 Cmd.exe 的准确命令行。  
%CMDEXTVERSION% 系统 返回当前的“命令处理程序扩展”的版本号。  
%COMPUTERNAME% 系统 返回计算机的名称。  
%COMSPEC% 系统 返回命令行解释器可执行程序的准确路径。  
%DATE% 系统 返回当前日期。使用与 date /t 命令相同的格式。由 Cmd.exe 生成。有关 date 命令的详细信息，请参阅 Date。  
%ERRORLEVEL% 系统 返回最近使用过的命令的错误代码。通常用非零值表示错误。  
%HOMEDRIVE% 系统 返回连接到用户主目录的本地工作站驱动器号。基于主目录值的设置。用户主目录是在“本地用户和组”中指定的。  
%HOMEPATH% 系统 返回用户主目录的完整路径。基于主目录值的设置。用户主目录是在“本地用户和组”中指定的。  
%HOMESHARE% 系统 返回用户的共享主目录的网络路径。基于主目录值的设置。用户主目录是在“本地用户和组”中指定的。  
%LOGONSEVER% 局部 返回验证当前登录会话的域控制器的名称。  
%NUMBER_OF_PROCESSORS% 系统 指定安装在计算机上的处理器的数目。  
%OS% 系统 返回操作系统的名称。Windows 2000 将操作系统显示为 Windows_NT。  
%PATH% 系统 指定可执行文件的搜索路径 (sysin)。  
%PATHEXT% 系统 返回操作系统认为可执行的文件扩展名的列表。  
%PROCESSOR_ARCHITECTURE% 系统 返回处理器的芯片体系结构。值: x86，IA64。  
%PROCESSOR_IDENTFIER% 系统 返回处理器说明。  
%PROCESSOR_LEVEL% 系统 返回计算机上安装的处理器的型号。  
%PROCESSOR_REVISION% 系统 返回处理器修订号的系统变量。  
%PROMPT% 局部 返回当前解释程序的命令提示符设置。由 Cmd.exe 生成。  
%RANDOM% 系统 返回 0 到 32767 之间的任意十进制数字。由 Cmd.exe 生成。  
%SYSTEMDRIVE% 系统 返回包含 Windows XP 根目录（即系统根目录）的驱动器。  
%SYSTEMROOT% 系统 返回 Windows XP 根目录的位置。  
%TEMP% and %TMP% 系统和用户 返回对当前登录用户可用的应用程序所使用的默认临时目录。有些应用程序需要 TEMP，而其它应用程序则需要 TMP。  
%TIME% 系统 返回当前时间。使用与 time /t 命令相同的格式。由 Cmd.exe 生成。有关 time 命令的详细信息，请参阅 Time。  
%USERDOMAIN% 局部 返回包含用户帐户的域的名称。  
%USERNAME% 局部 返回当前登录的用户的名称。  
%UserProfile% 局部 返回当前用户的配置文件的位置。  
%WINDIR% 系统 返回操作系统目录的位置。
```

# Powershell在线教程

https://www.pstips.net/powershell-online-tutorials#Powershell%E7%AE%A1%E9%81%93

# Powershell 常用命令

https://juejin.cn/post/6854573219232350221

##### 设置变量

```bash
PS D:\> $a = 5
PS D:\> $a * 4
20
PS D:\> "sdf".substring(2)
f
PS D:\> $str = "powershell"
PS D:\> $str.substring($str.indexOf("s"))
shell
```

powershell支持.net的方法：截取字符串、if语句判断等这些都没有问题。可以根据自己需要来去使用。

常用的变量知识大概这么多，其它更多变量知识可以参考[Powershell 定义变量](https://link.juejin.cn/?target=https%3A%2F%2Fwww.pstips.net%2Fpowershell-define-variable.html "https://www.pstips.net/powershell-define-variable.html")。

#### PowerShell常用命令

##### Start-Process，别名：start + 路径， 打开当前文件夹

```bash
# 把路径存储为变量，打开路径地址的文件
PS D:\> $variable = "D:\"
PS D:\> start $variable

# 以管理员身份启动Powershell
PS D:\> start -FilePath "powershell" -Verb RunAs
复制代码
```

##### Get-ChildItem，别名：ls、dir, 列出文件夹下所有文件

```lua
PS D:\> ls


    目录: D:\


Mode                LastWriteTime         Length Name
----                -------------         ------ ----
d-----        2020/1/20     11:58                Downloads
...
复制代码
```

##### 🎉Get-History，别名：history、h，列出之前的操作命令

```sql
PS D:\> history

  Id CommandLine
  -- -----------
   1 Start D:
   2 Start D
   3 ${"I"like $}="test"
   ...
复制代码
```

##### 🎉Get-Process， 别名：ps，查找进程, 可以通过进程名称或者进程ID来获取特定进程

```scss
 ➜  ~  ps wechat

Handles  NPM(K)    PM(K)      WS(K)     CPU(s)     Id  SI ProcessName
-------  ------    -----      -----     ------     --  -- -----------
   1255      81   153952     183020     220.53   5832   2 WeChat
复制代码
```

##### new-item，别名：ni <filename.txt>, 创建一个新的文本文件

```diff
PS D:\> ni test.txt
    目录: D:\
Mode                LastWriteTime         Length Name
----                -------------         ------ ----
-a----        2020/7/25     16:29              0 test.txt
复制代码
```

##### remove-item，别名： rm、del, 删除或删除文件

```bash
PS D:\> rm test.txt
复制代码
```

##### copy-item，别名：cp,\copy, 复制文件

```bash
PS D:\> cp test.txt test1.txt
复制代码
```

##### get-location，别名：pwd, 当前目录位置

```markdown
PS D:\> pwd
Path
----
D:\
复制代码
```

##### 🎉GET-HELP, 缩写help，查看命令的帮助

```arduino
get-help get-process
复制代码
```

##### 其他命令

-   get-date, 别名：date，获取系统当前时间
-   Get-Command，别名：gcm， 查找所有命令，可以通过通配符查找，如：get-command *process
-   write-output, 别名：echo、 write, 把东西输出发送到管道，从那里它可以通过管道传输到另一个cmdlet或者变量
-   write-host 直接输出到控制台
-   get-content，别名：cat, 输出文件内容到控制台
-   Install-Module module-name -Scope CurrentUser，从在线库中找到包并安装包到本地
-   Import-Module module-name，导入包
-   chcp, 修改当前的编码方式，默认 936 (GB2312)，可以通过`chcp 65001`设置为 UTF-8 格式
