package entity

type Version struct {
	Id int
	Version int
}

func (o *Version) TableName() string {
	return "version"
}