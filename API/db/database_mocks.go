package db

import "errors"

type TestDBRepo struct {
}

func (m *TestDBRepo) Connect() {
	m.Ping()
}

func (m *TestDBRepo) Close() {
}

func (m *TestDBRepo) Ping() {
}

func (m *TestDBRepo) ExistTable(tableName string) bool {
	return true
}

func (m *TestDBRepo) InsertChampion(name, title, lore string) (int64, error) {
	if name == "Non-existant-fake-name" {
		return 0, errors.New("error no exits champion name")
	}
	if name == "Rigth-fake-name" {
		return 0, nil
	}
	if name == "Real-name" {
		return 1, nil
	}
	return 1, nil
}

func (m *TestDBRepo) GetTagID(tag string) (int, error) {
	if tag == "Non-existant-fake-tag" {
		return 0, errors.New("error no exits tag")
	}
	if tag == "Rigth-fake-tag" {
		return 0, nil
	}
	return 1, nil // ID of the found tag
}

func (m *TestDBRepo) GetChampionID(champion string) (int, error) {
	if champion == "Wrong-fake-name" {
		return 0, errors.New("there is an unexpected error")
	}
	if champion == "Non-existant-fake-name" || champion == "Rigth-fake-name" {
		return 0, nil
	}
	if champion == "Real-name" {
		return 1, nil
	}
	return 1, nil

}

func (m *TestDBRepo) InsertTag(championID int, tag string) error {
	if championID == 1 && tag == "tag1" {
		return errors.New("repeated tag")
	}
	return nil

}

func (m *TestDBRepo) GetSkinID(Id_Num string) (int64, error) {

	return 1, nil

}

func (m *TestDBRepo) InsertSkins(Id_Num string, num, championID int, name string) error {
	return nil
}
