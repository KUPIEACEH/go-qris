package services

import (
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/pkg/models"
)

func TestMapQRISEntityToModel(t *testing.T) {
	type args struct {
		qris *entities.QRIS
	}

	tests := []struct {
		name string
		args args
		want *models.QRIS
	}{
		{
			name: "Success",
			args: args{
				qris: &testQRISEntity,
			},
			want: &testQRISModel,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapQRISEntityToModel(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "mapQRISEntityToModel()", test.want, got)
			}
		})
	}
}

func TestMapQRISModelToEntity(t *testing.T) {
	type args struct {
		qris *models.QRIS
	}

	tests := []struct {
		name string
		args args
		want *entities.QRIS
	}{
		{
			name: "Success",
			args: args{
				qris: &testQRISModel,
			},
			want: &testQRISEntity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapQRISModelToEntity(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "mapQRISModelToEntity()", test.want, got)
			}
		})
	}
}
