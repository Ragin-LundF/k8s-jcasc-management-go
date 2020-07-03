package kubectl

import (
	"testing"
)

func TestCheckIfKubectlOutputContainsValueForFieldWithSpaces(t *testing.T) {
	var kubeCtlOutput = `LOG    NAMER    TEST    NAME
logvalue    namerentry    test    namespace-name
logvalue2    namerentry2    test2    namespace
logvalue3    namerentry3    test3    namespace-value`

	exists := CheckIfKubectlOutputContainsValueForField(kubeCtlOutput, "NAME", "namespace")
	if exists {
		t.Log("Success. Found output with spaces")
	} else {
		t.Error("Failed. Can not find value with spaces")
	}
}

func TestCheckIfKubectlOutputContainsValueForFieldWithTabs(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace
logvalue3	namerentry3	test3	namespace-value`

	exists := CheckIfKubectlOutputContainsValueForField(kubeCtlOutput, "NAME", "namespace")
	if exists {
		t.Log("Success. Found output with tabs")
	} else {
		t.Error("Failed. Can not find value with tabs")
	}
}

func TestCheckIfKubectlOutputContainsValueForFieldWithTabsAndWrongNamespace(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	exists := CheckIfKubectlOutputContainsValueForField(kubeCtlOutput, "NAME", "namespace")
	if exists {
		t.Error("Failed. Found value that should not exist.")
	} else {
		t.Log("Success. Can not find wrong value")
	}
}

func TestFindFieldValuesInKubectlOutput(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	fieldValues, err := FindFieldValuesInKubectlOutput(kubeCtlOutput, "NAME")
	if err != nil {
		t.Error("Failed. An error happened.")
	}
	if len(fieldValues) == 3 && fieldValues[0] == "namespace-name" && fieldValues[2] == "namespace-value" {
		t.Log("Success. Found expected data.")
	} else {
		t.Errorf("Failed. Did not find expected values. Len [%v] Value1 [%v].", len(fieldValues), fieldValues[0])
	}
}

func TestFindFieldIndexInKubectlOutput(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	lineIndex, fieldIndex, err := FindFieldIndexInKubectlOutput(kubeCtlOutput, "NAME")
	if err != nil {
		t.Error("Failed. An error happened.")
	}
	if lineIndex == 0 && fieldIndex == 3 {
		t.Log("Success. Found expected data.")
	} else {
		t.Errorf("Failed. Can not find expected data. LineIndex [%v] FieldIndex [%v]", lineIndex, fieldIndex)
	}
}

func TestFindFieldIndexInKubectlOutputWithAdditionalLines(t *testing.T) {
	var kubeCtlOutput = `Kubectl can write additional things here.

LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	lineIndex, fieldIndex, err := FindFieldIndexInKubectlOutput(kubeCtlOutput, "NAME")
	if err != nil {
		t.Error("Failed. An error happened.")
	}
	if lineIndex == 2 && fieldIndex == 3 {
		t.Log("Success. Found expected data.")
	} else {
		t.Errorf("Failed. Can not find expected data. LineIndex [%v] FieldIndex [%v]", lineIndex, fieldIndex)
	}
}
