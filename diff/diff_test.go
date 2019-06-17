package diff

import (
	"testing"
)

func TestDiffStructures(t *testing.T) {

	t.Run("It should correctly diff simple map", func(t *testing.T) {
		mockLeft := map[string]interface{}{"foo": "bar", "fizz": 5}
		mockRight := map[string]interface{}{"foo": "bar", "fizz": 5}
		d, err := DiffStructures(mockLeft, mockRight)
		if d != true {
			t.Error("should be true", err)
		}
	})

	t.Run("It should correctly diff simple nested map", func(t *testing.T) {
		mockLeft := map[string]interface{}{
			"foo": "back",
			"boo": map[string]interface{}{
				"fizz": 5,
			},
		}
		mockRight := map[string]interface{}{
			"foo": "back",
			"boo": map[string]interface{}{
				"fizz": 5,
			},
		}
		d, err := DiffStructures(mockLeft, mockRight)
		if d != true {
			t.Error("should be true", err)
		}
	})

	t.Run("It should correctly diff nested map with varying properties", func(t *testing.T) {
		mockLeft := map[string]interface{}{
			"foo": "back",
			"boo": map[string]interface{}{
				"fizz": 5,
			},
		}
		mockRight := map[string]interface{}{
			"foo": "back",
			"boo": map[string]interface{}{
				"fizz": 5,
			},
			"moo": []string{"cake", "bake", "makes"},
		}
		d, err := DiffStructures(mockLeft, mockRight)
		if d != false {
			t.Error("should be false", err)
		}
	})

	t.Run("It should correctly diff multi nested map with varying properties", func(t *testing.T) {
		mockLeft := map[string]interface{}{
			"foo": "back",
			"boo": map[string]interface{}{
				"fizz": 5,
			},
			"goo": map[string]interface{}{
				"cake": map[string]interface{}{
					"dootie": 5,
					"yes":    true,
					"no":     false,
				},
			},
		}
		mockRight := map[string]interface{}{
			"foo": "back",
			"boo": map[string]interface{}{
				"fizz": 5,
			},
			"moo": []string{"cake", "bake", "makes"},
		}
		d, err := DiffStructures(mockLeft, mockRight)
		if d != false {
			t.Error("should be false", err)
		}
	})

	t.Run("It should correctly diff maps with arrays", func(t *testing.T) {
		mockLeft := map[string]interface{}{
			"foo": "back",
			"boo": []interface{}{
				"fizz", "buzz", "putz",
			},
		}
		mockRight := map[string]interface{}{
			"foo": "back",
			"boo": []interface{}{
				"fizz", "buzz", "putz",
			},
		}
		d, err := DiffStructures(mockLeft, mockRight)
		if d != true {
			t.Error("should be true", err)
		}
	})

	t.Run("It should correctly diff maps with arrays", func(t *testing.T) {
		mockLeft := map[string]interface{}{
			"foo": "back",
			"boo": []interface{}{
				"fizz", "buzz", "putz",
			},
		}
		mockRight := map[string]interface{}{
			"foo": "back",
			"boo": []interface{}{
				"cutz", "buzz", "putz",
			},
		}
		d, err := DiffStructures(mockLeft, mockRight)
		if d != false {
			t.Error("should be false", err)
		}
	})

	t.Run("It should correctly diff maps with arrays with the correct elements but the wrong orders", func(t *testing.T) {
		mockLeft := map[string]interface{}{
			"foo": "back",
			"boo": []interface{}{
				"fizz", "buzz", "putz",
			},
		}
		mockRight := map[string]interface{}{
			"foo": "back",
			"boo": []interface{}{
				"buzz", "fizz", "putz",
			},
		}
		d, err := DiffStructures(mockLeft, mockRight)
		if d != false {
			t.Error("should be false", err)
		}
	})

	// t.Run("It should correctly diff a massive object", func(t *testing.T) {
	// 	mockLeft := map[string]interface{}{
	// 		"foo": "back",
	// 		"boo": []interface{}{
	// 			"fizz", "buzz", "putz",
	// 		},
	// 		"moo": "boo",
	// 	}
	// })
}
