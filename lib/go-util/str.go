package util

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// RemoveStrDuplicates removes duplicates from strings, considering a map of already-seen duplicates.
// Returns the strings which are unique, and not already present in seen; and a map of the unique strings in inputStrings and seenStrings.
//
// This can be used, for example, to remove duplicates from multiple lists of strings, in order, using a shared map of seen strings.
func RemoveStrDuplicates(inputStrings []string, seenStrings map[string]struct{}) ([]string, map[string]struct{}) {
	uniqueStrings := []string{}
	for _, str := range inputStrings {
		if _, ok := seenStrings[str]; !ok {
			uniqueStrings = append(uniqueStrings, str)
			seenStrings[str] = struct{}{}
		}
	}
	return uniqueStrings, seenStrings
}
