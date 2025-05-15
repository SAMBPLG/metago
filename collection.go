package metago

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Collection represents metabase collection sdk
type collection struct {
	sdk *Metago
}

// Collection represents metabase collection model
type Collection struct {
	ParentID           any          `json:"parent_id,omitempty"`
	Name               string       `json:"name"`
	AuthorityLevel     string       `json:"authority_level"`
	Description        string       `json:"description"`
	Archived           bool         `json:"archived"`
	Slug               string       `json:"slug"`
	ArchiveOperationID string       `json:"archive_operation_id"`
	PersonalOwnerID    uint64       `json:"personal_owner_id"`
	Type               string       `json:"type"`
	IsSample           bool         `json:"is_sample"`
	ID                 any          `json:"id"`
	ArchivedDirectly   bool         `json:"archived_directly"`
	EntityID           string       `json:"entity_id"`
	Location           string       `json:"location"`
	EffectiveLocation  string       `json:"effective_location"`
	Namespace          string       `json:"namespace"`
	IsPersonal         bool         `json:"is_personal"`
	IsRoot             bool         `json:"metabase.models.collection.root/is-root?"`
	CreatedAt          time.Time    `json:"created_at"`
	EffectiveAchestors []Collection `json:"effective_ancestors"`
	Children           []Collection `json:"children,omitempty"`
	Here               []string     `json:"here,omitempty"`
	Below              []string     `json:"below,omitempty"`
	UserPermission
}

func (c *Collection) MarshalJSON() ([]byte, error) {
	m := map[string]any{"name": c.Name}
	if c.ParentID != nil {
		m["parent_id"] = c.ParentID
	}
	if c.AuthorityLevel != "" {
		m["authority_level"] = c.AuthorityLevel
	}
	if c.Description != "" {
		m["description"] = c.Description
	}
	if c.Namespace != "" {
		m["namespace"] = c.Namespace
	}
	return json.Marshal(m)
}

// CollectionParameters represents metabase collection api query parameters
type CollectionParameters struct {
	Archived                    bool
	ExcludeOtherUserCollections bool
	Namespace                   string
	PersonalOnly                bool
}

// Values returns url.Values representation of CollectionParameters
func (c CollectionParameters) Values() url.Values {
	q := url.Values{}
	q.Set("archived", strconv.FormatBool(c.Archived))
	q.Set("exclude-other-user-collections", strconv.FormatBool(c.ExcludeOtherUserCollections))
	q.Set("personal-only", strconv.FormatBool(c.PersonalOnly))
	if c.Namespace != "" {
		q.Set("namespace", c.Namespace)
	}
	return q
}

// TreeParameters represents Collection tree request parameters
type TreeParameters struct {
	ExcludeArchived             bool
	ExcludeOtherUserCollections bool
	Namespace                   string
	Shallow                     bool
	CollectionID                uint64
}

// Values returns url.Values representation of CollectionParameters
func (c TreeParameters) Values() url.Values {
	q := url.Values{}
	q.Set("exclude-archived", strconv.FormatBool(c.ExcludeArchived))
	q.Set("exclude-other-user-collections", strconv.FormatBool(c.ExcludeOtherUserCollections))
	q.Set("shallow", strconv.FormatBool(c.Shallow))
	if c.Namespace != "" {
		q.Set("namespace", c.Namespace)
	}
	if c.CollectionID > 0 {
		q.Set("collection-id", fmt.Sprintf("%v", c.CollectionID))
	}
	return q
}

// Collections fetch a list of all Collections that the current user has read permissions for (:can_write is returned as an additional property of each Collection so you can tell which of these you have write permissions for.)
// https://www.metabase.com/docs/latest/api#tag/apicollection
func (m *collection) Collections(ctx context.Context, opts CollectionParameters, result *[]Collection) error {
	var errResp MetabaseErr
	resp, err := m.sdk.getRequestWithAuth(ctx).
		SetError(&errResp).
		SetResult(result).
		SetQueryParamsFromValues(opts.Values()).
		Get(m.sdk.endpoint.CollectionEndpoint)
	if err := checkForError(resp, err, "failed to fetch collections"); err != nil {
		return err
	}
	return nil
}

// Return the 'Root' Collection object with standard details added
// https://www.metabase.com/docs/latest/api#tag/apicollection/GET/api/collection/root
func (m *collection) Root(ctx context.Context, opts CollectionParameters, result *Collection) error {
	var errResp MetabaseErr
	resp, err := m.sdk.getRequestWithAuth(ctx).
		SetError(&errResp).
		SetResult(result).
		SetQueryParamsFromValues(opts.Values()).
		Get(JoinPath(URLSeparator, m.sdk.endpoint.CollectionEndpoint, "root"))
	if err := checkForError(resp, err, "failed to fetch collection root"); err != nil {
		return err
	}
	return nil
}

// Tree is similar to Collections, but returns Collections in a tree structure.
// https://www.metabase.com/docs/latest/api#tag/apicollection/GET/api/collection/tree
func (m *collection) Tree(ctx context.Context, opts TreeParameters, result *[]Collection) error {
	var errResp MetabaseErr
	resp, err := m.sdk.getRequestWithAuth(ctx).
		SetError(&errResp).
		SetResult(result).
		SetQueryParamsFromValues(opts.Values()).
		Get(JoinPath(URLSeparator, m.sdk.endpoint.CollectionEndpoint, "tree"))
	if err := checkForError(resp, err, "failed to fetch collections tree"); err != nil {
		return err
	}
	return nil
}

// Get fetch a specific Collection with standard details added.
// https://www.metabase.com/docs/latest/api#tag/apicollection/GET/api/collection/{id}
func (m *collection) Get(ctx context.Context, id string, result *Collection) error {
	var errResp MetabaseErr
	resp, err := m.sdk.getRequestWithAuth(ctx).
		SetError(&errResp).
		SetResult(result).
		SetPathParam("collection_id", id).
		Get(JoinPath(URLSeparator, m.sdk.endpoint.CollectionEndpoint, "{collection_id}"))
	if err := checkForError(resp, err, "failed to fetch collection"); err != nil {
		return err
	}
	return nil
}

// Create create a new Collection.
// https://www.metabase.com/docs/latest/api#tag/apicollection/POST/api/collection/
func (m *collection) Create(ctx context.Context, req *Collection) error {
	var errResp MetabaseErr
	resp, err := m.sdk.getRequestWithAuth(ctx).
		SetBody(req).
		SetError(&errResp).
		SetResult(req).
		Post(m.sdk.endpoint.CollectionEndpoint)
	if err := checkForError(resp, err, "failed to create collection"); err != nil {
		return err
	}
	return nil
}
