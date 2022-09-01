package cat

import (
	"catinator-backend/pkg/db/ent"
	"catinator-backend/pkg/model"
	"catinator-backend/pkg/rfctime"
)

func MapDBCatToModel(dbCat *ent.Cat) *model.Cat {
	if dbCat == nil {
		return nil
	}
	return &model.Cat{
		ID:          dbCat.ID,
		Name:        dbCat.Name,
		Description: dbCat.Description,
		Tags:        dbCat.Tags,
		ImageID:     dbCat.ImageID,
		CreatedAt: rfctime.RFC3339Time{
			Time: dbCat.CreateTime,
		},
		UpdatedAt: rfctime.RFC3339Time{
			Time: dbCat.UpdateTime,
		},
	}
}

func MapDBCatsToModel(dbCats []*ent.Cat) []*model.Cat {
	if dbCats == nil {
		return nil
	}
	cats := make([]*model.Cat, 0, len(dbCats))
	for _, dbCat := range dbCats {
		cats = append(cats, MapDBCatToModel(dbCat))
	}
	return cats
}
