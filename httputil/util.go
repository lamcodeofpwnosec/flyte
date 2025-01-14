/*
Copyright (C) 2018 Expedia Group.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package httputil

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/rs/zerolog/log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	switch r.Header.Get(HeaderAccept) {
	case MediaTypeYaml:
		writeResponseAsYAML(w, v)
	default:
		writeResponseAsJSON(w, v)
	}
}

func writeResponseAsYAML(w http.ResponseWriter, v interface{}) {
	data, err := yaml.Marshal(v)
	if err != nil {
		log.Err(err).Msg("cannot convert to yaml")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeResponse(w, ContentTypeYaml, data)
}

func writeResponseAsJSON(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Err(err).Msg("cannot convert to JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeResponse(w, ContentTypeJson, data)
}

func writeResponse(w http.ResponseWriter, contentType string, data []byte) {
	w.Header().Set(HeaderContentType, contentType)
	_, err := w.Write(data)
	if err != nil {
		log.Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
	}
}
