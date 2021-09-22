package maputil

func MapStringSlice(origin[]string,  apply func (string)string)  []string{
	var res []string
	for _, v := range origin {
		res = append(res, apply(v))
	}
	return res
}