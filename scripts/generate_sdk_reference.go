//go:build ignore

// generate_sdk_reference.go regenerates docs/sdk-reference from the bundled
// contract artifacts.
//
// The pages describe the Go surface this repository actually ships:
// client.<Namespace>.<Method> for Golden, client.Silver.<Namespace>.<Method>
// for Silver, and the shared (ctx, RequestOptions, out) call signature. Only
// the route and prose metadata come from the synced artifacts, so a contract
// refresh updates the docs without reintroducing the source SDK's language.
//
// exportedName is duplicated from scripts/generate_wrappers.go on purpose: both
// files are build-ignored generator programs, and the duplication keeps the
// published module free of a naming package that exists only for tooling. The
// two copies must stay identical, and docs_reference_test.go fails if the
// documented method names drift from the generated wrappers.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

const referenceDir = "docs/sdk-reference"

type goldenOperation struct {
	Method      string `json:"method"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	OperationID string `json:"operation_id"`
	Path        string `json:"path"`
}

type silverOperation struct {
	HTTPMethod string   `json:"http_method"`
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	Path       string   `json:"path"`
	Sources    []string `json:"sources"`
}

type silverDetail struct {
	Description    string            `json:"description"`
	HTTPMethod     string            `json:"http_method"`
	MethodName     string            `json:"method_name"`
	NamespacePath  []string          `json:"namespace_path"`
	Parameters     []silverParameter `json:"parameters"`
	Route          string            `json:"route"`
	Sources        []string          `json:"sources"`
	StatusCodes    []int             `json:"status_codes"`
	Summary        string            `json:"summary"`
	UsesAppHeaders bool              `json:"uses_app_headers"`
}

type silverParameter struct {
	APIName     string `json:"api_name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Required    bool   `json:"required"`
	TypeDisplay string `json:"type_display"`
}

type legacyAlias struct {
	LegacyName      string `json:"legacy_name"`
	LegacyNamespace string `json:"legacy_namespace"`
	Surface         string `json:"surface"`
	TargetName      string `json:"target_name"`
	TargetNamespace string `json:"target_namespace"`
}

type typedSilverMethod struct {
	Namespace  string `json:"namespace"`
	Name       string `json:"name"`
	GoName     string `json:"go_name"`
	HTTPMethod string `json:"http_method"`
	Route      string `json:"route"`
	Signature  string `json:"signature"`
	Summary    string `json:"summary"`
}

// docAlias is a deprecated forwarder as it appears on a namespace page.
type docAlias struct {
	Name   string
	Target string
	Route  string
}

// namespacePage is one generated Markdown page.
type namespacePage struct {
	// Namespace is the contract namespace, such as "custom_fields" or
	// "apps.microsoft_intune".
	Namespace string
	// ClientPath is the Go expression that reaches the service, such as
	// "client.Silver.Apps.MicrosoftIntune".
	ClientPath string
	// FileName is the page file name inside docs/sdk-reference.
	FileName string
	Silver   bool
	Methods  []docMethod
	Aliases  []docAlias
	// Typed lists hand-written helpers that replace a generic wrapper on
	// this namespace.
	Typed []typedSilverMethod
}

type docMethod struct {
	// Name is the exported Go method name.
	Name string
	// SourceName is the contract operation name the wrapper was generated
	// from, kept so a reader can find the same route in the bundled
	// inventory files.
	SourceName     string
	HTTPMethod     string
	Path           string
	OperationID    string
	Summary        string
	Description    string
	Parameters     []docParameter
	StatusCodes    []string
	ResponseModel  string
	Sources        []string
	UsesAppHeaders bool
	// Typed is true when a hand-written typed helper replaces the generic
	// wrapper for this route.
	Typed bool
}

type docParameter struct {
	APIName     string
	In          string
	Required    bool
	GoType      string
	Field       string
	Description string
}

var nonIdentifier = regexp.MustCompile(`[^0-9A-Za-z]+`)

// pythonClientPath matches the source SDK's snake_case accessor as it appears
// in prose carried by the bundled artifacts, such as
// "client.silver.alerts.queue_notification".
var pythonClientPath = regexp.MustCompile(`client\.[a-z0-9_]+(?:\.[a-z0-9_]+)*`)

