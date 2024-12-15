package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewQRCode(t *testing.T) {
	tests := []struct {
		name   string
		fields QRCode
		want   QRCodeInterface
	}{
		{
			name:   "Success",
			fields: QRCode{},
			want:   &QRCode{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewQRCode()

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewQRCode", "QRCodeInterface")
			}

			got, ok := uc.(*QRCode)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*QRCode")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*QRCode", test.want, got)
			}
		})
	}
}

func TestStringToImageBase64(t *testing.T) {
	type args struct {
		qrString   string
		qrCodeSize int
	}

	tests := []struct {
		name      string
		fields    QRCode
		args      args
		want      string
		wantError error
	}{
		{
			name:   "Error: QR Code Scale",
			fields: QRCode{},
			args: args{
				qrString:   testQRISStaticString,
				qrCodeSize: -1,
			},
			want:      "",
			wantError: fmt.Errorf("can not scale barcode to an image smaller than 53x53"),
		},
		{
			name:   "Success",
			fields: QRCode{},
			args: args{
				qrString:   testQRISStaticString,
				qrCodeSize: 125,
			},
			want:      "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAH0AAAB9EAAAAAD6e++6AAAFC0lEQVR4nOxcy5LjNgyEU/v/v+zUHJSC2/0A9zJVIfqyHoqkKDYANiAlf97vuhT//PYCfg/76DdiH/1G7KPfiIsf/c930+ulu/8IoJ/rjxDqfafXep+n39Oe+rJ5p+tGLOsItkvICvZnu97nYez9tKk+fT5mBT+/0SrcuhHLOkPyKeav+LsIY8nvC6wIWcN5putGLOunwN10UbiMr6rrrF+yplMs638DFuXL+FfyXfa38uUU1SdY1hnUbiK7E5+bRHWm5kpohlPlxrCsI9yuJpankZddZ9aBGr/Iue9ig8Ky3pF8RUVj9W8f5xQi83FEiicnuJj11/deqfN6km9PYwCC6X/2O2WPuH63notZN/n6afUkZXepb2K5rwHbp9lax7KOcNqZ5d8lon1nZ1KlqYEVqLNexSBlBRezTiL8f5dM/Y3l1BPllXzZRXCcLym45PPL+leziOo12E1lLe4EcBFaWZoap857xLL+1RyUG/apYQ18yphivkycYb/ZfR8s6x9NQmOfaG923Y1lcIqSrStZJ+Ji1ge1OberqbrKxk98U62lRO0O7z/R8ct6h4vUqvKSLECpPqXimD7v90wWmKo9tawTKC3OfmMb87dUuZn6O8vqWBxI961lXUCdnwmquuLmV5HZMZ3UWlrvss7Q/c9p+SInwPvNxyVM8/4KtULWhljWHZBBdn6XeV9WobbvskR1f7QuphX6PKvhPyC+lmQamymzIrXxPk+HmrOPdbm/UobKr1OMWdY7lN5mOvr18rk0+/1gouRUJshiAluHy+KW9Q4VEVVejGqLqTRkQbUXYZFlcCzaK/aX9S+Y2pzzQbXjbBzzVTafWoOqEuF1hlVzFOadW5m6XJn3Z6y/swRXu8O1KNWXqjar5j5gIryrtTEoJtSYpNmZFbg1uDEMy3qHU3Dq90nETTo9nc/TNdXW5hQGb1oZUgVF9anBOzZswzXhOhU2c5MIb1rTzuG4dL7Xwfc0HUkhurWsr38hvHNT/uvyd/ZvBeamjLtzO50eiItZH35Lo/yyjGXU4MsIVYFxCm5y+rBxiGX9o8lkTg8cWwhlBWxOF/Unvq1U5WZuHxi+X3+bLxlwR1OtjPWtYd2+35dZ0lQn1N2sD9Vc8hvls/36JCdwuYDSBUrbr4aXGL5ffw/ejNaA8RSlk8Wp0wHvkTK+WtYJJv6l2hWzropyktfjOlnfzdwsgoZX7c5Xk1LD+Z2vJ3Wn4gq7P+Ji1sO3NMzfkAHm/6isEMxanG+zPng/pwYZlvUOtZvTnLvC21n2tzqrGdSpU+E0QSzrDEwXqz6nmp2NKWNZ7NzGNU4UXMey3qEiLsutlYpi8/UTwZ3xypJUno5rm7K/rHcodeUUFI53yqxMbc/N7+KJywUULmZ98F8ys/YHKYqr6otiRGVwKj6cZHeIZV1eHu40sqv082R8v3eF+uDU4hguZj1kbuq8dLbCamQq766Dryxc3FFnuWN+We9I/uP8jOXqfQzq8kklSF1P86HlIZZ1xCQHd1nWxNeY1mdItTs23sWWB8s6Q4quTI+zrGuq+4vEhRIaflLXc4zXsn4IxYKL1ie5+YRdd9/E9oNl/RTKb5VaU5Ge1QNcbo/5Oc6F5/hqeIq/+H9Gl4jg6vxV40rk86rq47JCpeQclnVEqn6oHFpFYcesqtdh2ySmVH3nEAoXsx6qNP9nXMz6PvqN2Ee/EfvoN2If/Ub8GwAA//8EcncY468vnQAAAABJRU5ErkJggg==",
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := test.fields

			got, err := uc.StringToImageBase64(test.args.qrString, test.args.qrCodeSize)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "StringToImageBase64()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "StringToImageBase64()", test.want, got)
			}
		})
	}
}
