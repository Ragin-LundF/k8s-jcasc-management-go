package kubectl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfKubectlOutputContainsValueForFieldWithSpaces(t *testing.T) {
	var kubeCtlOutput = `LOG    NAMER    TEST    NAME
logvalue    namerentry    test    namespace-name
logvalue2    namerentry2    test2    namespace
logvalue3    namerentry3    test3    namespace-value`

	exists := CheckIfKubectlOutputContainsValueForField(kubeCtlOutput, "NAME", "namespace")

	assert.True(t, exists)
}

func TestCheckIfKubectlOutputContainsValueForFieldWithTabs(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace
logvalue3	namerentry3	test3	namespace-value`

	exists := CheckIfKubectlOutputContainsValueForField(kubeCtlOutput, "NAME", "namespace")

	assert.True(t, exists)
}

func TestCheckIfKubectlOutputContainsValueForFieldWithTabsAndWrongNamespace(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	exists := CheckIfKubectlOutputContainsValueForField(kubeCtlOutput, "NAME", "namespace")

	assert.False(t, exists)
}

func TestFindFieldValuesInKubectlOutput(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	fieldValues, err := FindFieldValuesInKubectlOutput(kubeCtlOutput, "NAME")

	assert.NoError(t, err)
	assert.Len(t, fieldValues, 3)
	assert.Equal(t, "namespace-name", fieldValues[0])
	assert.Equal(t, "namespace-test", fieldValues[1])
	assert.Equal(t, "namespace-value", fieldValues[2])
}

func TestFindFieldIndexInKubectlOutput(t *testing.T) {
	var kubeCtlOutput = `LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	lineIndex, fieldIndex, err := FindFieldIndexInKubectlOutput(kubeCtlOutput, "NAME")

	assert.NoError(t, err)
	assert.Equal(t, 0, lineIndex)
	assert.Equal(t, 3, fieldIndex)
}

func TestFindFieldIndexInKubectlOutputWithAdditionalLines(t *testing.T) {
	var kubeCtlOutput = `Kubectl can write additional things here.

LOG	NAMER	TEST	NAME
logvalue	namerentry	test	namespace-name
logvalue2	namerentry2	test2	namespace-test
logvalue3	namerentry3	test3	namespace-value`

	lineIndex, fieldIndex, err := FindFieldIndexInKubectlOutput(kubeCtlOutput, "NAME")

	assert.NoError(t, err)
	assert.Equal(t, 2, lineIndex)
	assert.Equal(t, 3, fieldIndex)
}
