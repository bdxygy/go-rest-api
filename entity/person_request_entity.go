package entity

type PersonCreateOrUpdateRequestEntity struct {
	Name, Address, City string
	Age                 int8
}

type PersonFindByIdOrDeleteEntity struct {
	UUID string
}
