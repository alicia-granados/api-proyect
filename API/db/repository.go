package db

import "database/sql"

type DatabaseRepo interface {
	Connect()
	Close()
	Ping()
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	InsertChampion(name, title, lore string) (int64, error)
	GetTagID(tag string) (int, error)
	GetChampionID(champion string) (int, error)
	InsertTag(championID int, tag string) error
	GetSkinID(Id_Num string) (int64, error)
	InsertSkins(Id_Num string, num, championID int, name string) error
}
