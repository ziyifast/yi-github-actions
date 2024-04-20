package service

import "testing"

func Test_userService_GetUserName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test", "tom"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{}
			if got := u.GetUserName(); got != tt.want {
				t.Errorf("GetUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
