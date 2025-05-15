package metago

import (
	"context"

	"github.com/go-resty/resty/v2"
)

var URLSeparator string = "/"

// Option represents new metabase client instance option
type Option struct {
	// Your metabase address : http://my-metabase.com
	BasePath   string
	AuthMethod AuthMethod
	APIKey     string
	Username   string
	Password   string
	Context    context.Context
}

// Metago represents metabase API client SDK
type Metago struct {
	basePath    string // your metabase address
	authMethod  AuthMethod
	session     *string // session id or api key
	restyClient *resty.Client
	endpoint    Endpoint
	Action      *action
	Session     *session
	Collection  *collection
}

type Endpoint struct {
	ActionEndpoint             string
	ActivityEndpoint           string
	AlertEndpoint              string
	ApiKeyEndpoint             string
	AutomagicDashboardEndpoint string
	BookmarkEndpoint           string
	CacheEndpoint              string
	CardEnpoint                string
	CardsEdnpoint              string
	ChannelEndpoint            string
	CloudMigrationEndpoint     string
	CollectionEndpoint         string
	DashboardEndpoint          string
	DatabaseEndpoint           string
	DatasetEndpoint            string
	EmailEndpoint              string
	EmbedEndpoint              string
	FieldEndpoint              string
	GeojsonEndpoint            string
	GoogleEndpoint             string
	LdapEndpoint               string
	LoginHistoryEndpoint       string
	ModelIndexEndpoint         string
	NativeQuerySnippetEndpoint string
	NotificationEndpoint       string
	NotifyEndpoint             string
	PermissionsEndpoint        string
	PersistEndpoint            string
	PremiumFeaturesEndpoint    string
	PreviewEmbedEndpoint       string
	PublicEndpoint             string
	PulseEndpoint              string
	RevisionEndpoint           string
	SearchEndpoint             string
	SegmentEndpoint            string
	SessionEndpoint            string
	SettingEndpoint            string
	SetupEndpoint              string
	SlackEndpoint              string
	TableEndpoint              string
	TaskEndpoint               string
	TilesEndpoint              string
	TimelineEventEndpoint      string
	TimelineEndpoint           string
	UserKeyValueEndpoint       string
	UserEndpoint               string
	UtilEndpoint               string
}

// New returns new Metago instance
func New(opt Option, opts ...func(*Metago)) (*Metago, error) {
	m := &Metago{
		basePath:    opt.BasePath,
		authMethod:  opt.AuthMethod,
		restyClient: resty.New(),
		endpoint: Endpoint{
			ActionEndpoint:             JoinPath(URLSeparator, opt.BasePath, "api", "action"),
			ActivityEndpoint:           JoinPath(URLSeparator, opt.BasePath, "api", "activity"),
			AlertEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "alert"),
			ApiKeyEndpoint:             JoinPath(URLSeparator, opt.BasePath, "api", "api-key"),
			AutomagicDashboardEndpoint: JoinPath(URLSeparator, opt.BasePath, "api", "automagic-dashboard"),
			BookmarkEndpoint:           JoinPath(URLSeparator, opt.BasePath, "api", "bookmark"),
			CacheEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "cache"),
			CardEnpoint:                JoinPath(URLSeparator, opt.BasePath, "api", "card"),
			CardsEdnpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "cards"),
			ChannelEndpoint:            JoinPath(URLSeparator, opt.BasePath, "api", "channel"),
			CloudMigrationEndpoint:     JoinPath(URLSeparator, opt.BasePath, "api", "cloud-migration"),
			CollectionEndpoint:         JoinPath(URLSeparator, opt.BasePath, "api", "collection"),
			DashboardEndpoint:          JoinPath(URLSeparator, opt.BasePath, "api", "dashboard"),
			DatabaseEndpoint:           JoinPath(URLSeparator, opt.BasePath, "api", "database"),
			DatasetEndpoint:            JoinPath(URLSeparator, opt.BasePath, "api", "dataset"),
			EmailEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "email"),
			EmbedEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "embed"),
			FieldEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "field"),
			GeojsonEndpoint:            JoinPath(URLSeparator, opt.BasePath, "api", "geojson"),
			GoogleEndpoint:             JoinPath(URLSeparator, opt.BasePath, "api", "google"),
			LdapEndpoint:               JoinPath(URLSeparator, opt.BasePath, "api", "ldap"),
			LoginHistoryEndpoint:       JoinPath(URLSeparator, opt.BasePath, "api", "login-history"),
			ModelIndexEndpoint:         JoinPath(URLSeparator, opt.BasePath, "api", "model-index"),
			NativeQuerySnippetEndpoint: JoinPath(URLSeparator, opt.BasePath, "api", "native-query-snippet"),
			NotificationEndpoint:       JoinPath(URLSeparator, opt.BasePath, "api", "notification"),
			NotifyEndpoint:             JoinPath(URLSeparator, opt.BasePath, "api", "notify"),
			PermissionsEndpoint:        JoinPath(URLSeparator, opt.BasePath, "api", "permissions"),
			PersistEndpoint:            JoinPath(URLSeparator, opt.BasePath, "api", "persist"),
			PremiumFeaturesEndpoint:    JoinPath(URLSeparator, opt.BasePath, "api", "premium-features"),
			PreviewEmbedEndpoint:       JoinPath(URLSeparator, opt.BasePath, "api", "preview_embed"),
			PublicEndpoint:             JoinPath(URLSeparator, opt.BasePath, "api", "public"),
			PulseEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "pulse"),
			RevisionEndpoint:           JoinPath(URLSeparator, opt.BasePath, "api", "revision"),
			SearchEndpoint:             JoinPath(URLSeparator, opt.BasePath, "api", "search"),
			SessionEndpoint:            JoinPath(URLSeparator, opt.BasePath, "api", "session"),
			SettingEndpoint:            JoinPath(URLSeparator, opt.BasePath, "api", "setting"),
			SetupEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "setup"),
			SlackEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "slack"),
			TableEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "table"),
			TaskEndpoint:               JoinPath(URLSeparator, opt.BasePath, "api", "task"),
			TilesEndpoint:              JoinPath(URLSeparator, opt.BasePath, "api", "tiles"),
			TimelineEndpoint:           JoinPath(URLSeparator, opt.BasePath, "api", "timeline"),
			TimelineEventEndpoint:      JoinPath(URLSeparator, opt.BasePath, "api", "timeline-event"),
			UserKeyValueEndpoint:       JoinPath(URLSeparator, opt.BasePath, "api", "user-key-value"),
			UserEndpoint:               JoinPath(URLSeparator, opt.BasePath, "api", "user"),
			UtilEndpoint:               JoinPath(URLSeparator, opt.BasePath, "api", "util"),
		},
	}
	if err := m.getSession(opt.Context, opt.APIKey, opt.Username, opt.Password); err != nil {
		return nil, err
	}

	m.Action = &action{sdk: m}
	m.Session = &session{sdk: m}
	m.Collection = &collection{sdk: m}

	for _, option := range opts {
		option(m)
	}

	return m, nil
}

// NewWithRestyClient returns new Metago instance with by setting
// up resty client from the given parameter.
func NewWithRestyClient(opt Option, r *resty.Client) (*Metago, error) {
	return New(opt, func(m *Metago) {
		m.restyClient = r
	})
}
