package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/marcosavieira/simple-api-login/security"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

func (server *Server) userLogin(ctx *gin.Context) {
	var req loginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = security.CheckPasswordHash(user.Password, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 100).Unix()
	claims["username"] = req.Username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := &loginResponse{
		UserID: user.ID,
		Token:  tokenString,
	}

	ctx.JSON(http.StatusOK, arg)
}
