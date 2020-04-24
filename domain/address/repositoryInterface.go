package address

type AddressRepositoryInterface interface {
	GetByFormalname(term string) (*AddrObject, error)
	GetCityByFormalname(term string) (*AddrObject, error)
	BatchInsert() bool
}