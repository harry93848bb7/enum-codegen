# SQL Type and Enum Code Generator
This package serves the purpose to generate go types and enums to be used when dealing with sql drivers. This package allows you to define a `yaml` specification with your sql enums and will generate the go types, enums and tests with attached methods appropriate when dealing with sql or apis:

- `String() string`
- `Scan(value interface{}) error`
- `Value() (driver.Value, error)`
- `MarshalJSON() ([]byte, error)`
- `UnmarshalJSON(b []byte) error`

### Example
1. `go get github.com/harry93848bb7/enum-codegen/cmd/enum-codegen`
2. Define a enum-codegen specification file with the types and enums
```yaml
types: 
  - name: 'BusinessRole'
    description: 'defines a role of a user to a business'
    enums:
      - BusinessRoleUnknown
      - BusinessRoleAdministrator
      - BusinessRoleDeveloper
      - BusinessRoleAnalysis
      - BusinessRoleViewOnly

  - name: 'UserRole'
    description: 'defines a role of a user within the system'
    enums:
      - 'UserRoleUnknown'
      - 'UserRoleVerified'
      - 'UserRoleSuspended'
```
3. `enum-codegen --in=example.yaml --out=types.go --package=types`
<details><summary>Generated Output</summary>
<p>

```go
// Code generated by github.com/harry93848bb7/enum-codegen DO NOT EDIT.
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// BusinessRole defines a role of a user to a business
type BusinessRole int64

const (
	BusinessRoleUnknown       BusinessRole = 0
	BusinessRoleAdministrator BusinessRole = 1
	BusinessRoleDeveloper     BusinessRole = 2
	BusinessRoleAnalysis      BusinessRole = 3
	BusinessRoleViewOnly      BusinessRole = 4
)

// String returns the string form of BusinessRole, or "" if BusinessRole is invalid.
func (x BusinessRole) String() string {
	return map[BusinessRole]string{
		BusinessRoleUnknown:       "BusinessRoleUnknown",
		BusinessRoleAdministrator: "BusinessRoleAdministrator",
		BusinessRoleDeveloper:     "BusinessRoleDeveloper",
		BusinessRoleAnalysis:      "BusinessRoleAnalysis",
		BusinessRoleViewOnly:      "BusinessRoleViewOnly",
	}[x]
}

// Scan implements sql.Scanner so BusinessRole can be read from databases transparently.
func (x *BusinessRole) Scan(value interface{}) error {
	db := value.(string)
	*x = map[string]BusinessRole{
		"BusinessRoleUnknown":       BusinessRoleUnknown,
		"BusinessRoleAdministrator": BusinessRoleAdministrator,
		"BusinessRoleDeveloper":     BusinessRoleDeveloper,
		"BusinessRoleAnalysis":      BusinessRoleAnalysis,
		"BusinessRoleViewOnly":      BusinessRoleViewOnly,
	}[string(db)]
	return nil
}

// Value implements sql.Valuer so that BusinessRole can be written to databases transparently.
func (x BusinessRole) Value() (driver.Value, error) {
	return x.String(), nil
}

// MarshalJSON implements the json.Marshaler interface. Converting BusinessRole into a string, striping the prefix.
func (x BusinessRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ReplaceAll(x.String(), "BusinessRole", ""))
}

// UnmarshalJSON UnmarshalJSON implements the json.Unmarshaler interface. Adding the prefix and convert into BusinessRole.
func (x *BusinessRole) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	return x.Scan([]uint8((fmt.Sprintf("BusinessRole%s", strings.Title(j)))))
}

// UserRole defines a role of a user within the system
type UserRole int64

const (
	UserRoleUnknown   UserRole = 0
	UserRoleVerified  UserRole = 1
	UserRoleSuspended UserRole = 2
)

// String returns the string form of UserRole, or "" if UserRole is invalid.
func (x UserRole) String() string {
	return map[UserRole]string{
		UserRoleUnknown:   "UserRoleUnknown",
		UserRoleVerified:  "UserRoleVerified",
		UserRoleSuspended: "UserRoleSuspended",
	}[x]
}

// Scan implements sql.Scanner so UserRole can be read from databases transparently.
func (x *UserRole) Scan(value interface{}) error {
	db := value.(string)
	*x = map[string]UserRole{
		"UserRoleUnknown":   UserRoleUnknown,
		"UserRoleVerified":  UserRoleVerified,
		"UserRoleSuspended": UserRoleSuspended,
	}[string(db)]
	return nil
}

// Value implements sql.Valuer so that UserRole can be written to databases transparently.
func (x UserRole) Value() (driver.Value, error) {
	return x.String(), nil
}

// MarshalJSON implements the json.Marshaler interface. Converting UserRole into a string, striping the prefix.
func (x UserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ReplaceAll(x.String(), "UserRole", ""))
}

// UnmarshalJSON UnmarshalJSON implements the json.Unmarshaler interface. Adding the prefix and convert into UserRole.
func (x *UserRole) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	return x.Scan([]uint8((fmt.Sprintf("UserRole%s", strings.Title(j)))))
}
```

</p>
</details>
<details><summary>Generated Output Tests</summary>
<p>

```go
// Code generated by github.com/harry93848bb7/enum-codegen DO NOT EDIT.
package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BusinessRoleType(t *testing.T) {
	type object struct {
		Type BusinessRole `json:"type"`
	}

	tests := []struct {
		name     string
		arg      object
		expected BusinessRole
	}{
		{
			name: "Should match BusinessRoleUnknown marshal and unmarshal",
			arg: object{
				Type: BusinessRoleUnknown,
			},
			expected: BusinessRoleUnknown,
		},
		{
			name: "Should match BusinessRoleAdministrator marshal and unmarshal",
			arg: object{
				Type: BusinessRoleAdministrator,
			},
			expected: BusinessRoleAdministrator,
		},
		{
			name: "Should match BusinessRoleDeveloper marshal and unmarshal",
			arg: object{
				Type: BusinessRoleDeveloper,
			},
			expected: BusinessRoleDeveloper,
		},
		{
			name: "Should match BusinessRoleAnalysis marshal and unmarshal",
			arg: object{
				Type: BusinessRoleAnalysis,
			},
			expected: BusinessRoleAnalysis,
		},
		{
			name: "Should match BusinessRoleViewOnly marshal and unmarshal",
			arg: object{
				Type: BusinessRoleViewOnly,
			},
			expected: BusinessRoleViewOnly,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.arg)
			assert.NoError(t, err)

			var obj object
			assert.NoError(t, json.Unmarshal(b, &obj))
			assert.Equal(t, tt.expected, obj.Type)
		})
	}
}

func Test_UserRoleType(t *testing.T) {
	type object struct {
		Type UserRole `json:"type"`
	}

	tests := []struct {
		name     string
		arg      object
		expected UserRole
	}{
		{
			name: "Should match UserRoleUnknown marshal and unmarshal",
			arg: object{
				Type: UserRoleUnknown,
			},
			expected: UserRoleUnknown,
		},
		{
			name: "Should match UserRoleVerified marshal and unmarshal",
			arg: object{
				Type: UserRoleVerified,
			},
			expected: UserRoleVerified,
		},
		{
			name: "Should match UserRoleSuspended marshal and unmarshal",
			arg: object{
				Type: UserRoleSuspended,
			},
			expected: UserRoleSuspended,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.arg)
			assert.NoError(t, err)

			var obj object
			assert.NoError(t, json.Unmarshal(b, &obj))
			assert.Equal(t, tt.expected, obj.Type)
		})
	}
}
```

</p>
</details>
