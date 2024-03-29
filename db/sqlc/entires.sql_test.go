// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: entires.sql

package backend_masterclass

import (
	"context"
	"reflect"
	"testing"
)

func TestQueries_CreateEntery(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreateEnteryParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.CreateEntery(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CreateEntery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.CreateEntery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_GetEntry(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.GetEntry(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.GetEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_ListEntries(t *testing.T) {
	type args struct {
		ctx context.Context
		arg ListEntriesParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    []Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.ListEntries(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.ListEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.ListEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}
