package maths

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	sectorID = 10.0

	ms IMathService
)

func TestMain(m *testing.M) {
	code := 1
	defer func() {
		os.Exit(code)
	}()

	ms = New(sectorID)
	code = m.Run()
}

func TestCalculate(t *testing.T) {
	table := []struct {
		name string
		args *Request
		exp  float64
	}{
		{
			name: "all_zero",
			args: &Request{},
			exp:  0.0,
		},
	}

	for _, tt := range table {
		ctx := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			got := ms.Calculate(ctx, tt.args)
			require.Equal(t, tt.exp, got)
		})
	}
}
