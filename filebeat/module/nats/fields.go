// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package nats

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "nats", asset.ModuleFieldsPri, AssetNats); err != nil {
		panic(err)
	}
}

// AssetNats returns asset data.
// This is the base64 encoded gzipped contents of module/nats.
func AssetNats() string {
	return "eJy0lEFP3DAQhe/5FU/c4QfkUAkJVeqhXOCOnGSSDDieYI+7pL++coi3y5LAVmXntutk3jfvTXyJJ5pKOKOhAJTVUomL9POiABoKtedRWVyJbwUA/JQmWkIrHqPxgV2H2+v7O1jp0LKlcFUAniyZQCUqUlMALZNtQjk3uIQzA+0lU+k0UonOSxyXf1aEU32f+6D1MuxVZ8FUhyKHQla6/X9rWh/opbo1Gv5Od3B0PGSuY45DltoyOX1ztIX0CdamHYvG1dHTa1SHZNy8O8pk7JQ68ivnn/Cluu8J3EBaaE9r82eAIXTn9WWgEExH/2pMNemb2N/SWXHH2HNT8YPR7ZdPMO6Of1O2bTSTFdOA3UbHDJugNlmfaNqJX8v5xBxHLyq12GzlulxmCbF6pPp42b8I5+61+SwF7XkfL3YmwFNN/IsaiNvGO+fGR8fPkWDs2BsXB/JcJz/27+LHTc52y6bM6Wm004PK+WJlV8lL5oA47Hqu+9fFi5Xl0JMHB1gOSi7d+On29xRGceGDZRzMy8OSyvYH9F9WXzvI/ICxcHGoyCdXsyZUsDOsM25FrXiCiSqDUa6NtROiW0KpeOUzznOQ9+Kvlq6bgyi9rO36CVPckBq2AaaSqLPtsyKkrqP39D7ZDPYcKdLD2gX5FVhpNWaF1yt42YpsGHns2Fo8CrviTwAAAP//7bga1A=="
}