// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileconsumer

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/operatortest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/testutil"
)

func TestUnmarshal(t *testing.T) {
	operatortest.ConfigUnmarshalTests{
		DefaultConfig: newMockOperatorConfig(NewConfig()),
		TestsFile:     filepath.Join(".", "testdata", "config.yaml"),
		Tests: []operatortest.ConfigUnmarshalTest{
			{
				Name: "include_one",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "one.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_multi",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "one.log", "two.log", "three.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob_double_asterisk",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "**.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob_double_asterisk_nested",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "directory/**/*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_glob_double_asterisk_prefix",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "**/directory/**/*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_inline",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "a.log", "b.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "include_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "aString")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_one",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "one.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_multi",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "one.log", "two.log", "three.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "not*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob_double_asterisk",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "not**.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob_double_asterisk_nested",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "directory/**/not*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_glob_double_asterisk_prefix",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "**/directory/**/not*.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_inline",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "a.log", "b.log")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "exclude_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.Include = append(cfg.Include, "*.log")
					cfg.Exclude = append(cfg.Exclude, "aString")
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_no_units",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Second
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_1s",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Second
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_1ms",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Millisecond
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "poll_interval_1000ms",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.PollInterval = time.Second
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_no_units",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1000)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1kb_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1000)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1KB",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1000)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1kib_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1024)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_1KiB",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1024)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "fingerprint_size_float",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.FingerprintSize = helper.ByteSize(1100)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "multiline_line_start_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newMultiline := helper.NewMultilineConfig()
					newMultiline.LineStartPattern = "Start"
					return newMockOperatorConfigWithMultiline(cfg, newMultiline)
				}(),
			},
			{
				Name: "multiline_line_start_special",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newMultiline := helper.NewMultilineConfig()
					newMultiline.LineStartPattern = "%"
					return newMockOperatorConfigWithMultiline(cfg, newMultiline)
				}(),
			},
			{
				Name: "multiline_line_end_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newMultiline := helper.NewMultilineConfig()
					newMultiline.LineEndPattern = "Start"
					return newMockOperatorConfigWithMultiline(cfg, newMultiline)
				}(),
			},
			{
				Name: "multiline_line_end_special",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					newMultiline := helper.NewMultilineConfig()
					newMultiline.LineEndPattern = "%"
					return newMockOperatorConfigWithMultiline(cfg, newMultiline)
				}(),
			},
			{
				Name: "start_at_string",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.StartAt = "beginning"
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_concurrent_large",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxConcurrentFiles = 9223372036854775807
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mib_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mib_upper",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mb_upper",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "max_log_size_mb_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.MaxLogSize = helper.ByteSize(1048576)
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "encoding_lower",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.EncodingConfig = helper.EncodingConfig{Encoding: "utf-16le"}
					return newMockOperatorConfig(cfg)
				}(),
			},
			{
				Name: "encoding_upper",
				Expect: func() *mockOperatorConfig {
					cfg := NewConfig()
					cfg.EncodingConfig = helper.EncodingConfig{Encoding: "UTF-16lE"}
					return newMockOperatorConfig(cfg)
				}(),
			},
		},
	}.Run(t)
}

