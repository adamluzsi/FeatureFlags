package specs

import (
	"context"
	"testing"

	"github.com/adamluzsi/frameless"
	"github.com/adamluzsi/frameless/fixtures"
	"github.com/adamluzsi/frameless/iterators"
	"github.com/adamluzsi/frameless/resources"
	"github.com/adamluzsi/frameless/resources/specs"
	"github.com/adamluzsi/testcase"
	"github.com/stretchr/testify/require"

	"github.com/toggler-io/toggler/domains/release"
)

type FlagFinderSpec struct {
	Subject interface {
		release.FlagFinder
		resources.Creator
		resources.Finder
		resources.Deleter
	}

	specs.FixtureFactory
}

func (spec FlagFinderSpec) Benchmark(b *testing.B) {
	b.Run(`FlagFinderSpec`, func(b *testing.B) {
		b.Skip(`TODO`)
	})
}

func (spec FlagFinderSpec) Test(t *testing.T) {
	s := testcase.NewSpec(t)

	featureName := fixtures.Random.String()

	s.Describe(`FlagFinderSpec`, func(s *testcase.Spec) {
		s.Around(func(t *testcase.T) func() {
			require.Nil(t, spec.Subject.DeleteAll(spec.ctx(), release.Flag{}))
			return func() {
				require.Nil(t, spec.Subject.DeleteAll(spec.ctx(), release.Flag{}))
			}
		})

		s.Describe(`FindReleaseFlagByName`, func(s *testcase.Spec) {
			subject := func(t *testcase.T) *release.Flag {
				ff, err := spec.Subject.FindReleaseFlagByName(spec.ctx(), featureName)
				require.Nil(t, err)
				return ff
			}

			s.Before(func(t *testcase.T) {
				require.Nil(t, spec.Subject.DeleteAll(spec.ctx(), release.Flag{}))
			})

			s.Then(`we receive back nil pointer`, func(t *testcase.T) {
				require.Nil(t, subject(t))
			})

			s.When(`we have a release flag already set`, func(s *testcase.Spec) {
				s.Let(`ff`, func(t *testcase.T) interface{} {
					flag := &release.Flag{Name: featureName}
					require.Nil(t, spec.Subject.Create(spec.ctx(), flag))
					t.Defer(spec.Subject.DeleteByID, spec.ctx(), *flag, flag.ID)
					return flag
				})

				s.Then(`searching for it returns the flag entity`, func(t *testcase.T) {
					ff := t.I(`ff`).(*release.Flag)
					actually, err := spec.Subject.FindReleaseFlagByName(spec.ctx(), ff.Name)
					require.Nil(t, err)
					require.Equal(t, ff, actually)
				})
			})
		})

		s.Describe(`FindReleaseFlagsByName`, func(s *testcase.Spec) {
			var subject = func(t *testcase.T) frameless.Iterator {
				flagEntriesIter := spec.Subject.FindReleaseFlagsByName(spec.ctx(), t.I(`flag names`).([]string)...)
				t.Defer(flagEntriesIter.Close)
				return flagEntriesIter
			}

			s.Before(func(t *testcase.T) {
				ctx := spec.ctx()
				for _, name := range []string{`A`, `B`, `C`} {
					var flag = release.Flag{Name: name}
					require.Nil(t, spec.Subject.Create(ctx, &flag))
					t.Defer(spec.Subject.DeleteByID, ctx, flag, flag.ID)
				}
			})

			mustContainName := func(t *testcase.T, ffs []release.Flag, name string) {
				for _, ff := range ffs {
					if ff.Name == name {
						return
					}
				}

				t.Fatalf(`flag name could not be found: %s`, name)
			}

			s.When(`we request flags that can be found`, func(s *testcase.Spec) {
				s.Let(`flag names`, func(t *testcase.T) interface{} {
					return []string{`A`, `B`, `C`}
				})

				s.Then(`it will return all of them`, func(t *testcase.T) {
					flagsIter := subject(t)

					var flags []release.Flag
					require.Nil(t, iterators.Collect(flagsIter, &flags))

					require.Equal(t, 3, len(flags))
					mustContainName(t, flags, `A`)
					mustContainName(t, flags, `B`)
					mustContainName(t, flags, `C`)
				})
			})

			s.When(`the requested flags only partially found`, func(s *testcase.Spec) {
				s.Let(`flag names`, func(t *testcase.T) interface{} {
					return []string{`A`, `B`, `D`}
				})

				s.Then(`it will return existing flags`, func(t *testcase.T) {
					flagsIter := subject(t)

					var flags []release.Flag
					require.Nil(t, iterators.Collect(flagsIter, &flags))

					t.Logf(`%#v`, flags)

					require.Equal(t, 2, len(flags))
					mustContainName(t, flags, `A`)
					mustContainName(t, flags, `B`)
				})
			})

			s.When(`none of the requested flags found`, func(s *testcase.Spec) {
				s.Let(`flag names`, func(t *testcase.T) interface{} {
					return []string{`R`, `O`, `F`, `L`}
				})

				s.Then(`it will return an empty iterator`, func(t *testcase.T) {
					flagsIter := subject(t)

					count, err := iterators.Count(flagsIter)
					require.Nil(t, err)
					require.Equal(t, 0, count)
				})
			})
		})
	})
}

func (spec FlagFinderSpec) ctx() context.Context {
	return spec.FixtureFactory.Context()
}
