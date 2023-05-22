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
	"time"
)

// NoCache sets no-cache headers.
func NoCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		rw.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		rw.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

		next.ServeHTTP(rw, r)
	})
}
