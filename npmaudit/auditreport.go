package npmaudit

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os/exec"
)

type AuditReport struct {
	AuditReportVersion int                      `json:"auditReportVersion"`
	Vulnerabilities    map[string]Vulnerability `json:"vulnerabilities"`
	Metadata           Metadata                 `json:"metadata"`
}

type Vulnerability struct {
	Name         string
	Severity     string
	Range        string
	Nodes        []string
	Cause        []Via
	CausedBy     []string
	FixAvailable *FixAvailable
}

type vulnerability struct {
	Name         string          `json:"name"`
	Severity     string          `json:"severity"`
	Range        string          `json:"range"`
	Nodes        []string        `json:"nodes"`
	Via          []json.RawMessage `json:"via"`
	FixAvailable json.RawMessage `json:"fixAvailable"`
}

var falseBytes = []byte("false")

func (v *Vulnerability) UnmarshalJSON(data []byte) error {
	var raw vulnerability
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	v.Name = raw.Name
	v.Severity = raw.Severity
	v.Range = raw.Range
	v.Nodes = raw.Nodes
	if !bytes.Equal(falseBytes, raw.FixAvailable) {
		var fix FixAvailable
		json.Unmarshal(raw.FixAvailable, &fix)
		v.FixAvailable = &fix
	}
	var causedBy []string
	var cause []Via
	for _, v := range raw.Via {
		var cby string
		var c Via
		if json.Unmarshal(v, &cby) == nil {
			causedBy = append(causedBy, cby)
		} else if json.Unmarshal(v, &c) == nil {
			cause = append(cause, c)
		}
	}
	v.Cause = cause
	v.CausedBy = causedBy
	return nil
}

type FixAvailable struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	IsSemVerMajor bool   `json:"isSemVerMajor"`
}

type Vulnerabilities struct {
	Info     int `json:"info"`
	Low      int `json:"low"`
	Moderate int `json:"moderate"`
	High     int `json:"high"`
	Critical int `json:"critical"`
	Total    int `json:"total"`
}
type Dependencies struct {
	Prod         int `json:"prod"`
	Dev          int `json:"dev"`
	Optional     int `json:"optional"`
	Peer         int `json:"peer"`
	PeerOptional int `json:"peerOptional"`
	Total        int `json:"total"`
}

type Metadata struct {
	Vulnerabilities Vulnerabilities `json:"vulnerabilities"`
	Dependencies    Dependencies    `json:"dependencies"`
}

func parseAuditReport(r io.Reader) (*AuditReport, error) {
	var result AuditReport
	dec := json.NewDecoder(r)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ExecNpmAudit(ctx context.Context, root string) (*AuditReport, error) {
	cmd := exec.CommandContext(ctx, "npm", "audit", "--json")
	cmd.Dir = root
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer r.Close()
	var result *AuditReport
	var parseErr error
	go func() {
		result, parseErr = parseAuditReport(r)
	}()
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	if parseErr != nil {
		return nil, parseErr
	}
	return result, nil
}

type Via struct {
	Source     int    `json:"source"`
	Name       string `json:"name"`
	Dependency string `json:"dependency"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Severity   string `json:"severity"`
	Range      string `json:"range"`
}
