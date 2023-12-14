package handler

import (
	"encoding/json"
	"errors"
	myError "lms_try/helper/error"
	"lms_try/model/dto"
	"lms_try/model/global"
	interfaceService "lms_try/service/interface"
	"net/http"
)

type UserHandler struct {
	UserService interfaceService.IUserService
}

// function provider
func NewuserHandler(userService interfaceService.IUserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// handler GetAll
func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// call procedure in service
	users, err := u.UserService.GetAll(r.Context())
	if err != nil {
		switch {
		case errors.As(err, &myError.NotFoundError{}):
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("error not found"))
			return
		case errors.As(err, &myError.BadRequestError{}):
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("error bad request"))
			return
		case errors.As(err, &myError.ServerError{}):
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("error internal server"))
			return
		}
	}

	// success get all data users
	response := global.ApiResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success get all data users",
		Data:       users,
	}

	responseJson, _ := json.Marshal(&response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseJson)
}

// handler get user by id
func (u *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var request dto.UserIdRequest

	// if error when decode request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := global.ApiResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
			Message:    err.Error(),
		}
		_ = json.NewEncoder(w).Encode(&response)
		return
	}

	// call procedure in service
	users, err := u.UserService.GetById(r.Context(), &request)
	if err != nil {
		switch {
		case errors.As(err, &myError.NotFound):
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(global.ApiResponse{
				StatusCode: http.StatusNotFound,
				Status:     "not found",
				Message:    err.Error(),
			})
			return

		case errors.As(err, &myError.BadRequest):
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(global.ApiResponse{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    err.Error(),
			})
			return

		case errors.As(err, &myError.InternalServer):
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(global.ApiResponse{
				StatusCode: http.StatusInternalServerError,
				Status:     "internal server error",
				Message:    err.Error(),
			})
			return
		}
	}

	// success get data by id
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(global.ApiResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success get data by id",
		Data:       users,
	})
}
