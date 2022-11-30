package oty

import (
	"encoding/json"
	"fmt"
	"sigs.k8s.io/yaml"
)

func ObjectTOYaml(data interface{}) (string, error) {
	if data == nil {
		return "", nil
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("json serialization failed err:%v", err)
	}
	dataYaml, err := yaml.JSONToYAML(dataJSON)
	if err != nil {
		return "", fmt.Errorf("yaml serialization failed err:%v", err)
	}
	return string(dataYaml), nil
}
