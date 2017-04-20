package i18n

func init() {

	registerLanguage("en_us", map[StringID]string{
		ConvertValue_EnumTypeNil:                              "[TT101] ConvertValue: Enum type nil",
		ConvertValue_StructTypeNil:                            "[TT102] ConvertValue: Struct type nil",
		ConvertValue_EnumValueNotFound:                        "[TT103] ConvertValue: Enum value not found",
		ConvertValue_UnknownFieldType:                         "[TT104] ConvertValue: Unknown field type",
		StructParser_LexerError:                               "[TT201] StructParser: Lexer error",
		StructParser_ExpectField:                              "[TT202] StructParser: Expect field",
		StructParser_UnexpectedSpliter:                        "[TT203] StructParser: Unexpected k-v spliter",
		StructParser_FieldNotFound:                            "[TT204] StructParser: Field not found",
		StructParser_DuplicateFieldInCell:                     "[TT205] StructParser: Duplicate field",
		Run_CacheFile:                                         "Run: Caching file",
		Run_CollectTypeInfo:                                   "Run: Collect Type Info",
		Run_ExportSheetData:                                   "Run: Export Sheet Data",
		Globals_OutputCombineData:                             "Globals: Merge Combined Data",
		Globals_CombineNameLost:                               "[TT301] Globals: Please specify 'combinename' params",
		Globals_PackageNameDiff:                               "[TT302] Globals: Keep all type in same package",
		Globals_TableNameDuplicated:                           "[TT303] Globals: Duplicate table name",
		Globals_DuplicateTypeName:                             "[TT304] Globals: Duplicate type name( table name ?)",
		File_TypeSheetKeepSingleton:                           "[TT401] File: Type sheet only need ONE in a file",
		DataSheet_ValueConvertError:                           "[TT501] DataSheet: Cell value convert error",
		DataSheet_ValueRepeated:                               "[TT502] DataSheet: Duplicated cell value",
		DataSheet_MustFill:                                    "[TT503] DataSheet: cell value must fill",
		DataSheet_RowDataSplitedByEmptyLine:                   "[TT503] DataSheet: Row data splited by empty line",
		DataHeader_TypeNotFound:                               "[TT602] DataHeader: Type not found",
		DataHeader_MetaParseFailed:                            "[TT603] DataHeader: Meta parse failed",
		DataHeader_DuplicateFieldName:                         "[TT604] DataHeader: Duplicated field name",
		DataHeader_RepeatedFieldTypeNotSameInMultiColumn:      "[TT605] DataHeader: Repeated field type not same in columns",
		DataHeader_RepeatedFieldMetaNotSameInMultiColumn:      "[TT606] DataHeader: Repeated field meta not same in columns",
		DataHeader_UseReservedTypeName:                        "[TT607] DataHeader: Use reserved type name, like TableName+'Define' ",
		DataHeader_NotMatch:                                   "[TT608] DataHeader: Multi sheet data header not match",
		DataHeader_FieldNotDefinedInMainTableInMultiTableMode: "[TT609] DataHeader: Field not defined in main table, in multi table mode",
		TypeSheet_PragmaParseFailed:                           "[TT701] TypeSheet: File pragma parse failed",
		TypeSheet_TableNameIsEmpty:                            "[TT702] TypeSheet: Table name is empty",
		TypeSheet_PackageIsEmpty:                              "[TT703] TypeSheet: Package is empty",
		TypeSheet_FieldTypeNotFound:                           "[TT704] TypeSheet: Field type not found",
		TypeSheet_EnumValueParseFailed:                        "[TT705] TypeSheet: Enum value parse failed",
		TypeSheet_DescriptorKindNotSame:                       "[TT706] TypeSheet: Descriptor kind not the same, due to enum value",
		TypeSheet_FieldMetaParseFailed:                        "[TT707] TypeSheet: Field meta parse failed",
		TypeSheet_StructFieldCanNotBeStruct:                   "[TT708] TypeSheet: Struct field can not be struct kind",
		TypeSheet_FirstEnumValueShouldBeZero:                  "[TT709] TypeSheet: First enum value should be zero",
		TypeSheet_UnexpectedTypeHeader:                        "[TT710] TypeSheet: Unexpected type header",
		TypeSheet_DuplicatedEnumValue:                         "[TT711] TypeSheet: Duplicated enum value",
		TypeSheet_RowDataSplitedByEmptyLine:                   "[TT712] TypeSheet: Row data splited by empty line",
		TypeSheet_ObjectNameEmpty:                             "[TT713] TypeSheet: 'ObjectName' is empty",
		TypeSheet_DuplicateFieldName:                          "[TT714] TypeSheet: Duplicate field name",
		Printer_IgnoredByOutputTag:                            "[TT801] Printer: Ignored by 'OutputTag' in @Types",
		Printer_OpenWriteOutputFileFailed:                     "[TT802] Printer: Open write output file failed,",
		System_OpenReadXlsxFailed:                             "[TT901] Open read Xlsx failed, ",
	})
}
