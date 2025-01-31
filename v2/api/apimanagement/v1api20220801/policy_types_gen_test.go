// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"encoding/json"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Policy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Policy to hub returns original",
		prop.ForAll(RunResourceConversionTestForPolicy, PolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPolicy tests if a specific instance of Policy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPolicy(subject Policy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220801s.Policy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Policy
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Policy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Policy to Policy via AssignProperties_To_Policy & AssignProperties_From_Policy returns original",
		prop.ForAll(RunPropertyAssignmentTestForPolicy, PolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPolicy tests if a specific instance of Policy can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForPolicy(subject Policy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Policy
	err := copied.AssignProperties_To_Policy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Policy
	err = actual.AssignProperties_From_Policy(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Policy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Policy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicy, PolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicy runs a test to see if a specific instance of Policy round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicy(subject Policy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Policy
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Policy instances for property testing - lazily instantiated by PolicyGenerator()
var policyGenerator gopter.Gen

// PolicyGenerator returns a generator of Policy instances for property testing.
func PolicyGenerator() gopter.Gen {
	if policyGenerator != nil {
		return policyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPolicy(generators)
	policyGenerator = gen.Struct(reflect.TypeOf(Policy{}), generators)

	return policyGenerator
}

// AddRelatedPropertyGeneratorsForPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_Policy_SpecGenerator()
	gens["Status"] = Service_Policy_STATUSGenerator()
}

func Test_Service_Policy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_Policy_Spec to Service_Policy_Spec via AssignProperties_To_Service_Policy_Spec & AssignProperties_From_Service_Policy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_Policy_Spec, Service_Policy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_Policy_Spec tests if a specific instance of Service_Policy_Spec can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForService_Policy_Spec(subject Service_Policy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Service_Policy_Spec
	err := copied.AssignProperties_To_Service_Policy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_Policy_Spec
	err = actual.AssignProperties_From_Service_Policy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Service_Policy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Policy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Policy_Spec, Service_Policy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Policy_Spec runs a test to see if a specific instance of Service_Policy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Policy_Spec(subject Service_Policy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Policy_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Service_Policy_Spec instances for property testing - lazily instantiated by
// Service_Policy_SpecGenerator()
var service_Policy_SpecGenerator gopter.Gen

// Service_Policy_SpecGenerator returns a generator of Service_Policy_Spec instances for property testing.
func Service_Policy_SpecGenerator() gopter.Gen {
	if service_Policy_SpecGenerator != nil {
		return service_Policy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Policy_Spec(generators)
	service_Policy_SpecGenerator = gen.Struct(reflect.TypeOf(Service_Policy_Spec{}), generators)

	return service_Policy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_Policy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Policy_Spec(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.OneConstOf(
		PolicyContractProperties_Format_Rawxml,
		PolicyContractProperties_Format_RawxmlLink,
		PolicyContractProperties_Format_Xml,
		PolicyContractProperties_Format_XmlLink))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_Service_Policy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Service_Policy_STATUS to Service_Policy_STATUS via AssignProperties_To_Service_Policy_STATUS & AssignProperties_From_Service_Policy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForService_Policy_STATUS, Service_Policy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForService_Policy_STATUS tests if a specific instance of Service_Policy_STATUS can be assigned to v1api20220801storage and back losslessly
func RunPropertyAssignmentTestForService_Policy_STATUS(subject Service_Policy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220801s.Service_Policy_STATUS
	err := copied.AssignProperties_To_Service_Policy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Service_Policy_STATUS
	err = actual.AssignProperties_From_Service_Policy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Service_Policy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_Policy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_Policy_STATUS, Service_Policy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_Policy_STATUS runs a test to see if a specific instance of Service_Policy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_Policy_STATUS(subject Service_Policy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_Policy_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Service_Policy_STATUS instances for property testing - lazily instantiated by
// Service_Policy_STATUSGenerator()
var service_Policy_STATUSGenerator gopter.Gen

// Service_Policy_STATUSGenerator returns a generator of Service_Policy_STATUS instances for property testing.
func Service_Policy_STATUSGenerator() gopter.Gen {
	if service_Policy_STATUSGenerator != nil {
		return service_Policy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_Policy_STATUS(generators)
	service_Policy_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_Policy_STATUS{}), generators)

	return service_Policy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_Policy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_Policy_STATUS(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.OneConstOf(
		PolicyContractProperties_Format_STATUS_Rawxml,
		PolicyContractProperties_Format_STATUS_RawxmlLink,
		PolicyContractProperties_Format_STATUS_Xml,
		PolicyContractProperties_Format_STATUS_XmlLink))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
