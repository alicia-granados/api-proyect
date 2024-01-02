package handlers

import (
	"API/db"
	"fmt"
	"testing"
)

func TestProcessChampions(t *testing.T) {
	type args struct {
		dbRepo     db.DatabaseRepo
		infCampeon Champion
	}
	dbRepo := &db.TestDBRepo{}
	tests := []struct {
		name string
		args args
		msg  error
	}{
		{
			name: "should panic when champion could not be found",
			args: args{
				dbRepo: dbRepo,
				infCampeon: Champion{
					Id:    "Fake-id",
					Name:  "Wrong-fake-name",
					Title: "Fake-title",
					Skins: []Skins{},
					Lore:  "Fake-lore",
					Tags:  []string{},
				},
			},
			msg: fmt.Errorf("error getting the champion ID:there is an unexpected error"),
		},
		{
			name: "found champion",
			args: args{
				dbRepo: dbRepo,
				infCampeon: Champion{
					Id:    "Fake-id",
					Name:  "Non-existant-fake-name",
					Title: "Fake-title",
					Skins: []Skins{},
					Lore:  "Fake-lore",
					Tags:  []string{},
				},
			},
			msg: fmt.Errorf("error inserting the champion: error no exits champion name"),
		},
		{
			name: "found champion",
			args: args{
				dbRepo: dbRepo,
				infCampeon: Champion{
					Id:    "Fake-id",
					Name:  "Rigth-fake-name",
					Title: "Fake-title",
					Skins: []Skins{},
					Lore:  "Fake-lore",
					Tags:  []string{},
				},
			},
			msg: nil,
		},
		{
			name: "skin inserted correctly",
			args: args{
				dbRepo: dbRepo,
				infCampeon: Champion{
					Id:    "Real-skin-id",
					Name:  "Real-name",
					Title: "Real-title",
					Skins: []Skins{},
					Lore:  "Real-lore",
					Tags:  []string{},
				},
			},
			msg: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ProcessChampions(tt.args.dbRepo, tt.args.infCampeon)

			if err != nil {
				if err.Error() != tt.msg.Error() {
					t.Errorf("%v: expected returned error to be %v", tt.name, tt.msg)
				}
			}

		})
	}
}

func TestProcessTags(t *testing.T) {
	type args struct {
		dbRepo       db.DatabaseRepo
		tags         []string
		championName string
	}
	dbRepo := &db.TestDBRepo{}
	tests := []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "should panic when champion could not be found",
			args: args{
				dbRepo:       dbRepo,
				tags:         []string{"tag1", "tag1"},
				championName: "Wrong-fake-name",
			},
			expectedError: fmt.Errorf("error getting the champion ID:there is an unexpected error"),
		},
		{
			name: "champion has been found but the tags are empty",
			args: args{
				dbRepo:       dbRepo,
				tags:         []string{"tag1"},
				championName: "Real-name",
			},
			expectedError: fmt.Errorf("error inserting the tag:repeated tag"),
		},
		{
			name: "tag inserted correctly",
			args: args{
				dbRepo:       dbRepo,
				tags:         []string{"tagtest", "tagtest2"},
				championName: "Real-name",
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ProcessTags(tt.args.dbRepo, tt.args.tags, tt.args.championName)
			if err != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("%v: expected returned error to be %v", err, tt.expectedError)
				}
			}
		})
	}
}

func TestProcesSkins(t *testing.T) {
	type args struct {
		dbRepo       db.DatabaseRepo
		skins        []Skins
		championName string
	}
	dbRepo := &db.TestDBRepo{}
	tests := []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "should panic when skin could not be found",
			args: args{
				dbRepo: dbRepo,
				skins: []Skins{
					{
						Id_Num: "Wrong-fake-id",
						Num:    0,
						Name:   "Non-existant-fake-skin",
					},
				},
				championName: "fake-name",
			},
			expectedError: fmt.Errorf("error getting the skin ID:there is an unexpected error"),
		},
		{
			name: "found skin but no found champion",
			args: args{
				dbRepo: dbRepo,
				skins: []Skins{
					{
						Id_Num: "Non-existant-fake-skin",
						Num:    0,
						Name:   "Non-existant-fake-name",
					},
				},
				championName: "Wrong-fake-name",
			},
			expectedError: fmt.Errorf("error getting the champion ID:there is an unexpected error"),
		},
		{
			name: "skins no exits ",
			args: args{
				dbRepo: dbRepo,
				skins: []Skins{
					{
						Id_Num: "Non-existant-fake-skin",
						Num:    0,
						Name:   "skin-default",
					},
				},
				championName: "Non-existant-fake-name",
			},
			expectedError: fmt.Errorf("error getting the skin ID:there is an unexpected error"),
		},
		{
			name: "Champion has been found but the skin is repeated",
			args: args{
				dbRepo: dbRepo,
				skins: []Skins{
					{
						Id_Num: "Non-existant-fake-skin",
						Num:    0,
						Name:   "Repeated-skin",
					},
				},
				championName: "Real-name",
			},
			expectedError: fmt.Errorf("error inserting the skin: repeated skin"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ProcesSkins(tt.args.dbRepo, tt.args.skins, tt.args.championName)
			if err != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("%v: expected returned error to be %v", err, tt.expectedError)
				}
			}
		})
	}
}
