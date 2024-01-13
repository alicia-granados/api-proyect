package db

import (
	"ApiRest/models"
)

type Skins []models.Skins

func (r *RealDBRepo) GetSkins() (Skins, error) {
	sql := "SELECT Skins.Id, Skins.Id_Num, Skins.Num, Skins.Id_Champion, Skins.Name FROM Skins"

	Skinss := Skins{}
	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Skins := models.Skins{}
		err := rows.Scan(&Skins.Id, &Skins.IdNum, &Skins.Num, &Skins.IdChampion, &Skins.Name)
		if err != nil {
			return nil, err
		}

		// AAdd the data object to the Skins slice
		Skinss = append(Skinss, Skins)
	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Skinss, nil
}

func (r *RealDBRepo) GetSkinId(SkinsId int) (Skins, error) {
	sql := "SELECT Skins.Id, Skins.Id_Num, Skins.Num, Skins.Id_Champion, Skins.Name FROM Skins WHERE Skins.Id = ? "

	Skinss := Skins{}
	rows, err := r.DB.Query(sql, SkinsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Make sure to close the rows at the end

	for rows.Next() {
		Skins := models.Skins{}

		err := rows.Scan(&Skins.Id, &Skins.IdNum, &Skins.Num, &Skins.IdChampion, &Skins.Name)

		if err != nil {
			return nil, err
		}

		Skinss = append(Skinss, Skins)

	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return Skinss, nil
}

// InsertSkin inserts a new tag and returns its ID.
func (r *RealDBRepo) InsertSkin(Id_Num string, num, championID int, name string) error {
	_, err := r.DB.Exec("INSERT INTO Skins (Id_Num, Num, Id_Champion, Name) VALUES (?,?, ?,?)", Id_Num, num, championID, name)
	return err
}

func (r *RealDBRepo) UpdateSkin(skinID int, Skin models.Skins) error {
	sql := "UPDATE Skins SET  Id_Num=?, Num= ? , Id_Champion=?,  Name=? WHERE Id=?"
	_, err := r.DB.Exec(sql, Skin.IdNum, Skin.Num, Skin.IdChampion, Skin.Name, skinID)
	return err
}

func (r *RealDBRepo) DeleteSkin(skinID int) error {
	sql := "DELETE FROM Skins WHERE Id=?"
	_, err := r.DB.Exec(sql, skinID)
	return err
}
