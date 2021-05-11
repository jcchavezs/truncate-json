package truncate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testPayload = `{"name": "Nitin","language": "python","repositories": ["pythonagent"]}`

func TestTruncateJSONBasicTruncation(t *testing.T) {
	tp, err := truncateJSON(testPayload, 3)
	require.NoError(t, err)
	assert.Equal(t, "{\"hypertrace\": \"truncated\", \"n\": \"\"}", tp)
}

func TestTruncateJSONBeforeSemicolon(t *testing.T) {
	tp, err := truncateJSON(testPayload, 7)
	require.NoError(t, err)
	assert.Equal(t, "{\"hypertrace\": \"truncated\", \"name\": \"\"}", tp)
}

func TestTruncateAfterSemicolon(t *testing.T) {
	tp, err := truncateJSON(testPayload, 10)
	require.NoError(t, err)
	assert.Equal(t, "{\"hypertrace\": \"truncated\", \"name\": \"\"}", tp)
}

func TestTruncateEmptyArray(t *testing.T) {
	tp, err := truncateJSON(testPayload, 55)
	require.NoError(t, err)
	assert.Equal(t, "{\"hypertrace\": \"truncated\", \"name\": \"Nitin\",\"language\": \"python\",\"repositories\": []}", tp)
}

func TestTruncateOneQuoteArray(t *testing.T) {
	tp, err := truncateJSON(testPayload, 56)
	require.NoError(t, err)
	assert.Equal(t, "{\"hypertrace\": \"truncated\", \"name\": \"Nitin\",\"language\": \"python\",\"repositories\": [\"\"]}", tp)
}
