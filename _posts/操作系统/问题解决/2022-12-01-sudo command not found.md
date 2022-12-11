https://blog.csdn.net/hello_1995/article/details/109222650
[# Fixing the “Command Not Found” Error When Using Sudo](https://www.baeldung.com/linux/sudo-command-not-found)

在 Linux 系统中，使用 sudo 命令时如果提示 **“command not found”**，我们应该如何解决呢？  
首先执行以下命令看一下 /etc/sudoers.d 文件是否存在，

```bash
find /etc/sudoers.d
```

1、如果返回 **No such file or directory**，就说明你的系统没有安装sudo，下面是安装命令 ：

```bash
apt-get install sudo
```

如果提示 **E: Unable to locate package sudo** 的错误，先执行下面的命令：

```bash
apt-get update
```

再执行 sudo apt-get install 就可以了。

2、如果 /etc/sudoers.d 文件存在则说明系统已经安装了 sudo，只不过没有配置环境。

当你使用 sudo 去执行一个程序时，处于安全的考虑，这个程序将在一个新的、最小化的环境中执行，也就是说，诸如PATH这样的环境变量，在 sudo 命令下已经被重置成默认状态了。所以当一个刚初始化的 PATH 变量中不包含你所要运行的程序所在的目录，用 sudo 去执行，你就会得到 **“command not found”** 的错误提示。

要想改变 PATH 在 sudo 会话中的初始值，使用以下命令打开/etc/sudoers文件，

```bash
vi /etc/sudoers
```

找到"secure_path"一行，当你执行 sudo 命令时，"secure_path"中包含的路径将被当做默认 PATH 变量使用。

按 “i” 键就可以进入插入模式（Insert mode），添加所需要的路径(如 /usr/local/bin）到"secure_path"下，  
Defaults secure_path = /sbin:/bin:/usr/sbin:/usr/bin:/usr/local/bin

编辑完成后，按 “ESC” 键切换为命令行模式，输入 “: wq” 保存并退出。

解决！！！