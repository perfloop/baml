package baml

import (
	"testing"

	"github.com/boundaryml/baml/engine/language_client_go/baml_go/serde"
)

func BenchmarkEncodeMap(b *testing.B) {
	test_value := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
		"k5": "v5",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := serde.EncodeValue(test_value)
		if err != nil {
			b.Fatalf("failed to encode map: %v", err)
		}
	}
}

func BenchmarkEncodeEmptyMap(b *testing.B) {
	test_value := map[string]string{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := serde.EncodeValue(test_value)
		if err != nil {
			b.Fatalf("failed to encode map: %v", err)
		}
	}
}

func BenchmarkEncodeSingleMap(b *testing.B) {
	test_value := map[string]string{
		"k1": "v1",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := serde.EncodeValue(test_value)
		if err != nil {
			b.Fatalf("failed to encode map: %v", err)
		}
	}
}
