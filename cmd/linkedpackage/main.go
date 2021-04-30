package main

import (
	"context"
	"fmt"
	"github.com/future-architect/linkedpackage"
	"github.com/future-architect/linkedpackage/npmaudit"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	app = kingpin.New("license", "dump linked package's information from compiled application")

	jsFolders       = app.Flag("js-dist", "JavaScript application dist folder").ExistingDirs()
	jsRoot          = app.Flag("js-root", "JavaScript project root folder").ExistingDir()
	jsExtraPackages = app.Flag("js-extra-package", "JavaScript extra package").Strings()

	licenseCmd   = app.Command("license", "dump license")
	licenseTitle = licenseCmd.Flag("title", "report title").Default("Used OSS Licenses").String()

	auditCmd = app.Command("audit", "audit check")
	auditOutputFormat = auditCmd.Flag("audit-format", "export format").Default("plain").Enum("plain", "json")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("🐙")
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case licenseCmd.FullCommand():
		dumpLicense(*jsRoot, *jsFolders, *jsExtraPackages, *licenseTitle, os.Stdout)
	case auditCmd.FullCommand():
		checkAudit(*jsRoot, *jsFolders, *jsExtraPackages, *auditOutputFormat, os.Stdout)
	}
}

func checkAudit(jsRoot string, jsFolders, jsExtraPackages []string, format string, writer io.Writer) {
	parsedModules := readJSPackages(jsFolders, jsExtraPackages, jsRoot)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	auditReports, err := npmaudit.ExecNpmAudit(ctx, jsRoot)
	if err != nil {
		log.Fatal(err)
	}
	audits := map[string]npmaudit.Vulnerability{}
	for _, r := range auditReports.Vulnerabilities {
		/*if len(r.Cause) == 0 {
			continue
		}*/
		audits[r.Name] = r
	}
	for _, m := range parsedModules {
		v, ok := audits[m.Name]
		if ok {
			fmt.Printf("------\n")
			fmt.Printf("[%s] %s: %s\n", v.Severity, m.Name, m.Version)
			for i, c := range v.Cause {
				if i != 0 {
					fmt.Printf("    ------\n")
				}
				fmt.Printf("    [%s] %s @ %s\n",  c.Severity, c.Name, c.Range)
				fmt.Printf("    %s\n", c.Title)
				fmt.Printf("    %s\n", c.URL)
			}
		}
	}
}

func dumpLicense(jsRoot string, jsFolders, jsExtraPackages []string, title string, writer io.Writer) {
	parsedModules := readJSPackages(jsFolders, jsExtraPackages, jsRoot)

	groups := linkedpackage.GroupingModulesByLicense(parsedModules)

	fmt.Fprintf(writer, "# %s\n\n", title)

	for _, group := range groups {
		var projects []string
		for _, module := range group.Modules {
			projects = append(projects, fmt.Sprintf("%s@%s", module.Name, module.Version))
		}
		fmt.Fprintf(writer, "## %s\n\n", strings.Join(projects, ", "))
		fmt.Fprintf(writer, "* 作者: %s\n", group.Author)
		fmt.Fprintf(writer, "* ライセンス: %s\n", group.License)
		if group.Modules[0].LicenseContent != "" {
			fmt.Fprintf(writer, "\n```\n%s\n```\n\n\n", group.Modules[0].LicenseContent)
		} else {
			fmt.Fprintf(writer, "\n\n")
		}
	}
}

func readJSPackages(folders []string, extraPackages []string, root string) []linkedpackage.Module {
	var modules []linkedpackage.Module
	for _, folder := range folders {
		sourceMapPaths := linkedpackage.Search(folder, ".js.map")
		for _, sourceMapPath := range sourceMapPaths {
			smModules, err := linkedpackage.ParseJSSourcemapFile(sourceMapPath)
			if err != nil {
				log.Println(err)
				continue
			}
			modules = append(modules, smModules...)
		}

		sourcePaths := linkedpackage.Search(folder, ".js")
		for _, sourcePath := range sourcePaths {
			smModules, err := linkedpackage.ParseJSWebPack(sourcePath)
			if err != nil {
				log.Println(err)
				continue
			}
			modules = append(modules, smModules...)
		}
	}
	for _, extra := range extraPackages {
		modules = append(modules, linkedpackage.Module{
			Lang: "js",
			Name: extra,
			Path: "/node_modules/" + extra,
		})
	}
	modules = linkedpackage.UniqueModules(modules)
	parsedModules := []linkedpackage.Module{}
	for _, module := range modules {
		err := linkedpackage.ReadProjectData(&module, root)
		if err != nil {
			log.Println(err)
			continue
		}
		parsedModules = append(parsedModules, module)
	}
	return parsedModules
}
