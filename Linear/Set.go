package Linear

type Set struct {
	Lookup map[int]bool
}

func(set *Set) New() {
	set.Lookup = make(map[int]bool)
}

func(set *Set) AddElement(val int) {
	set.Lookup[val] = true
}

func(set *Set) DeleteElement(val int) {
	delete(set.Lookup, val)
}

func (set *Set) ContainsElement(val int) bool {
	var exists bool
	_, exists = set.Lookup[val]
	return exists
}