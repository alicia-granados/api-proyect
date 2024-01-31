package db

import "ApiRest/models"

type DatabaseRepo interface {
	Connect()
	Close()
	Ping()
	ExistsID(table string, id int) bool

	ChampionRead
	ChampionWriten
	SkinRepoRead
	SkinRepoWriten
	TagRepoRead
	TagRepoWriten
}

type ChampionRead interface {
	GetChampions() (Champion, error)
	GetChampionsDetailList() (Champion, error)
	GetChampionByID(championId int) (Champion, error)
	GetChampionDetailsByID(championId int) (Champion, error)
}

type ChampionWriten interface {
	DeleteChampionByID(championID int) error
	InsertChampion(name, title, lore string) error
	UpdateChampionByID(championID int, champion models.Champion) error
}

type SkinRepoRead interface {
	GetSkinList() (Skins, error)
	GetSkinByID(skinID int) (Skins, error)
}

type SkinRepoWriten interface {
	DeleteSkinByID(skinID int) error
	UpdateSkinByID(skinID int, skin models.Skins) error
	InsertSkin(idNum string, num, championID int, name string) error
}

type TagRepoRead interface {
	GetTagsList() (Tags, error)
	GetTagByID(tagID int) (Tags, error)
}

type TagRepoWriten interface {
	DeleteTagByID(tagID int) error
	InsertTag(championID int, tag string) error
	UpdateTagByID(tagID int, tag models.Tags) error
}
