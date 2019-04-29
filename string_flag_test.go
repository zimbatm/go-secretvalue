package secretvalue_test

import (
	"testing"
	"flag"

	"github.com/zimbatm/go-secretvalue"
	"github.com/stretchr/testify/assert"
)

func TestStringFlag(t *testing.T) {
	token := secretvalue.New("token")
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&secretvalue.StringFlag{token}, "token", "Secret token")
	fs.Parse([]string{"-token", "1234567"})

	assert.Equal(t, "1234567", token.GetString())
	assert.Equal(t, "<secret:token>", token.String())
}
