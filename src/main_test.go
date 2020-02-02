package main

import "testing"

func TestGenerateFloor(t *testing.T) {
	res := generateFloor()
	if 16 < len(res) {
		t.Errorf("len(generateFloor()) = %d; want <= 16", len(res))
	} else {
		//t.Logf("len(generateFloor()) = %d; want <= 16", len(res))
	}

}

func TestGenerateEncounter(t *testing.T) {
	res := generateEncounter()
	if 0 >= len(res) {
		t.Errorf("len(generateEncounter) = %d; want > 0", len(res))
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
