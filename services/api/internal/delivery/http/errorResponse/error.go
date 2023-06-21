package errorResponse


type Error struct{

	Message string `json:"message"`
	StatusCode int `json:"status"`
	Desscription string `json:"desscription"`

}