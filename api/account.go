package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/yogesh-k64/middleware-simple-bank/db/sqlc"
	"github.com/yogesh-k64/middleware-simple-bank/token"
)

type createAccountReq struct {
	Currency string `json:"currency" binding:"required,oneof=INR USD"`
}

type getAccountReq struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccountsReq struct {
	PageNo int32 `form:"pageNo" binding:"required,min=1"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

func (server Server) createAccount(ctx *gin.Context) {
	var req createAccountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorHandler(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateAccountParams{
		Owner:    authPayload.UserName,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorHandler(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorHandler(err))
		return
	}

	ctx.JSON(http.StatusAccepted, account)
}

func (server Server) getAccount(ctx *gin.Context) {
	var req getAccountReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorHandler(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorHandler(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorHandler(err))
		return
	}

	if account.Owner != authPayload.UserName {
		err := errors.New("account doesn't belong to the authenticated owner")
		ctx.JSON(http.StatusUnauthorized, errorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server Server) listAccounts(ctx *gin.Context) {
	var req listAccountsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorHandler(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListAccountsParams{
		Owner:  authPayload.UserName,
		Limit:  req.Limit,
		Offset: (req.PageNo - 1) * req.Limit,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorHandler(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
