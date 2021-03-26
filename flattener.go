package flatten

type FlattenerFunc func(nested map[string]interface{}, prefix string, style SeparatorStyle) (map[string]interface{}, error)

type Flattener struct {
	f FlattenerFunc
}

func DefaultFlattener() *Flattener {
	return &Flattener{f: Flatten}
}

// Slices are ignored
func ObjectFlattener() *Flattener {
	return &Flattener{f: FlattenObjects}
}

func (f *Flattener) Flatten(nested map[string]interface{}, prefix string, style SeparatorStyle) (map[string]interface{}, error) {
	return f.f(nested, prefix, style)
}
