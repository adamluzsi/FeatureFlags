package security_test

import (
	"testing"

	"github.com/adamluzsi/testcase"
	"github.com/stretchr/testify/require"

	"github.com/toggler-io/toggler/domains/security"
	sh "github.com/toggler-io/toggler/spechelper"
)

func TestDoorkeeper(t *testing.T) {
	s := sh.NewSpec(t)
	s.Parallel()

	s.Let(`doorkeeper`, func(t *testcase.T) interface{} {
		return security.NewDoorkeeper(sh.StorageGet(t))
	})

	s.Let(`token`, func(t *testcase.T) interface{} {
		issuer := security.NewIssuer(sh.StorageGet(t))
		textToken, token, err := issuer.CreateNewToken(sh.ContextGet(t), sh.ExampleUniqueUserID(t), nil, nil)
		t.Set(`text token`, textToken)
		require.Nil(t, err)
		return token
	})

	SpecDoorkeeperVerifyTextToken(s)
}

func SpecDoorkeeperVerifyTextToken(s *testcase.Spec) {
	var getToken = func(t *testcase.T) *security.Token {
		return t.I(`token`).(*security.Token)
	}

	var getTextToken = func(t *testcase.T) string {
		getToken(t) // trigger setup
		return t.I(`text token`).(string)
	}

	s.Describe(`VerifyTextToken`, func(s *testcase.Spec) {
		subject := func(t *testcase.T) (bool, error) {
			return doorkeeper(t).VerifyTextToken(sh.ContextGet(t), getTextToken(t))
		}

		onSuccess := func(t *testcase.T) bool {
			accepted, err := subject(t)
			require.Nil(t, err)
			return accepted
		}

		s.When(`token is a known resource`, func(s *testcase.Spec) {
			s.Before(func(t *testcase.T) {
				persisted, err := sh.StorageGet(t).SecurityToken(sh.ContextGet(t)).FindByID(sh.ContextGet(t), &security.Token{}, getToken(t).ID)
				require.Nil(t, err)
				require.True(t, persisted)
			})

			s.Then(`it will verify and accept it`, func(t *testcase.T) {
				require.True(t, onSuccess(t))
			})
		})

		s.When(`token is unknown`, func(s *testcase.Spec) {
			s.Before(func(t *testcase.T) {
				require.Nil(t, sh.StorageGet(t).SecurityToken(sh.ContextGet(t)).DeleteByID(sh.ContextGet(t), getToken(t).ID))
			})

			s.Then(`it will reject it`, func(t *testcase.T) {
				require.False(t, onSuccess(t))
			})
		})

	})
}

func doorkeeper(t *testcase.T) *security.Doorkeeper {
	return t.I(`doorkeeper`).(*security.Doorkeeper)
}
