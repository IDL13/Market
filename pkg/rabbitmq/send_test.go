package rabbitmq

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	t1 string
	t2 int
	t3 int
	t4 string
}

func TestWriteMessage(t *testing.T) {
	tests := []struct {
		name      string
		inputBody TestData
		response  int
	}{{
		name: "OK",
		inputBody: TestData{
			t1: "test1",
			t2: 1,
			t3: 1,
			t4: "test1",
		},
		response: 1,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rabbitMq := New("key_stream")
			marshalJson, err := json.Marshal(test.inputBody)
			if err != nil {
				panic(err)
			}
			res := rabbitMq.WriteData(marshalJson)
			assert.Equal(t, test.response, res)
		})
	}
}
