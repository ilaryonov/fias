package version

type Option struct {
	Id int
	Name string
	Value int
}

func(o *Option) GetValueByName() string {
	return o.Name
}