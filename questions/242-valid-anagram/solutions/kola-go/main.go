package main

func isAnagram(s string, t string) bool {
if(len(s) != len(t)) { return false }
set := newMultiSet()
set.addStrings(s)
for _, v := range t {
	if !set.find(v) {
		return false
	}
}
return true
}

type multiSet struct {
	items map[rune]int
}

func newMultiSet() *multiSet {
	return &multiSet{
		items: map[rune]int{},
	}
}
func (m *multiSet)addStrings(s string){
	for _, v := range s {
	m.addString(v)
	}
}

func (m *multiSet)addString(v rune){
	m.items[v]++
}

func (m *multiSet)find(v rune)bool{
	if vv, ok := m.items[v]; ok && vv > 0 {
		m.items[v]--
		return true
	}
	return false
}