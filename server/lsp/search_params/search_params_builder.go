package search_params

import (
	"github.com/pherrymason/c3-lsp/lsp/indexables"
	"github.com/pherrymason/c3-lsp/option"
)

type SearchParamsBuilder struct {
	params *SearchParams
}

func NewSearchParamsBuilder() *SearchParamsBuilder {
	return &SearchParamsBuilder{
		params: &SearchParams{},
	}
}

// Métodos para configurar los campos de SearchParams
func (b *SearchParamsBuilder) WithSymbol(param1 string) *SearchParamsBuilder {
	b.params.symbol = param1
	return b
}

func (b *SearchParamsBuilder) WithModule(moduleName string) *SearchParamsBuilder {
	b.params.symbolModulePath = indexables.NewModulePathFromString(moduleName)
	return b
}

func (b *SearchParamsBuilder) WithSymbolModule(modulePath indexables.ModulePath) *SearchParamsBuilder {
	b.params.symbolModulePath = modulePath
	return b
}

func (b *SearchParamsBuilder) WithTrackedModules(trackedModules TrackedModules) *SearchParamsBuilder {
	b.params.trackedModules = trackedModules
	return b
}

func (b *SearchParamsBuilder) WithDocId(docId string) *SearchParamsBuilder {
	b.params.docId = option.Some(docId)

	return b
}

func (b *SearchParamsBuilder) WithExcludedDocs(excludedDocId option.Option[string]) *SearchParamsBuilder {
	b.params.excludedDocId = excludedDocId

	return b
}

// Otros métodos withXXX según sea necesario

// Método para construir el objeto final
func (b *SearchParamsBuilder) Build() SearchParams {

	if b.params.trackedModules == nil {
		b.params.trackedModules = make(TrackedModules)
	}

	return *b.params
}