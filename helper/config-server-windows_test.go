package helper

import "testing"

func TestConfigServer(t *testing.T) {
	type args struct {
		homeDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"demo", args{"D:\\gopath\\src\\github.com\\yiwenlong\\launchduidemo\\windows"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConfigServer(tt.args.homeDir); (err != nil) != tt.wantErr {
				t.Errorf("ConfigServerOnWindows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
