package main

import (
	"testing"
)

func Test123(t *testing.T) {
	finder := Constructor()
	finder.AddNum(1)
	finder.AddNum(2)

	actual1 := finder.FindMedian()

	if actual1 != 1.5 {
		t.Fatalf("actual = %f, expected = 1.5", actual1)
	}

	finder.AddNum(3)

	actual2 := finder.FindMedian()

	if actual2 != 2.0 {
		t.Fatalf("actual = %f, expected = 2.0", actual2)
	}
}

func TestM12345(t *testing.T) {
	finder := Constructor()
	finder.AddNum(-1)

	actual1 := finder.FindMedian()
	finder.AddNum(-2)

	actual2 := finder.FindMedian()
	finder.AddNum(-3)

	actual3 := finder.FindMedian()
	finder.AddNum(-4)

	actual4 := finder.FindMedian()
	finder.AddNum(-5)

	actual5 := finder.FindMedian()

	if actual1 != -1.00000 {
		t.Fatalf("actual = %f, expected = -1.00000", actual1)
	}
	if actual2 != -1.50000 {
		t.Fatalf("actual = %f, expected = -1.50000", actual2)
	}
	if actual3 != -2.00000 {
		t.Fatalf("actual = %f, expected = -2.00000", actual3)
	}
	if actual4 != -2.50000 {
		t.Fatalf("actual = %f, expected = -2.50000", actual4)
	}
	if actual5 != -3.00000 {
		t.Fatalf("actual = %f, expected = -3.00000", actual5)
	}
}
