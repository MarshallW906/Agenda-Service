package model

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		noToken string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "GetToken: Get Token of user_Rtest1",
			args:    args{"user_Rtest1"},
			noToken: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetToken(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.noToken {
				t.Errorf("GetToken() = nil string, want a non-nil string generated from md5")
			}
		})
	}
}

func TestDeleteToken(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "DeleteToken: WrongToken_ReturnErr",
			args:    args{"wrong_token"},
			wantErr: true,
		},
		{
			name:    "DeleteToken: Genereated just by GetToken, OK",
			args:    args{""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.tokenStr != "wrong_token" {
				token, err := GetToken("user_Rtest2")
				if err != nil {
					t.Errorf("DeleteToken() error = %v, when calling GetToken()", err)
				}
				tt.args.tokenStr = token
			}
			if err := DeleteToken(tt.args.tokenStr); (err != nil) != tt.wantErr {
				t.Errorf("DeleteToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckTokenValid(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "CheckTokenValid: Call GetToken(), got OK",
			args:    args{""},
			want:    true,
			wantErr: false,
		},
		{
			name:    "CheckTokenValid: WrongToken(), got false, no err",
			args:    args{"wrong_token"},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.tokenStr != "wrong_token" {
				token, err := GetToken("user_Rtest2")
				if err != nil {
					t.Errorf("CheckTokenValid() error = %v, when calling GetToken()", err)
				}
				tt.args.tokenStr = token
			}
			got, err := CheckTokenValid(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckTokenValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckTokenValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsernameByToken(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "CheckTokenValid: Call GetToken(\"user_Rtest2\"), got OK",
			args:    args{""},
			want:    "user_Rtest2",
			wantErr: false,
		},
		{
			name:    "CheckTokenValid: WrongToken(), got false, no err",
			args:    args{"wrong_token"},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.tokenStr != "wrong_token" {
				token, err := GetToken("user_Rtest2")
				if err != nil {
					t.Errorf("GetUsernameByToken() error = %v, when calling GetToken()", err)
				}
				tt.args.tokenStr = token
			}
			got, err := GetUsernameByToken(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsernameByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUsernameByToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genToken(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		noToken string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "genToken: no error",
			args:    args{"user_Rtest1"},
			noToken: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genToken(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("genToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.noToken {
				t.Errorf("genToken() = \"\", want a string generated by md5")
			}
		})
	}
}