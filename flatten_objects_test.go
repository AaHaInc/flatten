package flatten

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFlattenObjects(t *testing.T) {
	cases := []struct {
		test   string
		want   map[string]interface{}
		prefix string
		style  SeparatorStyle
	}{
		{
			`{
				"foo": {
					"jim":"bean"
				},
				"fee": "bar",
				"objects":[{"a":{"b":{"c":"abc"}}},{"d":{"e":{"f":"def"}}}],
				"mixed": [{
					"foo":"bar",
					"deeper":{"foo":"baz"},
					"inner_list": [
						"a",
						"b",
						"c",
						{
							"d": "other",
							"e": "another"
						}
					]
				}],
				"scalars":["a","b"],
				"number": 1.4567,
				"bool":   true
			}`,
			map[string]interface{}{
				"foo.jim": "bean",
				"fee":     "bar",
				"number":  1.4567,
				"bool":    true,
				"scalars": []interface{}{"a", "b"},
				"objects": []interface{}{
					map[string]interface{}{"a.b.c": "abc"},
					map[string]interface{}{"d.e.f": "def"},
				},
				"mixed": []interface{}{
					map[string]interface{}{
						"foo":        "bar",
						"deeper.foo": "baz",
						"inner_list": []interface{}{
							"a",
							"b",
							"c",
							map[string]interface{}{
								"d": "other",
								"e": "another",
							},
						},
					},
				},
			},
			"",
			DotStyle,
		},
	}

	for i, test := range cases {
		var m interface{}
		err := json.Unmarshal([]byte(test.test), &m)
		if err != nil {
			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
			continue
		}
		got, err := FlattenObjects(m.(map[string]interface{}), test.prefix, test.style)
		if err != nil {
			t.Errorf("%d: failed to flatten: %v", i+1, err)
			continue
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: mismatch, got: %v wanted: %v", i+1, got, test.want)
		}
	}
}
