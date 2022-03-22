package collections

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	set := NewSet("A", "B", "C")

	require.True(t, set.Contains("B"))
	require.False(t, set.Contains("D"))

	set.Remove("A")
	require.False(t, set.Contains("A"))

	set.Add("D")
	require.True(t, set.Contains("D"))
}

func TestSet_Marshal(t *testing.T) {
	var data Set[string]

	err := json.Unmarshal([]byte(`["A", "B", "C"]`), &data)
	require.Nil(t, err)

	require.True(t, data.Contains("A"))
	require.True(t, data.Contains("B"))
	require.True(t, data.Contains("C"))
	require.False(t, data.Contains("D"))

	serialized, err := json.Marshal(data)
	require.Nil(t, err)

	re := regexp.MustCompile(`^\["(\w)","(\w)","(\w)"]$`)
	matches := re.FindStringSubmatch(string(serialized))

	require.Equal(t, 4, len(matches))
	for i := range matches[1:] {
		require.Contains(t, []string{"A", "B", "C"}, matches[1:][i])
	}
}
