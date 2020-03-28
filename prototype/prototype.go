package prototype

type ProtoType struct {
	Value0 string
	Value1 *DeepValue
	Value2 []int
	Value4 map[string]int
}

type DeepValue struct {
	Value6 string
}

// 根据字段需求选择深拷贝浅拷贝
func (p *ProtoType) Clone() *ProtoType {
	protoType := *p
	pv1 := *(p.Value1)
	protoType.Value1 = &pv1

	v2 := make([]int, len(p.Value2))
	copy(v2, p.Value2)
	protoType.Value2 = v2

	mp := make(map[string]int)
	for s, i := range p.Value4 {
		mp[s] = i
	}
	protoType.Value4 = mp

	return &protoType
}
