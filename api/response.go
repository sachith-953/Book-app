package api

import "encoding/json"

// Helper function to format response in JSON
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
