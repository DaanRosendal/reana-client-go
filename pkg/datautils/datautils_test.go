/*
This file is part of REANA.
Copyright (C) 2022 CERN.

REANA is free software; you can redistribute it and/or modify it
under the terms of the MIT License; see LICENSE file for more details.
*/

package datautils

import (
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

func TestHasAnyPrefix(t *testing.T) {
	tests := map[string]struct {
		prefixes []string
		str      string
		want     bool
	}{
		"one prefix":               {prefixes: []string{"foo"}, str: "foobar", want: true},
		"wrong prefix":             {prefixes: []string{"foo"}, str: "bar", want: false},
		"exact word":               {prefixes: []string{"foo", "bar"}, str: "foo", want: true},
		"two prefix options":       {prefixes: []string{"foo", "bar"}, str: "foobar", want: true},
		"wrong prefix two options": {prefixes: []string{"foo", "bar"}, str: "baz", want: false},
		"no options":               {prefixes: []string{}, str: "foobar", want: false},
		"empty string":             {prefixes: []string{"foo", "bar"}, str: "", want: false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := HasAnyPrefix(test.str, test.prefixes)
			if got != test.want {
				t.Errorf("Expected %t, got %t", test.want, got)
			}
		})
	}
}

func TestFromIsoToTimestamp(t *testing.T) {
	now := time.Now().UTC()
	tests := map[string]struct {
		dateIso   string
		wantError bool
		want      time.Time
	}{
		"regular date": {
			dateIso: "2020-01-01T03:16:45",
			want:    time.Date(2020, 1, 1, 3, 16, 45, 0, time.UTC),
		},
		"current timestamp": {
			dateIso: now.Format("2006-01-02T15:04:05"),
			want:    now.Truncate(time.Second),
		},
		"wrong format": {
			dateIso:   "2020-01-01T00:00:00Z",
			wantError: true,
		},
		"only time": {
			dateIso:   "09:30:00Z",
			wantError: true,
		},
		"empty string": {
			dateIso:   "",
			wantError: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := FromIsoToTimestamp(test.dateIso)
			if test.wantError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if got != test.want {
					t.Errorf("Expected %v, got %v", test.want, got)
				}
			}
		})
	}
}

func TestSplitLinesNoEmpty(t *testing.T) {
	tests := map[string]struct {
		arg  string
		want []string
	}{
		"empty string":        {arg: "", want: []string{}},
		"only one line":       {arg: "a", want: []string{"a"}},
		"two lines":           {arg: "a\nb", want: []string{"a", "b"}},
		"three lines":         {arg: "a\nb\nc", want: []string{"a", "b", "c"}},
		"ending with newline": {arg: "a\nb\nc\n", want: []string{"a", "b", "c"}},
		"with empty lines":    {arg: "a\n\nb\n\nc", want: []string{"a", "b", "c"}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := SplitLinesNoEmpty(test.arg)
			if !slices.Equal(got, test.want) {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
