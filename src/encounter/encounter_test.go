package encounter

import "testing"

func TestCreate(t *testing.T) {
	res := Create()
	if (0 >= len(res.Floor)) && (16 < len(res.Floor)) {
		t.Errorf("len(generateEncounter) = %d; want: 1..16", len(res.Floor))
	}
}

func TestGetWhereAmI(t *testing.T) {
	res := getWhereAmI()
	if "" == res {
		t.Errorf("OLD_TARGET_ENV = <empty>; want <not empty>")
	} else {
		t.Log("OLD_TARGET_ENV = " + res)
	}
}