package generator

type Controller interface {
	Expose()
}

type Usecase interface {
	Generate() error
}
