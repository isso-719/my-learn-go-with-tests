package StructsMethodsAndInterfaces

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}

//
// 最初に作ったやつ
//

//func TestPerimeter(t *testing.T) {
//	rectangle := Rectangle{10.0, 10.0}
//	got := Perimeter(rectangle)
//	want := 40.0
//
//	if got != want {
//		t.Errorf("got %.2f want %.2f", got, want)
//	}
//}

//func TestArea(t *testing.T) {
//	checkArea := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %.2g want %.2g", got, want)
//		}
//	}
//
//	t.Run("rectangle", func(t *testing.T) {
//		rectangle := Rectangle{12, 6}
//		want := 72.0
//		checkArea(t, rectangle, want)
//	})
//
//	t.Run("circle", func(t *testing.T) {
//		circle := Circle{10}
//		want := 314.1592653589793
//		checkArea(t, circle, want)
//	})
//}
