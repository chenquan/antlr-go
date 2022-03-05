package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		val := Calculate(`a=1`)
		assert.EqualValues(t, "1", val)

		val = Calculate(`1+1`)
		assert.EqualValues(t, "2", val)

		val = Calculate(`1+1+2`)
		assert.EqualValues(t, "4", val)

		val = Calculate(`1+1*3`)
		assert.EqualValues(t, "4", val)

		val = Calculate(`1+1*3-(2-1)`)
		assert.EqualValues(t, "3", val)

		val = Calculate(`1+1*3-(2-1)*2`)
		assert.EqualValues(t, "2", val)

		val = Calculate(`1+1*3-(2-2)/2`)
		assert.EqualValues(t, "4", val)

	})
	t.Run("2", func(t *testing.T) {
		val := Calculate(`a=1.1
b=1+a
a+b



`)
		assert.EqualValues(t, "3.2", val)
	})
}
