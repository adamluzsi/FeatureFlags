package rollouts_test

import (
	. "github.com/adamluzsi/FeatureFlags/testing"
	"github.com/adamluzsi/FeatureFlags/services/rollouts"
	"github.com/adamluzsi/testcase"
	"math/rand"
	"testing"
	"time"

	"github.com/adamluzsi/frameless/iterators"
	"github.com/stretchr/testify/require"
)

func TestRolloutManager(t *testing.T) {
	s := testcase.NewSpec(t)
	s.Parallel()
	SetupSpecCommonVariables(s)

	GetGeneratedRandomSeed := func(t *testcase.T) int64 {
		return t.I(`GeneratedRandomSeed`).(int64)
	}

	s.Let(`GeneratedRandomSeed`, func(t *testcase.T) interface{} {
		return time.Now().Unix()
	})

	manager := func(t *testcase.T) *rollouts.RolloutManager {
		return &rollouts.RolloutManager{
			Storage: GetStorage(t),
			RandSeedGenerator: func() int64 {
				return GetGeneratedRandomSeed(t)
			},
		}
	}

	s.Before(func(t *testcase.T) {
		require.Nil(t, GetStorage(t).Truncate(rollouts.FeatureFlag{}))
		require.Nil(t, GetStorage(t).Truncate(rollouts.Pilot{}))
	})

	s.Describe(`SetPilotEnrollmentForFeature`, func(s *testcase.Spec) {

		GetNewEnrollment := func(t *testcase.T) bool {
			return t.I(`NewEnrollment`).(bool)
		}

		subject := func(t *testcase.T) error {
			return manager(t).SetPilotEnrollmentForFeature(
				GetFeatureFlagName(t),
				GetExternalPilotID(t),
				GetNewEnrollment(t),
			)
		}

		s.Let(`NewEnrollment`, func(t *testcase.T) interface{} {
			return rand.Intn(2) == 0
		})

		findFlag := func(t *testcase.T) *rollouts.FeatureFlag {
			iter := GetStorage(t).FindAll(&rollouts.FeatureFlag{})
			require.NotNil(t, iter)
			require.True(t, iter.Next())
			var ff rollouts.FeatureFlag
			require.Nil(t, iter.Decode(&ff))
			require.False(t, iter.Next())
			require.Nil(t, iter.Err())
			return &ff
		}

		s.When(`no feature flag is seen ever before`, func(s *testcase.Spec) {
			s.Before(func(t *testcase.T) {
				require.Nil(t, GetStorage(t).Truncate(rollouts.FeatureFlag{}))
			})

			s.Then(`feature flag created`, func(t *testcase.T) {
				require.Nil(t, subject(t))

				flag := findFlag(t)
				require.Equal(t, GetFeatureFlagName(t), flag.Name)
				require.Equal(t, "", flag.Rollout.Strategy.URL)
				require.Equal(t, 0, flag.Rollout.Strategy.Percentage)
				require.Equal(t, GetGeneratedRandomSeed(t), flag.Rollout.RandSeedSalt)
			})

			s.Then(`pilot is enrollment is set for the feature is set`, func(t *testcase.T) {
				require.Nil(t, subject(t))

				flag := findFlag(t)
				pilot, err := GetStorage(t).FindFlagPilotByExternalPilotID(flag.ID, GetExternalPilotID(t))
				require.Nil(t, err)
				require.NotNil(t, pilot)
				require.Equal(t, GetNewEnrollment(t), pilot.Enrolled)
				require.Equal(t, GetExternalPilotID(t), pilot.ExternalID)
			})
		})

		s.When(`feature flag already configured`, func(s *testcase.Spec) {
			s.Before(func(t *testcase.T) {
				require.Nil(t, GetStorage(t).Save(GetFeatureFlag(t)))
			})

			s.Then(`flag is will not be recreated`, func(t *testcase.T) {
				require.Nil(t, subject(t))

				count, err := iterators.Count(GetStorage(t).FindAll(rollouts.FeatureFlag{}))
				require.Nil(t, err)
				require.Equal(t, 1, count)

				flag := findFlag(t)
				require.Equal(t, GetFeatureFlag(t), flag)
			})

			s.And(`pilot already exists`, func(s *testcase.Spec) {
				s.Before(func(t *testcase.T) {
					require.Nil(t, GetStorage(t).Save(GetPilot(t)))
				})

				s.And(`and pilot is has the opposite enrollment status`, func(s *testcase.Spec) {
					s.Let(`PilotEnrollment`, func(t *testcase.T) interface{} {
						return !GetNewEnrollment(t)
					})

					s.Then(`the original pilot is updated to the new enrollment status`, func(t *testcase.T) {
						require.Nil(t, subject(t))
						flag := findFlag(t)

						pilot, err := GetStorage(t).FindFlagPilotByExternalPilotID(flag.ID, GetExternalPilotID(t))
						require.Nil(t, err)

						require.NotNil(t, pilot)
						require.Equal(t, GetNewEnrollment(t), pilot.Enrolled)
						require.Equal(t, GetExternalPilotID(t), pilot.ExternalID)
						require.Equal(t, GetPilot(t), pilot)

						count, err := iterators.Count(GetStorage(t).FindAll(rollouts.Pilot{}))
						require.Nil(t, err)
						require.Equal(t, 1, count)
					})
				})

				s.And(`pilot already has the same enrollment status`, func(s *testcase.Spec) {
					s.Let(`PilotEnrollment`, func(t *testcase.T) interface{} {
						return GetNewEnrollment(t)
					})

					s.Then(`pilot remain the same`, func(t *testcase.T) {

						require.Nil(t, subject(t))
						ff := findFlag(t)

						pilot, err := GetStorage(t).FindFlagPilotByExternalPilotID(ff.ID, GetExternalPilotID(t))
						require.Nil(t, err)

						require.NotNil(t, pilot)
						require.Equal(t, GetNewEnrollment(t), pilot.Enrolled)
						require.Equal(t, GetExternalPilotID(t), pilot.ExternalID)

						count, err := iterators.Count(GetStorage(t).FindAll(rollouts.Pilot{}))
						require.Nil(t, err)
						require.Equal(t, 1, count)

					})
				})
			})

		})
	})

	s.Describe(`UpdateFeatureFlagRolloutPercentage`, func(s *testcase.Spec) {
		GetNewRolloutPercentage := func(t *testcase.T) int {
			return t.I(`NewRolloutPercentage`).(int)
		}

		subject := func(t *testcase.T) error {
			return manager(t).UpdateFeatureFlagRolloutPercentage(GetFeatureFlagName(t), GetNewRolloutPercentage(t))
		}

		s.When(`percentage less than 0`, func(s *testcase.Spec) {
			s.Let(`NewRolloutPercentage`, func(t *testcase.T) interface{} { return -1 + (rand.Intn(1024) * -1) })

			s.Then(`it will report error`, func(t *testcase.T) {
				require.Error(t, subject(t))
			})
		})

		s.When(`percentage greater than 100`, func(s *testcase.Spec) {
			s.Let(`NewRolloutPercentage`, func(t *testcase.T) interface{} { return 101 + rand.Intn(1024) })

			s.Then(`it will report error`, func(t *testcase.T) {
				require.Error(t, subject(t))
			})
		})

		s.When(`percentage is a number between 0 and 100`, func(s *testcase.Spec) {
			s.Let(`NewRolloutPercentage`, func(t *testcase.T) interface{} { return rand.Intn(101) })

			s.And(`feature flag was undefined until now`, func(s *testcase.Spec) {
				s.Before(func(t *testcase.T) {
					require.Nil(t, GetStorage(t).Truncate(rollouts.FeatureFlag{}))
				})

				s.Then(`feature flag entry created with the percentage`, func(t *testcase.T) {
					require.Nil(t, subject(t))
					flag, err := GetStorage(t).FindByFlagName(GetFeatureFlagName(t))
					require.Nil(t, err)
					require.NotNil(t, flag)

					require.Equal(t, GetFeatureFlagName(t), flag.Name)
					require.Equal(t, "", flag.Rollout.Strategy.URL)
					require.Equal(t, GetNewRolloutPercentage(t), flag.Rollout.Strategy.Percentage)
					require.Equal(t, GetGeneratedRandomSeed(t), flag.Rollout.RandSeedSalt)
				})
			})

			s.And(`feature flag is already exist with a different percentage`, func(s *testcase.Spec) {
				s.Let(`RolloutPercentage`, func(t *testcase.T) interface{} {
					for {
						roll := rand.Intn(101)
						if roll != GetNewRolloutPercentage(t) {
							return roll
						}
					}
				})

				s.Before(func(t *testcase.T) {
					require.Nil(t, GetStorage(t).Save(GetFeatureFlag(t)))
				})

				s.Then(`the same feature flag kept but updated to the new percentage`, func(t *testcase.T) {
					require.Nil(t, subject(t))
					flag, err := GetStorage(t).FindByFlagName(GetFeatureFlagName(t))
					require.Nil(t, err)
					require.NotNil(t, flag)

					require.Equal(t, GetFeatureFlag(t).ID, flag.ID)
					require.Equal(t, GetFeatureFlagName(t), flag.Name)
					require.Equal(t, "", flag.Rollout.Strategy.URL)
					require.Equal(t, GetNewRolloutPercentage(t), flag.Rollout.Strategy.Percentage)
					require.Equal(t, GetRolloutSeedSalt(t), flag.Rollout.RandSeedSalt)
				})
			})
		})
	})
}