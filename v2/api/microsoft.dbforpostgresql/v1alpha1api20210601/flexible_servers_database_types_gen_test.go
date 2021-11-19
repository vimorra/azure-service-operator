// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.dbforpostgresql/v1alpha1api20210601storage"
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

func Test_FlexibleServersDatabase_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersDatabase to FlexibleServersDatabase via AssignPropertiesToFlexibleServersDatabase & AssignPropertiesFromFlexibleServersDatabase returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersDatabase, FlexibleServersDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersDatabase tests if a specific instance of FlexibleServersDatabase can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersDatabase(subject FlexibleServersDatabase) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210601storage.FlexibleServersDatabase
	err := copied.AssignPropertiesToFlexibleServersDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersDatabase
	err = actual.AssignPropertiesFromFlexibleServersDatabase(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabase, FlexibleServersDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabase runs a test to see if a specific instance of FlexibleServersDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabase(subject FlexibleServersDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabase
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

// Generator of FlexibleServersDatabase instances for property testing - lazily instantiated by
//FlexibleServersDatabaseGenerator()
var flexibleServersDatabaseGenerator gopter.Gen

// FlexibleServersDatabaseGenerator returns a generator of FlexibleServersDatabase instances for property testing.
func FlexibleServersDatabaseGenerator() gopter.Gen {
	if flexibleServersDatabaseGenerator != nil {
		return flexibleServersDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersDatabase(generators)
	flexibleServersDatabaseGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase{}), generators)

	return flexibleServersDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersDatabasesSpecGenerator()
	gens["Status"] = DatabaseStatusGenerator()
}

func Test_Database_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Database_Status to Database_Status via AssignPropertiesToDatabaseStatus & AssignPropertiesFromDatabaseStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseStatus, DatabaseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseStatus tests if a specific instance of Database_Status can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForDatabaseStatus(subject Database_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210601storage.Database_Status
	err := copied.AssignPropertiesToDatabaseStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Database_Status
	err = actual.AssignPropertiesFromDatabaseStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Database_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseStatus, DatabaseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseStatus runs a test to see if a specific instance of Database_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseStatus(subject Database_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_Status
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

// Generator of Database_Status instances for property testing - lazily instantiated by DatabaseStatusGenerator()
var databaseStatusGenerator gopter.Gen

// DatabaseStatusGenerator returns a generator of Database_Status instances for property testing.
// We first initialize databaseStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseStatusGenerator() gopter.Gen {
	if databaseStatusGenerator != nil {
		return databaseStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatus(generators)
	databaseStatusGenerator = gen.Struct(reflect.TypeOf(Database_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatus(generators)
	AddRelatedPropertyGeneratorsForDatabaseStatus(generators)
	databaseStatusGenerator = gen.Struct(reflect.TypeOf(Database_Status{}), generators)

	return databaseStatusGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseStatus(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseStatus(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_FlexibleServersDatabases_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersDatabases_Spec to FlexibleServersDatabases_Spec via AssignPropertiesToFlexibleServersDatabasesSpec & AssignPropertiesFromFlexibleServersDatabasesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersDatabasesSpec, FlexibleServersDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersDatabasesSpec tests if a specific instance of FlexibleServersDatabases_Spec can be assigned to v1alpha1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersDatabasesSpec(subject FlexibleServersDatabases_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210601storage.FlexibleServersDatabases_Spec
	err := copied.AssignPropertiesToFlexibleServersDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersDatabases_Spec
	err = actual.AssignPropertiesFromFlexibleServersDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlexibleServersDatabases_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabases_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabasesSpec, FlexibleServersDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabasesSpec runs a test to see if a specific instance of FlexibleServersDatabases_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabasesSpec(subject FlexibleServersDatabases_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabases_Spec
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

// Generator of FlexibleServersDatabases_Spec instances for property testing - lazily instantiated by
//FlexibleServersDatabasesSpecGenerator()
var flexibleServersDatabasesSpecGenerator gopter.Gen

// FlexibleServersDatabasesSpecGenerator returns a generator of FlexibleServersDatabases_Spec instances for property testing.
func FlexibleServersDatabasesSpecGenerator() gopter.Gen {
	if flexibleServersDatabasesSpecGenerator != nil {
		return flexibleServersDatabasesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabasesSpec(generators)
	flexibleServersDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabases_Spec{}), generators)

	return flexibleServersDatabasesSpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersDatabasesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersDatabasesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}