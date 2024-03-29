// Code generated by protoc-gen-goext. DO NOT EDIT.

package endpoint

type Parser_Parser = isParser_Parser

func (m *Parser) SetParser(v Parser_Parser) {
	m.Parser = v
}

func (m *Parser) SetJsonParser(v *GenericParserCommon) {
	m.Parser = &Parser_JsonParser{
		JsonParser: v,
	}
}

func (m *Parser) SetAuditTrailsV1Parser(v *AuditTrailsV1Parser) {
	m.Parser = &Parser_AuditTrailsV1Parser{
		AuditTrailsV1Parser: v,
	}
}

func (m *Parser) SetCloudLoggingParser(v *CloudLoggingParser) {
	m.Parser = &Parser_CloudLoggingParser{
		CloudLoggingParser: v,
	}
}

func (m *Parser) SetTskvParser(v *GenericParserCommon) {
	m.Parser = &Parser_TskvParser{
		TskvParser: v,
	}
}

func (m *GenericParserCommon) SetDataSchema(v *DataSchema) {
	m.DataSchema = v
}

func (m *GenericParserCommon) SetNullKeysAllowed(v bool) {
	m.NullKeysAllowed = v
}

func (m *GenericParserCommon) SetAddRestColumn(v bool) {
	m.AddRestColumn = v
}

func (m *GenericParserCommon) SetUnescapeStringValues(v bool) {
	m.UnescapeStringValues = v
}
