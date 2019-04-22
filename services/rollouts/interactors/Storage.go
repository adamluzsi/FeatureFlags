package interactors

import (
	"github.com/adamluzsi/FeatureFlags/services/rollouts"
	"github.com/adamluzsi/frameless"
	"github.com/adamluzsi/frameless/resources/specs"
)

type Storage interface {
	specs.Save
	specs.FindByID
	specs.Truncate
	specs.DeleteByID

	FlagFinder
	PilotFinder
}

type FlagFinder interface {
	FindByFlagName(name string) (*rollouts.FeatureFlag, error)
}

type PilotFinder interface {
	FindPilotByFeatureFlagIDAndPublicPilotID(FeatureFlagID, ExternalPublicPilotID string) (*rollouts.Pilot, error)
	FindPilotsByFeatureFlag(*rollouts.FeatureFlag) frameless.Iterator
}
