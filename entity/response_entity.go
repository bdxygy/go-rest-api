package entity

type ResponseEntity struct {
	Code    int16
	Status  bool
	Data    interface{}
	Message string
}
