package db

import "ApiRest/models"

type DatabaseRepo interface {
	Connect()
	Close()
	Ping()
	ExistsID(table string, id int) (int, error)
	ListChampions() (Champion, error)
	InsertChampion(name, title, lore string) error
	GetTagID(tag string) (int, error)
	GetChampionId(champion string) (int, error)
	UpdateChampion(championID int, Champion models.Champion) error
	DeleteChampion(championID int) error
	InsertTag(championID int, tag string) error
	GetSkinID(Id_Num string) (int64, error)
	InsertSkins(Id_Num string, num, championID int, name string) error
}
