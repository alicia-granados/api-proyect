package db

import (
	"ApiRest/models"
)

type Champion []models.Champion

func (r *RealDBRepo) AllInfoChampions() (Champion, error) {
	sql := "SELECT Champion.Id, Champion.Name, Champion.Title, Champion.Lore, Skins.Id, Skins.Id_Num, Skins.Num, Skins.Id_Champion, Skins.Name, Tags.Id, Tags.Id_Champion, Tags.Name FROM Champion LEFT JOIN Skins ON Champion.Id = Skins.Id_Champion LEFT JOIN Tags ON Champion.Id = Tags.Id_Champion"

	champions := Champion{}
	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		champion := models.Champion{}
		skin := models.Skins{}
		tag := models.Tags{}

		err := rows.Scan(&champion.Id, &champion.Name, &champion.Title, &champion.Lore, &skin.Id, &skin.Id_Num, &skin.Num, &skin.Id_Champion, &skin.Name, &tag.Id, &tag.Id_Champion, &tag.Name)
		if err != nil {
			return nil, err
		}

		champion.Skins = []models.Skins{skin}
		champion.Tags = []models.Tags{tag}
		// AAdd the data object to the champion slice
		champions = append(champions, champion)
	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return champions, nil
}

// InsertChampion inserts a champion into the database and returns its ID
func (r *RealDBRepo) InsertChampion(name, title, lore string) error {
	_, err := r.DB.Exec("INSERT INTO Champion (Name, Title, Lore) VALUES (?, ?, ?)", name, title, lore)
	return err
}

func (r *RealDBRepo) GetInfoChampionId(championId int) (Champion, error) {
	sql := "SELECT Champion.Id, Champion.Name, Champion.Title, Champion.Lore, Skins.Id, Skins.Id_Num,  Skins.Num,  Skins.Id_Champion,  Skins.Name,  Tags.Id, Tags.Id_Champion, Tags.Name  FROM Champion LEFT JOIN Skins ON Champion.Id = Skins.Id_Champion LEFT JOIN Tags ON Champion.Id = Tags.Id_Champion WHERE Champion.Id = ? "

	champions := Champion{}
	rows, err := r.DB.Query(sql, championId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Make sure to close the rows at the end

	for rows.Next() {
		champion := models.Champion{}
		skin := models.Skins{}
		tag := models.Tags{}

		err := rows.Scan(&champion.Id, &champion.Name, &champion.Title, &champion.Lore, &skin.Id, &skin.Id_Num, &skin.Num, &skin.Id_Champion, &skin.Name, &tag.Id, &tag.Id_Champion, &tag.Name)

		if err != nil {
			return nil, err
		}

		champion.Skins = []models.Skins{}
		champion.Tags = []models.Tags{}

		champion.Skins = append(champion.Skins, skin)
		champion.Tags = append(champion.Tags, tag)

		champions = append(champions, champion)

	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return champions, nil
}

// UpdateChampion update a champion into the database and returns err
func (r *RealDBRepo) UpdateChampion(championID int, Champion models.Champion) error {
	sql := "UPDATE Champion SET  Name=?, Title= ? , Lore=? WHERE Id=?"
	_, err := r.DB.Exec(sql, Champion.Name, Champion.Title, Champion.Lore, championID)
	return err
}

// DeleteChampion delete a champion into the database
func (r *RealDBRepo) DeleteChampion(championID int) error {
	sql := "DELETE FROM Champion  WHERE Id=?"
	_, err := r.DB.Exec(sql, championID)
	return err
}
