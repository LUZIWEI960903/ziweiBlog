package models

type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryRes struct {
	*HomeRes
	CategoryName string
}
