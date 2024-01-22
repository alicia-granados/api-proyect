package db

import (
	"ApiRest/models"
)

type Tags []models.Tags

func (r *RealDBRepo) GetTagsList() (Tags, error) {
	sql := "SELECT Tags.Id, Tags.Id_Champion, Tags.Name FROM Tags"

	Tagss := Tags{}
	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Tags := models.Tags{}
		err := rows.Scan(&Tags.Id, &Tags.IdChampion, &Tags.Name)
		if err != nil {
			return nil, err
		}

		// AAdd the data object to the Tags slice
		Tagss = append(Tagss, Tags)
	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Tagss, nil
}

func (r *RealDBRepo) GetTagByID(TagsId int) (Tags, error) {
	sql := "SELECT Tags.Id, Tags.Id_Champion, Tags.Name FROM Tags WHERE Tags.Id = ? "

	Tagss := Tags{}
	rows, err := r.DB.Query(sql, TagsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Make sure to close the rows at the end

	for rows.Next() {
		Tags := models.Tags{}

		err := rows.Scan(&Tags.Id, &Tags.IdChampion, &Tags.Name)

		if err != nil {
			return nil, err
		}

		Tagss = append(Tagss, Tags)

	}

	// Check for errors after exiting the loop
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return Tagss, nil
}

// InsertTag inserts a new tag
func (r *RealDBRepo) InsertTag(championID int, tag string) error {
	_, err := r.DB.Exec("INSERT INTO Tags (Id_Champion, Name) VALUES (?, ?)", championID, tag)
	return err
}

func (r *RealDBRepo) UpdateTagByID(tagID int, Tag models.Tags) error {
	sql := "UPDATE Tags SET Id_Champion=?,  Name=? WHERE Id=?"
	_, err := r.DB.Exec(sql, Tag.IdChampion, Tag.Name, tagID)
	return err
}

func (r *RealDBRepo) DeleteTagByID(tagID int) error {
	sql := "DELETE FROM Tags WHERE Id=?"
	_, err := r.DB.Exec(sql, tagID)
	return err
}
