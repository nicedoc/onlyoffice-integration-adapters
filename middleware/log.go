/**
 *
 * (c) Copyright Ascensio System SIA 2023
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package middleware provides http middlewares
//
// The middleware package's functions get added to http services automatically.
package middleware

import (
	"net/http"

	"github.com/ONLYOFFICE/onlyoffice-integration-adapters/log"
)

// Log creates a new debug logging middleware.
func Log(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if logger != nil {
				logger.Debugf("calling [%s] route %s", r.Method, r.URL.Path)
			}
			next.ServeHTTP(w, r)
		})
	}
}
