package metago

import (
	"context"
	"net/url"
	"strconv"
)

type CardF uint8

const (
	CardFilterArchived CardF = iota
	CardFilterTable
	CardFilterUsingModel
	CardFilterBookmarked
	CardFilterUsingSegment
	CardFilterAll
	CardFilterMine
	CardFilterUsingMetric
	CardFilterDatabase
)

func (c CardF) String() string {
	return [...]string{
		"archived", "table", "using_model",
		"bookmarked", "using_segment", "all",
		"mine", "using_metric", "database",
	}[c]
}

func (c CardF) EnumIndex() uint8 {
	return uint8(c)
}

// card represents metabase card api collection
type card struct {
	sdk *Metago
}

// Card represents metabase card model
type Card struct {
}

// CardParameters represents metabase fetch card collection query parameters
type CardParameters struct {
	F       CardF
	ModelID uint64
}

// Values returns url.Values representation of CardParameters
func (c CardParameters) Values() url.Values {
	q := url.Values{}
	q.Set("f", c.F.String())
	q.Set("model_id", strconv.FormatUint(c.ModelID, 32))
	return q
}

// Collections returns all the Cards. Option filter param f can be used to change the set of Cards that are returned.
// https://www.metabase.com/docs/latest/api#tag/apicard
func (m *card) Collections(ctx context.Context, opts CardParameters, result *[]Card) error {
	var errResp MetabaseErr
	resp, err := m.sdk.getRequestWithAuth(ctx).
		SetError(&errResp).
		SetResult(result).
		SetQueryParamsFromValues(opts.Values()).
		Get(m.sdk.endpoint.CardEnpoint)
	if err := checkForError(resp, err, "failed to fetch cards"); err != nil {
		return err
	}
	return nil
}
