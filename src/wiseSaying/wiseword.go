package wisesaying

type Item struct {
	Id     int    `json:"id"`
	Word   string `json:"word"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type Wisewordes struct {
	Items []Item
}

func New() *Wisewordes {
	return &Wisewordes{}
}

func (r *Wisewordes) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Wisewordes) GetAll() []Item {
	return r.Items
}
