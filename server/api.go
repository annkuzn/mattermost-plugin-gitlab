package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
	"golang.org/x/oauth2"

	"github.com/annkuzn/mattermost-plugin-gitlab/server/gitlab"
)

const (
	APIErrorIDNotConnected = "not_connected"
)

type APIErrorResponse struct {
	ID         string `json:"id"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (p *Plugin) writeAPIError(w http.ResponseWriter, err *APIErrorResponse) {
	b, _ := json.Marshal(err)
	w.WriteHeader(err.StatusCode)
	if _, err := w.Write(b); err != nil {
		p.API.LogError("can't write api error http response", "err", err.Error())
	}
}

func (p *Plugin) writeAPIResponse(w http.ResponseWriter, resp interface{}) {
	b, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		p.API.LogError("Error encoding JSON response", "err", jsonErr.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Encountered an unexpected error. Please try again.", StatusCode: http.StatusInternalServerError})
	}
	if _, err := w.Write(b); err != nil {
		p.API.LogError("can't write response user to http", "err", err.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Encountered an unexpected error. Please try again.", StatusCode: http.StatusInternalServerError})
	}
}

type ConnectedResponse struct {
	Connected      bool                 `json:"connected"`
	GitlabUsername string               `json:"gitlab_username"`
	GitlabClientID string               `json:"gitlab_client_id"`
	GitlabURL      string               `json:"gitlab_url,omitempty"`
	Organization   string               `json:"organization"`
	Settings       *gitlab.UserSettings `json:"settings"`
}

type GitlabUserRequest struct {
	UserID string `json:"user_id"`
}

type GitlabUserResponse struct {
	Username string `json:"username"`
}

func (p *Plugin) getGitlabUser(w http.ResponseWriter, r *http.Request) {
	requestorID := r.Header.Get("Mattermost-User-ID")
	if requestorID == "" {
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Not authorized.", StatusCode: http.StatusUnauthorized})
		return
	}

	req := &GitlabUserRequest{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil || req.UserID == "" {
		if err != nil {
			p.API.LogError("Error decoding JSON body", "err", err.Error())
		}
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Please provide a JSON object with a non-blank user_id field.", StatusCode: http.StatusBadRequest})
		return
	}

	userInfo, apiErr := p.getGitlabUserInfoByMattermostID(req.UserID)
	if apiErr != nil {
		if apiErr.ID == APIErrorIDNotConnected {
			p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "User is not connected to a GitLab account.", StatusCode: http.StatusNotFound})
		} else {
			p.writeAPIError(w, apiErr)
		}
		return
	}

	if userInfo == nil {
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "User is not connected to a GitLab account.", StatusCode: http.StatusNotFound})
		return
	}

	p.writeAPIResponse(w, &GitlabUserResponse{Username: userInfo.GitlabUsername})
}

func (p *Plugin) getUnreads(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user, err := p.getGitlabUserInfoByMattermostID(userID)
	if err != nil {
		p.writeAPIError(w, err)
		return
	}

	result, errRequest := p.GitlabClient.GetUnreads(user)
	if errRequest != nil {
		p.API.LogError("unable to list unreads in GitLab API", "err", errRequest.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Unable to list unreads in GitLab API.", StatusCode: http.StatusInternalServerError})
		return
	}

	p.writeAPIResponse(w, result)
}

func (p *Plugin) getReviews(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user, err := p.getGitlabUserInfoByMattermostID(userID)

	if err != nil {
		p.writeAPIError(w, err)
		return
	}

	result, errRequest := p.GitlabClient.GetReviews(user)

	if errRequest != nil {
		p.API.LogError("unable to list merge-request where assignee in GitLab API", "err", errRequest.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Unable to list merge-request in GitLab API.", StatusCode: http.StatusInternalServerError})
		return
	}

	p.writeAPIResponse(w, result)
}

func (p *Plugin) getYourPrs(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user, err := p.getGitlabUserInfoByMattermostID(userID)

	if err != nil {
		p.writeAPIError(w, err)
		return
	}

	result, errRequest := p.GitlabClient.GetYourPrs(user)

	if errRequest != nil {
		p.API.LogError("can't list merge-request where author in GitLab API", "err", errRequest.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Unable to list merge-request in GitLab API.", StatusCode: http.StatusInternalServerError})
		return
	}

	p.writeAPIResponse(w, result)
}

func (p *Plugin) getYourAssignments(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user, err := p.getGitlabUserInfoByMattermostID(userID)

	if err != nil {
		p.writeAPIError(w, err)
		return
	}

	result, errRequest := p.GitlabClient.GetYourAssignments(user)

	if errRequest != nil {
		p.API.LogError("unable to list issue where assignee in GitLab API", "err", errRequest.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Unable to list issue in GitLab API.", StatusCode: http.StatusInternalServerError})
		return
	}

	p.writeAPIResponse(w, result)
}

func (p *Plugin) postToDo(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Not authorized.", StatusCode: http.StatusUnauthorized})
		return
	}

	user, err := p.getGitlabUserInfoByMattermostID(userID)

	if err != nil {
		p.writeAPIError(w, err)
		return
	}

	_, text, errRequest := p.GetToDo(user)
	if errRequest != nil {
		p.API.LogError("can't get todo", "err", errRequest.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Encountered an error getting the to do items.", StatusCode: http.StatusUnauthorized})
		return
	}

	if err := p.CreateBotDMPost(userID, text, "custom_git_todo"); err != nil {
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Encountered an error posting the to do items.", StatusCode: http.StatusUnauthorized})
	}

	p.writeAPIResponse(w, struct{ status string }{status: "OK"})
}

func (p *Plugin) updateSettings(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	var settings *gitlab.UserSettings
	err := json.NewDecoder(r.Body).Decode(&settings)
	if settings == nil || err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	info, errGitlab := p.getGitlabUserInfoByMattermostID(userID)
	if errGitlab != nil {
		p.writeAPIError(w, errGitlab)
		return
	}

	info.Settings = settings

	if err := p.storeGitlabUserInfo(info); err != nil {
		p.API.LogError("can't store GitLab user info when update settings", "err", err.Error())
		http.Error(w, "Encountered error updating settings", http.StatusInternalServerError)
	}

	p.writeAPIResponse(w, info.Settings)
}
