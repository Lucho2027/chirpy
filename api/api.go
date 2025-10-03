package api

import (
	"fmt"
	"net/http"
)


const filepathRoot="."

func RegisterRoutes(mux *http.ServeMux, apiCfg *ApiConfig ){
		
	mux.Handle("/app/", apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz",  HandleReadiness)
	
	mux.HandleFunc("GET /admin/metrics", apiCfg.HandlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandleReset)
	// Users
	mux.HandleFunc("POST /api/users", apiCfg.HandleCreateUser )
	mux.HandleFunc("POST /api/login", apiCfg.HandleLogin)

	mux.HandleFunc("PUT /api/users", apiCfg.HandleUpdateUser)
	// Chirps
	mux.HandleFunc("POST /api/chirps", apiCfg.HandleCreateChirp)
	mux.HandleFunc("GET /api/chirps", apiCfg.HandleGetAll)
    mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.HandleGetChirpById)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.HandleDeleteChirpById)
    //Refresh Tokens
	mux.HandleFunc("POST /api/refresh", apiCfg.RefreshToken)
	mux.HandleFunc("POST /api/revoke", apiCfg.RevokeToken)

	// Webhooks
	mux.HandleFunc("POST /api/polka/webhooks", apiCfg.HandleWebhook)

}


func (cfg *ApiConfig) HandlerMetrics(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, cfg.FileserverHits.Load())))
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cfg.FileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}