// goifyProse rewrites source-SDK accessors in contract prose into the Go
// accessors this SDK actually exposes, so a reader never has to translate a
// Python path into a Go one.
func goifyProse(value string) string {
	return pythonClientPath.ReplaceAllStringFunc(value, func(match string) string {
		segments := strings.Split(match, ".")[1:]
		for i, segment := range segments {
			segments[i] = exportedName(segment)
		}
		return "client." + strings.Join(segments, ".")
	})
}

func main() {
	golden := readGolden("testdata/contract/golden_sdk_inventory.json")
	silver := readSilver("testdata/contract/silver_sdk_inventory.json")
	silverDetails := readSilverDetails("data/silver_inventory.json")
	spec := readSpec("data/openapi/openapi-spec.json")
	aliases := readAliases("data/legacy/aliases.json")
	typed := readTypedSilverMethods("data/typed_silver_methods.json")

	goldenPages := buildGoldenPages(golden, spec)
	silverPages := buildSilverPages(silver, silverDetails, typed)
	goldenPages = attachAliases(goldenPages, silverPages, aliases)

	if err := os.RemoveAll(referenceDir); err != nil {
		fatal("remove %s: %v", referenceDir, err)
	}
	if err := os.MkdirAll(referenceDir, 0o755); err != nil {
		fatal("create %s: %v", referenceDir, err)
	}

	for _, page := range append(append([]namespacePage(nil), goldenPages...), silverPages...) {
		writePage(page)
	}
	writeGoldenIndex(goldenPages, silverPages)
	writeSilverIndex(silverPages)

	fmt.Printf("wrote %d reference pages to %s\n", len(goldenPages)+len(silverPages)+2, referenceDir)
}

// -- inputs -----------------------------------------------------------------

func readJSON(path string, out any) {
	payload, err := os.ReadFile(path)
	if err != nil {
		fatal("read %s: %v", path, err)
	}
	if err := json.Unmarshal(payload, out); err != nil {
		fatal("parse %s: %v", path, err)
	}
}

func readGolden(path string) []goldenOperation {
	var ops []goldenOperation
	readJSON(path, &ops)
	return ops
}

func readSilver(path string) []silverOperation {
	var ops []silverOperation
	readJSON(path, &ops)
	return ops
}

func readSilverDetails(path string) map[string]silverDetail {
	var bundle struct {
		Endpoints []silverDetail `json:"endpoints"`
	}
	readJSON(path, &bundle)
	details := make(map[string]silverDetail, len(bundle.Endpoints))
	for _, endpoint := range bundle.Endpoints {
		key := strings.Join(endpoint.NamespacePath, ".") + "." + endpoint.MethodName
		details[key] = endpoint
	}
	return details
}

func readSpec(path string) map[string]any {
	var spec map[string]any
	readJSON(path, &spec)
	return spec
}

func readAliases(path string) []legacyAlias {
	var bundle struct {
		Aliases []legacyAlias `json:"aliases"`
	}
	readJSON(path, &bundle)
	return bundle.Aliases
}

func readTypedSilverMethods(path string) []typedSilverMethod {
	var raw []typedSilverMethod
	readJSON(path, &raw)
	sort.Slice(raw, func(i, j int) bool {
		if raw[i].Namespace != raw[j].Namespace {
			return raw[i].Namespace < raw[j].Namespace
		}
		return raw[i].GoName < raw[j].GoName
	})
	return raw
}

// -- page construction ------------------------------------------------------

func buildGoldenPages(ops []goldenOperation, spec map[string]any) []namespacePage {
	operations := indexSpecOperations(spec)
	byNamespace := map[string][]docMethod{}
	for _, op := range ops {
		method := docMethod{
			Name:        exportedName(op.Name),
			SourceName:  op.Name,
			HTTPMethod:  strings.ToUpper(op.Method),
			Path:        op.Path,
			OperationID: op.OperationID,
		}
		if entry, ok := operations[strings.ToUpper(op.Method)+" "+op.Path]; ok {
			applySpecDetail(&method, entry, spec)
		}
		byNamespace[op.Namespace] = append(byNamespace[op.Namespace], method)
	}
	return assemblePages(byNamespace, false)
}

