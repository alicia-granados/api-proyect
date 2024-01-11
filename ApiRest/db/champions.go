package db

import (
	"ApiRest/entities"
)

type Champion []entities.Champion

func (r *RealDBRepo) ListChampions() (Champion, error) {
	sql := "SELECT Champion.Id, Champion.Name, Champion.Title, Champion.Lore, Skins.Id, Skins.Id_Num,  Skins.Num,  Skins.Id_Champion,  Skins.Name,  Tags.Id, Tags.Id_Champion, Tags.Name  FROM Champion LEFT JOIN Skins ON Champion.Id = Skins.Id_Champion LEFT JOIN Tags ON Champion.Id = Tags.Id_Champion"

	champions := Champion{}
	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Make sure to close the rows at the end

	for rows.Next() {
		champion := entities.Champion{}
		skin := entities.Skins{}
		tag := entities.Tags{}

		err := rows.Scan(&champion.Id, &champion.Name, &champion.Title, &champion.Lore, &skin.Id, &skin.Id_Num, &skin.Num, &skin.Id_Champion, &skin.Name, &tag.Id, &tag.Id_Champion, &tag.Name)

		if err != nil {
			return nil, err
		}

		champion.Skins = []entities.Skins{}
		champion.Tags = []entities.Tags{}

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
