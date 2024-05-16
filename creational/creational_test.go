package creational

import "testing"

func TestSingleton(t *testing.T) {
	// write a test where two instances of a singleton are created
	// and the values are changed but they should not be different

	i1 := GetInstance()
	i2 := GetInstance()

	i1.value = 3
	i2.value = 4

	t.Run("test", func(t *testing.T) {
		if i1.value != i2.value {
			t.Errorf("expected %d got %d", i1.value, i2.value)
		}
	})
}
