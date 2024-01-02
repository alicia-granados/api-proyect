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
