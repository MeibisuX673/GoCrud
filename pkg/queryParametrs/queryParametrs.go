package queryParametrs

type QueryParams struct{

	Page int
	Limit int

}

func New() *QueryParams{

	return &QueryParams{
		Page: 1,
		Limit: 3,
	}
	
}