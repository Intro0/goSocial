package store

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginatedFeedQuery struct {
	Limit  int      `json:"limit" validate:"gte=1,lte=20"`
	Offset int      `json:"offset" validate:"gte=0"`
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Tags   []string `json:"tags" validate:"max=5"`
	Search string   `json:"search" validate:"max=100"`
	Since  string   `json:"since"`
	Until  string   `json:"until"`
}

func (fq PaginatedFeedQuery) Parse(r *http.Request) (PaginatedFeedQuery, error) {
	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, fmt.Errorf("parsing limit: %w", err)
		}

		fq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return fq, fmt.Errorf("parsing offset: %w", err)
		}

		fq.Offset = l
	}

	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	tags := qs.Get("tags")
	if tags != "" {
		fq.Tags = strings.Split(tags, ",")
	}

	search := qs.Get("search")
	if search != "" {
		fq.Search = search
	}

	since := qs.Get("since")
	if since != "" {
		parsedSince, err := parseTime(since)
		if err != nil {
			return fq, fmt.Errorf("parsing since: %w", err)
		}
		fq.Since = parsedSince
	}

	until := qs.Get("until")
	if until != "" {
		parsedUntil, err := parseTime(until)
		if err != nil {
			return fq, fmt.Errorf("parsing until: %w", err)
		}
		fq.Until = parsedUntil
	}

	return fq, nil
}

func parseTime(s string) (string, error) {
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return "", err
	}

	return t.Format(time.DateTime), nil
}
