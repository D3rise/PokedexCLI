package context

import (
	"errors"
	"testing"
)

func TestSet(t *testing.T) {
	context := NewContext()

	cases := []struct {
		input struct {
			key   ContextKey
			value string
		}
		expected struct {
			key   ContextKey
			value string
		}
	}{
		{
			input: struct {
				key   ContextKey
				value string
			}{
				key:   "testKey",
				value: "testValue",
			},
			expected: struct {
				key   ContextKey
				value string
			}{
				key:   "testKey",
				value: "testValue",
			},
		},
		{
			input: struct {
				key   ContextKey
				value string
			}{
				key:   "testKey",
				value: "newTestValue",
			},
			expected: struct {
				key   ContextKey
				value string
			}{
				key:   "testKey",
				value: "newTestValue",
			},
		},
	}

	for _, c := range cases {
		context.Set(c.input.key, c.input.value)

		if context.kv[c.expected.key] != c.expected.value {
			t.Errorf("Expected key %s: %v, got %s: %v", c.expected.key, c.expected.value, c.expected.key, context.kv[c.expected.key])
		}
	}
}

func TestGet(t *testing.T) {
	context := NewContext()
	context.kv["testKey"] = "testValue"

	cases := []struct {
		input    ContextKey
		expected any
	}{
		{
			input:    "testKey",
			expected: "testValue",
		},
		{
			input:    "nonExistentKey",
			expected: nil,
		},
	}

	for _, c := range cases {
		value := context.Get(c.input)

		if c.expected != value {
			t.Errorf("Expected %v, got %v", c.expected, value)
		}
	}
}

func TestDel(t *testing.T) {
	context := NewContext()
	context.kv["testKey"] = "testValue"

	cases := []struct {
		input            ContextKey
		expected         any
		expectedKeyValue any
	}{
		{
			input:            "testKey",
			expected:         nil,
			expectedKeyValue: nil,
		},
		{
			input:            "nonExistentKey",
			expected:         errors.New("key does not exist"),
			expectedKeyValue: nil,
		},
	}

	for _, c := range cases {
		value := context.Del(c.input)
		var ok bool

		switch e := c.expected.(type) {
		case error:
			ok = e.Error() == value.Error()
			ok = ok && context.kv[c.input] == c.expectedKeyValue
		default:
			ok = e == value
			ok = ok && context.kv[c.input] == c.expectedKeyValue
		}

		if !ok {
			t.Errorf("Expected %v, got %v", c.expected, value)
		}
	}
}
