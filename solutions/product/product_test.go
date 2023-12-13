package product_test

import (
	"math"
	"testing"

	"github.com/bitfield/product"
)

func TestProductOfInts2And3Is6(t *testing.T) {
	t.Parallel()
	want := 6
	got := product.Product(2, 3)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestProductOfFloats1Point6And2Point3Is3Point68(t *testing.T) {
	t.Parallel()
	want := 3.68
	got := product.Product(1.6, 2.3)
	if math.Abs(want-got) > 0.0001 {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestProductOfComplex2Plus3iAnd3Plus2iIs0Plus13i(t *testing.T) {
	t.Parallel()
	want := 0 + 13i
	got := product.Product(2+3i, 3+2i)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
