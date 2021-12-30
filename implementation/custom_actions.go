package implementation

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/implementation/custom_actions/effective_permissions"
)

var customActionsMap = map[string]func(map[string]interface{}) (map[string]interface{}, error){
	"GetUserEffectivePermissions":    GetUserEffectivePermissions,
	"GetRoleEffectivePermissions":    GetRoleEffectivePermissions,
	"GetUserOrgEffectivePermissions": GetOrgEffectivePermissions,
}

func GetUserEffectivePermissions(parameters map[string]interface{}) (map[string]interface{}, error) {
	client, err := effective_permissions.NewEffectivePermissions(parameters)
	if err != nil {
		return nil, err
	}
	out, err := client.GetUserEffectivePermissions()
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

func GetRoleEffectivePermissions(parameters map[string]interface{}) (map[string]interface{}, error) {
	client, err := effective_permissions.NewEffectivePermissions(parameters)
	if err != nil {
		return nil, err
	}

	out, err := client.GetRoleEffectivePermissions()
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
	client, err := effective_permissions.NewEffectivePermissions(parameters)
	if err != nil {
		return nil, err
	}

	out, err := client.GetOrgEffectivePermissions()
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
