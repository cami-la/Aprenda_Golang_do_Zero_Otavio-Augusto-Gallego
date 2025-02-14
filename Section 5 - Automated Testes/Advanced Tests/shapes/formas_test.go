package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		t.Run("Rectangle", func(t *testing.T) {
			r := Rectangle{10, 12}
			expectedArea := float64(120)
			receivedArea := r.area()
			if expectedArea != receivedArea {
				t.Fatalf("Received %f but expected %f", receivedArea, expectedArea)
			}
		})

		t.Run("Circle", func(t *testing.T) {
			c := Circle{10}
			expectedArea := math.Pi * 100
			receivedArea := c.area()
			if expectedArea != receivedArea {
				t.Fatalf("Received %f but expected %f", receivedArea, expectedArea)
			}
		})
	})
}
