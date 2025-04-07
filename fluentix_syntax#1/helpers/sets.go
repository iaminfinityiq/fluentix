package helpers

type IntSet struct {
	Set map[int]bool
}

func (i *IntSet) Add(value int) {
	i.Set[value] = true
}

func (i *IntSet) Remove(value int) {
	i.Set[value] = false
}

func (i *IntSet) Contains(value int) bool {
	return i.Set[value]
}