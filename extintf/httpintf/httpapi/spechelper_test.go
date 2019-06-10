package httpapi_test

import (
	"encoding/json"
	"github.com/adamluzsi/toggler/services/security"
	"github.com/adamluzsi/testcase"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"

	. "github.com/adamluzsi/toggler/testing"
)


func UpdateFeatureFlagRolloutPercentage(t *testcase.T, featureFlagName string, rolloutPercentage int) {
	EnsureFlag(t, featureFlagName, rolloutPercentage)
}

func CreateSecurityTokenString(t *testcase.T) string {
	token, err := security.NewIssuer(GetStorage(t)).CreateNewToken(GetUniqUserID(t), nil, nil)
	require.Nil(t, err)
	return token.Token
}

func IsJsonRespone(t *testcase.T, r *httptest.ResponseRecorder, ptr interface{}) {
	require.Equal(t, "application/json", r.Header().Get(`Content-Type`))
	require.Nil(t, json.NewDecoder(r.Body).Decode(ptr))
}

func newRequest(t *testcase.T) *http.Request {

	return nil
}

// r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")