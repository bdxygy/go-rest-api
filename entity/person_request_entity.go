package entity

type PersonCreateOrUpdateRequestEntity struct {
	UUID, Name, Address, City string
	Age                       int8
}

type PersonFindByIdOrDeleteEntity struct {
	UUID string
}
