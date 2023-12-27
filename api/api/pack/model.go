package pack

type Pack struct {
	Size int `json:"size"  db:"size"`
}

const PackKey string = "pack"
