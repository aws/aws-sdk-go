package effective_permissions

type FullStatement struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Effect    string                            `json:"Effect"`
	Action    interface{}                       `json:"Action"`
	Resource  interface{}                       `json:"Resource"`
	Condition map[string]map[string]interface{} `json:"Condition,omitempty"`
}
