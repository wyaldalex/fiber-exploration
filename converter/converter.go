package converter

import (
	"FiberAuth/models"
	"fmt"
)

func ParseRole(s string) (models.Role, error) {
	switch s {
	case "User":
		return models.BaseRole, nil
	case "Admin":
		return models.AdminRole, nil
	default:
		return 0, fmt.Errorf("invalid role value: %s", s)
	}
}
