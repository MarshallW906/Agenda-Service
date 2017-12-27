# Agenda-Service: Client

## Introdution
这是Agenda-Service的客户端，是一个基于命令行界面的用户相关操作的客户端程序，使用golang实现。

## Usage
```
go build
./Agenda [subcommand] [flags]
```

## Subcommands
- help
  - usage: 列出命令说明
  - args: command string (= "all")
  - notes: 参数默认为all，打印所有命令的说明
- register
  - usage: 用户注册
  - args: username string, password string, email string, phone string
  - notes: None
- login
  - usage: 用户登录
  - args: username string, password string
  - notes: 若已有登录用户，则先登出，无论后续是否能登录成功
- logout
  - usage: 用户登出
  - args: None
  - notes: 若未登录，则静默
- listUsers
  - usage: 列出所有用户
  - args: None
  - notes: 要求已登录
- removeUser
  - usage: 用户删除
  - args: username string
  - notes: 要求已登录
- findUser
  - usage: 用户查询
  - args: username string
  - notes: 要求已登录


## Test Subcommands
本地数据会话：本地登录数据会话保存在`session.json`文件中，且放在`$homedir/.agenda/`文件夹下（例如，在UNIX系统下的`$HOME/.agenda/`）。

`user_test.go`测试结果：
```powershell
PS D:\college\Junior\Fuwujisuan\gp\src\github.com\ZM-J\Agenda-Service\Client\vendor\service> go test
Testing Register...
Register successful.
Testing Login...
Login successful.
Testing Logout...
Login successful.
Testing List All Users...
Login successful.
USERNAME             EMAIL                PHONE
Chuansao             mail                 13463246324
Testing Find User...
Login successful.
USERNAME             EMAIL                PHONE
Chuansao             mail                 13463246324
Testing Remove User...
Login successful.
Delete successful.
PASS
ok      github.com/ZM-J/Agenda-Service/Client/vendor/service    0.906s
```

其他的可选测试命令：
```powershell
go run main.go register --username SunXiaoChuan --password 258daidai --email 6324@douyu.com --phone 13838383838
go run main.go register --username SunYaLong --password deyunse --email deyunse@douyu.com --phone 13666666666
go run main.go login --username SunXiaoChuan --password 258daidai # right
go run main.go login --username SunXiaoChuan --password daidai258 # wrong
go run main.go listUsers
go run main.go findUser --username SunYaLong
go run main.go removeUser --username SunYaLong
go run main.go logout
```