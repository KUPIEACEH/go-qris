package usecases

import (
	"reflect"
	"testing"
)

func TestCommonInRange(t *testing.T) {
	type args struct {
		value string
		start string
		end   string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success: True",
			args: args{
				value: "27",
				start: "27",
				end:   "49",
			},
			want: true,
		},
		{
			name: "Success: False",
			args: args{
				value: "22",
				start: "27",
				end:   "49",
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := inRange(test.args.value, test.args.start, test.args.end)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Expected %v = %v, but got = %v", "inRange()", test.want, got)
			}
		})
	}
}
