package npmaudit

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAuditReport(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    *AuditReport
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				src: `
{
	"auditReportVersion": 2,
	"vulnerabilities": {
	},
	"metadata": {
		"vulnerabilities": {
			"info": 0,
			"low": 6,
			"moderate": 4,
			"high": 0,
			"critical": 0,
			"total": 10
		},
		"dependencies": {
			"prod": 388,
			"dev": 2044,
			"optional": 243,
			"peer": 0,
			"peerOptional": 0,
			"total": 2612
		}
	}
}`,
			},
			want: &AuditReport{
				AuditReportVersion: 2,
				Vulnerabilities:    map[string]Vulnerability{},
				Metadata: Metadata{
					Vulnerabilities: Vulnerabilities{
						Info:     0,
						Low:      6,
						Moderate: 4,
						High:     0,
						Critical: 0,
						Total:    10,
					},
					Dependencies: Dependencies{
						Prod:         388,
						Dev:          2044,
						Optional:     243,
						Peer:         0,
						PeerOptional: 0,
						Total:        2612,
					},
				},
			},
		},
		{
			name: "fix available exists",
			args: args{
				src: `
{
	"auditReportVersion": 2,
	"vulnerabilities": {
		"@vue/cli-plugin-e2e-cypress": {
			"name": "@vue/cli-plugin-e2e-cypress",
			"severity": "low",
			"via": [],
			"effects": [],
			"range": "<=4.5.12",
			"nodes": [
				"node_modules/@vue/cli-plugin-e2e-cypress"
			],
			"fixAvailable": false
		},
		"@vue/cli-service": {
			"name": "@vue/cli-service",
			"severity": "moderate",
			"via": [],
			"effects": [],
			"range": "4.0.0-alpha.0 - 4.5.12",
			"nodes": [
				"node_modules/@vue/cli-service"
			],
			"fixAvailable": {
				"name": "@vue/cli-service",
				"version": "4.1.1",
				"isSemVerMajor": true
			}
		}
	},
	"metadata": {
		"vulnerabilities": {
			"info": 0,
			"low": 6,
			"moderate": 4,
			"high": 0,
			"critical": 0,
			"total": 10
		},
		"dependencies": {
			"prod": 388,
			"dev": 2044,
			"optional": 243,
			"peer": 0,
			"peerOptional": 0,
			"total": 2612
		}
	}
}`,
			},
			want: &AuditReport{
				AuditReportVersion: 2,
				Vulnerabilities: map[string]Vulnerability{
					"@vue/cli-plugin-e2e-cypress": {
						Name:         "@vue/cli-plugin-e2e-cypress",
						Severity:     "low",
						Range:        "<=4.5.12",
						Nodes:        []string{"node_modules/@vue/cli-plugin-e2e-cypress"},
						FixAvailable: nil,
					},
					"@vue/cli-service": {
						Name:     "@vue/cli-service",
						Severity: "moderate",
						Range:    "4.0.0-alpha.0 - 4.5.12",
						Nodes:    []string{"node_modules/@vue/cli-service"},
						FixAvailable: &FixAvailable{
							Name:          "@vue/cli-service",
							Version:       "4.1.1",
							IsSemVerMajor: true,
						},
					},
				},
				Metadata: Metadata{
					Vulnerabilities: Vulnerabilities{
						Info:     0,
						Low:      6,
						Moderate: 4,
						High:     0,
						Critical: 0,
						Total:    10,
					},
					Dependencies: Dependencies{
						Prod:         388,
						Dev:          2044,
						Optional:     243,
						Peer:         0,
						PeerOptional: 0,
						Total:        2612,
					},
				},
			},
		},
		{
			name: "via exists",
			args: args{
				src: `
{
	"auditReportVersion": 2,
	"vulnerabilities": {
		"extract-zip": {
			"name": "extract-zip",
			"severity": "low",
			"via": [
				"mkdirp"
			],
			"effects": [
				"cypress"
			],
			"range": "<=1.6.7",
			"nodes": [
				"node_modules/@vue/cli-plugin-e2e-cypress/node_modules/extract-zip"
			],
			"fixAvailable": false
    	},
		"lodash": {
			"name": "lodash",
			"severity": "low",
			"via": [
				{
					"source": 1523,
					"name": "lodash",
					"dependency": "lodash",
					"title": "Prototype Pollution",
					"url": "https://npmjs.com/advisories/1523",
					"severity": "low",
					"range": "<4.17.19"
				}
			],
			"effects": [
				"cypress"
			],
			"range": "<4.17.19",
			"nodes": [
				"node_modules/@vue/cli-plugin-e2e-cypress/node_modules/lodash"
			],
			"fixAvailable": false
		}
	},
	"metadata": {
		"vulnerabilities": {
			"info": 0,
			"low": 6,
			"moderate": 4,
			"high": 0,
			"critical": 0,
			"total": 10
		},
		"dependencies": {
			"prod": 388,
			"dev": 2044,
			"optional": 243,
			"peer": 0,
			"peerOptional": 0,
			"total": 2612
		}
	}
}`,
			},
			want: &AuditReport{
				AuditReportVersion: 2,
				Vulnerabilities: map[string]Vulnerability{
					"extract-zip": {
						Name:         "extract-zip",
						Severity:     "low",
						Range:        "<=1.6.7",
						CausedBy:     []string{"mkdirp"},
						Nodes:        []string{"node_modules/@vue/cli-plugin-e2e-cypress/node_modules/extract-zip"},
						FixAvailable: nil,
					},
					"lodash": {
						Name:     "lodash",
						Severity: "low",
						Range:    "<4.17.19",
						Nodes:    []string{"node_modules/@vue/cli-plugin-e2e-cypress/node_modules/lodash"},
						Cause: []Via{
							{
								Source:     1523,
								Name:       "lodash",
								Dependency: "lodash",
								Title:      "Prototype Pollution",
								URL:        "https://npmjs.com/advisories/1523",
								Severity:   "low",
								Range:      "<4.17.19",
							},
						},
						FixAvailable: nil,
					},
				},
				Metadata: Metadata{
					Vulnerabilities: Vulnerabilities{
						Info:     0,
						Low:      6,
						Moderate: 4,
						High:     0,
						Critical: 0,
						Total:    10,
					},
					Dependencies: Dependencies{
						Prod:         388,
						Dev:          2044,
						Optional:     243,
						Peer:         0,
						PeerOptional: 0,
						Total:        2612,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAuditReport(strings.NewReader(tt.args.src))
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAuditReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
