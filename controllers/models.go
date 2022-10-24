package controllers

type passedBook struct {
	Name   string `form:"name" json:"name" xml:"name"  binding:"required"`
	Author string `form:"author" json:"author" xml:"author"  binding:"required"`
}
