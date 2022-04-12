package api

type HomePage struct {
	Time string
}

type DataInput struct {
	From string
	To string
}

type DataOutput struct {
	Result string
	Text string
	Time string
	Duration int
}

type RespToArray struct {
	DiffMin int
	DepartureTime string
}
