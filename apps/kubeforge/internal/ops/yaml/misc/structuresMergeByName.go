// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

package yaml

func StructuresMergeByName(existing, incoming interface{}) interface{} {
	switch ex := existing.(type) {
	case map[string]interface{}:
		if in, ok := incoming.(map[string]interface{}); ok {
			merged := make(map[string]interface{})
			// Merge maps by iterating over existing and incoming
			for key, value := range ex {
				merged[key] = value
			}
			// Traverse over the incoming map
			for key, value := range in {
				if existingValue, ok := merged[key]; ok {
					// Recursively merge values if they exist in both
					merged[key] = StructuresMergeByName(existingValue, value)
				} else {
					merged[key] = value
				}
			}
			return merged
		}
	case []interface{}:
		if in, ok := incoming.([]interface{}); ok {
			// Merge slices recursively
			return mergeSlices(ex, in)
		}
	}
	// If the types don't match, return the incoming directly
	return incoming
}

func mergeSlices(existing, incoming []interface{}) []interface{} {
	// Create a new slice to store the merged result
	merged := append([]interface{}{}, existing...)
	existingKeyIndex := map[string]int{} // Track existing items by their keys.

	// Detect key for matching (e.g., "name", "metadata.name", etc.)
	keyExtractor := identifyKey(existing)

	// Populate existingKeyIndex with indices from the existing slice
	for i, item := range existing {
		if key := keyExtractor(item); key != "" {
			existingKeyIndex[key] = i
		}
	}

	// Merge or append items from the incoming slice
	for _, incomingItem := range incoming {
		if key := keyExtractor(incomingItem); key != "" {
			// If the key exists, merge the items recursively
			if index, exists := existingKeyIndex[key]; exists {
				existingItem := merged[index]
				// Merge the matching item from existing and incoming
				merged[index] = StructuresMergeByName(existingItem, incomingItem)
			} else {
				// If no matching key, add the new item
				merged = append(merged, incomingItem)
			}
		} else {
			// If no key, append the item directly
			merged = append(merged, incomingItem)
		}
	}

	return merged
}

// IdentifyKey identifies the key used for matching items in a slice.
// It dynamically detects keys such as "name", "metadata.name", or other common identifiers.
func identifyKey(slice []interface{}) func(interface{}) string {
	// Iterate through the slice to find a sample item.
	for _, item := range slice {
		if itemMap, ok := item.(map[string]interface{}); ok {
			// Check for common keys like "name"
			if _, exists := itemMap["name"]; exists {
				return func(i interface{}) string {
					if m, ok := i.(map[string]interface{}); ok {
						if name, ok := m["name"].(string); ok {
							return name
						}
					}
					return ""
				}
			}
			// Check for nested key: "metadata.name"
			if metadata, exists := itemMap["metadata"]; exists {
				if metaMap, ok := metadata.(map[string]interface{}); ok {
					if _, exists := metaMap["name"]; exists {
						return func(i interface{}) string {
							if m, ok := i.(map[string]interface{}); ok {
								if meta, ok := m["metadata"].(map[string]interface{}); ok {
									if name, ok := meta["name"].(string); ok {
										return name
									}
								}
							}
							return ""
						}
					}
				}
			}
		}
	}
	// Default to no key function.
	return func(i interface{}) string { return "" }
}
