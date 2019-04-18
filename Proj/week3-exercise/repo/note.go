package repo

import (
	"github.com/Golang/Proj/week3-exercise/helper"
	"github.com/Golang/Proj/week3-exercise/model"
	"github.com/jinzhu/gorm"
)

type NoteRepo interface {
	Find(int) (*model.Note, error)
	List(helper.Pagination) ([]model.Note, error)
	Update(int, model.Note) error
	Delete(int) error
	Create(model.Note) (*model.Note, error)
}

type NoteRepoImpl struct {
	DB *gorm.DB
}

func (self *NoteRepoImpl) Create(note model.Note) (*model.Note, error) {

	err := self.DB.Create(&note).Error
	return &note, err
}

func (self NoteRepoImpl) Find(id int) (*model.Note, error) {
	note := &model.Note{}
	err := self.DB.Where("id = ?", id).First(note).Error
	return note, err
}

func (self *NoteRepoImpl) List(pagination helper.Pagination) ([]model.Note, error) {
	notes := []model.Note{}
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	error := self.DB.Offset(offset).Limit(limit).Find(&notes).Error
	return notes, error
}

func (self *NoteRepoImpl) Update(id int, note model.Note) error {

	error := self.DB.Where("id = ?", id).Delete(&model.Note{}).Error
	return error
}
