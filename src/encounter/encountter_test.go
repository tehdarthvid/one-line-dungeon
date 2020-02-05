package encounter

import "testing"

func TestGenerateEncounter(t *testing.T) {
	res := GenerateEncounter()
	if (0 >= len(res.Floor)) && (16 < len(res.Floor)) {
		t.Errorf("len(generateEncounter) = %d; want: 1..16", len(res.Floor))
	} else {
		//t.Logf("len(generateFloor()) = %d; want <= 16", len(res))
	}

}

func TestDeploymentEnv(t *testing.T) {
	res := getWhereAmI()
	if "" == res {
		t.Errorf("OLD_DEPLOYMENT_ENV = <empty>; want <not empty>")
	} else {
		//t.Logf("len(generateFloor()) = %d; want <= 16", len(res))
	}

}
