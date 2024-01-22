package db

import "ApiRest/models"

type DatabaseRepo interface {
	Connect()
	Close()
	Ping()
	ExistsID(table string, id int) (int, error)

	ChampionRead
	ChampionWriten
	SkinRepoRead
	SkinRepoWriten
	TagRepoRead
	TagRepoWriten
}

type ChampionRead interface {
	GetChampions() (models.Champion, error)
	GetChampionsDetailList() (models.Champion, error)
	GetChampionByID(championId int) (models.Champion, error)
	GetChampionDetailsByID(championId int) (models.Champion, error)
}

type ChampionWriten interface {
	DeleteChampionByID(championID int) error
	InsertChampion(name, title, lore string) error
	UpdateChampionByID(championID int, champion models.Champion) error
}

type SkinRepoRead interface {
	GetSkinList() (models.Skins, error)
	GetSkinByID(skinID int) (models.Skins, error)
}

type SkinRepoWriten interface {
	DeleteSkinByID(skinID int) error
	UpdateSkinByID(skinID int, skin models.Skins) error
	InsertSkin(idNum string, num, championID int, name string) error
}

type TagRepoRead interface {
	GetTagsList() (models.Tags, error)
	GetTagByID(tagID int) (models.Tags, error)
}

type TagRepoWriten interface {
	DeleteTagByID(tagID int) error
	InsertTag(championID int, tag string) error
	UpdateTagByID(tagID int, tag models.Tags) error
}
