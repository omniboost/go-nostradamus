package nostradamus_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestCreateArticleGroups(t *testing.T) {
	b := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<nostradamus>
 <articlegroups>
   <group id="1" office_id="1" name="Dranken" />
   <group id="2" office_id="1" name="Bijgerechten" />
   <group id="3" office_id="1" name="Voorgerechten" />
   <group id="4" office_id="1" name="Hoofdgerechten" />
   <group id="5" office_id="1" name="Nagerechten" />
   <group id="1" office_id="2" name="Dranken" />
   <group id="2" office_id="2" name="Bijgerechten" />
   <group id="3" office_id="2" name="Voorgerechten" />
   <group id="4" office_id="2" name="Hoofdgerechten" />
   <group id="5" office_id="2" name="Nagerechten" />
 </articlegroups>
</nostradamus>`)

	req := client.NewCreateArticleGroupsRequest()
	xml.Unmarshal(b, req.RequestBody())

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	t.Log(string(b))
}
