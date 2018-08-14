package sdk_go

type TypeString string

type Conf interface {
	Get(a TypeString, b unit32) error
}
