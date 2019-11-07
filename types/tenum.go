package types

type TEnumItem struct {
	Index int
	Name  string
}
type TEnum struct {
	Items []TEnumItem
}

func (p *TEnum) GetNameByIndex(index int) string {
	for i := 0; i < len(p.Items); i++ {
		if p.Items[i].Index == index {
			return p.Items[i].Name
		}
	}
	return ""
}
func (p *TEnum) Add(index int, name string) {
	p.Items = append(p.Items, TEnumItem{index, name})
}
func (p *TEnum) GetIndexByName(name string) int {
	for i := 0; i < len(p.Items); i++ {
		if p.Items[i].Name == name {
			return p.Items[i].Index
		}
	}
	return -1
}

func init() {
	MARKETYPES := new(TEnum)
	MARKETYPES.Add(0, CategoryMarketCN)
	MARKETYPES.Add(1, CategoryMarketHK)
	MARKETYPES.Add(2, CategoryMarketUS)
}