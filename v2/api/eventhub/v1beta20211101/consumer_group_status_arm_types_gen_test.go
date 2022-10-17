// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
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

func Test_ConsumerGroup_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConsumerGroup_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConsumerGroup_STATUS_ARM, ConsumerGroup_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConsumerGroup_STATUS_ARM runs a test to see if a specific instance of ConsumerGroup_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConsumerGroup_STATUS_ARM(subject ConsumerGroup_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConsumerGroup_STATUS_ARM
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

// Generator of ConsumerGroup_STATUS_ARM instances for property testing - lazily instantiated by
// ConsumerGroup_STATUS_ARMGenerator()
var consumerGroup_STATUS_ARMGenerator gopter.Gen

// ConsumerGroup_STATUS_ARMGenerator returns a generator of ConsumerGroup_STATUS_ARM instances for property testing.
// We first initialize consumerGroup_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConsumerGroup_STATUS_ARMGenerator() gopter.Gen {
	if consumerGroup_STATUS_ARMGenerator != nil {
		return consumerGroup_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroup_STATUS_ARM(generators)
	consumerGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroup_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForConsumerGroup_STATUS_ARM(generators)
	consumerGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_STATUS_ARM{}), generators)

	return consumerGroup_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConsumerGroup_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConsumerGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConsumerGroup_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConsumerGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ConsumerGroup_Properties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_ConsumerGroup_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConsumerGroup_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConsumerGroup_Properties_STATUS_ARM, ConsumerGroup_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConsumerGroup_Properties_STATUS_ARM runs a test to see if a specific instance of ConsumerGroup_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConsumerGroup_Properties_STATUS_ARM(subject ConsumerGroup_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConsumerGroup_Properties_STATUS_ARM
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

// Generator of ConsumerGroup_Properties_STATUS_ARM instances for property testing - lazily instantiated by
// ConsumerGroup_Properties_STATUS_ARMGenerator()
var consumerGroup_Properties_STATUS_ARMGenerator gopter.Gen

// ConsumerGroup_Properties_STATUS_ARMGenerator returns a generator of ConsumerGroup_Properties_STATUS_ARM instances for property testing.
func ConsumerGroup_Properties_STATUS_ARMGenerator() gopter.Gen {
	if consumerGroup_Properties_STATUS_ARMGenerator != nil {
		return consumerGroup_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroup_Properties_STATUS_ARM(generators)
	consumerGroup_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_Properties_STATUS_ARM{}), generators)

	return consumerGroup_Properties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConsumerGroup_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConsumerGroup_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}