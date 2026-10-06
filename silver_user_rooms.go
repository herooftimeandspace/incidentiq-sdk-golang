package incidentiq

import (
	"context"
	"net/http"
	"strings"
)

const (
	userRoomsPath = "/api/v1.0/users/{user_id}/rooms"
	userRoomPath  = "/api/v1.0/users/{user_id}/rooms/{location_room_id}"
)

// AddUserRoom adds one location-room association to an Incident IQ user.
//
// This is a Silver live-client capability rather than a Golden Stoplight
// operation. It deliberately disables Silver's automatic retry-without-Client
// fallback because replaying this POST after an ambiguous rejection could
// duplicate a write. Callers can explicitly set OmitClientHeader when they have
// tenant-specific evidence that the route requires the browser header shape.
func (s *SilverUsersService) AddUserRoom(ctx context.Context, userID string, locationRoomID string, opts RequestOptions, out any) error {
	if err := validateUserRoomIDs(userID, locationRoomID); err != nil {
		return err
	}
	if opts.JSON != nil || opts.Body != nil {
		return &ValidationError{Message: "AddUserRoom does not accept a request body"}
	}
	opts.PathParams = userRoomPathParams(opts.PathParams, userID, locationRoomID)
	return s.client.request(ctx, http.MethodPost, userRoomPath, opts, out, false)
}

// SetUserRooms replaces every location-room association for an Incident IQ
// user with locationRoomIDs.
//
// The method owns the JSON request body so unrelated user fields cannot be
// mutated accidentally. A nil slice is rejected; an empty non-nil slice is an
// explicit request to clear all room associations. Like AddUserRoom, this POST
// does not use Silver's automatic Client-header fallback.
func (s *SilverUsersService) SetUserRooms(ctx context.Context, userID string, locationRoomIDs []string, opts RequestOptions, out any) error {
	if strings.TrimSpace(userID) == "" {
		return &ValidationError{Message: "user ID is required"}
	}
	if locationRoomIDs == nil {
		return &ValidationError{Message: "location room IDs must be a non-nil slice"}
	}
	for _, locationRoomID := range locationRoomIDs {
		if strings.TrimSpace(locationRoomID) == "" {
			return &ValidationError{Message: "location room IDs cannot contain an empty ID"}
		}
	}
	if opts.JSON != nil || opts.Body != nil {
		return &ValidationError{Message: "SetUserRooms owns the JSON request body"}
	}
	opts.PathParams = userRoomPathParams(opts.PathParams, userID, "")
	opts.JSON = locationRoomIDs
	return s.client.request(ctx, http.MethodPost, userRoomsPath, opts, out, false)
}

// RemoveUserRoom removes one location-room association from an Incident IQ
// user. DELETE follows the shared idempotent retry policy, but does not use the
// Silver retry-without-Client fallback.
func (s *SilverUsersService) RemoveUserRoom(ctx context.Context, userID string, locationRoomID string, opts RequestOptions, out any) error {
	if err := validateUserRoomIDs(userID, locationRoomID); err != nil {
		return err
	}
	if opts.JSON != nil || opts.Body != nil {
		return &ValidationError{Message: "RemoveUserRoom does not accept a request body"}
	}
	opts.PathParams = userRoomPathParams(opts.PathParams, userID, locationRoomID)
	return s.client.request(ctx, http.MethodDelete, userRoomPath, opts, out, false)
}

// validateUserRoomIDs rejects missing identifiers before transport work begins.
func validateUserRoomIDs(userID string, locationRoomID string) error {
	if strings.TrimSpace(userID) == "" {
		return &ValidationError{Message: "user ID is required"}
	}
	if strings.TrimSpace(locationRoomID) == "" {
		return &ValidationError{Message: "location room ID is required"}
	}
	return nil
}

// userRoomPathParams copies caller-owned parameters before adding route IDs.
func userRoomPathParams(pathParams map[string]any, userID string, locationRoomID string) map[string]any {
	result := make(map[string]any, len(pathParams)+2)
	for key, value := range pathParams {
		result[key] = value
	}
	result["user_id"] = userID
	if locationRoomID != "" {
		result["location_room_id"] = locationRoomID
	}
	return result
}
