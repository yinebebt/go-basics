package methods

import (
	"testing"
)

func TestMethod(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.4f want %.4f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rec := Rectangle{10, 4}
		checkArea(t, &rec, 40.00)
	})

	t.Run("circle", func(t *testing.T) { // $go test -run TestMethod/circel | to run this subtest alone
		cir := Circle{10}
		checkArea(t, &cir, 314.1592653589793) // for this long precision number, use %g formatter
	})

	t.Run("test via table driven", func(t *testing.T) {
		// table driven test
		TestTable := []struct { // anonymous struct slice
			shp Shape
			wnt float64
		}{
			{&Rectangle{6, 8}, 48.0}, //we used pointer, since onluy *Rectangle does implemen the method 'Area', not Rectangle.
			{&Circle{2}, 6.28},
		}

		for _, val := range TestTable {
			got := val.shp.Area()
			if got != val.wnt {
				t.Errorf("got %.2f want %.2f", got, val.wnt)
			}
		}

	})

}
