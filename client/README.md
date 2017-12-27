# AgendaService: Client

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
`user_test.go`测试结果：
```
=== RUN   TestRegister
Testing Register...
[INFO] 2017/12/28 02:06:25 CreateUser [user]
[negroni] 2017-12-28T02:06:24+08:00 | 201 |      120.922093ms | localhost:8080 | PUT /users
[INFO] 2017/12/28 02:06:25 Register successful.
--- PASS: TestRegister (0.15s)
        user_test.go:18: Test Register OK.
=== RUN   TestLogin
Testing Login...
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      155.156805ms | localhost:8080 | GET /token
[INFO] 2017/12/28 02:06:25 Login successful.
--- PASS: TestLogin (0.16s)
        user_test.go:24: Test Login OK.
=== RUN   TestLogout
Testing Logout...
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      420.441µs | localhost:8080 | GET /token
[INFO] 2017/12/28 02:06:25 Login successful.
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      120.860638ms | localhost:8080 | DELETE /token
[INFO] 2017/12/28 02:06:25 Logout successful.
--- PASS: TestLogout (0.12s)
        user_test.go:31: Test Logout OK.
=== RUN   TestListAllUsers
Testing List All Users...
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      111.516715ms | localhost:8080 | GET /token
[INFO] 2017/12/28 02:06:25 Login successful.
[INFO] 2017/12/28 02:06:25 ListAllUsers By User:[user] with Token:[824a80b12b90df82400b7e08cc1dd7ee]
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      654.737µs | localhost:8080 | GET /users
USERNAME             EMAIL                PHONE
haha                 haha                 haha
user_Rtest1          email_Rtest1@example.com 13399991111
user_Rtest2          email_Rtest2@example.com 13399992222
user_Rtest4          email_Rtest4@example.com 13399994444
user                 mail                 13333333333
--- PASS: TestListAllUsers (0.11s)
        user_test.go:38: Test List All Users OK.
=== RUN   TestFindUser
Testing Find User...
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      436.066µs | localhost:8080 | GET /token
[INFO] 2017/12/28 02:06:25 Login successful.
[INFO] 2017/12/28 02:06:25 Retrieve user [user]
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      312.394µs | localhost:8080 | GET /users/user
USERNAME             EMAIL                PHONE
user                 mail                 13333333333
--- PASS: TestFindUser (0.00s)
        user_test.go:45: Test Find User OK.
=== RUN   TestRemoveUser
Testing Remove User...
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      293.871µs | localhost:8080 | GET /token
[INFO] 2017/12/28 02:06:25 Login successful.
[negroni] 2017-12-28T02:06:25+08:00 | 200 |      284.892655ms | localhost:8080 | DELETE /users/user
[INFO] 2017/12/28 02:06:25 Delete successful.
--- PASS: TestRemoveUser (0.29s)
        user_test.go:52: Test Remove User OK.
PASS
ok      github.com/MarshallW906/Agenda-Service/client/vendor/service    0.834s
```

其他的可选测试命令：
```
go run main.go register --username user --password pass --email user@example.com --phone 13333333333
go run main.go login --username user --password pass
go run main.go login --username user --password wrong
go run main.go listUsers
go run main.go findUser --username user
go run main.go removeUser --username user
go run main.go logout
```