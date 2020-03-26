package engine

type Request struct {
	URL       string
	ParseFunc func([]byte) ParseRes
}

type ParseRes struct {
	Requests []Request
	Items    []interface{}
}

func NilParse(content []byte) ParseRes {
	return ParseRes{}
}
