package nostradamus_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestCreateArticles(t *testing.T) {
	b := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<nostradamus>
 <articles>
   <article id="1" group_id="1" subgroup_id="3" office_id="1" name="Fristi" price="2.05" tax="6" state="1" />
   <article id="2" group_id="1" subgroup_id="3" office_id="1" name="Chocomel" price="2.05" tax="6" state="1" />
   <article id="3" group_id="1" subgroup_id="3" office_id="1" name="Appelsap" price="2.05" tax="6" state="0" />
   <article id="4" group_id="1" subgroup_id="1" office_id="1" name="Amaretto" price="3.95" tax="19" state="1" />
   <article id="5" group_id="1" subgroup_id="1" office_id="1" name="Tia Maria" price="3.95" tax="19" state="1" />
 </articles>
</nostradamus>`)

	req := client.NewCreateArticlesRequest()
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
