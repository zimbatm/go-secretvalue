package secretvalue_test

import (
	"fmt"
	"testing"
	"github.com/zimbatm/go-secretvalue"
	"github.com/stretchr/testify/assert"
)

func TestSecretValue(t *testing.T) {
	secret := secretvalue.New("oauth-token")
	assert.False(t, secret.IsSet())
	secret.SetString("this is a secret")
	assert.True(t, secret.IsSet())

	// Test that logging doesn't leak the secret
	assert.Equal(t, "<secret:oauth-token>", fmt.Sprintf("%s", secret))
	assert.Equal(t, "<secret:oauth-token>", fmt.Sprintf("%v", secret))
	assert.Equal(t, "<secret:oauth-token>", fmt.Sprintf("%+v", secret))

	// Test that we get the same secret out
	assert.Equal(t, "this is a secret", secret.GetString())
}
