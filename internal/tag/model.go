package tag

import (
	"github.com/google/uuid"
	"inventory/ent/entgen"
)

type RQTag struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	IsActive bool   `json:"isActive"`
}

type Tag struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	IsActive bool   `json:"isActive"`
}

type RSTag struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	IsActive bool   `json:"isActive"`
}

func (rs *RSTag) MapFrom(tag *Tag) {
	rs.Id = tag.Id
	rs.Name = tag.Name
	rs.Desc = tag.Desc
	rs.IsActive = tag.IsActive
}

func (tag *Tag) MapFrom(rq *RQTag) {
	tag.Id = uuid.NewString()

	tag.Name = rq.Name
	tag.Desc = rq.Desc
	tag.IsActive = rq.IsActive

}

func (tag *Tag) MapFromEnt(entTag *entgen.TblTag) {
	tag.Id = entTag.ID
	tag.Name = entTag.Name
	tag.IsActive = entTag.IsActive
	tag.Desc = entTag.Description
}
