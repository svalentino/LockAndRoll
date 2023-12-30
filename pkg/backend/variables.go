package backend

import (
	"encoding/json"
	"fmt"
)

type Variables map[string]string

// Custom errors
type VariableNameNotFoundError struct {
	Name string
}

type VariableNameAlreadyExistError struct {
	Name string
}

func (e *VariableNameNotFoundError) Error() string {
	return fmt.Sprintf("variable name not found: %s", e.Name)
}

func (e *VariableNameAlreadyExistError) Error() string {
	return fmt.Sprintf("variable name already exists: %s", e.Name)
}

// Methods
func (v *Variables) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(&v)
	if err != nil {
		return "", fmt.Errorf("failed to encode the variables into JSON: %w", err)
	}
	return string(jsonBytes), nil
}

func (v *Variables) FromJSON(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), &v)
	if err != nil {
		return fmt.Errorf("failed to decode the variables from JSON: %w", err)
	}
	return nil
}

func (v *Variables) Add(name string, value string) error {
	if _, ok := (*v)[name]; ok {
		return &VariableNameAlreadyExistError{Name: name}
	}
	(*v)[name] = value
	return nil
}

func (v *Variables) Remove(name string) error {
	if _, ok := (*v)[name]; ok {
		delete(*v, name)
		return nil
	}
	return &VariableNameNotFoundError{Name: name}
}

func (v *Variables) Set(name string, value string) error {
	if _, ok := (*v)[name]; ok {
		(*v)[name] = value
		return nil
	}
	return &VariableNameNotFoundError{Name: name}
}

func (v *Variables) Get(name string) (string, error) {
	if value, ok := (*v)[name]; ok {
		return value, nil
	}
	return "", &VariableNameNotFoundError{Name: name}
}
