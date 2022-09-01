package model

import "catinator-backend/pkg/rfctime"

type Cat struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Tags        []string            `json:"tags"`
	ImageID     string              `json:"imageid"`
	CreatedAt   rfctime.RFC3339Time `json:"createdAt"`
	UpdatedAt   rfctime.RFC3339Time `json:"updatedAt"`
}

type AddCat struct {
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	Tags        []string `json:"tags" form:"tags"`
}

type UpdateCat struct {
	Name        *string   `json:"name" form:"name"`
	Description *string   `json:"description" form:"description"`
	Tags        *[]string `json:"tags" form:"tags"`
}
