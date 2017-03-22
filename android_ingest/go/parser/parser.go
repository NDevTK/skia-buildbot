// Parser parses incoming JSON files from Android Testing and converts them
// into a format acceptable to Skia Perf.
package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"go.skia.org/infra/go/sklog"

	"go.skia.org/infra/perf/go/ingestcommon"
)

// Incoming is the JSON structure of the data sent to us from the Android
// testing infrastructure.
type Incoming struct {
	BuildId     string `json:"build_id"`
	BuildFlavor string `json:"build_flavor"`
	Branch      string `json:"branch"`

	// Metrics is a map[test name]map[metric]value, where value
	// is a string encoded float, thus the use of json.Number.
	Metrics map[string]map[string]json.Number `json:"metrics"`
}

// Parse the 'incoming' stream into an *Incoming struct.
func Parse(incoming io.Reader) (*Incoming, error) {
	ret := &Incoming{}
	if err := json.NewDecoder(incoming).Decode(ret); err != nil {
		return nil, fmt.Errorf("Failed to decode incoming JSON: %s", err)
	}
	return ret, nil
}

// An interface for looking up a git hashe from a buildid.
//
// The *lookup.Cache satisfies this interface.
type Lookup interface {
	Lookup(buildid int64) (string, error)
}

// Converter converts a serialized *Incoming into
// an *ingestcommon.BenchData.
type Converter struct {
	lookup Lookup
	branch string
}

// New creates a new *Converter.
//
// Only data from the given branch are successfully converted.
func New(lookup Lookup, branch string) *Converter {
	return &Converter{
		lookup: lookup,
		branch: branch,
	}
}

// Covert the serialize *Incoming JSON into an *ingestcommon.BenchData.
func (c *Converter) Convert(incoming io.Reader) (*ingestcommon.BenchData, error) {
	b, err := ioutil.ReadAll(incoming)
	if err != nil {
		return nil, fmt.Errorf("Failed to read during convert: %s", err)
	}
	reader := bytes.NewReader(b)
	in, err := Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse during convert: %s", err)
	}
	sklog.Infof("POST for buildid: %s branch: %s", in.BuildId, in.Branch)
	if in.Branch != c.branch {
		return nil, fmt.Errorf("Found data for a branch we weren't expecting: Got %q Want %q at BuildID: %s", in.Branch, c.branch, in.BuildId)
	}
	buildid, err := strconv.ParseInt(in.BuildId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse buildid %q: %s", in.BuildId, err)
	}
	hash, err := c.lookup.Lookup(buildid)
	if err != nil {
		return nil, fmt.Errorf("Failed to find matching hash for buildid %d: %s", buildid, err)
	}

	// Convert Incoming into ingestcommon.BenchData, i.e. convert the following:
	//
	//		{
	//			"build_id": "3567162",
	//			"build_flavor": "marlin-userdebug",
	//			"metrics": {
	//				"android.platform.systemui.tests.jank.LauncherJankTests#testAppSwitchGMailtoHome": {
	//				"frame-fps": "9.328892269753897",
	//				"frame-avg-jank": "8.4",
	//				"frame-max-frame-duration": "7.834711093388444",
	//				"frame-max-jank": "10"
	//			},
	//	    ...
	//    }
	//  }
	//
	//  into
	//
	//  {
	//    "gitHash" : "8dcc84f7dc8523dd90501a4feb1f632808337c34",
	//    "key" : {
	//      "build_flavor" : "marlin-userdebug"
	//    },
	//    "results" : {
	//      "android.platform.systemui.tests.jank.LauncherJankTests#testAppSwitchGMailtoHome" : {
	//        "default" : {
	//          "frame-fps": 9.328892269753897,
	//          "frame-avg-jank": 8.4,
	//          "frame-max-frame-duration": 7.834711093388444,
	//          "frame-max-jank": 10
	//          "options" : {
	//            "name" : "android.platform.systemui.tests.jank.LauncherJankTests",
	//            "subtest" : "testAppSwitchGMailtoHome",
	//          },
	//        },
	//      }
	//    }
	//  }
	//
	// Note that the incoming data doesn't have a concept similar to "config" so we just
	// use a value of "default" for config for now.
	benchData := &ingestcommon.BenchData{
		Hash: hash,
		Key: map[string]string{
			"build_flavor": in.BuildFlavor,
		},
		Results: map[string]ingestcommon.BenchResults{},
	}
	for test, metrics := range in.Metrics {
		benchData.Results[test] = ingestcommon.BenchResults{}
		benchData.Results[test]["default"] = ingestcommon.BenchResult{}
		for key, value := range metrics {
			f, err := value.Float64()
			if err != nil {
				sklog.Errorf("Couldn't parse %q as a float64: %s", value.String(), err)
				continue
			}
			benchData.Results[test]["default"][key] = f
		}
		parts := strings.Split(test, "#")
		if len(parts) != 2 {
			sklog.Errorf("Test name didn't have a single # separator: %q", test)
			continue
		}
		benchData.Results[test]["default"]["options"] = map[string]string{
			"name":    parts[0],
			"subtest": parts[1],
		}
	}
	if len(benchData.Results) == 0 {
		sklog.Warningf("Failed to extract any data from incoming file: %q", string(b))
	}

	return benchData, nil
}
