// Code generated by github.com/DIMO-Network/pkg/ruptela/codegen DO NOT EDIT.
package ruptela

import (
	"fmt"
	"slices"
	"strconv"
)

const bitsInByte = 8

// Convert102 converts the given raw value to a float64.
// Unit: 'km' Min: '0' Max: '65535'.
func Convert102(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65535 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert103 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '255'.
func Convert103(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.39215686274509803)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 255 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert107 converts the given raw value to a float64.
// Min: '0' Max: '65535'.
func Convert107(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65535 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert114 converts the given raw value to a float64.
// Unit: 'm' Min: '0' Max: '4211081215'.
func Convert114(rawValue string) (float64, error) {
	const byteSize = 4
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(5)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is within the error range.
	if 4261412864 <= rawInt && rawInt <= 4278190079 {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 4211081215 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert1148 converts the given raw value to a float64.
// Unit: 'liters' Min: '0' Max: '254'.
func Convert1148(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{255}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 254 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert1149 converts the given raw value to a float64.
// Unit: 'liters' Min: '0' Max: '254'.
func Convert1149(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{255}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 254 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert1150 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '250'.
func Convert1150(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.4)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{251}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 250 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert205 converts the given raw value to a float64.
// Unit: 'l' Min: '0' Max: '65535'.
func Convert205(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65535 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert207 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '250'.
func Convert207(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.4)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{254}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 250 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert29 converts the given raw value to a float64.
// Unit: 'mV' Min: '0' Max: '65535'.
func Convert29(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{65535}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65535 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert483 converts the given raw value to a float64.
// Unit: '-' Min: '0' Max: '250'.
func Convert483(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{254}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 250 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert642 converts the given raw value to a float64.
// Unit: 'l' Min: '0' Max: '0xFFFF or 65535'.
func Convert642(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{65535}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65535 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert645 converts the given raw value to a float64.
// Unit: 'km' Min: '0' Max: '0xFFFFFFFF'.
func Convert645(rawValue string) (float64, error) {
	const byteSize = 4
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{4294967295}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 4294967295 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert722 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '255'.
func Convert722(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 255 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert723 converts the given raw value to a float64.
// Unit: 'km' Min: '0' Max: '65535'.
func Convert723(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65535 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert94 converts the given raw value to a float64.
// Unit: 'RPM' Min: '0' Max: '65,535'.
func Convert94(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.25)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65535 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert95 converts the given raw value to a float64.
// Unit: 'km/h' Min: '0' Max: '255'.
func Convert95(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 255 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert96 converts the given raw value to a float64.
// Unit: '°C' Min: '0' Max: '255'.
func Convert96(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(-40)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{255}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 255 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert960 converts the given raw value to a float64.
// Unit: 'PSI' Min: '0' Max: '65534'.
func Convert960(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.05)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{65535}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65534 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert961 converts the given raw value to a float64.
// Unit: 'PSI' Min: '0' Max: '65534'.
func Convert961(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.05)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{65535}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65534 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert962 converts the given raw value to a float64.
// Unit: 'PSI' Min: '0' Max: '65534'.
func Convert962(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.05)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{65535}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65534 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert963 converts the given raw value to a float64.
// Unit: 'PSI' Min: '0' Max: '65534'.
func Convert963(rawValue string) (float64, error) {
	const byteSize = 2
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.05)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{65535}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 65534 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert964 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '254'.
func Convert964(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.393)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{255}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 254 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert97 converts the given raw value to a float64.
// Unit: '°C' Min: '0' Max: '255'.
func Convert97(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(-40)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{255}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 255 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert98 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '255'.
func Convert98(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(0.39215686274509803)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	// Check if the value is in the error set.
	if slices.Contains([]uint64{255}, rawInt) {
		return 0, errNotFound
	}
	// Check if the value is less than the minimum value.
	if rawInt < 0 {
		return 0, errNotFound
	}
	// Check if the value is greater than the maximum value.
	if rawInt > 255 {
		return 0, errNotFound
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert99 converts the given raw value to a float64.
// Unit: '-'.
func Convert99(rawValue string) (float64, error) {
	const byteSize = 1
	const offset = float64(0)
	const maxSize = 1<<(byteSize*bitsInByte) - 1
	const multiplier = float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}

	// Check if the value is equal to the maximum value for the given size.
	if rawInt == maxSize {
		return 0, errNotFound
	}

	return float64(rawInt)*multiplier + offset, nil
}
