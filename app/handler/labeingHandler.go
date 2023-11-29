package handler

// import (
// 	"GoogleProjectBackend/util"
// 	"encoding/json"
// 	"net/http"
// )

// func GetPatchSizeHandler(w http.ResponseWriter, r *http.Request) {
// 	util.EnableCors(&w)
// 	if r.Method == http.MethodGet {
// 		//각 URL에 알맞는 비디오 지정
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(models.PatchSize{
// 			TotalPatch:  100,
// 			WidthPatch:  10,
// 			HeightPatch: 10,
// 		})
// 	}
// }
