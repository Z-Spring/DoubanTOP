package movie

type ById []*Movie

func (id ById) Len() int {
	return len(id)
}

func (id ById) Less(i int, j int) bool {
	return id[i].Id < id[j].Id
}

func (id ById) Swap(i int, j int) {
	id[i], id[j] = id[j], id[i]
}
