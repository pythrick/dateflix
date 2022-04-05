package repository_test

import (
	"github.com/pythrick/dateflix/pkg/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericErr(t *testing.T) {
	given := repository.ErrEmptyConfig
	wantMsg := "required config is empty"
	assert.Equal(t, given.Error(), wantMsg)
}
