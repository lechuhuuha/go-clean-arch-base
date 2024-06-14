package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	A string `validate:"default=hello,required"`
	B int    `validate:"default=1,required"`
	C int8   `validate:"default=2"`
	D int16  `validate:"default=3"`
	E int32  `validate:"default=4"`
	F int64  `validate:"default=5"`
	G uint   `validate:"default=6"`
	H uint8  `validate:"default=7"`
	I uint16 `validate:"default=8"`
	J uint32 `validate:"default=9"`
	K uint64 `validate:"default=10"`
	L bool   `validate:"default=true"`
}

func TestDefaultValidator(t *testing.T) {
	var vv args
	if err := Validate(&vv); err != nil {
		t.Errorf("DefaultValidator() error = %v", err)
		return
	}

	assert.Equal(t, &args{
		A: "hello",
		B: 1,
		C: 2,
		D: 3,
		E: 4,
		F: 5,
		G: 6,
		H: 7,
		I: 8,
		J: 9,
		K: 10,
		L: true,
	}, &vv)
}
