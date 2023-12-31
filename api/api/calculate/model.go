package calculate

type Response struct {
	Packs []string `json:"packs,omitempty"`
	Total int      `json:"total,omitempty"`
}
