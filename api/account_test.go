package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/edarha/simplebank/db/mock"
	db "github.com/edarha/simplebank/db/sqlc"
	"github.com/edarha/simplebank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	account := createAccountRandom()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// build stubs
	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// start test server and send request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	path := fmt.Sprintf("/accounts/%d", account.ID)
	req, err := http.NewRequest(http.MethodGet, path, nil)

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	// test
	require.Equal(t, http.StatusOK, recorder.Code)

	requireBodyMatchAccount(t, recorder.Body, account)
}

func createAccountRandom() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)

	require.NoError(t, err)
	var acc db.Account
	err = json.Unmarshal(data, &acc)
	require.NoError(t, err)
	require.Equal(t, account, acc)
}