func buildSilverPages(ops []silverOperation, details map[string]silverDetail, typed []typedSilverMethod) []namespacePage {
	replaced := map[string]bool{}
	for _, method := range typed {
		replaced[method.Namespace+"."+method.Name] = true
	}
	byNamespace := map[string][]docMethod{}
	for _, op := range ops {
		method := docMethod{
			Name:       exportedName(op.Name),
			SourceName: op.Name,
			HTTPMethod: strings.ToUpper(op.HTTPMethod),
			Path:       op.Path,
			Sources:    op.Sources,
			Typed:      replaced[op.Namespace+"."+op.Name],
		}
		if detail, ok := details[op.Namespace+"."+op.Name]; ok {
			applySilverDetail(&method, detail)
		}
		byNamespace[op.Namespace] = append(byNamespace[op.Namespace], method)
	}
	// A typed helper may be the only Go surface for its namespace, so make
	// sure the namespace gets a page either way.
	for _, method := range typed {
		if _, ok := byNamespace[method.Namespace]; !ok {
			byNamespace[method.Namespace] = nil
		}
	}

	pages := assemblePages(byNamespace, true)
	for i, page := range pages {
		for _, method := range typed {
			if method.Namespace == page.Namespace {
				pages[i].Typed = append(pages[i].Typed, method)
			}
		}
	}
	return pages
}

func assemblePages(byNamespace map[string][]docMethod, silver bool) []namespacePage {
	namespaces := make([]string, 0, len(byNamespace))
	for namespace := range byNamespace {
		namespaces = append(namespaces, namespace)
	}
	sort.Strings(namespaces)

	pages := make([]namespacePage, 0, len(namespaces))
	for _, namespace := range namespaces {
		methods := byNamespace[namespace]
		sort.Slice(methods, func(i, j int) bool { return methods[i].Name < methods[j].Name })
		pages = append(pages, namespacePage{
			Namespace:  namespace,
			ClientPath: clientPath(namespace, silver),
			FileName:   pageFileName(namespace, silver),
			Silver:     silver,
			Methods:    methods,
		})
	}
	return pages
}

// clientPath renders the Go accessor for a namespace. Silver app routes are
// nested one level deeper, matching the generated SilverAppsClient.
func clientPath(namespace string, silver bool) string {
	parts := strings.Split(namespace, ".")
	for i, part := range parts {
		parts[i] = exportedName(part)
	}
	joined := strings.Join(parts, ".")
	if silver {
		return "client.Silver." + joined
	}
	return "client." + joined
}

func pageFileName(namespace string, silver bool) string {
	parts := strings.Split(namespace, ".")
	for i, part := range parts {
		parts[i] = strings.ToLower(exportedNameToSlug(exportedName(part)))
	}
	slug := strings.Join(parts, "-")
	if silver {
		return "silver-" + slug + ".md"
	}
	return slug + ".md"
}

// exportedNameToSlug converts GoCamelCase into kebab-case for file names.
func exportedNameToSlug(name string) string {
	var builder strings.Builder
	for i, r := range name {
		if unicode.IsUpper(r) && i > 0 {
			builder.WriteRune('-')
		}
		builder.WriteRune(unicode.ToLower(r))
	}
	return builder.String()
}

