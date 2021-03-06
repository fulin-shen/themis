package pdp

/* AUTOMATICALLY GENERATED FROM errors.yaml - DO NOT EDIT */

import (
	"github.com/google/uuid"
	"math"
	"net"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Numeric identifiers of errors.
const (
	externalErrorID                                       = 0
	multiErrorID                                          = 1
	missingAttributeErrorID                               = 2
	MissingValueErrorID                                   = 3
	unknownTypeStringCastErrorID                          = 4
	invalidTypeStringCastErrorID                          = 5
	notImplementedStringCastErrorID                       = 6
	invalidBooleanStringCastErrorID                       = 7
	invalidIntegerStringCastErrorID                       = 8
	invalidFloatStringCastErrorID                         = 9
	invalidAddressStringCastErrorID                       = 10
	invalidNetworkStringCastErrorID                       = 11
	invalidAddressNetworkStringCastErrorID                = 12
	invalidDomainNameStringCastErrorID                    = 13
	attributeValueTypeErrorID                             = 14
	attributeValueFlagsTypeErrorID                        = 15
	attributeValueFlagsBitsErrorID                        = 16
	duplicateAttributeValueErrorID                        = 17
	unknownTypeSerializationErrorID                       = 18
	invalidTypeSerializationErrorID                       = 19
	notMatchingTypeRebindErrorID                          = 20
	unknownMetaTypeID                                     = 21
	assignmentTypeMismatchID                              = 22
	mapperArgumentTypeErrorID                             = 23
	flagsMapperRCAArgumentTypeErrorID                     = 24
	UntaggedPolicyModificationErrorID                     = 25
	MissingPolicyTagErrorID                               = 26
	PolicyTagsNotMatchErrorID                             = 27
	emptyPathModificationErrorID                          = 28
	invalidRootPolicyItemTypeErrorID                      = 29
	hiddenRootPolicyAppendErrorID                         = 30
	invalidRootPolicyErrorID                              = 31
	hiddenPolicySetModificationErrorID                    = 32
	invalidPolicySetItemTypeErrorID                       = 33
	tooShortPathPolicySetModificationErrorID              = 34
	missingPolicySetChildErrorID                          = 35
	hiddenPolicyAppendErrorID                             = 36
	policyTransactionTagsNotMatchErrorID                  = 37
	failedPolicyTransactionErrorID                        = 38
	unknownPolicyUpdateOperationErrorID                   = 39
	hiddenPolicyModificationErrorID                       = 40
	tooLongPathPolicyModificationErrorID                  = 41
	tooShortPathPolicyModificationErrorID                 = 42
	invalidPolicyItemTypeErrorID                          = 43
	hiddenRuleAppendErrorID                               = 44
	missingPolicyChildErrorID                             = 45
	MissingContentErrorID                                 = 46
	invalidContentStorageItemID                           = 47
	MissingContentItemErrorID                             = 48
	invalidContentItemErrorID                             = 49
	invalidContentItemTypeErrorID                         = 50
	invalidSelectorPathErrorID                            = 51
	invalidSelectorValuePathErrorID                       = 52
	networkMapKeyValueTypeErrorID                         = 53
	mapContentSubitemErrorID                              = 54
	invalidContentModificationErrorID                     = 55
	missingPathContentModificationErrorID                 = 56
	tooLongPathContentModificationErrorID                 = 57
	invalidContentValueModificationErrorID                = 58
	UntaggedContentModificationErrorID                    = 59
	MissingContentTagErrorID                              = 60
	ContentTagsNotMatchErrorID                            = 61
	unknownContentUpdateOperationErrorID                  = 62
	failedContentTransactionErrorID                       = 63
	contentTransactionIDNotMatchErrorID                   = 64
	contentTransactionTagsNotMatchErrorID                 = 65
	tooShortRawPathContentModificationErrorID             = 66
	tooLongRawPathContentModificationErrorID              = 67
	invalidContentUpdateDataErrorID                       = 68
	invalidContentUpdateResultTypeErrorID                 = 69
	invalidContentUpdateKeysErrorID                       = 70
	unknownContentItemResultTypeErrorID                   = 71
	invalidContentItemResultTypeErrorID                   = 72
	invalidContentKeyTypeErrorID                          = 73
	invalidContentStringMapErrorID                        = 74
	invalidContentNetworkMapErrorID                       = 75
	invalidContentDomainMapErrorID                        = 76
	invalidContentValueErrorID                            = 77
	invalidContentValueTypeErrorID                        = 78
	invalidContentStringFlags8MapValueErrorID             = 79
	invalidContentStringFlags16MapValueErrorID            = 80
	invalidContentStringFlags32MapValueErrorID            = 81
	invalidContentStringFlags64MapValueErrorID            = 82
	invalidContentNetworkFlags8MapValueErrorID            = 83
	invalidContentNetworkFlags16MapValueErrorID           = 84
	invalidContentNetworkFlags32MapValueErrorID           = 85
	invalidContentNetworkFlags64MapValueErrorID           = 86
	invalidContentDomainFlags8MapValueErrorID             = 87
	invalidContentDomainFlags16MapValueErrorID            = 88
	invalidContentDomainFlags32MapValueErrorID            = 89
	invalidContentDomainFlags64MapValueErrorID            = 90
	integerDivideByZeroErrorID                            = 91
	floatDivideByZeroErrorID                              = 92
	floatNanErrorID                                       = 93
	floatInfErrorID                                       = 94
	ReadOnlySymbolsChangeErrorID                          = 95
	nilTypeErrorID                                        = 96
	builtinCustomTypeErrorID                              = 97
	duplicateCustomTypeErrorID                            = 98
	duplicatesBuiltinTypeErrorID                          = 99
	duplicateFlagNameID                                   = 100
	noTypedAttributeErrorID                               = 101
	undefinedAttributeTypeErrorID                         = 102
	unknownAttributeTypeErrorID                           = 103
	duplicateAttributeErrorID                             = 104
	noFlagsDefinedErrorID                                 = 105
	tooManyFlagsDefinedErrorID                            = 106
	listOfStringsTypeErrorID                              = 107
	concatTypeErrorID                                     = 108
	unsupportedSelectorSchemeErrorID                      = 109
	disabledSelectorErrorID                               = 110
	marshalInvalidDepthErrorID                            = 111
	invalidHeaderErrorID                                  = 112
	nonMarshableErrorID                                   = 113
	nilRootErrorID                                        = 114
	PathNotFoundErrorID                                   = 115
	requestBufferOverflowErrorID                          = 116
	requestBufferUnderflowErrorID                         = 117
	requestVersionErrorID                                 = 118
	requestInvalidAttributeCountErrorID                   = 119
	requestTooManyAttributesErrorID                       = 120
	requestAttributeCountErrorID                          = 121
	requestAttributeMarshallingTypeErrorID                = 122
	requestAttributeUnmarshallingTypeErrorID              = 123
	requestAttributeUnmarshallingFlagsSizeErrorID         = 124
	requestAttributeUnmarshallingBooleanTypeErrorID       = 125
	requestAttributeUnmarshallingStringTypeErrorID        = 126
	requestAttributeUnmarshallingIntegerTypeErrorID       = 127
	requestAttributeUnmarshallingFloatTypeErrorID         = 128
	requestAttributeUnmarshallingAddressTypeErrorID       = 129
	requestAttributeUnmarshallingNetworkTypeErrorID       = 130
	requestAttributeUnmarshallingDomainTypeErrorID        = 131
	requestAttributeUnmarshallingSetOfStringsTypeErrorID  = 132
	requestAttributeUnmarshallingSetOfNetworksTypeErrorID = 133
	requestAttributeUnmarshallingSetOfDomainsTypeErrorID  = 134
	requestAttributeUnmarshallingListOfStringsTypeErrorID = 135
	requestAttributeMarshallingNotImplementedErrorID      = 136
	requestAttributeUnmarshallingNotImplementedErrorID    = 137
	requestTooLongAttributeNameErrorID                    = 138
	requestTooLongStringValueErrorID                      = 139
	requestAddressValueErrorID                            = 140
	requestInvalidNetworkValueErrorID                     = 141
	requestIPv4InvalidMaskErrorID                         = 142
	requestIPv6InvalidMaskErrorID                         = 143
	requestTooLongCollectionValueErrorID                  = 144
	requestInvalidExpressionErrorID                       = 145
	requestAssignmentsOverflowErrorID                     = 146
	requestValuesOverflowErrorID                          = 147
	requestUnmarshalEffectConstErrorID                    = 148
	requestUnmarshalEffectTypeErrorID                     = 149
	requestUnmarshalStatusConstErrorID                    = 150
	requestUnmarshalStatusTypeErrorID                     = 151
	requestUnmarshalBooleanConstErrorID                   = 152
	requestUnmarshalBooleanTypeErrorID                    = 153
	requestUnmarshalStringConstErrorID                    = 154
	requestUnmarshalStringTypeErrorID                     = 155
	requestUnmarshalIntegerConstErrorID                   = 156
	requestUnmarshalIntegerTypeErrorID                    = 157
	requestUnmarshalIntegerOverflowErrorID                = 158
	requestUnmarshalIntegerUnderflowErrorID               = 159
	requestUnmarshalFloatConstErrorID                     = 160
	requestUnmarshalFloatTypeErrorID                      = 161
	requestUnmarshalAddressConstErrorID                   = 162
	requestUnmarshalAddressTypeErrorID                    = 163
	requestUnmarshalNetworkConstErrorID                   = 164
	requestUnmarshalNetworkTypeErrorID                    = 165
	requestUnmarshalDomainConstErrorID                    = 166
	requestUnmarshalDomainTypeErrorID                     = 167
	requestUnmarshalSetOfStringsConstErrorID              = 168
	requestUnmarshalSetOfStringsTypeErrorID               = 169
	requestUnmarshalSetOfNetworksConstErrorID             = 170
	requestUnmarshalSetOfNetworksTypeErrorID              = 171
	requestUnmarshalSetOfDomainsConstErrorID              = 172
	requestUnmarshalSetOfDomainsTypeErrorID               = 173
	requestUnmarshalListOfStringsConstErrorID             = 174
	requestUnmarshalListOfStringsTypeErrorID              = 175
	responseEffectErrorID                                 = 176
	ResponseServerErrorID                                 = 177
	policyCalculationErrorID                              = 178
	obligationCalculationErrorID                          = 179
	noInformationalErrorID                                = 180
)

type externalError struct {
	errorLink
	err error
}

func newExternalError(err error) *externalError {
	return &externalError{
		errorLink: errorLink{id: externalErrorID},
		err:       err}
}

func (e *externalError) Error() string {
	return e.errorf("%s", e.err)
}

type multiError struct {
	errorLink
	errs []error
}

func newMultiError(errs []error) *multiError {
	return &multiError{
		errorLink: errorLink{id: multiErrorID},
		errs:      errs}
}

func (e *multiError) Error() string {
	msgs := make([]string, len(e.errs))
	for i, err := range e.errs {
		msgs[i] = err.Error()
	}
	msg := strings.Join(msgs, ", ")

	return e.errorf("multiple errors: %s", msg)
}

type missingAttributeError struct {
	errorLink
}

func newMissingAttributeError() *missingAttributeError {
	return &missingAttributeError{
		errorLink: errorLink{id: missingAttributeErrorID}}
}

func (e *missingAttributeError) Error() string {
	return e.errorf("Missing attribute")
}

// MissingValueError indicates that content doesn't have desired value.
type MissingValueError struct {
	errorLink
}

func newMissingValueError() *MissingValueError {
	return &MissingValueError{
		errorLink: errorLink{id: MissingValueErrorID}}
}

// Error implements error interface.
func (e *MissingValueError) Error() string {
	return e.errorf("Missing value")
}

type unknownTypeStringCastError struct {
	errorLink
	t Type
}

func newUnknownTypeStringCastError(t Type) *unknownTypeStringCastError {
	return &unknownTypeStringCastError{
		errorLink: errorLink{id: unknownTypeStringCastErrorID},
		t:         t}
}

func (e *unknownTypeStringCastError) Error() string {
	return e.errorf("Unknown type id %q", e.t)
}

type invalidTypeStringCastError struct {
	errorLink
	t Type
}

func newInvalidTypeStringCastError(t Type) *invalidTypeStringCastError {
	return &invalidTypeStringCastError{
		errorLink: errorLink{id: invalidTypeStringCastErrorID},
		t:         t}
}

func (e *invalidTypeStringCastError) Error() string {
	return e.errorf("Can't convert string to value of %q type", e.t)
}

type notImplementedStringCastError struct {
	errorLink
	t Type
}

func newNotImplementedStringCastError(t Type) *notImplementedStringCastError {
	return &notImplementedStringCastError{
		errorLink: errorLink{id: notImplementedStringCastErrorID},
		t:         t}
}

func (e *notImplementedStringCastError) Error() string {
	return e.errorf("Conversion from string to value of %q type hasn't been implemented", e.t)
}

type invalidBooleanStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidBooleanStringCastError(s string, err error) *invalidBooleanStringCastError {
	return &invalidBooleanStringCastError{
		errorLink: errorLink{id: invalidBooleanStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidBooleanStringCastError) Error() string {
	return e.errorf("Can't treat %q as boolean (%s)", e.s, e.err)
}

type invalidIntegerStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidIntegerStringCastError(s string, err error) *invalidIntegerStringCastError {
	return &invalidIntegerStringCastError{
		errorLink: errorLink{id: invalidIntegerStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidIntegerStringCastError) Error() string {
	return e.errorf("Can't treat %q as integer (%s)", e.s, e.err)
}

type invalidFloatStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidFloatStringCastError(s string, err error) *invalidFloatStringCastError {
	return &invalidFloatStringCastError{
		errorLink: errorLink{id: invalidFloatStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidFloatStringCastError) Error() string {
	return e.errorf("Can't treat %q as float (%s)", e.s, e.err)
}

type invalidAddressStringCastError struct {
	errorLink
	s string
}

func newInvalidAddressStringCastError(s string) *invalidAddressStringCastError {
	return &invalidAddressStringCastError{
		errorLink: errorLink{id: invalidAddressStringCastErrorID},
		s:         s}
}

func (e *invalidAddressStringCastError) Error() string {
	return e.errorf("Can't treat %q as IP address", e.s)
}

type invalidNetworkStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidNetworkStringCastError(s string, err error) *invalidNetworkStringCastError {
	return &invalidNetworkStringCastError{
		errorLink: errorLink{id: invalidNetworkStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidNetworkStringCastError) Error() string {
	return e.errorf("Can't treat %q as network address (%s)", e.s, e.err)
}

type invalidAddressNetworkStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidAddressNetworkStringCastError(s string, err error) *invalidAddressNetworkStringCastError {
	return &invalidAddressNetworkStringCastError{
		errorLink: errorLink{id: invalidAddressNetworkStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidAddressNetworkStringCastError) Error() string {
	return e.errorf("Can't treat %q as address or network (%s)", e.s, e.err)
}

type invalidDomainNameStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidDomainNameStringCastError(s string, err error) *invalidDomainNameStringCastError {
	return &invalidDomainNameStringCastError{
		errorLink: errorLink{id: invalidDomainNameStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidDomainNameStringCastError) Error() string {
	return e.errorf("Can't treat %q as domain name (%s)", e.s, e.err)
}

type attributeValueTypeError struct {
	errorLink
	expected Type
	actual   Type
}

func newAttributeValueTypeError(expected, actual Type) *attributeValueTypeError {
	return &attributeValueTypeError{
		errorLink: errorLink{id: attributeValueTypeErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *attributeValueTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", e.expected, e.actual)
}

type attributeValueFlagsTypeError struct {
	errorLink
	t Type
}

func newAttributeValueFlagsTypeError(t Type) *attributeValueFlagsTypeError {
	return &attributeValueFlagsTypeError{
		errorLink: errorLink{id: attributeValueFlagsTypeErrorID},
		t:         t}
}

func (e *attributeValueFlagsTypeError) Error() string {
	return e.errorf("Expected flags value but got %q", e.t)
}

type attributeValueFlagsBitsError struct {
	errorLink
	t        Type
	expected int
	actual   int
}

func newAttributeValueFlagsBitsError(t Type, expected, actual int) *attributeValueFlagsBitsError {
	return &attributeValueFlagsBitsError{
		errorLink: errorLink{id: attributeValueFlagsBitsErrorID},
		t:         t,
		expected:  expected,
		actual:    actual}
}

func (e *attributeValueFlagsBitsError) Error() string {
	return e.errorf("Expected %d bits flags value but got %q value with %d positions", e.expected, e.t, e.actual)
}

type duplicateAttributeValueError struct {
	errorLink
	ID   string
	t    Type
	curr AttributeValue
	prev AttributeValue
}

func newDuplicateAttributeValueError(ID string, t Type, curr, prev AttributeValue) *duplicateAttributeValueError {
	return &duplicateAttributeValueError{
		errorLink: errorLink{id: duplicateAttributeValueErrorID},
		ID:        ID,
		t:         t,
		curr:      curr,
		prev:      prev}
}

func (e *duplicateAttributeValueError) Error() string {
	return e.errorf("Duplicate attribute %q of type %q in request %s - %s", e.ID, e.t, e.curr.describe(), e.prev.describe())
}

type unknownTypeSerializationError struct {
	errorLink
	t Type
}

func newUnknownTypeSerializationError(t Type) *unknownTypeSerializationError {
	return &unknownTypeSerializationError{
		errorLink: errorLink{id: unknownTypeSerializationErrorID},
		t:         t}
}

func (e *unknownTypeSerializationError) Error() string {
	return e.errorf("Unknown type id %q", e.t)
}

type invalidTypeSerializationError struct {
	errorLink
	t Type
}

func newInvalidTypeSerializationError(t Type) *invalidTypeSerializationError {
	return &invalidTypeSerializationError{
		errorLink: errorLink{id: invalidTypeSerializationErrorID},
		t:         t}
}

func (e *invalidTypeSerializationError) Error() string {
	return e.errorf("Can't serialize value of %q type", e.t)
}

type notMatchingTypeRebindError struct {
	errorLink
	actual   Type
	expected Type
}

func newNotMatchingTypeRebindError(actual, expected Type) *notMatchingTypeRebindError {
	return &notMatchingTypeRebindError{
		errorLink: errorLink{id: notMatchingTypeRebindErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *notMatchingTypeRebindError) Error() string {
	return e.errorf("Can't rebind value to type %q as it doesn't match original type %q", e.expected, e.actual)
}

type unknownMetaType struct {
	errorLink
	t Type
}

func newUnknownMetaType(t Type) *unknownMetaType {
	return &unknownMetaType{
		errorLink: errorLink{id: unknownMetaTypeID},
		t:         t}
}

func (e *unknownMetaType) Error() string {
	return e.errorf("Unknown metatype %T of type %q", e.t, e.t)
}

type assignmentTypeMismatch struct {
	errorLink
	a Attribute
	t Type
}

func newAssignmentTypeMismatch(a Attribute, t Type) *assignmentTypeMismatch {
	return &assignmentTypeMismatch{
		errorLink: errorLink{id: assignmentTypeMismatchID},
		a:         a,
		t:         t}
}

func (e *assignmentTypeMismatch) Error() string {
	return e.errorf("Can't assign %q value to attribute %q of type %q", e.t, e.a.id, e.a.t)
}

type mapperArgumentTypeError struct {
	errorLink
	actual Type
}

func newMapperArgumentTypeError(actual Type) *mapperArgumentTypeError {
	return &mapperArgumentTypeError{
		errorLink: errorLink{id: mapperArgumentTypeErrorID},
		actual:    actual}
}

func (e *mapperArgumentTypeError) Error() string {
	return e.errorf("Expected %s, %s or %s as argument but got %s", TypeString, TypeSetOfStrings, TypeListOfStrings, e.actual)
}

type flagsMapperRCAArgumentTypeError struct {
	errorLink
	t Type
}

func newFlagsMapperRCAArgumentTypeError(t Type) *flagsMapperRCAArgumentTypeError {
	return &flagsMapperRCAArgumentTypeError{
		errorLink: errorLink{id: flagsMapperRCAArgumentTypeErrorID},
		t:         t}
}

func (e *flagsMapperRCAArgumentTypeError) Error() string {
	return e.errorf("Expected flags type for the mapper algorithm but got %q", e.t)
}

// UntaggedPolicyModificationError indicates attempt to modify incrementally a policy which has no tag.
type UntaggedPolicyModificationError struct {
	errorLink
}

func newUntaggedPolicyModificationError() *UntaggedPolicyModificationError {
	return &UntaggedPolicyModificationError{
		errorLink: errorLink{id: UntaggedPolicyModificationErrorID}}
}

// Error implements error interface.
func (e *UntaggedPolicyModificationError) Error() string {
	return e.errorf("Can't modify policies with no tag")
}

// MissingPolicyTagError indicates that update has no tag to match policy before modification.
type MissingPolicyTagError struct {
	errorLink
}

func newMissingPolicyTagError() *MissingPolicyTagError {
	return &MissingPolicyTagError{
		errorLink: errorLink{id: MissingPolicyTagErrorID}}
}

// Error implements error interface.
func (e *MissingPolicyTagError) Error() string {
	return e.errorf("Update has no previous policy tag")
}

// PolicyTagsNotMatchError indicates that update tag doesn't match policy before modification.
type PolicyTagsNotMatchError struct {
	errorLink
	cntTag *uuid.UUID
	updTag *uuid.UUID
}

func newPolicyTagsNotMatchError(cntTag, updTag *uuid.UUID) *PolicyTagsNotMatchError {
	return &PolicyTagsNotMatchError{
		errorLink: errorLink{id: PolicyTagsNotMatchErrorID},
		cntTag:    cntTag,
		updTag:    updTag}
}

// Error implements error interface.
func (e *PolicyTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match policies tag %s", e.updTag.String(), e.cntTag.String())
}

type emptyPathModificationError struct {
	errorLink
}

func newEmptyPathModificationError() *emptyPathModificationError {
	return &emptyPathModificationError{
		errorLink: errorLink{id: emptyPathModificationErrorID}}
}

func (e *emptyPathModificationError) Error() string {
	return e.errorf("Can't modify items by empty path")
}

type invalidRootPolicyItemTypeError struct {
	errorLink
	item interface{}
}

func newInvalidRootPolicyItemTypeError(item interface{}) *invalidRootPolicyItemTypeError {
	return &invalidRootPolicyItemTypeError{
		errorLink: errorLink{id: invalidRootPolicyItemTypeErrorID},
		item:      item}
}

func (e *invalidRootPolicyItemTypeError) Error() string {
	return e.errorf("Expected policy or policy set as new root policy but got %T", e.item)
}

type hiddenRootPolicyAppendError struct {
	errorLink
}

func newHiddenRootPolicyAppendError() *hiddenRootPolicyAppendError {
	return &hiddenRootPolicyAppendError{
		errorLink: errorLink{id: hiddenRootPolicyAppendErrorID}}
}

func (e *hiddenRootPolicyAppendError) Error() string {
	return e.errorf("Can't append hidden policy to as root policy")
}

type invalidRootPolicyError struct {
	errorLink
	actual   string
	expected string
}

func newInvalidRootPolicyError(actual, expected string) *invalidRootPolicyError {
	return &invalidRootPolicyError{
		errorLink: errorLink{id: invalidRootPolicyErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *invalidRootPolicyError) Error() string {
	return e.errorf("Root policy is %q but got %q as first path element", e.expected, e.actual)
}

type hiddenPolicySetModificationError struct {
	errorLink
}

func newHiddenPolicySetModificationError() *hiddenPolicySetModificationError {
	return &hiddenPolicySetModificationError{
		errorLink: errorLink{id: hiddenPolicySetModificationErrorID}}
}

func (e *hiddenPolicySetModificationError) Error() string {
	return e.errorf("Can't modify hidden policy set")
}

type invalidPolicySetItemTypeError struct {
	errorLink
	item interface{}
}

func newInvalidPolicySetItemTypeError(item interface{}) *invalidPolicySetItemTypeError {
	return &invalidPolicySetItemTypeError{
		errorLink: errorLink{id: invalidPolicySetItemTypeErrorID},
		item:      item}
}

func (e *invalidPolicySetItemTypeError) Error() string {
	return e.errorf("Expected policy or policy set to append but got %T", e.item)
}

type tooShortPathPolicySetModificationError struct {
	errorLink
}

func newTooShortPathPolicySetModificationError() *tooShortPathPolicySetModificationError {
	return &tooShortPathPolicySetModificationError{
		errorLink: errorLink{id: tooShortPathPolicySetModificationErrorID}}
}

func (e *tooShortPathPolicySetModificationError) Error() string {
	return e.errorf("Path to item to delete is too short")
}

type missingPolicySetChildError struct {
	errorLink
	ID string
}

func newMissingPolicySetChildError(ID string) *missingPolicySetChildError {
	return &missingPolicySetChildError{
		errorLink: errorLink{id: missingPolicySetChildErrorID},
		ID:        ID}
}

func (e *missingPolicySetChildError) Error() string {
	return e.errorf("Policy set has no child policy or policy set with id %q", e.ID)
}

type hiddenPolicyAppendError struct {
	errorLink
}

func newHiddenPolicyAppendError() *hiddenPolicyAppendError {
	return &hiddenPolicyAppendError{
		errorLink: errorLink{id: hiddenPolicyAppendErrorID}}
}

func (e *hiddenPolicyAppendError) Error() string {
	return e.errorf("Can't append hidden policy or policy set")
}

type policyTransactionTagsNotMatchError struct {
	errorLink
	tTag uuid.UUID
	uTag uuid.UUID
}

func newPolicyTransactionTagsNotMatchError(tTag, uTag uuid.UUID) *policyTransactionTagsNotMatchError {
	return &policyTransactionTagsNotMatchError{
		errorLink: errorLink{id: policyTransactionTagsNotMatchErrorID},
		tTag:      tTag,
		uTag:      uTag}
}

func (e *policyTransactionTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match policies transaction tag %s", e.uTag.String(), e.tTag.String())
}

type failedPolicyTransactionError struct {
	errorLink
	t   uuid.UUID
	err error
}

func newFailedPolicyTransactionError(t uuid.UUID, err error) *failedPolicyTransactionError {
	return &failedPolicyTransactionError{
		errorLink: errorLink{id: failedPolicyTransactionErrorID},
		t:         t,
		err:       err}
}

func (e *failedPolicyTransactionError) Error() string {
	return e.errorf("Can't operate with failed transaction on policies %s. Previous failure %s", e.t, e.err)
}

type unknownPolicyUpdateOperationError struct {
	errorLink
	op int
}

func newUnknownPolicyUpdateOperationError(op int) *unknownPolicyUpdateOperationError {
	return &unknownPolicyUpdateOperationError{
		errorLink: errorLink{id: unknownPolicyUpdateOperationErrorID},
		op:        op}
}

func (e *unknownPolicyUpdateOperationError) Error() string {
	return e.errorf("Unknown operation %d", e.op)
}

type hiddenPolicyModificationError struct {
	errorLink
}

func newHiddenPolicyModificationError() *hiddenPolicyModificationError {
	return &hiddenPolicyModificationError{
		errorLink: errorLink{id: hiddenPolicyModificationErrorID}}
}

func (e *hiddenPolicyModificationError) Error() string {
	return e.errorf("Can't modify hidden policy")
}

type tooLongPathPolicyModificationError struct {
	errorLink
	path []string
}

func newTooLongPathPolicyModificationError(path []string) *tooLongPathPolicyModificationError {
	return &tooLongPathPolicyModificationError{
		errorLink: errorLink{id: tooLongPathPolicyModificationErrorID},
		path:      path}
}

func (e *tooLongPathPolicyModificationError) Error() string {
	return e.errorf("Trailing path \"%s\"", strings.Join(e.path, "/"))
}

type tooShortPathPolicyModificationError struct {
	errorLink
}

func newTooShortPathPolicyModificationError() *tooShortPathPolicyModificationError {
	return &tooShortPathPolicyModificationError{
		errorLink: errorLink{id: tooShortPathPolicyModificationErrorID}}
}

func (e *tooShortPathPolicyModificationError) Error() string {
	return e.errorf("Path to item to delete is too short")
}

type invalidPolicyItemTypeError struct {
	errorLink
	item interface{}
}

func newInvalidPolicyItemTypeError(item interface{}) *invalidPolicyItemTypeError {
	return &invalidPolicyItemTypeError{
		errorLink: errorLink{id: invalidPolicyItemTypeErrorID},
		item:      item}
}

func (e *invalidPolicyItemTypeError) Error() string {
	return e.errorf("Expected rule to append but got %T", e.item)
}

type hiddenRuleAppendError struct {
	errorLink
}

func newHiddenRuleAppendError() *hiddenRuleAppendError {
	return &hiddenRuleAppendError{
		errorLink: errorLink{id: hiddenRuleAppendErrorID}}
}

func (e *hiddenRuleAppendError) Error() string {
	return e.errorf("Can't append hidden rule")
}

type missingPolicyChildError struct {
	errorLink
	ID string
}

func newMissingPolicyChildError(ID string) *missingPolicyChildError {
	return &missingPolicyChildError{
		errorLink: errorLink{id: missingPolicyChildErrorID},
		ID:        ID}
}

func (e *missingPolicyChildError) Error() string {
	return e.errorf("Policy has no rule with id %q", e.ID)
}

// MissingContentError indicates that there is no desired content.
type MissingContentError struct {
	errorLink
	ID string
}

func newMissingContentError(ID string) *MissingContentError {
	return &MissingContentError{
		errorLink: errorLink{id: MissingContentErrorID},
		ID:        ID}
}

// Error implements error interface.
func (e *MissingContentError) Error() string {
	return e.errorf("Missing content %s", e.ID)
}

type invalidContentStorageItem struct {
	errorLink
	ID string
	v  interface{}
}

func newInvalidContentStorageItem(ID string, v interface{}) *invalidContentStorageItem {
	return &invalidContentStorageItem{
		errorLink: errorLink{id: invalidContentStorageItemID},
		ID:        ID,
		v:         v}
}

func (e *invalidContentStorageItem) Error() string {
	return e.errorf("Invalid value at %s (expected *LocalContent but got %T)", e.ID, e.v)
}

// MissingContentItemError indicates that content doesn't have desired item.
type MissingContentItemError struct {
	errorLink
	ID string
}

func newMissingContentItemError(ID string) *MissingContentItemError {
	return &MissingContentItemError{
		errorLink: errorLink{id: MissingContentItemErrorID},
		ID:        ID}
}

// Error implements error interface.
func (e *MissingContentItemError) Error() string {
	return e.errorf("Missing content item %q", e.ID)
}

type invalidContentItemError struct {
	errorLink
	v interface{}
}

func newInvalidContentItemError(v interface{}) *invalidContentItemError {
	return &invalidContentItemError{
		errorLink: errorLink{id: invalidContentItemErrorID},
		v:         v}
}

func (e *invalidContentItemError) Error() string {
	return e.errorf("Invalid value (expected *ContentItem but got %T)", e.v)
}

type invalidContentItemTypeError struct {
	errorLink
	expected Type
	actual   Type
}

func newInvalidContentItemTypeError(expected, actual Type) *invalidContentItemTypeError {
	return &invalidContentItemTypeError{
		errorLink: errorLink{id: invalidContentItemTypeErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *invalidContentItemTypeError) Error() string {
	return e.errorf("Invalid conent item type. Expected %q but got %q", e.expected, e.actual)
}

type invalidSelectorPathError struct {
	errorLink
	expected Signature
	actual   []Expression
}

func newInvalidSelectorPathError(expected Signature, actual []Expression) *invalidSelectorPathError {
	return &invalidSelectorPathError{
		errorLink: errorLink{id: invalidSelectorPathErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *invalidSelectorPathError) Error() string {
	actual := "nothing"
	if len(e.actual) > 0 {
		strs := make([]string, len(e.actual))
		for i, e := range e.actual {
			strs[i] = e.GetResultType().String()
		}
		actual = strings.Join(strs, "/")
	}

	return e.errorf("Invalid selector path. Expected %s path but got %s", e.expected, actual)
}

type invalidSelectorValuePathError struct {
	errorLink
	expected Signature
	actual   []AttributeValue
}

func newInvalidSelectorValuePathError(expected Signature, actual []AttributeValue) *invalidSelectorValuePathError {
	return &invalidSelectorValuePathError{
		errorLink: errorLink{id: invalidSelectorValuePathErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *invalidSelectorValuePathError) Error() string {
	actual := "nothing"
	if len(e.actual) > 0 {
		strs := make([]string, len(e.actual))
		for i, e := range e.actual {
			strs[i] = e.GetResultType().String()
		}
		actual = strings.Join(strs, "/")
	}

	return e.errorf("Invalid value path. Expected %s path but got %s", e.expected, actual)
}

type networkMapKeyValueTypeError struct {
	errorLink
	t Type
}

func newNetworkMapKeyValueTypeError(t Type) *networkMapKeyValueTypeError {
	return &networkMapKeyValueTypeError{
		errorLink: errorLink{id: networkMapKeyValueTypeErrorID},
		t:         t}
}

func (e *networkMapKeyValueTypeError) Error() string {
	return e.errorf("Expected %q or %q as network map key but got %q", TypeAddress, TypeNetwork, e.t)
}

type mapContentSubitemError struct {
	errorLink
}

func newMapContentSubitemError() *mapContentSubitemError {
	return &mapContentSubitemError{
		errorLink: errorLink{id: mapContentSubitemErrorID}}
}

func (e *mapContentSubitemError) Error() string {
	return e.errorf("Not a map of the content")
}

type invalidContentModificationError struct {
	errorLink
}

func newInvalidContentModificationError() *invalidContentModificationError {
	return &invalidContentModificationError{
		errorLink: errorLink{id: invalidContentModificationErrorID}}
}

func (e *invalidContentModificationError) Error() string {
	return e.errorf("Can't modify non-mapping content item")
}

type missingPathContentModificationError struct {
	errorLink
}

func newMissingPathContentModificationError() *missingPathContentModificationError {
	return &missingPathContentModificationError{
		errorLink: errorLink{id: missingPathContentModificationErrorID}}
}

func (e *missingPathContentModificationError) Error() string {
	return e.errorf("Missing path for content item change")
}

type tooLongPathContentModificationError struct {
	errorLink
	expected Signature
	actual   []AttributeValue
}

func newTooLongPathContentModificationError(expected Signature, actual []AttributeValue) *tooLongPathContentModificationError {
	return &tooLongPathContentModificationError{
		errorLink: errorLink{id: tooLongPathContentModificationErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *tooLongPathContentModificationError) Error() string {
	actStrs := make([]string, len(e.actual))
	for i, e := range e.actual {
		actStrs[i] = strconv.Quote(e.GetResultType().String())
	}
	actual := strings.Join(actStrs, "/")

	return e.errorf("Too long modification path. Expected %s path but got %s", e.expected, actual)
}

type invalidContentValueModificationError struct {
	errorLink
}

func newInvalidContentValueModificationError() *invalidContentValueModificationError {
	return &invalidContentValueModificationError{
		errorLink: errorLink{id: invalidContentValueModificationErrorID}}
}

func (e *invalidContentValueModificationError) Error() string {
	return e.errorf("Can't modify final content value")
}

// UntaggedContentModificationError indicates attempt to modify incrementally a content which has no tag.
type UntaggedContentModificationError struct {
	errorLink
	ID string
}

func newUntaggedContentModificationError(ID string) *UntaggedContentModificationError {
	return &UntaggedContentModificationError{
		errorLink: errorLink{id: UntaggedContentModificationErrorID},
		ID:        ID}
}

// Error implements error interface.
func (e *UntaggedContentModificationError) Error() string {
	return e.errorf("Can't modify content %q with no tag", e.ID)
}

// MissingContentTagError indicates that update has no tag to match content before modification.
type MissingContentTagError struct {
	errorLink
}

func newMissingContentTagError() *MissingContentTagError {
	return &MissingContentTagError{
		errorLink: errorLink{id: MissingContentTagErrorID}}
}

// Error implements error interface.
func (e *MissingContentTagError) Error() string {
	return e.errorf("Update has no previous content tag")
}

// ContentTagsNotMatchError indicates that update tag doesn't match content before modification.
type ContentTagsNotMatchError struct {
	errorLink
	ID     string
	cntTag *uuid.UUID
	updTag *uuid.UUID
}

func newContentTagsNotMatchError(ID string, cntTag, updTag *uuid.UUID) *ContentTagsNotMatchError {
	return &ContentTagsNotMatchError{
		errorLink: errorLink{id: ContentTagsNotMatchErrorID},
		ID:        ID,
		cntTag:    cntTag,
		updTag:    updTag}
}

// Error implements error interface.
func (e *ContentTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match content %q tag %s", e.cntTag.String(), e.ID, e.updTag.String())
}

type unknownContentUpdateOperationError struct {
	errorLink
	op int
}

func newUnknownContentUpdateOperationError(op int) *unknownContentUpdateOperationError {
	return &unknownContentUpdateOperationError{
		errorLink: errorLink{id: unknownContentUpdateOperationErrorID},
		op:        op}
}

func (e *unknownContentUpdateOperationError) Error() string {
	return e.errorf("Unknown operation %d", e.op)
}

type failedContentTransactionError struct {
	errorLink
	id  string
	t   uuid.UUID
	err error
}

func newFailedContentTransactionError(id string, t uuid.UUID, err error) *failedContentTransactionError {
	return &failedContentTransactionError{
		errorLink: errorLink{id: failedContentTransactionErrorID},
		id:        id,
		t:         t,
		err:       err}
}

func (e *failedContentTransactionError) Error() string {
	return e.errorf("Can't operate with failed transaction on content %q tagged %s. Previous failure %s", e.id, e.t, e.err)
}

type contentTransactionIDNotMatchError struct {
	errorLink
	tID string
	uID string
}

func newContentTransactionIDNotMatchError(tID, uID string) *contentTransactionIDNotMatchError {
	return &contentTransactionIDNotMatchError{
		errorLink: errorLink{id: contentTransactionIDNotMatchErrorID},
		tID:       tID,
		uID:       uID}
}

func (e *contentTransactionIDNotMatchError) Error() string {
	return e.errorf("Update content ID %q doesn't match transaction content ID %q", e.uID, e.tID)
}

type contentTransactionTagsNotMatchError struct {
	errorLink
	id   string
	tTag uuid.UUID
	uTag uuid.UUID
}

func newContentTransactionTagsNotMatchError(id string, tTag, uTag uuid.UUID) *contentTransactionTagsNotMatchError {
	return &contentTransactionTagsNotMatchError{
		errorLink: errorLink{id: contentTransactionTagsNotMatchErrorID},
		id:        id,
		tTag:      tTag,
		uTag:      uTag}
}

func (e *contentTransactionTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match content %q transaction tag %s", e.uTag.String(), e.id, e.tTag.String())
}

type tooShortRawPathContentModificationError struct {
	errorLink
}

func newTooShortRawPathContentModificationError() *tooShortRawPathContentModificationError {
	return &tooShortRawPathContentModificationError{
		errorLink: errorLink{id: tooShortRawPathContentModificationErrorID}}
}

func (e *tooShortRawPathContentModificationError) Error() string {
	return e.errorf("Expected at least content item ID in path but got nothing")
}

type tooLongRawPathContentModificationError struct {
	errorLink
	expected Signature
	actual   []string
}

func newTooLongRawPathContentModificationError(expected Signature, actual []string) *tooLongRawPathContentModificationError {
	return &tooLongRawPathContentModificationError{
		errorLink: errorLink{id: tooLongRawPathContentModificationErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *tooLongRawPathContentModificationError) Error() string {
	actStrs := make([]string, len(e.actual))
	for i, s := range e.actual {
		actStrs[i] = strconv.Quote(s)
	}
	actual := strings.Join(actStrs, "/")

	return e.errorf("Too long modification path. Expected %s path but got %s", e.expected, actual)
}

type invalidContentUpdateDataError struct {
	errorLink
	v interface{}
}

func newInvalidContentUpdateDataError(v interface{}) *invalidContentUpdateDataError {
	return &invalidContentUpdateDataError{
		errorLink: errorLink{id: invalidContentUpdateDataErrorID},
		v:         v}
}

func (e *invalidContentUpdateDataError) Error() string {
	return e.errorf("Expected content update data but got %T", e.v)
}

type invalidContentUpdateResultTypeError struct {
	errorLink
	actual   Type
	expected Type
}

func newInvalidContentUpdateResultTypeError(actual, expected Type) *invalidContentUpdateResultTypeError {
	return &invalidContentUpdateResultTypeError{
		errorLink: errorLink{id: invalidContentUpdateResultTypeErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *invalidContentUpdateResultTypeError) Error() string {
	return e.errorf("Expected %q as a result type but got %q", e.expected, e.actual)
}

type invalidContentUpdateKeysError struct {
	errorLink
	start    int
	actual   Signature
	expected Signature
}

func newInvalidContentUpdateKeysError(start int, actual, expected Signature) *invalidContentUpdateKeysError {
	return &invalidContentUpdateKeysError{
		errorLink: errorLink{id: invalidContentUpdateKeysErrorID},
		start:     start,
		actual:    actual,
		expected:  expected}
}

func (e *invalidContentUpdateKeysError) Error() string {
	return e.errorf("Expected %s path after position %d but got %s", e.expected, e.start, e.actual)
}

type unknownContentItemResultTypeError struct {
	errorLink
	t Type
}

func newUnknownContentItemResultTypeError(t Type) *unknownContentItemResultTypeError {
	return &unknownContentItemResultTypeError{
		errorLink: errorLink{id: unknownContentItemResultTypeErrorID},
		t:         t}
}

func (e *unknownContentItemResultTypeError) Error() string {
	return e.errorf("Unknown result type for content item: %q", e.t)
}

type invalidContentItemResultTypeError struct {
	errorLink
	t Type
}

func newInvalidContentItemResultTypeError(t Type) *invalidContentItemResultTypeError {
	return &invalidContentItemResultTypeError{
		errorLink: errorLink{id: invalidContentItemResultTypeErrorID},
		t:         t}
}

func (e *invalidContentItemResultTypeError) Error() string {
	return e.errorf("Invalid result type for content item: %q", e.t)
}

type invalidContentKeyTypeError struct {
	errorLink
	t        Type
	expected TypeSet
}

func newInvalidContentKeyTypeError(t Type, expected TypeSet) *invalidContentKeyTypeError {
	return &invalidContentKeyTypeError{
		errorLink: errorLink{id: invalidContentKeyTypeErrorID},
		t:         t,
		expected:  expected}
}

func (e *invalidContentKeyTypeError) Error() string {
	names := make([]string, len(e.expected))
	i := 0
	for t := range e.expected {
		names[i] = strconv.Quote(t.String())
		i++
	}
	s := strings.Join(names, ", ")

	return e.errorf("Invalid key type for content item: %q (expected %s)", e.t, s)
}

type invalidContentStringMapError struct {
	errorLink
	value interface{}
}

func newInvalidContentStringMapError(value interface{}) *invalidContentStringMapError {
	return &invalidContentStringMapError{
		errorLink: errorLink{id: invalidContentStringMapErrorID},
		value:     value}
}

func (e *invalidContentStringMapError) Error() string {
	return e.errorf("Expected string map but got %T", e.value)
}

type invalidContentNetworkMapError struct {
	errorLink
	value interface{}
}

func newInvalidContentNetworkMapError(value interface{}) *invalidContentNetworkMapError {
	return &invalidContentNetworkMapError{
		errorLink: errorLink{id: invalidContentNetworkMapErrorID},
		value:     value}
}

func (e *invalidContentNetworkMapError) Error() string {
	return e.errorf("Expected network map but got %T", e.value)
}

type invalidContentDomainMapError struct {
	errorLink
	value interface{}
}

func newInvalidContentDomainMapError(value interface{}) *invalidContentDomainMapError {
	return &invalidContentDomainMapError{
		errorLink: errorLink{id: invalidContentDomainMapErrorID},
		value:     value}
}

func (e *invalidContentDomainMapError) Error() string {
	return e.errorf("Expected domain map but got %T", e.value)
}

type invalidContentValueError struct {
	errorLink
	value interface{}
}

func newInvalidContentValueError(value interface{}) *invalidContentValueError {
	return &invalidContentValueError{
		errorLink: errorLink{id: invalidContentValueErrorID},
		value:     value}
}

func (e *invalidContentValueError) Error() string {
	return e.errorf("Expected value but got %T", e.value)
}

type invalidContentValueTypeError struct {
	errorLink
	value    interface{}
	expected Type
}

func newInvalidContentValueTypeError(value interface{}, expected Type) *invalidContentValueTypeError {
	return &invalidContentValueTypeError{
		errorLink: errorLink{id: invalidContentValueTypeErrorID},
		value:     value,
		expected:  expected}
}

func (e *invalidContentValueTypeError) Error() string {
	return e.errorf("Expected value of type %q but got %T", e.expected, e.value)
}

type invalidContentStringFlags8MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentStringFlags8MapValueError(value ContentValue) *invalidContentStringFlags8MapValueError {
	return &invalidContentStringFlags8MapValueError{
		errorLink: errorLink{id: invalidContentStringFlags8MapValueErrorID},
		value:     value}
}

func (e *invalidContentStringFlags8MapValueError) Error() string {
	return e.errorf("Expected uint8 value for string map but got %T", e.value.value)
}

type invalidContentStringFlags16MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentStringFlags16MapValueError(value ContentValue) *invalidContentStringFlags16MapValueError {
	return &invalidContentStringFlags16MapValueError{
		errorLink: errorLink{id: invalidContentStringFlags16MapValueErrorID},
		value:     value}
}

func (e *invalidContentStringFlags16MapValueError) Error() string {
	return e.errorf("Expected uint16 value for string map but got %T", e.value.value)
}

type invalidContentStringFlags32MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentStringFlags32MapValueError(value ContentValue) *invalidContentStringFlags32MapValueError {
	return &invalidContentStringFlags32MapValueError{
		errorLink: errorLink{id: invalidContentStringFlags32MapValueErrorID},
		value:     value}
}

func (e *invalidContentStringFlags32MapValueError) Error() string {
	return e.errorf("Expected uint32 value for string map but got %T", e.value.value)
}

type invalidContentStringFlags64MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentStringFlags64MapValueError(value ContentValue) *invalidContentStringFlags64MapValueError {
	return &invalidContentStringFlags64MapValueError{
		errorLink: errorLink{id: invalidContentStringFlags64MapValueErrorID},
		value:     value}
}

func (e *invalidContentStringFlags64MapValueError) Error() string {
	return e.errorf("Expected uint64 value for string map but got %T", e.value.value)
}

type invalidContentNetworkFlags8MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentNetworkFlags8MapValueError(value ContentValue) *invalidContentNetworkFlags8MapValueError {
	return &invalidContentNetworkFlags8MapValueError{
		errorLink: errorLink{id: invalidContentNetworkFlags8MapValueErrorID},
		value:     value}
}

func (e *invalidContentNetworkFlags8MapValueError) Error() string {
	return e.errorf("Expected uint8 value for network map but got %T", e.value.value)
}

type invalidContentNetworkFlags16MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentNetworkFlags16MapValueError(value ContentValue) *invalidContentNetworkFlags16MapValueError {
	return &invalidContentNetworkFlags16MapValueError{
		errorLink: errorLink{id: invalidContentNetworkFlags16MapValueErrorID},
		value:     value}
}

func (e *invalidContentNetworkFlags16MapValueError) Error() string {
	return e.errorf("Expected uint16 value for network map but got %T", e.value.value)
}

type invalidContentNetworkFlags32MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentNetworkFlags32MapValueError(value ContentValue) *invalidContentNetworkFlags32MapValueError {
	return &invalidContentNetworkFlags32MapValueError{
		errorLink: errorLink{id: invalidContentNetworkFlags32MapValueErrorID},
		value:     value}
}

func (e *invalidContentNetworkFlags32MapValueError) Error() string {
	return e.errorf("Expected uint32 value for network map but got %T", e.value.value)
}

type invalidContentNetworkFlags64MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentNetworkFlags64MapValueError(value ContentValue) *invalidContentNetworkFlags64MapValueError {
	return &invalidContentNetworkFlags64MapValueError{
		errorLink: errorLink{id: invalidContentNetworkFlags64MapValueErrorID},
		value:     value}
}

func (e *invalidContentNetworkFlags64MapValueError) Error() string {
	return e.errorf("Expected uint64 value for network map but got %T", e.value.value)
}

type invalidContentDomainFlags8MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentDomainFlags8MapValueError(value ContentValue) *invalidContentDomainFlags8MapValueError {
	return &invalidContentDomainFlags8MapValueError{
		errorLink: errorLink{id: invalidContentDomainFlags8MapValueErrorID},
		value:     value}
}

func (e *invalidContentDomainFlags8MapValueError) Error() string {
	return e.errorf("Expected uint8 value for domain map but got %T", e.value.value)
}

type invalidContentDomainFlags16MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentDomainFlags16MapValueError(value ContentValue) *invalidContentDomainFlags16MapValueError {
	return &invalidContentDomainFlags16MapValueError{
		errorLink: errorLink{id: invalidContentDomainFlags16MapValueErrorID},
		value:     value}
}

func (e *invalidContentDomainFlags16MapValueError) Error() string {
	return e.errorf("Expected uint16 value for domain map but got %T", e.value.value)
}

type invalidContentDomainFlags32MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentDomainFlags32MapValueError(value ContentValue) *invalidContentDomainFlags32MapValueError {
	return &invalidContentDomainFlags32MapValueError{
		errorLink: errorLink{id: invalidContentDomainFlags32MapValueErrorID},
		value:     value}
}

func (e *invalidContentDomainFlags32MapValueError) Error() string {
	return e.errorf("Expected uint32 value for domain map but got %T", e.value.value)
}

type invalidContentDomainFlags64MapValueError struct {
	errorLink
	value ContentValue
}

func newInvalidContentDomainFlags64MapValueError(value ContentValue) *invalidContentDomainFlags64MapValueError {
	return &invalidContentDomainFlags64MapValueError{
		errorLink: errorLink{id: invalidContentDomainFlags64MapValueErrorID},
		value:     value}
}

func (e *invalidContentDomainFlags64MapValueError) Error() string {
	return e.errorf("Expected uint64 value for domain map but got %T", e.value.value)
}

type integerDivideByZeroError struct {
	errorLink
}

func newIntegerDivideByZeroError() *integerDivideByZeroError {
	return &integerDivideByZeroError{
		errorLink: errorLink{id: integerDivideByZeroErrorID}}
}

func (e *integerDivideByZeroError) Error() string {
	return e.errorf("Integer divisor has a value of 0")
}

type floatDivideByZeroError struct {
	errorLink
}

func newFloatDivideByZeroError() *floatDivideByZeroError {
	return &floatDivideByZeroError{
		errorLink: errorLink{id: floatDivideByZeroErrorID}}
}

func (e *floatDivideByZeroError) Error() string {
	return e.errorf("Float divisor has a value of 0")
}

type floatNanError struct {
	errorLink
}

func newFloatNanError() *floatNanError {
	return &floatNanError{
		errorLink: errorLink{id: floatNanErrorID}}
}

func (e *floatNanError) Error() string {
	return e.errorf("Float result has a value of NaN")
}

type floatInfError struct {
	errorLink
}

func newFloatInfError() *floatInfError {
	return &floatInfError{
		errorLink: errorLink{id: floatInfErrorID}}
}

func (e *floatInfError) Error() string {
	return e.errorf("Float result has a value of Inf")
}

// ReadOnlySymbolsChangeError indicates attempt to modify read-only symbol table.
type ReadOnlySymbolsChangeError struct {
	errorLink
}

func newReadOnlySymbolsChangeError() *ReadOnlySymbolsChangeError {
	return &ReadOnlySymbolsChangeError{
		errorLink: errorLink{id: ReadOnlySymbolsChangeErrorID}}
}

// Error implements error interface.
func (e *ReadOnlySymbolsChangeError) Error() string {
	return e.errorf("Can't put symbol to read-only table")
}

type nilTypeError struct {
	errorLink
}

func newNilTypeError() *nilTypeError {
	return &nilTypeError{
		errorLink: errorLink{id: nilTypeErrorID}}
}

func (e *nilTypeError) Error() string {
	return e.errorf("Can't put nil type into custom types symbol table")
}

type builtinCustomTypeError struct {
	errorLink
	t Type
}

func newBuiltinCustomTypeError(t Type) *builtinCustomTypeError {
	return &builtinCustomTypeError{
		errorLink: errorLink{id: builtinCustomTypeErrorID},
		t:         t}
}

func (e *builtinCustomTypeError) Error() string {
	return e.errorf("Can't put built-in type %q into custom types symbol table", e.t)
}

type duplicateCustomTypeError struct {
	errorLink
	n Type
	p Type
}

func newDuplicateCustomTypeError(n, p Type) *duplicateCustomTypeError {
	return &duplicateCustomTypeError{
		errorLink: errorLink{id: duplicateCustomTypeErrorID},
		n:         n,
		p:         p}
}

func (e *duplicateCustomTypeError) Error() string {
	return e.errorf("Can't put type %q into symbol table as it already contains %q", e.n, e.p)
}

type duplicatesBuiltinTypeError struct {
	errorLink
	name string
}

func newDuplicatesBuiltinTypeError(name string) *duplicatesBuiltinTypeError {
	return &duplicatesBuiltinTypeError{
		errorLink: errorLink{id: duplicatesBuiltinTypeErrorID},
		name:      name}
}

func (e *duplicatesBuiltinTypeError) Error() string {
	return e.errorf("Can't create flags type %q. The name is taken by a built-in types", e.name)
}

type duplicateFlagName struct {
	errorLink
	name string
	flag string
	i    int
	j    int
}

func newDuplicateFlagName(name, flag string, i, j int) *duplicateFlagName {
	return &duplicateFlagName{
		errorLink: errorLink{id: duplicateFlagNameID},
		name:      name,
		flag:      flag,
		i:         i,
		j:         j}
}

func (e *duplicateFlagName) Error() string {
	return e.errorf("Can't create flags type %q. Flag %q at %d position duplicates flag at %d", e.name, e.flag, e.i, e.j)
}

type noTypedAttributeError struct {
	errorLink
	a Attribute
}

func newNoTypedAttributeError(a Attribute) *noTypedAttributeError {
	return &noTypedAttributeError{
		errorLink: errorLink{id: noTypedAttributeErrorID},
		a:         a}
}

func (e *noTypedAttributeError) Error() string {
	return e.errorf("Attribute %q has no type", e.a.id)
}

type undefinedAttributeTypeError struct {
	errorLink
	a Attribute
}

func newUndefinedAttributeTypeError(a Attribute) *undefinedAttributeTypeError {
	return &undefinedAttributeTypeError{
		errorLink: errorLink{id: undefinedAttributeTypeErrorID},
		a:         a}
}

func (e *undefinedAttributeTypeError) Error() string {
	return e.errorf("Attribute %q has type %q", e.a.id, TypeUndefined)
}

type unknownAttributeTypeError struct {
	errorLink
	a Attribute
}

func newUnknownAttributeTypeError(a Attribute) *unknownAttributeTypeError {
	return &unknownAttributeTypeError{
		errorLink: errorLink{id: unknownAttributeTypeErrorID},
		a:         a}
}

func (e *unknownAttributeTypeError) Error() string {
	return e.errorf("Attribute %q has unknown type %q", e.a.id, e.a.t)
}

type duplicateAttributeError struct {
	errorLink
	a Attribute
}

func newDuplicateAttributeError(a Attribute) *duplicateAttributeError {
	return &duplicateAttributeError{
		errorLink: errorLink{id: duplicateAttributeErrorID},
		a:         a}
}

func (e *duplicateAttributeError) Error() string {
	return e.errorf("Can't put attribute %q into symbol table as it already contains one with the same id", e.a.id)
}

type noFlagsDefinedError struct {
	errorLink
	name string
	n    int
}

func newNoFlagsDefinedError(name string, n int) *noFlagsDefinedError {
	return &noFlagsDefinedError{
		errorLink: errorLink{id: noFlagsDefinedErrorID},
		name:      name,
		n:         n}
}

func (e *noFlagsDefinedError) Error() string {
	return e.errorf("Required at least one flag to define flags type %q got %d", e.name, e.n)
}

type tooManyFlagsDefinedError struct {
	errorLink
	name string
	n    int
}

func newTooManyFlagsDefinedError(name string, n int) *tooManyFlagsDefinedError {
	return &tooManyFlagsDefinedError{
		errorLink: errorLink{id: tooManyFlagsDefinedErrorID},
		name:      name,
		n:         n}
}

func (e *tooManyFlagsDefinedError) Error() string {
	return e.errorf("Required no more than 64 flags to define flags type %q got %d", e.name, e.n)
}

type listOfStringsTypeError struct {
	errorLink
	t Type
}

func newListOfStringsTypeError(t Type) *listOfStringsTypeError {
	return &listOfStringsTypeError{
		errorLink: errorLink{id: listOfStringsTypeErrorID},
		t:         t}
}

func (e *listOfStringsTypeError) Error() string {
	return e.errorf("Can't convert %q to %q", e.t, TypeListOfStrings)
}

type concatTypeError struct {
	errorLink
	t Type
}

func newConcatTypeError(t Type) *concatTypeError {
	return &concatTypeError{
		errorLink: errorLink{id: concatTypeErrorID},
		t:         t}
}

func (e *concatTypeError) Error() string {
	return e.errorf("Can't concatenate %q with %q", e.t, TypeListOfStrings)
}

type unsupportedSelectorSchemeError struct {
	errorLink
	uri *url.URL
}

func newUnsupportedSelectorSchemeError(uri *url.URL) *unsupportedSelectorSchemeError {
	return &unsupportedSelectorSchemeError{
		errorLink: errorLink{id: unsupportedSelectorSchemeErrorID},
		uri:       uri}
}

func (e *unsupportedSelectorSchemeError) Error() string {
	return e.errorf("Unsupported selector scheme %q", e.uri.Scheme)
}

type disabledSelectorError struct {
	errorLink
	s Selector
}

func newDisabledSelectorError(s Selector) *disabledSelectorError {
	return &disabledSelectorError{
		errorLink: errorLink{id: disabledSelectorErrorID},
		s:         s}
}

func (e *disabledSelectorError) Error() string {
	return e.errorf("Selector for %q is disabled", e.s.Scheme())
}

type marshalInvalidDepthError struct {
	errorLink
	t int
}

func newMarshalInvalidDepthError(t int) *marshalInvalidDepthError {
	return &marshalInvalidDepthError{
		errorLink: errorLink{id: marshalInvalidDepthErrorID},
		t:         t}
}

func (e *marshalInvalidDepthError) Error() string {
	return e.errorf("Expecting depth >= 0, got %d", e.t)
}

type invalidHeaderError struct {
	errorLink
	head interface{}
}

func newInvalidHeaderError(head interface{}) *invalidHeaderError {
	return &invalidHeaderError{
		errorLink: errorLink{id: invalidHeaderErrorID},
		head:      head}
}

func (e *invalidHeaderError) Error() string {
	return e.errorf("Invalid marshaled format for head interface %v+", e.head)
}

type nonMarshableError struct {
	errorLink
	s string
}

func newNonMarshableError(s string) *nonMarshableError {
	return &nonMarshableError{
		errorLink: errorLink{id: nonMarshableErrorID},
		s:         s}
}

func (e *nonMarshableError) Error() string {
	return e.errorf("Ecountered non-marshalable node \"%s\"", e.s)
}

type nilRootError struct {
	errorLink
}

func newNilRootError() *nilRootError {
	return &nilRootError{
		errorLink: errorLink{id: nilRootErrorID}}
}

func (e *nilRootError) Error() string {
	return e.errorf("Storage root is nil")
}

// PathNotFoundError indicates a non-existent path when traversing path.
type PathNotFoundError struct {
	errorLink
	path []string
}

func newPathNotFoundError(path []string) *PathNotFoundError {
	return &PathNotFoundError{
		errorLink: errorLink{id: PathNotFoundErrorID},
		path:      path}
}

// Error implements error interface.
func (e *PathNotFoundError) Error() string {
	return e.errorf("Path %v not found", e.path)
}

type requestBufferOverflowError struct {
	errorLink
}

func newRequestBufferOverflowError() *requestBufferOverflowError {
	return &requestBufferOverflowError{
		errorLink: errorLink{id: requestBufferOverflowErrorID}}
}

func (e *requestBufferOverflowError) Error() string {
	return e.errorf("Reached end of buffer while marshalling request")
}

type requestBufferUnderflowError struct {
	errorLink
}

func newRequestBufferUnderflowError() *requestBufferUnderflowError {
	return &requestBufferUnderflowError{
		errorLink: errorLink{id: requestBufferUnderflowErrorID}}
}

func (e *requestBufferUnderflowError) Error() string {
	return e.errorf("Reached end of buffer while unmarshalling request")
}

type requestVersionError struct {
	errorLink
	actual   uint16
	expected uint16
}

func newRequestVersionError(actual, expected uint16) *requestVersionError {
	return &requestVersionError{
		errorLink: errorLink{id: requestVersionErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *requestVersionError) Error() string {
	return e.errorf("Got request of version %d while expected %d", e.actual, e.expected)
}

type requestInvalidAttributeCountError struct {
	errorLink
	n int
}

func newRequestInvalidAttributeCountError(n int) *requestInvalidAttributeCountError {
	return &requestInvalidAttributeCountError{
		errorLink: errorLink{id: requestInvalidAttributeCountErrorID},
		n:         n}
}

func (e *requestInvalidAttributeCountError) Error() string {
	return e.errorf("Invalid count of attributes %d", e.n)
}

type requestTooManyAttributesError struct {
	errorLink
	n int
}

func newRequestTooManyAttributesError(n int) *requestTooManyAttributesError {
	return &requestTooManyAttributesError{
		errorLink: errorLink{id: requestTooManyAttributesErrorID},
		n:         n}
}

func (e *requestTooManyAttributesError) Error() string {
	return e.errorf("Expected no more than %d attributes but got %d", math.MaxUint16, e.n)
}

type requestAttributeCountError struct {
	errorLink
	actual   uint16
	expected uint16
}

func newRequestAttributeCountError(actual, expected uint16) *requestAttributeCountError {
	return &requestAttributeCountError{
		errorLink: errorLink{id: requestAttributeCountErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *requestAttributeCountError) Error() string {
	return e.errorf("Expected exactly %d values but got %d", e.expected, e.actual)
}

type requestAttributeMarshallingTypeError struct {
	errorLink
	t int
}

func newRequestAttributeMarshallingTypeError(t int) *requestAttributeMarshallingTypeError {
	return &requestAttributeMarshallingTypeError{
		errorLink: errorLink{id: requestAttributeMarshallingTypeErrorID},
		t:         t}
}

func (e *requestAttributeMarshallingTypeError) Error() string {
	return e.errorf("Unknown attribute type for request %d", e.t)
}

type requestAttributeUnmarshallingTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingTypeError(t int) *requestAttributeUnmarshallingTypeError {
	return &requestAttributeUnmarshallingTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingTypeError) Error() string {
	return e.errorf("Unknown attribute type in request %d", e.t)
}

type requestAttributeUnmarshallingFlagsSizeError struct {
	errorLink
	s int
}

func newRequestAttributeUnmarshallingFlagsSizeError(s int) *requestAttributeUnmarshallingFlagsSizeError {
	return &requestAttributeUnmarshallingFlagsSizeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingFlagsSizeErrorID},
		s:         s}
}

func (e *requestAttributeUnmarshallingFlagsSizeError) Error() string {
	return e.errorf("Expected total number of flags from 1 to 64 but got %d", e.s)
}

type requestAttributeUnmarshallingBooleanTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingBooleanTypeError(t int) *requestAttributeUnmarshallingBooleanTypeError {
	return &requestAttributeUnmarshallingBooleanTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingBooleanTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingBooleanTypeError) Error() string {
	return e.errorf("Expected %q or %q value but got %q", getRequestWireTypeName(requestWireTypeBooleanFalse), getRequestWireTypeName(requestWireTypeBooleanTrue), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingStringTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingStringTypeError(t int) *requestAttributeUnmarshallingStringTypeError {
	return &requestAttributeUnmarshallingStringTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingStringTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingStringTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeString), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingIntegerTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingIntegerTypeError(t int) *requestAttributeUnmarshallingIntegerTypeError {
	return &requestAttributeUnmarshallingIntegerTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingIntegerTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingIntegerTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeInteger), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingFloatTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingFloatTypeError(t int) *requestAttributeUnmarshallingFloatTypeError {
	return &requestAttributeUnmarshallingFloatTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingFloatTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingFloatTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeFloat), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingAddressTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingAddressTypeError(t int) *requestAttributeUnmarshallingAddressTypeError {
	return &requestAttributeUnmarshallingAddressTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingAddressTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingAddressTypeError) Error() string {
	return e.errorf("Expected %q or %q value but got %q", getRequestWireTypeName(requestWireTypeIPv4Address), getRequestWireTypeName(requestWireTypeIPv6Address), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingNetworkTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingNetworkTypeError(t int) *requestAttributeUnmarshallingNetworkTypeError {
	return &requestAttributeUnmarshallingNetworkTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingNetworkTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingNetworkTypeError) Error() string {
	return e.errorf("Expected %q or %q value but got %q", getRequestWireTypeName(requestWireTypeIPv4Network), getRequestWireTypeName(requestWireTypeIPv6Network), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingDomainTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingDomainTypeError(t int) *requestAttributeUnmarshallingDomainTypeError {
	return &requestAttributeUnmarshallingDomainTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingDomainTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingDomainTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeDomain), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingSetOfStringsTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingSetOfStringsTypeError(t int) *requestAttributeUnmarshallingSetOfStringsTypeError {
	return &requestAttributeUnmarshallingSetOfStringsTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingSetOfStringsTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingSetOfStringsTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeSetOfStrings), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingSetOfNetworksTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingSetOfNetworksTypeError(t int) *requestAttributeUnmarshallingSetOfNetworksTypeError {
	return &requestAttributeUnmarshallingSetOfNetworksTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingSetOfNetworksTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingSetOfNetworksTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeSetOfNetworks), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingSetOfDomainsTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingSetOfDomainsTypeError(t int) *requestAttributeUnmarshallingSetOfDomainsTypeError {
	return &requestAttributeUnmarshallingSetOfDomainsTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingSetOfDomainsTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingSetOfDomainsTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeSetOfDomains), getRequestWireTypeName(e.t))
}

type requestAttributeUnmarshallingListOfStringsTypeError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingListOfStringsTypeError(t int) *requestAttributeUnmarshallingListOfStringsTypeError {
	return &requestAttributeUnmarshallingListOfStringsTypeError{
		errorLink: errorLink{id: requestAttributeUnmarshallingListOfStringsTypeErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingListOfStringsTypeError) Error() string {
	return e.errorf("Expected %q value but got %q", getRequestWireTypeName(requestWireTypeListOfStrings), getRequestWireTypeName(e.t))
}

type requestAttributeMarshallingNotImplementedError struct {
	errorLink
	t Type
}

func newRequestAttributeMarshallingNotImplementedError(t Type) *requestAttributeMarshallingNotImplementedError {
	return &requestAttributeMarshallingNotImplementedError{
		errorLink: errorLink{id: requestAttributeMarshallingNotImplementedErrorID},
		t:         t}
}

func (e *requestAttributeMarshallingNotImplementedError) Error() string {
	return e.errorf("Marshalling for type %q hasn't been implemented yet", e.t)
}

type requestAttributeUnmarshallingNotImplementedError struct {
	errorLink
	t int
}

func newRequestAttributeUnmarshallingNotImplementedError(t int) *requestAttributeUnmarshallingNotImplementedError {
	return &requestAttributeUnmarshallingNotImplementedError{
		errorLink: errorLink{id: requestAttributeUnmarshallingNotImplementedErrorID},
		t:         t}
}

func (e *requestAttributeUnmarshallingNotImplementedError) Error() string {
	return e.errorf("Unmarshalling for type %q hasn't been implemented yet", requestWireTypeNames[e.t])
}

type requestTooLongAttributeNameError struct {
	errorLink
	name string
}

func newRequestTooLongAttributeNameError(name string) *requestTooLongAttributeNameError {
	return &requestTooLongAttributeNameError{
		errorLink: errorLink{id: requestTooLongAttributeNameErrorID},
		name:      name}
}

func (e *requestTooLongAttributeNameError) Error() string {
	return e.errorf("Attribute name is too long %d (expected no more than %d bytes)", len(e.name), math.MaxUint8)
}

type requestTooLongStringValueError struct {
	errorLink
	s string
}

func newRequestTooLongStringValueError(s string) *requestTooLongStringValueError {
	return &requestTooLongStringValueError{
		errorLink: errorLink{id: requestTooLongStringValueErrorID},
		s:         s}
}

func (e *requestTooLongStringValueError) Error() string {
	return e.errorf("String value is too long %d (expected no more than %d bytes)", len(e.s), math.MaxUint16)
}

type requestAddressValueError struct {
	errorLink
	a []byte
}

func newRequestAddressValueError(a []byte) *requestAddressValueError {
	return &requestAddressValueError{
		errorLink: errorLink{id: requestAddressValueErrorID},
		a:         a}
}

func (e *requestAddressValueError) Error() string {
	return e.errorf("Can't treat [% x] as IPv4 or IPv6 address", e.a)
}

type requestInvalidNetworkValueError struct {
	errorLink
	n *net.IPNet
}

func newRequestInvalidNetworkValueError(n *net.IPNet) *requestInvalidNetworkValueError {
	return &requestInvalidNetworkValueError{
		errorLink: errorLink{id: requestInvalidNetworkValueErrorID},
		n:         n}
}

func (e *requestInvalidNetworkValueError) Error() string {
	return e.errorf("Can't treat %s as IPv4 or IPv6 network", e.n)
}

type requestIPv4InvalidMaskError struct {
	errorLink
	b byte
}

func newRequestIPv4InvalidMaskError(b byte) *requestIPv4InvalidMaskError {
	return &requestIPv4InvalidMaskError{
		errorLink: errorLink{id: requestIPv4InvalidMaskErrorID},
		b:         b}
}

func (e *requestIPv4InvalidMaskError) Error() string {
	return e.errorf("Invalid IPv4 CIDR: %d", e.b)
}

type requestIPv6InvalidMaskError struct {
	errorLink
	b byte
}

func newRequestIPv6InvalidMaskError(b byte) *requestIPv6InvalidMaskError {
	return &requestIPv6InvalidMaskError{
		errorLink: errorLink{id: requestIPv6InvalidMaskErrorID},
		b:         b}
}

func (e *requestIPv6InvalidMaskError) Error() string {
	return e.errorf("Invalid IPv6 CIDR: %d", e.b)
}

type requestTooLongCollectionValueError struct {
	errorLink
	t Type
	n int
}

func newRequestTooLongCollectionValueError(t Type, n int) *requestTooLongCollectionValueError {
	return &requestTooLongCollectionValueError{
		errorLink: errorLink{id: requestTooLongCollectionValueErrorID},
		t:         t,
		n:         n}
}

func (e *requestTooLongCollectionValueError) Error() string {
	return e.errorf("Number of elements in %s value is too big %d (expected no more than %d)", e.t, e.n, math.MaxUint16)
}

type requestInvalidExpressionError struct {
	errorLink
	a AttributeAssignment
}

func newRequestInvalidExpressionError(a AttributeAssignment) *requestInvalidExpressionError {
	return &requestInvalidExpressionError{
		errorLink: errorLink{id: requestInvalidExpressionErrorID},
		a:         a}
}

func (e *requestInvalidExpressionError) Error() string {
	return e.errorf("Expected value as right part of assignment expression for %s but got %T", e.a.a.describe(), e.a.e)
}

type requestAssignmentsOverflowError struct {
	errorLink
	actual   int
	expected int
}

func newRequestAssignmentsOverflowError(actual, expected int) *requestAssignmentsOverflowError {
	return &requestAssignmentsOverflowError{
		errorLink: errorLink{id: requestAssignmentsOverflowErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *requestAssignmentsOverflowError) Error() string {
	return e.errorf("Expected buffer for at least %d assignments but got %d", e.expected, e.actual)
}

type requestValuesOverflowError struct {
	errorLink
	actual   int
	expected int
}

func newRequestValuesOverflowError(actual, expected int) *requestValuesOverflowError {
	return &requestValuesOverflowError{
		errorLink: errorLink{id: requestValuesOverflowErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *requestValuesOverflowError) Error() string {
	return e.errorf("Expected buffer for at least %d values but got %d", e.expected, e.actual)
}

type requestUnmarshalEffectConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalEffectConstError(v reflect.Value) *requestUnmarshalEffectConstError {
	return &requestUnmarshalEffectConstError{
		errorLink: errorLink{id: requestUnmarshalEffectConstErrorID},
		v:         v}
}

func (e *requestUnmarshalEffectConstError) Error() string {
	return e.errorf("Can't unmarshal effect to unchengeable %s", e.v.Type())
}

type requestUnmarshalEffectTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalEffectTypeError(v reflect.Value) *requestUnmarshalEffectTypeError {
	return &requestUnmarshalEffectTypeError{
		errorLink: errorLink{id: requestUnmarshalEffectTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalEffectTypeError) Error() string {
	return e.errorf("Can't unmarshal effect to %s", e.v.Type())
}

type requestUnmarshalStatusConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalStatusConstError(v reflect.Value) *requestUnmarshalStatusConstError {
	return &requestUnmarshalStatusConstError{
		errorLink: errorLink{id: requestUnmarshalStatusConstErrorID},
		v:         v}
}

func (e *requestUnmarshalStatusConstError) Error() string {
	return e.errorf("Can't unmarshal status to unchengeable %s", e.v.Type())
}

type requestUnmarshalStatusTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalStatusTypeError(v reflect.Value) *requestUnmarshalStatusTypeError {
	return &requestUnmarshalStatusTypeError{
		errorLink: errorLink{id: requestUnmarshalStatusTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalStatusTypeError) Error() string {
	return e.errorf("Can't unmarshal status to %s", e.v.Type())
}

type requestUnmarshalBooleanConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalBooleanConstError(v reflect.Value) *requestUnmarshalBooleanConstError {
	return &requestUnmarshalBooleanConstError{
		errorLink: errorLink{id: requestUnmarshalBooleanConstErrorID},
		v:         v}
}

func (e *requestUnmarshalBooleanConstError) Error() string {
	return e.errorf("Can't unmarshal boolean to unchengeable %s", e.v.Type())
}

type requestUnmarshalBooleanTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalBooleanTypeError(v reflect.Value) *requestUnmarshalBooleanTypeError {
	return &requestUnmarshalBooleanTypeError{
		errorLink: errorLink{id: requestUnmarshalBooleanTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalBooleanTypeError) Error() string {
	return e.errorf("Can't unmarshal boolean to %s", e.v.Type())
}

type requestUnmarshalStringConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalStringConstError(v reflect.Value) *requestUnmarshalStringConstError {
	return &requestUnmarshalStringConstError{
		errorLink: errorLink{id: requestUnmarshalStringConstErrorID},
		v:         v}
}

func (e *requestUnmarshalStringConstError) Error() string {
	return e.errorf("Can't unmarshal string to unchengeable %s", e.v.Type())
}

type requestUnmarshalStringTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalStringTypeError(v reflect.Value) *requestUnmarshalStringTypeError {
	return &requestUnmarshalStringTypeError{
		errorLink: errorLink{id: requestUnmarshalStringTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalStringTypeError) Error() string {
	return e.errorf("Can't unmarshal string to %s", e.v.Type())
}

type requestUnmarshalIntegerConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalIntegerConstError(v reflect.Value) *requestUnmarshalIntegerConstError {
	return &requestUnmarshalIntegerConstError{
		errorLink: errorLink{id: requestUnmarshalIntegerConstErrorID},
		v:         v}
}

func (e *requestUnmarshalIntegerConstError) Error() string {
	return e.errorf("Can't unmarshal integer to unchengeable %s", e.v.Type())
}

type requestUnmarshalIntegerTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalIntegerTypeError(v reflect.Value) *requestUnmarshalIntegerTypeError {
	return &requestUnmarshalIntegerTypeError{
		errorLink: errorLink{id: requestUnmarshalIntegerTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalIntegerTypeError) Error() string {
	return e.errorf("Can't unmarshal integer to %s", e.v.Type())
}

type requestUnmarshalIntegerOverflowError struct {
	errorLink
	n int64
	v reflect.Value
}

func newRequestUnmarshalIntegerOverflowError(n int64, v reflect.Value) *requestUnmarshalIntegerOverflowError {
	return &requestUnmarshalIntegerOverflowError{
		errorLink: errorLink{id: requestUnmarshalIntegerOverflowErrorID},
		n:         n,
		v:         v}
}

func (e *requestUnmarshalIntegerOverflowError) Error() string {
	return e.errorf("Integer value %d overflows %s", e.n, e.v.Type())
}

type requestUnmarshalIntegerUnderflowError struct {
	errorLink
	n int64
	v reflect.Value
}

func newRequestUnmarshalIntegerUnderflowError(n int64, v reflect.Value) *requestUnmarshalIntegerUnderflowError {
	return &requestUnmarshalIntegerUnderflowError{
		errorLink: errorLink{id: requestUnmarshalIntegerUnderflowErrorID},
		n:         n,
		v:         v}
}

func (e *requestUnmarshalIntegerUnderflowError) Error() string {
	return e.errorf("Negative integer value %d underflows %s", e.n, e.v.Type())
}

type requestUnmarshalFloatConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalFloatConstError(v reflect.Value) *requestUnmarshalFloatConstError {
	return &requestUnmarshalFloatConstError{
		errorLink: errorLink{id: requestUnmarshalFloatConstErrorID},
		v:         v}
}

func (e *requestUnmarshalFloatConstError) Error() string {
	return e.errorf("Can't unmarshal float to unchengeable %s", e.v.Type())
}

type requestUnmarshalFloatTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalFloatTypeError(v reflect.Value) *requestUnmarshalFloatTypeError {
	return &requestUnmarshalFloatTypeError{
		errorLink: errorLink{id: requestUnmarshalFloatTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalFloatTypeError) Error() string {
	return e.errorf("Can't unmarshal float to %s", e.v.Type())
}

type requestUnmarshalAddressConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalAddressConstError(v reflect.Value) *requestUnmarshalAddressConstError {
	return &requestUnmarshalAddressConstError{
		errorLink: errorLink{id: requestUnmarshalAddressConstErrorID},
		v:         v}
}

func (e *requestUnmarshalAddressConstError) Error() string {
	return e.errorf("Can't unmarshal address to unchengeable %s", e.v.Type())
}

type requestUnmarshalAddressTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalAddressTypeError(v reflect.Value) *requestUnmarshalAddressTypeError {
	return &requestUnmarshalAddressTypeError{
		errorLink: errorLink{id: requestUnmarshalAddressTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalAddressTypeError) Error() string {
	return e.errorf("Can't unmarshal address to %s", e.v.Type())
}

type requestUnmarshalNetworkConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalNetworkConstError(v reflect.Value) *requestUnmarshalNetworkConstError {
	return &requestUnmarshalNetworkConstError{
		errorLink: errorLink{id: requestUnmarshalNetworkConstErrorID},
		v:         v}
}

func (e *requestUnmarshalNetworkConstError) Error() string {
	return e.errorf("Can't unmarshal network to unchengeable %s", e.v.Type())
}

type requestUnmarshalNetworkTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalNetworkTypeError(v reflect.Value) *requestUnmarshalNetworkTypeError {
	return &requestUnmarshalNetworkTypeError{
		errorLink: errorLink{id: requestUnmarshalNetworkTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalNetworkTypeError) Error() string {
	return e.errorf("Can't unmarshal network to %s", e.v.Type())
}

type requestUnmarshalDomainConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalDomainConstError(v reflect.Value) *requestUnmarshalDomainConstError {
	return &requestUnmarshalDomainConstError{
		errorLink: errorLink{id: requestUnmarshalDomainConstErrorID},
		v:         v}
}

func (e *requestUnmarshalDomainConstError) Error() string {
	return e.errorf("Can't unmarshal domain to unchengeable %s", e.v.Type())
}

type requestUnmarshalDomainTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalDomainTypeError(v reflect.Value) *requestUnmarshalDomainTypeError {
	return &requestUnmarshalDomainTypeError{
		errorLink: errorLink{id: requestUnmarshalDomainTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalDomainTypeError) Error() string {
	return e.errorf("Can't unmarshal domain to %s", e.v.Type())
}

type requestUnmarshalSetOfStringsConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalSetOfStringsConstError(v reflect.Value) *requestUnmarshalSetOfStringsConstError {
	return &requestUnmarshalSetOfStringsConstError{
		errorLink: errorLink{id: requestUnmarshalSetOfStringsConstErrorID},
		v:         v}
}

func (e *requestUnmarshalSetOfStringsConstError) Error() string {
	return e.errorf("Can't unmarshal set of strings to unchengeable %s", e.v.Type())
}

type requestUnmarshalSetOfStringsTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalSetOfStringsTypeError(v reflect.Value) *requestUnmarshalSetOfStringsTypeError {
	return &requestUnmarshalSetOfStringsTypeError{
		errorLink: errorLink{id: requestUnmarshalSetOfStringsTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalSetOfStringsTypeError) Error() string {
	return e.errorf("Can't unmarshal set of strings to %s", e.v.Type())
}

type requestUnmarshalSetOfNetworksConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalSetOfNetworksConstError(v reflect.Value) *requestUnmarshalSetOfNetworksConstError {
	return &requestUnmarshalSetOfNetworksConstError{
		errorLink: errorLink{id: requestUnmarshalSetOfNetworksConstErrorID},
		v:         v}
}

func (e *requestUnmarshalSetOfNetworksConstError) Error() string {
	return e.errorf("Can't unmarshal set of networks to unchengeable %s", e.v.Type())
}

type requestUnmarshalSetOfNetworksTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalSetOfNetworksTypeError(v reflect.Value) *requestUnmarshalSetOfNetworksTypeError {
	return &requestUnmarshalSetOfNetworksTypeError{
		errorLink: errorLink{id: requestUnmarshalSetOfNetworksTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalSetOfNetworksTypeError) Error() string {
	return e.errorf("Can't unmarshal set of networks to %s", e.v.Type())
}

type requestUnmarshalSetOfDomainsConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalSetOfDomainsConstError(v reflect.Value) *requestUnmarshalSetOfDomainsConstError {
	return &requestUnmarshalSetOfDomainsConstError{
		errorLink: errorLink{id: requestUnmarshalSetOfDomainsConstErrorID},
		v:         v}
}

func (e *requestUnmarshalSetOfDomainsConstError) Error() string {
	return e.errorf("Can't unmarshal set of domains to unchengeable %s", e.v.Type())
}

type requestUnmarshalSetOfDomainsTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalSetOfDomainsTypeError(v reflect.Value) *requestUnmarshalSetOfDomainsTypeError {
	return &requestUnmarshalSetOfDomainsTypeError{
		errorLink: errorLink{id: requestUnmarshalSetOfDomainsTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalSetOfDomainsTypeError) Error() string {
	return e.errorf("Can't unmarshal set of domains to %s", e.v.Type())
}

type requestUnmarshalListOfStringsConstError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalListOfStringsConstError(v reflect.Value) *requestUnmarshalListOfStringsConstError {
	return &requestUnmarshalListOfStringsConstError{
		errorLink: errorLink{id: requestUnmarshalListOfStringsConstErrorID},
		v:         v}
}

func (e *requestUnmarshalListOfStringsConstError) Error() string {
	return e.errorf("Can't unmarshal list of strings to unchengeable %s", e.v.Type())
}

type requestUnmarshalListOfStringsTypeError struct {
	errorLink
	v reflect.Value
}

func newRequestUnmarshalListOfStringsTypeError(v reflect.Value) *requestUnmarshalListOfStringsTypeError {
	return &requestUnmarshalListOfStringsTypeError{
		errorLink: errorLink{id: requestUnmarshalListOfStringsTypeErrorID},
		v:         v}
}

func (e *requestUnmarshalListOfStringsTypeError) Error() string {
	return e.errorf("Can't unmarshal list of strings to %s", e.v.Type())
}

type responseEffectError struct {
	errorLink
	effect int
}

func newResponseEffectError(effect int) *responseEffectError {
	return &responseEffectError{
		errorLink: errorLink{id: responseEffectErrorID},
		effect:    effect}
}

func (e *responseEffectError) Error() string {
	return e.errorf("Unknown effect %d", e.effect)
}

// ResponseServerError indicates that server returned an error message.
type ResponseServerError struct {
	errorLink
	msg string
}

func newResponseServerError(msg string) *ResponseServerError {
	return &ResponseServerError{
		errorLink: errorLink{id: ResponseServerErrorID},
		msg:       msg}
}

// Error implements error interface.
func (e *ResponseServerError) Error() string {
	return e.errorf("%s", e.msg)
}

type policyCalculationError struct {
	errorLink
	err error
}

func newPolicyCalculationError(err error) *policyCalculationError {
	return &policyCalculationError{
		errorLink: errorLink{id: policyCalculationErrorID},
		err:       err}
}

func (e *policyCalculationError) Error() string {
	return e.errorf("Failed to process request: %s", e.err)
}

type obligationCalculationError struct {
	errorLink
	a   Attribute
	err error
}

func newObligationCalculationError(a Attribute, err error) *obligationCalculationError {
	return &obligationCalculationError{
		errorLink: errorLink{id: obligationCalculationErrorID},
		a:         a,
		err:       err}
}

func (e *obligationCalculationError) Error() string {
	return e.errorf("Failed to calculate obligation for %s: %s", e.a.describe(), e.err)
}

type noInformationalError struct {
	errorLink
}

func newNoInformationalError() *noInformationalError {
	return &noInformationalError{
		errorLink: errorLink{id: noInformationalErrorID}}
}

func (e *noInformationalError) Error() string {
	return e.errorf("No information error providied to marshaller")
}
