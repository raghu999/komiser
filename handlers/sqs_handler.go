package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) SQSQueuesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sqs")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeQueues(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "sqs:ListQueues is missing")
		} else {
			handler.cache.Set("sqs", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
