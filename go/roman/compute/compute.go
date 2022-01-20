package compute

import (
	"github.com/emirpasic/gods/maps/treemap"
)

type Arab struct {
	value int
}

type Roman struct {
	value string
}

func recConvertToRoman(a Arab, m *treemap.Map) string {
	arabValue, romanValue := m.Floor(a.value)
	if a.value == arabValue {
		return romanValue.(string)
	}

	return romanValue.(string) + recConvertToRoman(Arab{a.value - arabValue.(int)}, m)
}

func (a Arab) ConvertToRoman() Roman {
	m := treemap.NewWithIntComparator()
	m.Put(1000, "M")
	m.Put(900, "CM")
	m.Put(500, "D")
	m.Put(400, "CD")
	m.Put(100, "C")
	m.Put(90, "XC")
	m.Put(50, "L")
	m.Put(40, "XL")
	m.Put(10, "X")
	m.Put(9, "IX")
	m.Put(5, "V")
	m.Put(4, "IV")
	m.Put(1, "I")

	return Roman{recConvertToRoman(a, m)}
}
