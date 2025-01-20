package nostradamus_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestCreateArticleSubGroups(t *testing.T) {
	b := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<nostradamus>
 <articlesubgroups>
   <group id="1" parent="1" office_id="1" name="Alcoholische Dranken" />
   <group id="2" parent="1" office_id="1" name="Warme Dranken" />
   <group id="3" parent="1" office_id="1" name="Koude Dranken" />
   <group id="4" parent="2" office_id="1" name="Salades" />
   <group id="5" parent="2" office_id="1" name="Snacks" />
   <group id="1" parent="1" office_id="2" name="Alcoholische Dranken" />
   <group id="2" parent="1" office_id="2" name="Warme Dranken" />
   <group id="3" parent="1" office_id="2" name="Koude Dranken" />
   <group id="4" parent="2" office_id="2" name="Salades" />
   <group id="5" parent="2" office_id="2" name="Snacks" />
 </articlesubgroups>
</nostradamus>`)

	req := client.NewCreateArticleSubGroupsRequest()
	xml.Unmarshal(b, req.RequestBody())

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	t.Log(string(b))

	// req := client.NewCreateArticlesRequest()
	// xml.Unmarshal(b, req.RequestBody())

	// reqBody := req.RequestBody()
	// bodyBytes, _ := xml.MarshalIndent(reqBody, "", "  ")
	// t.Logf("Request body: %s", string(bodyBytes))

	// resp, err := req.Do()
	// if err != nil {
	// 	t.Error(err)
	// }

	// b, _ = json.MarshalIndent(resp, "", "  ")
	// t.Log(string(b))
}
