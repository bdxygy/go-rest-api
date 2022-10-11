package entity

type PersonCreateOrUpdateRequestEntity struct {
	UUID    string
	Name    string `validate:"required"`
	Address string
	City    string
	Age     int8
}

type PersonFindByIdOrDeleteEntity struct {
	UUID string
}
