package nostradamus_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestCreateSales(t *testing.T) {
	b := []byte(`<nostradamus>
 <sales erase="true" business_date="2009-10-04">
  <sale office_id="1" enter="2009-10-04 19:18:00" ticket_id="20072" table="450" guests="2" employee_id="4">
   <product datetime="2009-10-04 19:18:00" pos_id="1" article_id="10" ordering="0" amount="2" value="3.6660" employee_id="4" />
   <product datetime="2009-10-04 19:18:00" pos_id="1" article_id="160" ordering="1" amount="2" value="11.1860" employee_id="4"/>
  </sale>
  <sale office_id="1" enter="2009-10-04 18:01:00" ticket_id="20055" table="340" guests="8" staff_id="96">
    <product datetime="2009-10-04 18:01:00" pos_id="1" article_id="906" ordering="0" amount="1" value="2.2090" staff_id="96" />
    <product datetime="2009-10-04 18:01:00" pos_id="1" article_id="905" ordering="1" amount="1" value="2.2090" staff_id="96" />
    <product datetime="2009-10-04 18:01:00" pos_id="1" article_id="904" ordering="2" amount="1" value="2.2090" staff_id="96" />
    <product datetime="2009-10-04 18:01:00" pos_id="1" article_id="9" ordering="3" amount="2" value="3.6660" staff_id="96" />
    <product datetime="2009-10-04 18:01:00" pos_id="1" article_id="7" ordering="4" amount="1" value="1.9270" staff_id="96" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="791" ordering="0" amount="1" value="9.3530" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="430" ordering="1" amount="2" value="0.0000" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="792" ordering="2" amount="1" value="9.3530" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="786" ordering="3" amount="1" value="4.6530" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="429" ordering="4" amount="1" value="0.0000" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="289" ordering="5" amount="1" value="12.1730" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="427" ordering="6" amount="3" value="0.0000" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="292" ordering="7" amount="1" value="12.1730" staff_id="97" />
    <product datetime="2009-10-04 18:09:00" pos_id="1" article_id="299" ordering="8" amount="1" value="12.1730" staff_id="97" />
    <product datetime="2009-10-04 18:48:00" pos_id="1" article_id="337" ordering="0" amount="3" value="16.7790" staff_id="97" />
    <product datetime="2009-10-04 18:48:00" pos_id="1" article_id="308" ordering="1" amount="1" value="3.7130" staff_id="97" />
    <product datetime="2009-10-04 18:48:00" pos_id="1" article_id="428" ordering="2" amount="1" value="0.0000" staff_id="97" />
    <product datetime="2009-10-04 18:48:00" pos_id="1" article_id="771" ordering="3" amount="1" value="0.0000" staff_id="97" />
  </sale>
 </sales>
</nostradamus>`)

	req := client.NewCreateSalesRequest()
	xml.Unmarshal(b, req.RequestBody())

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	t.Log(string(b))

}
