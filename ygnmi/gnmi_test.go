// Copyright 2023 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ygnmi

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestWrapJSONIETF(t *testing.T) {
	tests := []struct {
		desc               string
		in                 string
		inQualifiedRelPath []string
		want               string
		wantErr            bool
	}{{
		desc:               "one-level",
		inQualifiedRelPath: []string{"openconfig-withlistval:ordered-list"},
		in: `[
  {
    "openconfig-withlistval:config": {
      "another-mod:value": "42",
      "key": "foo"
    },
    "openconfig-withlistval:key": "foo"
  },
  {
    "openconfig-withlistval:config": {
      "another-mod:value": "43",
      "key": "bar"
    },
    "openconfig-withlistval:key": "bar"
  },
  {
    "openconfig-withlistval:config": {
      "another-mod:value": "44",
      "key": "baz"
    },
    "openconfig-withlistval:key": "baz"
  }
]`,
		want: `{
  "openconfig-withlistval:ordered-list": [
    {
      "config": {
        "another-mod:value": "42",
        "key": "foo"
      },
      "key": "foo"
    },
    {
      "config": {
        "another-mod:value": "43",
        "key": "bar"
      },
      "key": "bar"
    },
    {
      "config": {
        "another-mod:value": "44",
        "key": "baz"
      },
      "key": "baz"
    }
  ]
}`,
	}, {
		desc:               "two-levels",
		inQualifiedRelPath: []string{"openconfig-withlistval:ordered-lists", "openconfig-withlistval:ordered-list"},
		in: `[
  {
    "openconfig-withlistval:config": {
      "another-mod:value": "42",
      "key": "foo"
    },
    "openconfig-withlistval:key": "foo"
  },
  {
    "openconfig-withlistval:config": {
      "another-mod:value": "43",
      "key": "bar"
    },
    "openconfig-withlistval:key": "bar"
  },
  {
    "openconfig-withlistval:config": {
      "another-mod:value": "44",
      "key": "baz"
    },
    "openconfig-withlistval:key": "baz"
  }
]`,
		want: `{
  "openconfig-withlistval:ordered-lists": {
    "ordered-list": [
      {
        "config": {
          "another-mod:value": "42",
          "key": "foo"
        },
        "key": "foo"
      },
      {
        "config": {
          "another-mod:value": "43",
          "key": "bar"
        },
        "key": "bar"
      },
      {
        "config": {
          "another-mod:value": "44",
          "key": "baz"
        },
        "key": "baz"
      }
    ]
  }
}`,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			in := &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{JsonIetfVal: []byte(tt.in)}}
			err := wrapJSONIETF(in, tt.inQualifiedRelPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			want := &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{JsonIetfVal: []byte(strings.Join(strings.Fields(tt.want), ""))}}
			if diff := cmp.Diff(in, want, protocmp.Transform()); diff != "" {
				t.Errorf("(-got, +want):\n%s", diff)
			}
		})
	}
}
