package baml

import (
	"fmt"
	"testing"

	"github.com/boundaryml/baml/engine/language_client_go/baml_go/serde"
	"github.com/boundaryml/baml/engine/language_client_go/pkg/cffi"
	"github.com/ghetzel/testify/require"
)

func TestEncodeFunctionArguments(t *testing.T) {
	client := NewClientRegistry()
	client.AddLlmClient("a", "b", map[string]any{"a": "b", "c": 1, "d": 2.2, "e": true})
	client.SetPrimaryClient("a")

	tests := []BamlFunctionArguments{
		{
			Kwargs: map[string]any{"a": "b", "c": 1, "d": 2.2, "e": true},
		},
		{
			Kwargs:         map[string]any{"a": "b", "c": 1, "d": 2.2, "e": true},
			ClientRegistry: client,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("EncodeFunctionArguments(%v)", test), func(t *testing.T) {
			_, err := test.Encode()
			require.NoError(t, err)
		})
	}

	t.Run("EncodeMap", func(t *testing.T) {
		test_value := map[string]string{
			"a": "b",
			"c": "d",
			"e": "f",
		}

		res, err := serde.EncodeValue(test_value)
		require.NoError(t, err)
		require.NotNil(t, res)
		mapVal, ok := res.Value.(*cffi.HostValue_MapValue)
		require.True(t, ok)
		require.NotNil(t, mapVal.MapValue)
		
		require.Len(t, mapVal.MapValue.Entries, 3)
		expected := map[string]string{
			"a": "b",
			"c": "d",
			"e": "f",
		}
		for _, entry := range mapVal.MapValue.Entries {
			strKey, ok := entry.Key.(*cffi.HostMapEntry_StringKey)
			require.True(t, ok)
			valVal, ok := entry.Value.Value.(*cffi.HostValue_StringValue)
			require.True(t, ok)
			require.Equal(t, expected[strKey.StringKey], valVal.StringValue)
		}
	})

	t.Run("EncodeMapWithOptional", func(t *testing.T) {
		foo, bar := "foo", "bar"
		test_value := map[string]*string{
			"a": &foo,
			"b": &bar,
			"c": nil,
		}

		res, err := serde.EncodeValue(test_value)
		require.NoError(t, err)
		require.NotNil(t, res)
		mapVal, ok := res.Value.(*cffi.HostValue_MapValue)
		require.True(t, ok)
		require.NotNil(t, mapVal.MapValue)

		require.Len(t, mapVal.MapValue.Entries, 3)
		for _, entry := range mapVal.MapValue.Entries {
			strKey, ok := entry.Key.(*cffi.HostMapEntry_StringKey)
			require.True(t, ok)
			switch strKey.StringKey {
			case "a":
				valVal, ok := entry.Value.Value.(*cffi.HostValue_StringValue)
				require.True(t, ok)
				require.Equal(t, "foo", valVal.StringValue)
			case "b":
				valVal, ok := entry.Value.Value.(*cffi.HostValue_StringValue)
				require.True(t, ok)
				require.Equal(t, "bar", valVal.StringValue)
			case "c":
				require.Nil(t, entry.Value.Value)
			default:
				t.Fatalf("unexpected key: %s", strKey.StringKey)
			}
		}
	})
}
