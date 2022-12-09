package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	irepository "imersao11-consolidacao/internal/domain/repository"
	"imersao11-consolidacao/internal/infra/db"
	"imersao11-consolidacao/internal/infra/presenter"
)

func ListPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := queries.FindAllPlayers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(players)
	}
}

// list my team players
func ListMyTeamPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamID := strings.TrimPrefix(r.URL.Path, "/my-teams/")[0:1]
		players, err := queries.GetPlayersByMyTeamID(ctx, teamID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(players)
	}
}

// list all matches
func ListMatchesHandler(ctx context.Context, matchRepository irepository.MatchRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		matches, err := matchRepository.FindAll(ctx)
		var matchesPresenter presenter.Matches
		for _, match := range matches {
			matchesPresenter = append(matchesPresenter, presenter.NewMatchPresenter(match))
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(matchesPresenter)
	}
}

// list match by id
func ListMatchByIDHandler(ctx context.Context, matchRepository irepository.MatchRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchID := strings.TrimPrefix(r.URL.Path, "/matches/")[0:1]
		match, err := matchRepository.FindByID(ctx, matchID)
		matchPresenter := presenter.NewMatchPresenter(match)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(matchPresenter)
	}
}

// get my team balance
func GetMyTeamBalanceHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamID := strings.TrimPrefix(r.URL.Path, "/my-teams/")[0:1]
		balance, err := queries.GetMyTeamBalance(ctx, teamID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resultJson := map[string]float64{"balance": balance}
		json.NewEncoder(w).Encode(resultJson)
	}
}
