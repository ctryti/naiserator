package resourcecreator_test

import (
	"github.com/nais/naiserator/pkg/resourcecreator"
	"github.com/nais/naiserator/pkg/test/fixtures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServiceAccount(t *testing.T) {
	app := fixtures.Application()
	svcAcc := resourcecreator.ServiceAccount(app)

	assert.Equal(t, app.Name, svcAcc.Name)
	assert.Equal(t, app.Namespace, svcAcc.Namespace)
}
