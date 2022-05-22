package gateway_layer

import (
	research_works_info_controllers "Backend/controllers/research_works_info"
	research_works_info_models "Backend/models/research_works_info"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddResearchWorksInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-research-work", func(w http.ResponseWriter, r *http.Request) {
		var request research_works_info_models.CreateResearchWorkRequest
		json.NewDecoder(r.Body).Decode(&request)

		//isValidToken, _, role := helpers.ValidateToken(request.Token)
		//
		//if !isValidToken {
		//	w.WriteHeader(401)
		//	return
		//}
		//
		//if !(role == types.Administrator) {
		//	w.WriteHeader(403)
		//	return
		//}

		err := research_works_info_controllers.CreateResearchWork(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/get-user-research-works", func(w http.ResponseWriter, r *http.Request) {
		var request research_works_info_models.GetUserResearchWorksRequest
		json.NewDecoder(r.Body).Decode(&request)

		//isValidToken, login, role := helpers.ValidateToken(request.Token)
		//
		//if !isValidToken {
		//	w.WriteHeader(401)
		//	return
		//}
		//
		//if !(role == types.Administrator || login != request.Login) {
		//	w.WriteHeader(403)
		//	return
		//}

		response, err := research_works_info_controllers.GetUserResearchWorks(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})

	http.HandleFunc("/get-research-work-info", func(w http.ResponseWriter, r *http.Request) {
		var request research_works_info_models.GetResearchWorkInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		//isValidToken, _, role := helpers.ValidateToken(request.Token)
		//
		//if !isValidToken {
		//	w.WriteHeader(401)
		//	return
		//}
		//
		//if !(role == types.Administrator) {
		//	w.WriteHeader(403)
		//	return
		//}

		response, err := research_works_info_controllers.GetResearchWorkInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})

	http.HandleFunc("/update-research-work-info", func(w http.ResponseWriter, r *http.Request) {
		var request research_works_info_models.UpdateResearchWorkInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		//isValidToken, login, role := helpers.ValidateToken(request.Token)
		//
		//if !isValidToken {
		//	w.WriteHeader(401)
		//	return
		//}
		//
		//if !(role == types.Administrator || login != request.Login) {
		//	w.WriteHeader(403)
		//	return
		//}

		err := research_works_info_controllers.UpdateResearchWorkInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
	})
}
