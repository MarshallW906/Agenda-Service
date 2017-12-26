package model

import (
	"entity"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		newUser *entity.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "CreateUser: user_Rtest1",
			args: args{&entity.User{
				Username: "user_Rtest1",
				Password: "pass_Rtest1",
				Email:    "email_Rtest1@example.com",
				Phone:    "13399991111",
			}},
			wantErr: false,
		},
		{
			name: "CreateUser: user_Rtest2",
			args: args{&entity.User{
				Username: "user_Rtest2",
				Password: "pass_Rtest2",
				Email:    "email_Rtest2@example.com",
				Phone:    "13399992222",
			}},
			wantErr: false,
		},
		{
			name: "CreateUser: user_Rtest3",
			args: args{&entity.User{
				Username: "user_Rtest3",
				Password: "pass_Rtest3",
				Email:    "email_Rtest3@example.com",
				Phone:    "13399993333",
			}},
			wantErr: false,
		},
		{
			name: "CreateUser: user_Rtest4",
			args: args{&entity.User{
				Username: "user_Rtest4",
				Password: "pass_Rtest4",
				Email:    "email_Rtest4@example.com",
				Phone:    "13399994444",
			}},
			wantErr: false,
		},
		{
			name: "CreateUser: user_Rtest1_Conflict",
			args: args{&entity.User{
				Username: "user_Rtest1",
				Password: "pass_Rtest1",
				Email:    "email_Rtest1@example.com",
				Phone:    "13399991111",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.newUser); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "DeleteUser: Delete user_Rtest3",
			args:    args{"user_Rtest3"},
			wantErr: false,
		},
		{
			name:    "DeleteUser: Delete user_Rtest3 again",
			args:    args{"user_Rtest3"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRetriveUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "RetriveUser: Retrive user_Rtest2",
			args: args{"user_Rtest2"},
			want: &entity.User{
				Username: "user_Rtest2",
				Password: "pass_Rtest2",
				Email:    "email_Rtest2@example.com",
				Phone:    "13399992222",
			},
			wantErr: false,
		},
		{
			name:    "RetriveUser: Retrive a not-existed user",
			args:    args{"user_NotExist"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetriveUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetriveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RetriveUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllUsers(t *testing.T) {
	tests := []struct {
		name    string
		want    *entity.Users
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ListAllUsers: got 1,2,4",
			want: &entity.Users{&entity.User{
				Username: "user_Rtest1",
				Password: "pass_Rtest1",
				Email:    "email_Rtest1@example.com",
				Phone:    "13399991111",
			}, &entity.User{
				Username: "user_Rtest2",
				Password: "pass_Rtest2",
				Email:    "email_Rtest2@example.com",
				Phone:    "13399992222",
			}, &entity.User{
				Username: "user_Rtest4",
				Password: "pass_Rtest4",
				Email:    "email_Rtest4@example.com",
				Phone:    "13399994444",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckLoginInfo(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "CheckLoginInfo: user_Rtest1, got OK",
			args:    args{"user_Rtest1", "pass_Rtest1"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "CheckLoginInfo: user_Rtest2, got OK",
			args:    args{"user_Rtest2", "pass_Rtest2"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "CheckLoginInfo: user_WrongPassWord, got false, no err",
			args:    args{"user_Rtest1", "pass_Wrong"},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckLoginInfo(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLoginInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckLoginInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
