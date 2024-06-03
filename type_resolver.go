package digoflow

import (
	"fmt"
	"strconv"
)

func ConvertToType(typeName string, value any) (any, error) {
	switch typeName {
	case "string":
		return ConvertToString(value)
	case "int":
		return convertToInt(value)
	case "bool":
		return convertToBool(value)
	case "float":
		return convertToFloat(value)
	default:
		return value, nil
	}
}

func ConvertToString(value any) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("unsupported type %T", v)
	}
}

func convertToInt(value any) (int, error) {
	if v, ok := value.(string); ok {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("error converting string to int: %v", err)
		}
		return intVal, nil
	}
	return value.(int), nil
}

func convertToBool(value any) (bool, error) {
	if v, ok := value.(string); ok {
		boolVal, err := strconv.ParseBool(v)
		if err != nil {
			return false, fmt.Errorf("error converting string to bool: %v", err)
		}
		return boolVal, nil
	}
	return value.(bool), nil
}

func convertToFloat(value any) (float64, error) {
	if v, ok := value.(string); ok {
		floatVal, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("error converting string to float: %v", err)
		}
		return floatVal, nil
	}
	return value.(float64), nil
}
