package kata_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/kata"
	"github.com/stretchr/testify/assert"
)

func TestKata_Value_should_work_with_dependency_implementation(t *testing.T) {
	dependency := kata.NewDependencyImpl()
	k := kata.New(dependency)

	actual := k.Value()

	assert.Equal(t, 1, actual)
}
