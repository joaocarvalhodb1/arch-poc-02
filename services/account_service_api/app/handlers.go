package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/joaocarvalhodb1/arch-poc/shared/contracts/protog"
	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
)

func (app *AccountAppAPI) Home(w http.ResponseWriter, r *http.Request) {
	helpers.JsonResponse(w, http.StatusOK, "{service: 'Account API Service - version 0.0.1'}")
}

func (app *AccountAppAPI) FindOne(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.Params("id", w, r)
	if err != nil {
		helpers.JsonError(w, http.StatusBadRequest, err)
	}
	account, err := app.accountgRPCService.FindOne(r.Context(), &protog.AccountFilterRequest{
		AccountId: id,
	})
	if err != nil {
		helpers.JsonError(w, http.StatusServiceUnavailable, err)
	}
	helpers.JsonResponse(w, http.StatusOK, account)
}

func (app *AccountAppAPI) FindAll(w http.ResponseWriter, r *http.Request) {
	accountList, err := app.accountgRPCService.FindMany(r.Context(), &protog.AccountFilterRequest{
		AccountId:   "",
		AccountType: "",
	})
	if err != nil {
		helpers.JsonError(w, http.StatusServiceUnavailable, err)
	}
	helpers.JsonResponse(w, http.StatusOK, accountList)
}

func (app *AccountAppAPI) CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error reading request body: "+err.Error()))
		return
	}
	var account *AccountDTO
	if err = json.Unmarshal(body, &account); err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("JSON structure error: "+err.Error()))
		return
	}
	err = account.Validate()
	if err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error validate: "+err.Error()))
		return
	}
	createdAccount, err := app.accountgRPCService.CreateAccount(r.Context(),
		&protog.AccountRequest{
			Data: &protog.Account{
				Id:          account.ID,
				Name:        account.Name,
				Email:       account.Email,
				CellPhone:   account.CellPhone,
				AccountType: protog.AccountType(protog.AccountType_value[account.AccountType]),
				CreditLimit: account.CreditLimit,
			},
		})
	if err != nil {
		helpers.JsonError(w, http.StatusInternalServerError, err)
	}
	helpers.JsonResponse(w, http.StatusCreated, createdAccount)
}

func (app *AccountAppAPI) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.Params("id", w, r)
	if err != nil {
		helpers.JsonError(w, http.StatusBadRequest, err)
	}
	if id == "" {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error reading ID in body: "+err.Error()))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error reading request body: "+err.Error()))
		return
	}
	var account *AccountDTO
	if err = json.Unmarshal(body, &account); err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("JSON structure error: "+err.Error()))
		return
	}
	err = account.Validate()
	if err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error validate: "+err.Error()))
		return
	}
	_, err = app.accountgRPCService.UpdateAccount(r.Context(),
		&protog.AccountRequest{
			Data: &protog.Account{
				Id:          id,
				Name:        account.Name,
				Email:       account.Email,
				CellPhone:   account.CellPhone,
				AccountType: protog.AccountType(protog.AccountType_value[account.AccountType]),
				CreditLimit: account.CreditLimit,
			},
		})
	if err != nil {
		helpers.JsonError(w, http.StatusInternalServerError, err)
	}
	helpers.JsonResponse(w, http.StatusOK, nil)
}

func (app *AccountAppAPI) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.Params("id", w, r)
	if err != nil {
		helpers.JsonError(w, http.StatusBadRequest, err)
	}
	if id == "" {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error reading ID in body: "+err.Error()))
		return
	}
	_, err = app.accountgRPCService.DeleteAccount(r.Context(), &protog.AccountDeleteRequest{
		AccountId: id,
	})
	if err != nil {
		helpers.JsonError(w, http.StatusInternalServerError, err)
	}
	helpers.JsonResponse(w, http.StatusNoContent, nil)
}

func (app *AccountAppAPI) ApplyCreditLilite(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error reading request body: "+err.Error()))
		return
	}
	var creditLimite *CreditLimiteDTO
	if err = json.Unmarshal(body, &creditLimite); err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("JSON structure error: "+err.Error()))
		return
	}
	err = creditLimite.Validate()
	if err != nil {
		helpers.JsonError(w, http.StatusBadGateway, errors.New("Error validate: "+err.Error()))
		return
	}
	_, err = app.accountgRPCService.CreditLimitApply(r.Context(),
		&protog.CreditLimiteRequest{
			AccountType: creditLimite.AccountType,
			CreditLimit: creditLimite.CreditLimit,
		})
	if err != nil {
		helpers.JsonError(w, http.StatusInternalServerError, err)
	}
	helpers.JsonResponse(w, http.StatusOK, nil)
}
