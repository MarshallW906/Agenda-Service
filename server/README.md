# AgendaService: Server

## Introduction

The Server side of AgendaService. Using [negroni](github.com/codegangsta/negroni), [mux](github.com/gorilla/mux), [render](github.com/unrolled/render).

## Test with go test -v

```bash
Agenda-Service git:(cli-mock-test) ✗ go test -v ./server/vendor/model
=== RUN   TestGetToken
=== RUN   TestGetToken/GetToken:_Get_Token_of_user_Rtest1
--- PASS: TestGetToken (0.00s)
    --- PASS: TestGetToken/GetToken:_Get_Token_of_user_Rtest1 (0.00s)
=== RUN   TestDeleteToken
=== RUN   TestDeleteToken/DeleteToken:_WrongToken_ReturnErr
=== RUN   TestDeleteToken/DeleteToken:_Genereated_just_by_GetToken,_OK
--- PASS: TestDeleteToken (0.00s)
    --- PASS: TestDeleteToken/DeleteToken:_WrongToken_ReturnErr (0.00s)
    --- PASS: TestDeleteToken/DeleteToken:_Genereated_just_by_GetToken,_OK (0.00s)
=== RUN   TestCheckTokenValid
=== RUN   TestCheckTokenValid/CheckTokenValid:_Call_GetToken(),_got_OK
=== RUN   TestCheckTokenValid/CheckTokenValid:_WrongToken(),_got_false,_no_err
--- PASS: TestCheckTokenValid (0.00s)
    --- PASS: TestCheckTokenValid/CheckTokenValid:_Call_GetToken(),_got_OK (0.00s)
    --- PASS: TestCheckTokenValid/CheckTokenValid:_WrongToken(),_got_false,_no_err (0.00s)
=== RUN   TestGetUsernameByToken
=== RUN   TestGetUsernameByToken/CheckTokenValid:_Call_GetToken("user_Rtest2"),_got_OK
=== RUN   TestGetUsernameByToken/CheckTokenValid:_WrongToken(),_got_false,_no_err
--- PASS: TestGetUsernameByToken (0.00s)
    --- PASS: TestGetUsernameByToken/CheckTokenValid:_Call_GetToken("user_Rtest2"),_got_OK (0.00s)
    --- PASS: TestGetUsernameByToken/CheckTokenValid:_WrongToken(),_got_false,_no_err (0.00s)
=== RUN   Test_genToken
=== RUN   Test_genToken/genToken:_no_error
--- PASS: Test_genToken (0.00s)
    --- PASS: Test_genToken/genToken:_no_error (0.00s)
=== RUN   TestCreateUser
=== RUN   TestCreateUser/CreateUser:_user_Rtest1
=== RUN   TestCreateUser/CreateUser:_user_Rtest2
=== RUN   TestCreateUser/CreateUser:_user_Rtest3
=== RUN   TestCreateUser/CreateUser:_user_Rtest4
=== RUN   TestCreateUser/CreateUser:_user_Rtest1_Conflict
--- PASS: TestCreateUser (0.00s)
    --- PASS: TestCreateUser/CreateUser:_user_Rtest1 (0.00s)
    --- PASS: TestCreateUser/CreateUser:_user_Rtest2 (0.00s)
    --- PASS: TestCreateUser/CreateUser:_user_Rtest3 (0.00s)
    --- PASS: TestCreateUser/CreateUser:_user_Rtest4 (0.00s)
    --- PASS: TestCreateUser/CreateUser:_user_Rtest1_Conflict (0.00s)
=== RUN   TestDeleteUser
=== RUN   TestDeleteUser/DeleteUser:_Delete_user_Rtest3
=== RUN   TestDeleteUser/DeleteUser:_Delete_user_Rtest3_again
--- PASS: TestDeleteUser (0.00s)
    --- PASS: TestDeleteUser/DeleteUser:_Delete_user_Rtest3 (0.00s)
    --- PASS: TestDeleteUser/DeleteUser:_Delete_user_Rtest3_again (0.00s)
=== RUN   TestRetrieveUser
=== RUN   TestRetrieveUser/RetrieveUser:_Retrieve_user_Rtest2
=== RUN   TestRetrieveUser/RetrieveUser:_Retrieve_a_not-existed_user
--- PASS: TestRetrieveUser (0.00s)
    --- PASS: TestRetrieveUser/RetrieveUser:_Retrieve_user_Rtest2 (0.00s)
    --- PASS: TestRetrieveUser/RetrieveUser:_Retrieve_a_not-existed_user (0.00s)
=== RUN   TestListAllUsers
=== RUN   TestListAllUsers/ListAllUsers:_got_1,2,4
--- PASS: TestListAllUsers (0.00s)
    --- PASS: TestListAllUsers/ListAllUsers:_got_1,2,4 (0.00s)
=== RUN   TestCheckLoginInfo
=== RUN   TestCheckLoginInfo/CheckLoginInfo:_user_Rtest1,_got_OK
=== RUN   TestCheckLoginInfo/CheckLoginInfo:_user_Rtest2,_got_OK
=== RUN   TestCheckLoginInfo/CheckLoginInfo:_user_WrongPassWord,_got_false,_no_err
--- PASS: TestCheckLoginInfo (0.00s)
    --- PASS: TestCheckLoginInfo/CheckLoginInfo:_user_Rtest1,_got_OK (0.00s)
    --- PASS: TestCheckLoginInfo/CheckLoginInfo:_user_Rtest2,_got_OK (0.00s)
    --- PASS: TestCheckLoginInfo/CheckLoginInfo:_user_WrongPassWord,_got_false,_no_err (0.00s)
PASS
ok  	github.com/MarshallW906/Agenda-Service/server/vendor/model	0.042s
```