package fieldgroups

import (
	"errors"
	"github.com/creasty/defaults"
)

// TeamSyncingFieldGroup represents the TeamSyncingFieldGroup config fields
type TeamSyncingFieldGroup struct {
	FeatureNonsuperuserTeamSyncingSetup bool   `default:"false" validate:""`
	FeatureTeamSyncing                  bool   `default:"false" validate:""`
	TeamResyncStaleTime                 string `default:"30m" validate:"customValidateTimePattern"`
}

// NewTeamSyncingFieldGroup creates a new TeamSyncingFieldGroup
func NewTeamSyncingFieldGroup(fullConfig map[string]interface{}) (FieldGroup, error) {
	newTeamSyncingFieldGroup := &TeamSyncingFieldGroup{}
	defaults.Set(newTeamSyncingFieldGroup)

	if value, ok := fullConfig["FEATURE_NONSUPERUSER_TEAM_SYNCING_SETUP"]; ok {
		newTeamSyncingFieldGroup.FeatureNonsuperuserTeamSyncingSetup, ok = value.(bool)
		if !ok {
			return newTeamSyncingFieldGroup, errors.New("FEATURE_NONSUPERUSER_TEAM_SYNCING_SETUP must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_TEAM_SYNCING"]; ok {
		newTeamSyncingFieldGroup.FeatureTeamSyncing, ok = value.(bool)
		if !ok {
			return newTeamSyncingFieldGroup, errors.New("FEATURE_TEAM_SYNCING must be of type bool")
		}
	}
	if value, ok := fullConfig["TEAM_RESYNC_STALE_TIME"]; ok {
		newTeamSyncingFieldGroup.TeamResyncStaleTime, ok = value.(string)
		if !ok {
			return newTeamSyncingFieldGroup, errors.New("TEAM_RESYNC_STALE_TIME must be of type string")
		}
	}

	return newTeamSyncingFieldGroup, nil
}
