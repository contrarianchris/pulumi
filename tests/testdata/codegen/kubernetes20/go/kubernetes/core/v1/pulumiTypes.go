// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// ConfigMap holds configuration data for pods to consume.
type ConfigMapType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData map[string]string `pulumi:"binaryData"`
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data map[string]string `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable *bool `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
}

// ConfigMapTypeInput is an input type that accepts ConfigMapTypeArgs and ConfigMapTypeOutput values.
// You can construct a concrete instance of `ConfigMapTypeInput` via:
//
//	ConfigMapTypeArgs{...}
type ConfigMapTypeInput interface {
	pulumi.Input

	ToConfigMapTypeOutput() ConfigMapTypeOutput
	ToConfigMapTypeOutputWithContext(context.Context) ConfigMapTypeOutput
}

// ConfigMap holds configuration data for pods to consume.
type ConfigMapTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapInput `pulumi:"binaryData"`
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data pulumi.StringMapInput `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolPtrInput `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput `pulumi:"metadata"`
}

func (ConfigMapTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMapType)(nil)).Elem()
}

func (i ConfigMapTypeArgs) ToConfigMapTypeOutput() ConfigMapTypeOutput {
	return i.ToConfigMapTypeOutputWithContext(context.Background())
}

func (i ConfigMapTypeArgs) ToConfigMapTypeOutputWithContext(ctx context.Context) ConfigMapTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapTypeOutput)
}

// ConfigMapTypeArrayInput is an input type that accepts ConfigMapTypeArray and ConfigMapTypeArrayOutput values.
// You can construct a concrete instance of `ConfigMapTypeArrayInput` via:
//
//	ConfigMapTypeArray{ ConfigMapTypeArgs{...} }
type ConfigMapTypeArrayInput interface {
	pulumi.Input

	ToConfigMapTypeArrayOutput() ConfigMapTypeArrayOutput
	ToConfigMapTypeArrayOutputWithContext(context.Context) ConfigMapTypeArrayOutput
}

type ConfigMapTypeArray []ConfigMapTypeInput

func (ConfigMapTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigMapType)(nil)).Elem()
}

func (i ConfigMapTypeArray) ToConfigMapTypeArrayOutput() ConfigMapTypeArrayOutput {
	return i.ToConfigMapTypeArrayOutputWithContext(context.Background())
}

func (i ConfigMapTypeArray) ToConfigMapTypeArrayOutputWithContext(ctx context.Context) ConfigMapTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapTypeArrayOutput)
}

// ConfigMap holds configuration data for pods to consume.
type ConfigMapTypeOutput struct{ *pulumi.OutputState }

func (ConfigMapTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMapType)(nil)).Elem()
}

func (o ConfigMapTypeOutput) ToConfigMapTypeOutput() ConfigMapTypeOutput {
	return o
}

func (o ConfigMapTypeOutput) ToConfigMapTypeOutputWithContext(ctx context.Context) ConfigMapTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ConfigMapTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigMapType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
func (o ConfigMapTypeOutput) BinaryData() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConfigMapType) map[string]string { return v.BinaryData }).(pulumi.StringMapOutput)
}

// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
func (o ConfigMapTypeOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConfigMapType) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
func (o ConfigMapTypeOutput) Immutable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigMapType) *bool { return v.Immutable }).(pulumi.BoolPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ConfigMapTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigMapType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ConfigMapTypeOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v ConfigMapType) *metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

type ConfigMapTypeArrayOutput struct{ *pulumi.OutputState }

func (ConfigMapTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigMapType)(nil)).Elem()
}

func (o ConfigMapTypeArrayOutput) ToConfigMapTypeArrayOutput() ConfigMapTypeArrayOutput {
	return o
}

func (o ConfigMapTypeArrayOutput) ToConfigMapTypeArrayOutputWithContext(ctx context.Context) ConfigMapTypeArrayOutput {
	return o
}

func (o ConfigMapTypeArrayOutput) Index(i pulumi.IntInput) ConfigMapTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigMapType {
		return vs[0].([]ConfigMapType)[vs[1].(int)]
	}).(ConfigMapTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapTypeInput)(nil)).Elem(), ConfigMapTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapTypeArrayInput)(nil)).Elem(), ConfigMapTypeArray{})
	pulumi.RegisterOutputType(ConfigMapTypeOutput{})
	pulumi.RegisterOutputType(ConfigMapTypeArrayOutput{})
}