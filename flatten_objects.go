package flatten

func FlattenObjects(nested map[string]interface{}, prefix string, style SeparatorStyle) (map[string]interface{}, error) {
	flatmap := make(map[string]interface{})

	err := flattenObjects(true, flatmap, nested, prefix, style)
	if err != nil {
		return nil, err
	}

	return flatmap, nil
}

func flattenObjects(top bool, flatMap map[string]interface{}, nested interface{}, prefix string, style SeparatorStyle) error {
	assign := func(newKey string, v interface{}) error {
		switch vv := v.(type) {
		case map[string]interface{}:
			if err := flattenObjects(false, flatMap, v, newKey, style); err != nil {
				return err
			}
		case []interface{}:
			slice := make([]interface{}, len(vv), len(vv))
			for i, t := range vv {
				switch tt := t.(type) {
				case map[string]interface{}:
					innerFlatMap := make(map[string]interface{})
					if err := flattenObjects(true, innerFlatMap, tt, prefix, style); err != nil {
						return err
					}
					slice[i] = innerFlatMap
				default:
					slice[i] = tt
				}
			}
			flatMap[newKey] = slice
		default:
			flatMap[newKey] = v
		}

		return nil
	}

	switch nested.(type) {
	case map[string]interface{}:
		for k, v := range nested.(map[string]interface{}) {
			newKey := enkey(top, prefix, k, style)
			assign(newKey, v)
		}
	case []interface{}:
		newKey := enkey(top, prefix, "", style)
		assign(newKey, nested)
	default:
		return NotValidInputError
	}

	return nil
}
