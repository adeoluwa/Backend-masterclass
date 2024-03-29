// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transfers.sql

package backend_masterclass

import (
	"context"
	"reflect"
	"testing"
)

func TestQueries_CreateTransfer(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreateTransferParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.CreateTransfer(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CreateTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.CreateTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_GetTransfer(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.GetTransfer(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.GetTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_ListTransfer(t *testing.T) {
	type args struct {
		ctx context.Context
		arg ListTransferParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    []Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.ListTransfer(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.ListTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.ListTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
