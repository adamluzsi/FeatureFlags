package usecases_test

import (
	. "github.com/adamluzsi/FeatureFlags/testing"
	"github.com/adamluzsi/testcase"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUseCases_SetFeatureFlagRolloutStrategyToUseDecisionLogicAPI(t *testing.T) {
	s := testcase.NewSpec(t)
	SetupSpecCommonVariables(s)
	SetupSpec(s)
	s.Parallel()

	subject := func(t *testcase.T) error {
		return GetProtectedUsecases(t).SetFeatureFlag(GetFeatureFlag(t))
	}

	s.When(`with valid values`, func(s *testcase.Spec) {
		s.Then(`it will be set/persisted`, func(t *testcase.T) {
			require.Nil(t, subject(t))

			require.Equal(t, GetFeatureFlag(t), FindStoredFeatureFlag(t))
		})
	})

}