// attachAliases mirrors scripts/generate_wrappers.go: every legacy name becomes
// a deprecated forwarder on its original Golden namespace, dropped namespaces
// get an alias-only page, and a generated method always wins a name collision.
func attachAliases(golden []namespacePage, silver []namespacePage, aliases []legacyAlias) []namespacePage {
	silverPaths := map[string]string{}
	for _, page := range silver {
		silverPaths[page.Namespace] = page.ClientPath
	}
	goldenPaths := map[string]string{}
	index := map[string]int{}
	for i, page := range golden {
		index[page.Namespace] = i
		goldenPaths[page.Namespace] = page.ClientPath
	}
	routes := map[string]string{}
	for _, page := range append(append([]namespacePage(nil), golden...), silver...) {
		for _, method := range page.Methods {
			routes[page.Namespace+"."+method.SourceName] = method.HTTPMethod + " " + method.Path
		}
	}

	sorted := append([]legacyAlias(nil), aliases...)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].LegacyNamespace != sorted[j].LegacyNamespace {
			return sorted[i].LegacyNamespace < sorted[j].LegacyNamespace
		}
		return sorted[i].LegacyName < sorted[j].LegacyName
	})

	for _, alias := range sorted {
		position, ok := index[alias.LegacyNamespace]
		if !ok {
			golden = append(golden, namespacePage{
				Namespace:  alias.LegacyNamespace,
				ClientPath: clientPath(alias.LegacyNamespace, false),
				FileName:   pageFileName(alias.LegacyNamespace, false),
			})
			position = len(golden) - 1
			index[alias.LegacyNamespace] = position
			goldenPaths[alias.LegacyNamespace] = golden[position].ClientPath
		}

		page := golden[position]
		name := exportedName(alias.LegacyName)
		if aliasNameTaken(page, name) {
			continue
		}

		var target string
		switch alias.Surface {
		case "silver":
			target, ok = silverPaths[alias.TargetNamespace]
		case "golden":
			target, ok = goldenPaths[alias.TargetNamespace]
		default:
			fatal("alias %s.%s has unknown surface %q", alias.LegacyNamespace, alias.LegacyName, alias.Surface)
		}
		if !ok {
			fatal("alias %s.%s targets unknown %s namespace %q", alias.LegacyNamespace, alias.LegacyName, alias.Surface, alias.TargetNamespace)
		}

		page.Aliases = append(page.Aliases, docAlias{
			Name:   name,
			Target: target + "." + exportedName(alias.TargetName),
			Route:  routes[alias.TargetNamespace+"."+alias.TargetName],
		})
		golden[position] = page
	}

	sort.Slice(golden, func(i, j int) bool { return golden[i].Namespace < golden[j].Namespace })
	return golden
}

func aliasNameTaken(page namespacePage, name string) bool {
	for _, method := range page.Methods {
		if method.Name == name {
			return true
		}
	}
	for _, alias := range page.Aliases {
		if alias.Name == name {
			return true
		}
	}
	return false
}

// -- contract metadata ------------------------------------------------------

// specOperation is one OpenAPI operation together with the parameters its path
// item declares for every method on that path.
type specOperation struct {
	operation map[string]any
	// shared holds the path item's own "parameters" array. OpenAPI lets a
	// path declare parameters once for all its methods, and 13 paths in the
	// bundled contract do, so an operation's own list is not the whole set.
	shared []any
}

// indexSpecOperations flattens the OpenAPI paths object into a
// "METHOD /path" lookup.
func indexSpecOperations(spec map[string]any) map[string]specOperation {
	operations := map[string]specOperation{}
	paths, _ := spec["paths"].(map[string]any)
	for path, raw := range paths {
		item, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		shared, _ := item["parameters"].([]any)
		for method, rawOperation := range item {
			if method == "parameters" {
				continue
			}
			operation, ok := rawOperation.(map[string]any)
			if !ok {
				continue
			}
			operations[strings.ToUpper(method)+" "+path] = specOperation{
				operation: operation,
				shared:    shared,
			}
		}
	}
	return operations
}

