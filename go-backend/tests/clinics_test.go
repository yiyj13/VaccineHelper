package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"v-helper/internal/model"
	"v-helper/pkg/utils"
)

func TestVaccineClinicList(t *testing.T) {
	const url string = "http://localhost:80/api/clinics"
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	user := model.User{OpenID: "Admin"}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	// POST /clinics
	newVaccineClinicList := model.VaccineClinicList{
		VaccineName: timeStr,
		ClinicList:  "清华大学社区卫生服务中心;北京大学医院;北京大学第三医院;中关村医院;中国气象局医院;海淀区双榆树社区卫生服务中心",
	}
	jsonData, _ := json.Marshal(newVaccineClinicList)
	resp, err := sendRequestWithToken("POST", url, bytes.NewBuffer(jsonData), token)
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to create: %v", resp.StatusCode)
	}

	// GET /clinics
	resp, err = sendRequestWithToken("GET", url, nil, token)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	t.Logf("GET Response: %s", string(body))

	var vaccineClinicLists []model.VaccineClinicList
	if err := json.Unmarshal(body, &vaccineClinicLists); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	var targetID uint
	for _, vaccineClinicList := range vaccineClinicLists {
		if vaccineClinicList.VaccineName == timeStr {
			targetID = vaccineClinicList.ID
			break
		}
	}
	urlWithId := url + "/" + strconv.Itoa(int(targetID))

	// GET /clinics/:id
	resp, err = sendRequestWithToken("GET", urlWithId, nil, token)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get: %v", resp.StatusCode)
	}

	// PUT /clinics/:id
	newVaccineClinicList.VaccineName = "testUpdateVaccineName"
	jsonData, _ = json.Marshal(newVaccineClinicList)
	resp, err = sendRequestWithToken("PUT", urlWithId, bytes.NewBuffer(jsonData), token)
	if err != nil {
		t.Fatalf("Failed to send PUT request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to update: %v", resp.StatusCode)
	}

	// GET /clinics/vaccineName/:vaccineName
	resp, err = sendRequestWithToken("GET", "http://localhost:80/api/clinics/vaccineName/常规疫苗", nil, token)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get: %v", resp.StatusCode)
	}

	// GET /clinics/clinicName/:clinicName
	resp, err = sendRequestWithToken("GET", "http://localhost:80/api/clinics/clinicName/清华大学社区卫生服务中心", nil, token)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get: %v", resp.StatusCode)
	}

	// DELETE /clinics/:id
	resp, err = sendRequestWithToken("DELETE", urlWithId, nil, token)
	if err != nil {
		t.Fatalf("Failed to send DELETE request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get: %v", resp.StatusCode)
	}
}
