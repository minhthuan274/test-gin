package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewReviewButNotValidBody(t *testing.T) {
	router := setupRouter()
	body := bytes.NewBuffer([]byte("{\"foo\":\"bar\"}"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v3/reviews", body)

	req.Header.Set("Authorization", getToken())

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddNewReviewAndSucess(t *testing.T) {
	router := setupRouter()
	body := bytes.NewBuffer([]byte("{\"merchant\":\"56f8f6bdb7bfb8a979db12b5\", \"feedback\": \"Quán nấu ăn ngon\", \"point\": 4}"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v3/reviews", body)

	req.Header.Set("Authorization", getToken())

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	// var response map[string]string
	// _ = json.Unmarshal([]byte(w.Body.String()), &response)

}

func TestHomeWithUnauthorized(t *testing.T) {
	// Grab our router
	router := setupRouter()
	// Perform a GET request with that handler.
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v3/home", nil)
	router.ServeHTTP(w, req)
	// Assert we encoded correctly,
	// the request gives a 401
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestHomeWithAuthorize(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v3/home", nil)

	setHeaderAuthorization(*req)
	router.ServeHTTP(w, req)

	var response map[string]string
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	userID, exists := response["userID"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, exists)
	assert.Equal(t, userID, "571899623dc63af34d67a662")
}

func getToken() string {
	return "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI1NzE4OTk2MjNkYzYzYWYzNGQ2N2E2NjIiLCJuYW1lIjoiQuG7nW0iLCJjcm0iOnt9LCJzdGF0aXN0aWMiOnsiY2hlY2tpbiI6MTU2LCJiaWxsIjo5MCwiY29pbiI6MTQ1MSwiY2hlY2tpbkNvaW4iOjI4MzUsImJpbGxlZFpjb2luIjoxNzM2LCJleHBlbnNlIjoxMTY1MDAwLCJyZXdhcmQiOjYsImJhZGdlIjoxMCwibGFzdEJpbGxBdCI6IjIwMTgtMDQtMTNUMDQ6MzQ6MzYuMjMyWiIsImxhc3RDaGVja2luQXQiOiIyMDE3LTA1LTI5VDA3OjI2OjA5Ljk0N1oifSwic3RhdHVzZXMiOnsib25saW5lIjpmYWxzZSwidmVyaWZpZWQiOnRydWUsImJhbm5lZCI6ZmFsc2UsInN1c3BpY2lvbiI6e319LCJiaXJ0aGRheSI6IjE5ODktMDktMDFUMTc6MDA6MDAuMDAwWiIsInJvbGVzIjpbImF1dGhlbnRpY2F0ZWQiXSwibG9jYWxlIjoidmkiLCJ2ZXJzaW9uIjoiMi45LjIiLCJpbnRlZ3JhdGVkSWQiOnsidGNoIjoiTTE2MDYyOTg4In0sInJlZmVycmFsQ29kZSI6IiIsImF2YXRhciI6Imh0dHBzOi8vem9keWFwcC1kZXYuczMuYW1hem9uYXdzLmNvbS9zbV85MDA1NzExODg3NDBfMTQ5MDE1NTc3MTQ5My5wbmciLCJjaXR5IjoiZGEtbmFuZyIsImdlbmRlciI6Im1hbGUiLCJmYWNlYm9vayI6Imh0dHBzOi8vd3d3LmZhY2Vib29rLmNvbS9hcHBfc2NvcGVkX3VzZXJfaWQvMTIxNzEwNDgzODMxNzk4OC8iLCJwaG9uZSI6Iis4NDkzNDg3MTYyNyIsInJlZmVycmFsIjp7ImNvZGUiOiJCT01CT00iLCJjYW5FZGl0Ijp0cnVlLCJzaGFyZVRleHQiOiJUaGFtIGdpYSBab2R5IHTDrWNoIMSRaeG7g20sIMSR4buVaSB0aMaw4bufbmcga2hpIMSRaSDEg24gdeG7kW5nLiBU4bqjaSBhcHAgdOG6oWk6IGh0dHBzOi8vem9keS52bi9kbC9zaGFyZSwgbmjhuq1wIG3DoyBnaeG7m2kgdGhp4buHdTogQk9NQk9NIGtoaSDEkcSDbmcga8O9IMSR4buDIG5o4bqtbiDEkWnhu4NtIHRoxrDhu59uZy4iLCJob3N0IjoiaHR0cHM6Ly96b2R5LnZuL2RsL3NoYXJlIn0sImlhdCI6MTUyNzY1NDg4NCwiZXhwIjoxNTU5MjEyNDg0fQ.RdVrncFfsFAGRK5f-QsWmX2bkrfFTxPZa9YN-XcUzzk"
}

func setHeaderAuthorization(req http.Request) {
	req.Header.Set("Authorization", getToken())
}
