package Department

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T){
	num := NewOrganization().Count()
	assert.Equal(t, num ,20)
}
