package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandId(t *testing.T) {
	require := require.New(t)
	id, err := RandID()
	require.NoError(err)
	t.Log(id)
	require.Equal(16, len(id))
}

func TestGetAuthKey(t *testing.T) {
	require := require.New(t)
	key := GetAuthKey("1234", 1488720000)
	require.Equal("ff6216b8258ba3ebb33f4c3af86ba2544d43badd81fbd08efdbe26410790975a", key)
}

func TestParseRangeNumbers(t *testing.T) {
	require := require.New(t)
	numbers, err := ParseRangeNumbers("2-5")
	require.NoError(err)
	require.Equal([]int64{2, 3, 4, 5}, numbers)

	numbers, err = ParseRangeNumbers("1")
	require.NoError(err)
	require.Equal([]int64{1}, numbers)

	numbers, err = ParseRangeNumbers("3-5,8")
	require.NoError(err)
	require.Equal([]int64{3, 4, 5, 8}, numbers)

	numbers, err = ParseRangeNumbers(" 3-5,8, 10-12 ")
	require.NoError(err)
	require.Equal([]int64{3, 4, 5, 8, 10, 11, 12}, numbers)

	_, err = ParseRangeNumbers("3-a")
	require.Error(err)
}
