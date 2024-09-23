package xip

import "testing"

func TestGetServerIp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "local ip", want: "192.168.31.95"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetServerIp()
			if got != tt.want {
				t.Errorf("GetServerIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
