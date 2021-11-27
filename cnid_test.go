package cnid

import "testing"

func TestId17to18(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"first",args{"37148119850509001"},"371481198505090015"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Id17to18(tt.args.id); got != tt.want {
				t.Errorf("Id17to18() = %v, want %v", got, tt.want)
			}
		})
	}
}
