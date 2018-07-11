package values

type StoreRequest struct {
	First  Word
	Second Word
}

type GetRequest struct {
	Word       Word
	InLanguage string
}
