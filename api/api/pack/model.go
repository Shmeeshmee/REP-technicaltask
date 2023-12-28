package pack

type Pack struct {
	Size int `json:"size"  db:"size"`
}

const PackKey string = "pack"

type Elem struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