func TestBuild(t *testing.T) {
	t.Parallel()

	basicConfig := func() *Config {
		cfg := NewConfig()
		cfg.Include = []string{"/var/log/testpath.*"}
		cfg.Exclude = []string{"/var/log/testpath.ex*"}
		cfg.PollInterval = 10 * time.Millisecond
		return cfg
	}

	cases := []struct {
		name             string
		modifyBaseConfig func(*Config)
		MultilineConfig  helper.MultilineConfig
		errorRequirement require.ErrorAssertionFunc
		validate         func(*testing.T, *Manager)
	}{
		{
			"Basic",
			func(f *Config) {},
			helper.NewMultilineConfig(),
			require.NoError,
			func(t *testing.T, f *Manager) {
				require.Equal(t, f.finder.Include, []string{"/var/log/testpath.*"})
				require.Equal(t, f.pollInterval, 10*time.Millisecond)
			},
		},
		{
			"BadIncludeGlob",
			func(f *Config) {
				f.Include = []string{"["}
			},
			helper.NewMultilineConfig(),
			require.Error,
			nil,
		},
		{
			"BadExcludeGlob",
			func(f *Config) {
				f.Include = []string{"["}
			},
			helper.NewMultilineConfig(),
			require.Error,
			nil,
		},
		{
			"MultilineConfiguredStartAndEndPatterns",
			func(cfg *Config) {
				cfg.EncodingConfig = helper.NewEncodingConfig()
				cfg.Flusher = helper.NewFlusherConfig()
			},
			helper.MultilineConfig{
				LineEndPattern:   "Exists",
				LineStartPattern: "Exists",
			},
			require.Error,
			nil,
		},
		{
			"MultilineConfiguredStartPattern",
			func(f *Config) {
				f.EncodingConfig = helper.NewEncodingConfig()
				f.Flusher = helper.NewFlusherConfig()
			},
			helper.MultilineConfig{
				LineStartPattern: "START.*",
			},
			require.NoError,
			func(t *testing.T, f *Manager) {},
		},
		{
			"MultilineConfiguredEndPattern",
			func(f *Config) {
				f.EncodingConfig = helper.NewEncodingConfig()
				f.Flusher = helper.NewFlusherConfig()
			},
			helper.MultilineConfig{
				LineEndPattern: "END.*",
			},
			require.NoError,
			func(t *testing.T, f *Manager) {},
		},
		{
			"InvalidEncoding",
			func(f *Config) {
				f.EncodingConfig = helper.EncodingConfig{Encoding: "UTF-3233"}
			},
			helper.NewMultilineConfig(),
			require.Error,
			nil,
		},
		{
			"LineStartAndEnd",
			func(f *Config) {
				f.EncodingConfig = helper.NewEncodingConfig()
				f.Flusher = helper.NewFlusherConfig()
			},
			helper.MultilineConfig{
				LineStartPattern: ".*",
				LineEndPattern:   ".*",
			},
			require.Error,
			nil,
		},
		{
			"NoLineStartOrEnd",
			func(f *Config) {
				f.EncodingConfig = helper.NewEncodingConfig()
				f.Flusher = helper.NewFlusherConfig()
			},
			helper.NewMultilineConfig(),
			require.NoError,
			func(t *testing.T, f *Manager) {},
		},
		{
			"InvalidLineStartRegex",
			func(f *Config) {
				f.EncodingConfig = helper.NewEncodingConfig()
				f.Flusher = helper.NewFlusherConfig()
			},
			helper.MultilineConfig{
				LineStartPattern: "(",
			},
			require.Error,
			nil,
		},
		{
			"InvalidLineEndRegex",
			func(f *Config) {
				f.EncodingConfig = helper.NewEncodingConfig()
				f.Flusher = helper.NewFlusherConfig()
			},
			helper.MultilineConfig{
				LineEndPattern: "(",
			},
			require.Error,
			nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			cfg := basicConfig()
			tc.modifyBaseConfig(cfg)

			nopEmit := func(_ context.Context, _ *FileAttributes, _ []byte) {}

			enc, err := cfg.EncodingConfig.Build()
			if err != nil {
				tc.errorRequirement(t, err)
				return
			}
			flusher := cfg.Flusher.Build()
			splitter, err := tc.MultilineConfig.Build(enc.Encoding, false, flusher, int(cfg.MaxLogSize))
			if err != nil {
				tc.errorRequirement(t, err)
				return
			}
			input, err := cfg.Build(testutil.Logger(t), nopEmit, WithCustomizedSplitter(splitter))
			tc.errorRequirement(t, err)
			if err != nil {
				return
			}

			tc.validate(t, input)
		})
	}
}