func applySpecDetail(method *docMethod, entry specOperation, spec map[string]any) {
	operation := entry.operation
	method.Summary = goifyProse(stringField(operation, "summary"))
	method.Description = goifyProse(stringField(operation, "description"))

	rawParameters, _ := operation["parameters"].([]any)
	seen := map[string]bool{}
	for _, raw := range append(append([]any(nil), rawParameters...), entry.shared...) {
		parameter, ok := resolveRef(raw, spec)
		if !ok {
			continue
		}
		in := stringField(parameter, "in")
		name := stringField(parameter, "name")
		// The operation's own entry overrides the path item's for the
		// same name and location, so it is read first and wins here.
		if seen[in+" "+name] {
			continue
		}
		seen[in+" "+name] = true
		method.Parameters = append(method.Parameters, docParameter{
			APIName:     name,
			In:          in,
			Required:    boolField(parameter, "required"),
			GoType:      schemaGoType(parameter["schema"], spec),
			Field:       optionsField(in, name),
			Description: goifyProse(stringField(parameter, "description")),
		})
	}
	if body, ok := operation["requestBody"].(map[string]any); ok {
		resolved, _ := resolveRef(body, spec)
		mediaType, goType := bodyType(resolved, spec)
		// Only a JSON body goes in RequestOptions.JSON. A multipart or
		// binary body has to be encoded by the caller and sent as bytes.
		field := "`JSON`"
		if mediaType != "application/json" {
			field = fmt.Sprintf("`Body` with `ContentType: %q`", mediaType)
		}
		method.Parameters = append(method.Parameters, docParameter{
			APIName:     "request body",
			In:          "body",
			Required:    boolField(resolved, "required"),
			GoType:      goType,
			Field:       field,
			Description: goifyProse(stringField(resolved, "description")),
		})
	}

	responses, _ := operation["responses"].(map[string]any)
	codes := make([]string, 0, len(responses))
	for code := range responses {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	method.StatusCodes = codes
	for _, code := range codes {
		if !strings.HasPrefix(code, "2") {
			continue
		}
		response, ok := resolveRef(responses[code], spec)
		if !ok {
			continue
		}
		if model := responseModel(response); model != "" {
			method.ResponseModel = model
			break
		}
	}
}

func applySilverDetail(method *docMethod, detail silverDetail) {
	method.Summary = goifyProse(detail.Summary)
	method.Description = goifyProse(detail.Description)
	method.UsesAppHeaders = detail.UsesAppHeaders
	for _, code := range detail.StatusCodes {
		method.StatusCodes = append(method.StatusCodes, fmt.Sprintf("%d", code))
	}
	for _, parameter := range detail.Parameters {
		method.Parameters = append(method.Parameters, docParameter{
			APIName:     parameter.APIName,
			In:          parameter.Location,
			Required:    parameter.Required,
			GoType:      goTypeFromDisplay(parameter.TypeDisplay),
			Field:       optionsField(parameter.Location, parameter.APIName),
			Description: goifyProse(parameter.Description),
		})
	}
}

// optionsField names the RequestOptions field a caller populates for a
// parameter in the given location.
func optionsField(in string, name string) string {
	switch strings.ToLower(in) {
	case "path":
		return fmt.Sprintf("`PathParams[%q]`", name)
	case "query":
		return fmt.Sprintf("`Params[%q]`", name)
	case "header":
		return fmt.Sprintf("`Headers[%q]`", name)
	case "body":
		return "`JSON`"
	case "file":
		return "`Body` with `ContentType`"
	default:
		return "`RequestOptions`"
	}
}

func responseModel(response map[string]any) string {
	content, ok := response["content"].(map[string]any)
	if !ok {
		return ""
	}
	media, ok := content["application/json"].(map[string]any)
	if !ok {
		return ""
	}
	schema, ok := media["schema"].(map[string]any)
	if !ok {
		return ""
	}
	if ref := stringField(schema, "$ref"); ref != "" {
		return refName(ref)
	}
	if items, ok := schema["items"].(map[string]any); ok {
		if ref := stringField(items, "$ref"); ref != "" {
			return "[]" + refName(ref)
		}
	}
	return ""
}

// bodyType reports the media type a request body is sent as, and the Go type
// the caller supplies for it.
func bodyType(body map[string]any, spec map[string]any) (string, string) {
	content, ok := body["content"].(map[string]any)
	if !ok {
		return "application/json", "any"
	}
	for _, mediaType := range []string{"application/json", "multipart/form-data", "application/octet-stream"} {
		media, ok := content[mediaType].(map[string]any)
		if !ok {
			continue
		}
		return mediaType, schemaGoType(media["schema"], spec)
	}
	return "application/json", "any"
}

// schemaGoType renders the Go type a caller would supply or decode for a
// schema. It stays deliberately shallow: the SDK marshals whatever the caller
// passes, so the table documents intent rather than a generated type.
func schemaGoType(raw any, spec map[string]any) string {
	schema, ok := raw.(map[string]any)
	if !ok {
		return "any"
	}
	if ref := stringField(schema, "$ref"); ref != "" {
		return refName(ref)
	}
	switch stringField(schema, "type") {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		return "[]" + schemaGoType(schema["items"], spec)
	case "object":
		return "any"
	default:
		return "any"
	}
}

// goTypeFromDisplay converts the Python type rendering carried in the Silver
// inventory into the Go type a caller supplies through RequestOptions.
func goTypeFromDisplay(display string) string {
	trimmed := strings.TrimSpace(display)
	if trimmed == "" {
		return "any"
	}
	if strings.Contains(trimmed, "|") {
		return "any"
	}
	switch trimmed {
	case "str":
		return "string"
	case "bool":
		return "bool"
	case "int":
		return "int"
	case "float":
		return "float64"
	}
	if strings.HasPrefix(trimmed, "list[") {
		return "[]any"
	}
	if strings.HasPrefix(trimmed, "dict[") {
		return "any"
	}
	return "any"
}

func resolveRef(raw any, spec map[string]any) (map[string]any, bool) {
	node, ok := raw.(map[string]any)
	if !ok {
		return nil, false
	}
	ref := stringField(node, "$ref")
	if ref == "" {
		return node, true
	}
	cursor := any(spec)
	for _, segment := range strings.Split(strings.TrimPrefix(ref, "#/"), "/") {
		container, ok := cursor.(map[string]any)
		if !ok {
			return nil, false
		}
		cursor, ok = container[segment]
		if !ok {
			return nil, false
		}
	}
	resolved, ok := cursor.(map[string]any)
	return resolved, ok
}

func refName(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

func stringField(node map[string]any, key string) string {
	value, _ := node[key].(string)
	return value
}

func boolField(node map[string]any, key string) bool {
	value, _ := node[key].(bool)
	return value
}

// -- rendering --------------------------------------------------------------

func writePage(page namespacePage) {
	var out strings.Builder
	writeGeneratedHeader(&out)

	surface := "Golden"
	if page.Silver {
		surface = "Silver"
	}
	fmt.Fprintf(&out, "# `%s` %s Namespace\n\n", page.ClientPath, surface)
	fmt.Fprintf(&out, "Client access: `%s`\n\n", page.ClientPath)

	if page.Silver {
		out.WriteString("These methods are Silver: they come from routes observed in live site interaction HARs rather than from the published contract. They stay under `client.Silver` so quasi-supported behavior never stands in for the Golden SDK surface.\n\n")
		out.WriteString("Method prose on this page is carried over from the bundled inventory artifacts and describes the observed HTTP route. It is not a description of Go-side behavior: this SDK sends what you put in `RequestOptions` and decodes the response, and performs no payload transformation, no response validation, and no client-side file conversion.\n\n")
	} else {
		out.WriteString("These methods are Golden: they come from the bundled Incident IQ OpenAPI contract and are the correct default SDK path for supported calls.\n\n")
	}
	out.WriteString("Every generated method has the same signature:\n\n")
	out.WriteString("```go\nfunc (s *" + serviceTypeName(page) + ") " + exampleMethodName(page) + "(ctx context.Context, opts incidentiq.RequestOptions, out any) error\n```\n\n")
	out.WriteString("`opts` carries the path, query, header, and body values listed per method. `out` receives the decoded JSON response; pass a pointer to your own struct or to `map[string]any`, or pass `nil` to discard the body.\n\n")

	if len(page.Aliases) > 0 {
		out.WriteString("## Deprecated Aliases\n\n")
		out.WriteString("These pre-migration names still compile. Each one forwards to its new location and carries a `// Deprecated:` comment, which `go doc`, editors, and `staticcheck` surface. Nothing is logged at runtime.\n\n")
		out.WriteString("| Deprecated Method | Forwards To | Route |\n| --- | --- | --- |\n")
		for _, alias := range page.Aliases {
			route := alias.Route
			if route == "" {
				route = "-"
			} else {
				route = "`" + route + "`"
			}
			fmt.Fprintf(&out, "| `%s.%s` | `%s` | %s |\n", page.ClientPath, alias.Name, alias.Target, route)
		}
		out.WriteString("\n")
	}

	if len(page.Typed) > 0 {
		out.WriteString("## Typed Helpers\n\n")
		out.WriteString("These routes have hand-written helpers instead of generic wrappers, because they validate their arguments, own their request body, or narrow the shared retry policy. The generator never emits a generic wrapper over a typed helper.\n\n")
		for _, method := range page.Typed {
			fmt.Fprintf(&out, "### `%s`\n\n", method.GoName)
			if method.Summary != "" {
				fmt.Fprintf(&out, "%s\n\n", method.Summary)
			}
			fmt.Fprintf(&out, "```go\nfunc (s *%s) %s\n```\n\n", serviceTypeName(page), method.Signature)
			fmt.Fprintf(&out, "- HTTP route: `%s %s`\n\n", method.HTTPMethod, method.Route)
		}
	}

	if len(page.Methods) == 0 {
		out.WriteString("## Methods\n\nThe published contract documents no operations in this namespace. It exists only to keep the deprecated aliases above compiling.\n")
		writeFile(page.FileName, out.String())
		return
	}

	out.WriteString("## Methods\n\n")
	for _, method := range page.Methods {
		writeMethod(&out, page, method)
	}
	writeFile(page.FileName, out.String())
}

func writeMethod(out *strings.Builder, page namespacePage, method docMethod) {
	fmt.Fprintf(out, "### `%s`\n\n", method.Name)

	if method.Summary != "" {
		fmt.Fprintf(out, "%s\n\n", strings.TrimSpace(method.Summary))
	}
	fmt.Fprintf(out, "- Call: `%s.%s(ctx, opts, &out)`\n", page.ClientPath, method.Name)
	fmt.Fprintf(out, "- HTTP route: `%s %s`\n", method.HTTPMethod, method.Path)
	if method.OperationID != "" {
		fmt.Fprintf(out, "- Operation ID: `%s`\n", method.OperationID)
	}
	fmt.Fprintf(out, "- Inventory entry: `%s.%s`\n", page.Namespace, method.SourceName)
	if len(method.Sources) > 0 {
		fmt.Fprintf(out, "- Observed in: `%s`\n", strings.Join(method.Sources, "`, `"))
	}
	if method.UsesAppHeaders {
		out.WriteString("- Requires app headers: set `Config.AppHeaders` (or `INCIDENTIQ_APP_HEADERS_JSON`), or pass them in `opts.Headers`\n")
	}
	if method.Typed {
		out.WriteString("- A hand-written typed helper replaces the generic wrapper for this route; see the package documentation for its signature.\n")
	}
	out.WriteString("\n")

	if method.Description != "" {
		fmt.Fprintf(out, "%s\n\n", strings.TrimSpace(method.Description))
	}

	if len(method.Parameters) > 0 {
		out.WriteString("#### Request Options\n\n")
		out.WriteString("| API Name | In | Required | Go Type | RequestOptions Field | Description |\n| --- | --- | --- | --- | --- | --- |\n")
		for _, parameter := range method.Parameters {
			required := "no"
			if parameter.Required {
				required = "yes"
			}
			fmt.Fprintf(
				out,
				"| `%s` | `%s` | `%s` | `%s` | %s | %s |\n",
				parameter.APIName,
				parameter.In,
				required,
				parameter.GoType,
				parameter.Field,
				tableCell(parameter.Description),
			)
		}
		out.WriteString("\n")
	}

	out.WriteString("#### Response\n\n")
	out.WriteString("- Decoded with `encoding/json` into the `out` argument when `out` is non-nil.\n")
	if method.ResponseModel != "" {
		fmt.Fprintf(out, "- Contract response schema: `%s`. The SDK does not generate response structs, so model your own type on it or decode into `map[string]any`.\n", method.ResponseModel)
	} else {
		out.WriteString("- The contract declares no JSON response schema for this operation. Decode into `map[string]any` or `[]any` unless you know the shape.\n")
	}
	if len(method.StatusCodes) > 0 {
		fmt.Fprintf(out, "- Documented status codes: `%s`. Non-2xx responses return `*incidentiq.APIError`.\n", strings.Join(method.StatusCodes, "`, `"))
	} else {
		out.WriteString("- Non-2xx responses return `*incidentiq.APIError`.\n")
	}
	out.WriteString("\n---\n\n")
}

// serviceTypeName reports the generated service type that carries a page's
// methods, so the signature block names a type a reader can look up.
func serviceTypeName(page namespacePage) string {
	// exportedName flattens the dot in a nested namespace, which is exactly
	// how generate_wrappers.go names the service struct.
	name := exportedName(page.Namespace)
	if page.Silver {
		return "incidentiq.Silver" + name + "Service"
	}
	return "incidentiq.Golden" + name + "Service"
}

func exampleMethodName(page namespacePage) string {
	if len(page.Methods) > 0 {
		return page.Methods[0].Name
	}
	if len(page.Aliases) > 0 {
		return page.Aliases[0].Name
	}
	return "Method"
}

func writeGoldenIndex(golden []namespacePage, silver []namespacePage) {
	var out strings.Builder
	writeGeneratedHeader(&out)
	out.WriteString("# SDK Reference\n\n")
	out.WriteString("Generated reference for the Go surface in this repository. Golden methods come from the bundled Incident IQ OpenAPI contract and are reached directly as `client.<Namespace>.<Method>`. Silver methods come from routes observed in live site interaction HARs and stay under `client.Silver.<Namespace>.<Method>` so they never stand in for the Golden path.\n\n")
	out.WriteString("Every generated method shares one signature:\n\n")
	out.WriteString("```go\nerr := client.Assets.GetAssetById(ctx, incidentiq.RequestOptions{\n    PathParams: map[string]any{\"assetId\": assetID},\n}, &asset)\n```\n\n")
	out.WriteString("## Golden Namespaces\n\n")
	out.WriteString("| Namespace | Methods | Deprecated Aliases | Page |\n| --- | ---: | ---: | --- |\n")
	for _, page := range golden {
		fmt.Fprintf(
			&out,
			"| `%s` | %d | %d | [`%s`](%s) |\n",
			page.Namespace,
			len(page.Methods),
			len(page.Aliases),
			page.ClientPath,
			page.FileName,
		)
	}
	out.WriteString("\n## Silver\n\n")
	fmt.Fprintf(&out, "Silver spans %d namespaces. See the [Silver overview](silver.md).\n", len(silver))
	writeFile("index.md", out.String())
}

func writeSilverIndex(silver []namespacePage) {
	var out strings.Builder
	writeGeneratedHeader(&out)
	out.WriteString("# `client.Silver` Namespace\n\n")
	out.WriteString("Silver routes are quasi-supported Incident IQ APIs derived from live site interaction HARs. The SDK exposes them explicitly and separately, because the Golden contract is always preferred when it documents the same route. App-specific routes are nested one level deeper, as `client.Silver.Apps.<AppNamespace>.<Method>`.\n\n")
	out.WriteString("| Namespace | Methods | Page |\n| --- | ---: | --- |\n")
	for _, page := range silver {
		fmt.Fprintf(&out, "| `%s` | %d | [`%s`](%s) |\n", page.Namespace, len(page.Methods), page.ClientPath, page.FileName)
	}
	writeFile("silver.md", out.String())
}

func writeGeneratedHeader(out *strings.Builder) {
	out.WriteString("<!-- Code generated by scripts/generate_sdk_reference.go; DO NOT EDIT. -->\n\n")
}

// tableCell flattens prose into a single Markdown table cell.
func tableCell(value string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return "-"
	}
	replacer := strings.NewReplacer("|", "\\|", "\r\n", " ", "\n", " ", "\r", " ")
	return strings.Join(strings.Fields(replacer.Replace(trimmed)), " ")
}

func writeFile(name string, content string) {
	path := filepath.Join(referenceDir, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		fatal("write %s: %v", path, err)
	}
}

// exportedName mirrors scripts/generate_wrappers.go so documented names match
// the generated wrappers exactly.
func exportedName(value string) string {
	parts := strings.Fields(nonIdentifier.ReplaceAllString(value, " "))
	if len(parts) == 0 {
		return "Value"
	}
	for i, part := range parts {
		parts[i] = exportPart(part)
	}
	name := strings.Join(parts, "")
	if name == "" {
		return "Value"
	}
	if unicode.IsDigit(rune(name[0])) {
		name = "N" + name
	}
	return name
}

func exportPart(value string) string {
	if value == "" {
		return ""
	}
	runes := []rune(value)
	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
			continue
		}
		runes[i] = unicode.ToLower(r)
	}
	return string(runes)
}

func fatal(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
