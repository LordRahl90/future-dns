package requests

import (
	"os"
	"testing"

	"dns/domains/maths"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	code := 1
	defer func() {
		os.Exit(code)
	}()
	code = m.Run()
}

func TestConvertRequestToDomainEntity(t *testing.T) {
	table := []struct {
		name   string
		arg    Request
		expErr bool
		errMsg string
		exp    *maths.Request
	}{
		{
			name:   "all_empty",
			arg:    Request{},
			expErr: true,
			errMsg: `invalid X coordinate value: ()`,
		},
		{
			name: "invalid_X",
			arg: Request{
				CoordX: "x",
			},
			expErr: true,
			errMsg: `invalid X coordinate value: (x)`,
		},
		{
			name: "invalid_Y",
			arg: Request{
				CoordX: "0.0",
				CoordY: "y",
			},
			expErr: true,
			errMsg: `invalid Y coordinate value: (y)`,
		},
		{
			name: "invalid_Z",
			arg: Request{
				CoordX: "5.0",
				CoordY: "0.5",
				CoordZ: "z",
			},
			expErr: true,
			errMsg: `invalid Z coordinate value: (z)`,
		},
		{
			name: "invalid_velocity",
			arg: Request{
				CoordX:   "0.2",
				CoordY:   "5.0",
				CoordZ:   "2.2",
				Velocity: "v",
			},
			expErr: true,
			errMsg: `invalid velocity value: (v)`,
		},
		{
			name: "valid_request",
			arg: Request{
				CoordX:   "0.2",
				CoordY:   "5.0",
				CoordZ:   "2.2",
				Velocity: "10.4",
			},
			expErr: false,
			exp: &maths.Request{
				CoordX:   0.2,
				CoordY:   5.0,
				CoordZ:   2.2,
				Velocity: 10.4,
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.arg.ToDomainRequest()
			if tt.expErr {
				require.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.exp.CoordX, got.CoordX)
				require.Equal(t, tt.exp.CoordY, got.CoordY)
				require.Equal(t, tt.exp.CoordZ, got.CoordZ)
				require.Equal(t, tt.exp.Velocity, got.Velocity)
			}
		})
	}
}
