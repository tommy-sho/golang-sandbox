package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMessageUnmarshal(t *testing.T) {
	d := `{
 "insertId":  "1pvxdo2g1l3z1q9"  ,
   "jsonPayload": {
        "msg": "failed to expire delete job  [segmentID=f7adc6b3-9c81-4ddf-b972-d543a643c20d]:main.main.func5/go/src/github.com/abema/vega-dmp-segment/internal/adaptor/cmd/batch/batch.go:25faild [segmentID=f7adc6b3-9c81-4ddf-b972-d543a643c20d]github.com/abema/vega-dmp-segment/internal/application.SegmentDeleter.Delet/go/src/github.com/abema/vega-dmp-segment/internal/application/segment_delete.go:43BadRequest:depended campaign is activated campaignID :[[711 ON]] campaign.Activated:%!!(MISSING)v(MISSING)"
    },
    "ts": 1568963104.4556043,
 "resource": {
  "labels": {
   "cluster_name":  "stg-vega"    ,
   "container_name":  "vega-dmp-segment-batch"  ,
   "instance_id":  "7900413573781772793"    ,
   "namespace_id":  "default"    ,
   "pod_id":  "vega-dmp-segment-batch-5ff48c5ff7-znmt5"    ,
   "project_id":  "vega-177606"    ,
   "zone":  "asia-east1-a"
  }},
 "severity":  "ERROR",
 "timestamp": "2019-09-24T02:05:03.697532774Z"
}`
	t.Run("unmarshal", func(t *testing.T) {
		msg := Message{}
		err := json.Unmarshal([]byte(d), &msg)
		if err != nil {
			t.Errorf("failed to unmarshaling: %v", err)
		}
		if msg.JsonPayload.Msg != "failed to expire delete job  [segmentID=f7adc6b3-9c81-4ddf-b972-d543a643c20d]:main.main.func5/go/src/github.com/abema/vega-dmp-segment/internal/adaptor/cmd/batch/batch.go:25faild [segmentID=f7adc6b3-9c81-4ddf-b972-d543a643c20d]github.com/abema/vega-dmp-segment/internal/application.SegmentDeleter.Delet/go/src/github.com/abema/vega-dmp-segment/internal/application/segment_delete.go:43BadRequest:depended campaign is activated campaignID :[[711 ON]] campaign.Activated:%!!(MISSING)v(MISSING)" {
			t.Errorf("unmatch value[jsonpayload.msg]: %v", err)
		}
		if msg.Severity != "ERROR" {
			t.Errorf("unmatch value[severity]: %v", err)
		}
		fmt.Println(msg)
	})
}
