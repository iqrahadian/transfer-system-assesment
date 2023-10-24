package common

type Error struct {
	Error   error  `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
