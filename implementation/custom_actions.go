package implementation

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/implementation/custom_actions"
)

var customActionsMap = map[string]func(map[string]interface{}) (map[string]interface{}, error){
	"GetUserEffectivePermissions":    GetUserEffectivePermissions,
	"GetUserOrgEffectivePermissions": GetOrgEffectivePermissions,
}

func GetUserEffectivePermissions(parameters map[string]interface{}) (map[string]interface{}, error) {
	action := custom_actions.NewEffectivePermissions(parameters)
	out, err := action.GetUserEffectivePermissions()
	if err != nil {
		return nil, err
	}

	outMarshaled, err := json.Marshal(out)
	if err != nil {
		return nil, errors.New("failed to marshal output")
	}

	output := make(map[string]interface{})
	if err := json.Unmarshal(outMarshaled, &output); err != nil {
		return nil, errors.New("failed to unmarshal output")
	}
	return output, nil
}

func GetOrgEffectivePermissions(parameters map[string]interface{}) (map[string]interface{}, error) {
	action := custom_actions.NewEffectivePermissions(parameters)
	out, err := action.GetOrgEffectivePermissions()
	if err != nil {
		return nil, err
	}

	outMarshaled, err := json.Marshal(out)
	if err != nil {
		return nil, errors.New("failed to marshal output")
	}

	output := make(map[string]interface{})
	if err := json.Unmarshal(outMarshaled, &output); err != nil {
		return nil, errors.New("failed to unmarshal output")
	}
	return output, nil
}
