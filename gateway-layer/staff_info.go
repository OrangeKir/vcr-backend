package gateway_layer

import (
	staff_info_controllers "Backend/controllers/staff_info"
	"Backend/helpers"
	staff_info_models "Backend/models/staff_info"
	"Backend/types"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddStaffInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-staff", func(w http.ResponseWriter, r *http.Request) {
		var request staff_info_models.CreateStaffRequest
		json.NewDecoder(r.Body).Decode(&request)

		if request.PositionType != staff_info_models.Main {
			combinationRate, err := staff_info_controllers.GetSummaryCombinationRates(conn, request.PersonLogin)

			if err != nil {
				w.WriteHeader(500)
				return
			}

			if combinationRate+request.WorkingRate > 0.5 {
				w.WriteHeader(400)
				return
			}
		}

		isValidToken, _, role := helpers.ValidateToken(request.Token)

		if !isValidToken {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		if !(role == types.Administrator) {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		err := staff_info_controllers.CreateStaff(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})

	http.HandleFunc("/update-staff-info", func(w http.ResponseWriter, r *http.Request) {
		var request staff_info_models.UpdateStaffInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		isValidToken, _, role := helpers.ValidateToken(request.Token)

		if !isValidToken {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		if !(role == types.Administrator || role == types.DepartmentHead) {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		err := staff_info_controllers.UpdateStaffInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})

	http.HandleFunc("/get-staff-info", func(w http.ResponseWriter, r *http.Request) {
		var request staff_info_models.GetStaffInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		response, err := staff_info_controllers.GetStaffInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})
}
