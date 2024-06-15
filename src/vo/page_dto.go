package vo

type PageVO struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}
