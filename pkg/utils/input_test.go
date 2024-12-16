package utils

import (
	"reflect"
	"testing"
)

func TestNewInput(t *testing.T) {
	tests := []struct {
		name   string
		fields Input
		want   InputInterface
	}{
		{
			name:   "Success",
			fields: Input{},
			want:   &Input{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := NewInput()

			if u == nil {
				t.Errorf(expectedReturnNonNil, "NewInput", "InputInterface")
			}

			got, ok := u.(*Input)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*Input")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*Input", test.want, got)
			}
		})
	}
}

func TestInputSanitize(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name   string
		fields Input
		args   args
		want   string
	}{
		{
			name:   "Success: Trim Space",
			fields: Input{},
			args: args{
				input: "  hello world  ",
			},
			want: "hello world",
		},
		{
			name:   "Success: With Word",
			fields: Input{},
			args: args{
				input: "\nhello\nworld\r",
			},
			want: "helloworld",
		},
		{
			name:   "Success: No Word",
			fields: Input{},
			args: args{
				input: "   \n\r  ",
			},
			want: "",
		},
		{
			name:   "Success: No Replace",
			fields: Input{},
			args: args{
				input: "hello",
			},
			want: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := test.fields
			got := u.Sanitize(test.args.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Sanitize()", test.want, got)
			}
		})
	}
}
