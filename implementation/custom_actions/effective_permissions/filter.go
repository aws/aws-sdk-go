package effective_permissions

import (
	"encoding/json"
	"net/url"
	"strings"
)

func filterTagConditions(conditions map[string]interface{}) map[string][]string {
	tags := map[string][]string{}

	for key, val := range conditions {
		if strings.Contains(key, PrincipalTag) {
			switch val.(type) {
			case []interface{}:
				interfaceSlice := val.([]interface{})
				newArray := make([]string, len(interfaceSlice))
				for _, obj := range interfaceSlice {
					newArray = append(newArray, obj.(string))
				}
				tags[key] = newArray
			default:
				tags[key] = []string{val.(string)}
			}
		}
	}

	return tags
}

func anyFailsInTe(conditions map[string][]string, userTags map[string]string) bool {
	return !anyFailsInTne(conditions, userTags)
}

func anyFailsInTne(conditions map[string][]string, userTags map[string]string) bool {
	found := 0
	for name, tags := range conditions {
		for _, tag := range tags {
			if val, ok := userTags[name]; ok && tag == val {
				found++
			}
		}
	}

	return found > 0
}

func statementFailsAnyCondition(statement Statement, tags map[string]string) (bool, error) {
	var tagEqualsCondition map[string][]string
	if val, ok := statement.Condition[ExactMatch]; ok {
		tagEqualsCondition = filterTagConditions(val)
	}

	var tagNotEqualsCondition map[string][]string
	if val, ok := statement.Condition[NegatedMatch]; ok {
		tagNotEqualsCondition = filterTagConditions(val)
	}

	return !(anyFailsInTe(tagEqualsCondition, tags) || anyFailsInTne(tagNotEqualsCondition, tags)), nil
}

func getRelevantStatements(statements []Statement, tags map[string]string) ([]Statement, error) {
	RelevantStatements := []Statement{}

	for _, statement := range statements {
		val, err := statementFailsAnyCondition(statement, tags)
		if err != nil {
			return nil, err
		}

		if !val { // if the statement doesnt doesn't fail add it to the list
			RelevantStatements = append(RelevantStatements, statement)
		}
	}

	return RelevantStatements, nil
}

func filterStatements(userPolicyRows []*MyPolicyDetail, tags map[string]string) ([]*OutputPolicyDetail, error) {
	var final []*OutputPolicyDetail

	for _, row := range userPolicyRows {
		encodedValue := *row.PolicyDocument

		decodedValue, err := url.QueryUnescape(encodedValue)
		if err != nil {
			return nil, err
		}

		var data FullStatement
		if err := json.Unmarshal([]byte(decodedValue), &data); err != nil {
			return nil, err
		}

		output, err := getRelevantStatements(data.Statement, tags)
		if err != nil {
			return nil, err
		}

		final = append(final, &OutputPolicyDetail{
			PolicyDocument: output,
			PolicyName:     row.PolicyName,
		})
	}

	return final, nil
}

func formatStatements(userPolicyRows []*MyPolicyDetail) ([]*OutputPolicyDetail, error) {
	var final []*OutputPolicyDetail

	for _, row := range userPolicyRows {
		encodedValue := *row.PolicyDocument

		decodedValue, err := url.QueryUnescape(encodedValue)
		if err != nil {
			return nil, err
		}

		var data FullStatement
		if err := json.Unmarshal([]byte(decodedValue), &data); err != nil {
			return nil, err
		}

		final = append(final, &OutputPolicyDetail{
			PolicyDocument: data.Statement,
			PolicyName:     row.PolicyName,
		})
	}

	return final, nil
}
