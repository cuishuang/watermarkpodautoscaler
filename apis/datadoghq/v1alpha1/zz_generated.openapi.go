//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Code generated by openapi-gen. DO NOT EDIT.

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.CrossVersionObjectReference":  schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_CrossVersionObjectReference(ref),
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.ExternalMetricSource":         schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_ExternalMetricSource(ref),
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.MetricSpec":                   schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_MetricSpec(ref),
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.RecommenderSpec":              schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_RecommenderSpec(ref),
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.ResourceMetricSource":         schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_ResourceMetricSource(ref),
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.WatermarkPodAutoscaler":       schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_WatermarkPodAutoscaler(ref),
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.WatermarkPodAutoscalerSpec":   schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_WatermarkPodAutoscalerSpec(ref),
		"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.WatermarkPodAutoscalerStatus": schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_WatermarkPodAutoscalerStatus(ref),
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_CrossVersionObjectReference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CrossVersionObjectReference contains enough information to let you identify the referred resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds\"",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "API version of the referent",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"kind", "name"},
			},
		},
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_ExternalMetricSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). Exactly one \"target\" type should be set.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"metricName": {
						SchemaProps: spec.SchemaProps{
							Description: "metricName is the name of the metric in question.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metricSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "metricSelector is used to identify a specific time series within a given metric.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"highWatermark": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"lowWatermark": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
				},
				Required: []string{"metricName"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/api/resource.Quantity", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_MetricSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MetricSpec specifies how to scale based on a single metric (only `type` and one other matching field should be set at once).",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "type is the type of metric source.  It should be one of \"Object\", \"Pods\" or \"Resource\", each mapping to a matching field in the object.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"external": {
						SchemaProps: spec.SchemaProps{
							Description: "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
							Ref:         ref("github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.ExternalMetricSource"),
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
							Ref:         ref("github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.ResourceMetricSource"),
						},
					},
				},
				Required: []string{"type"},
			},
		},
		Dependencies: []string{
			"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.ExternalMetricSource", "github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.ResourceMetricSource"},
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_RecommenderSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RecommenderSpec indicates which recommender service to use to calculate the desired replica count\n\nSee https://github.com/DataDog/agent-payload/pull/348 for details about the API.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL of the recommender service to use",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"Settings": {
						SchemaProps: spec.SchemaProps{
							Description: "Settings to pass to the recommender service",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"targetType": {
						SchemaProps: spec.SchemaProps{
							Description: "TargetType is the type of target the recommender service should use.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"highWatermark": {
						SchemaProps: spec.SchemaProps{
							Description: "These will map to lowerBound/upperBound in the recommender service",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"lowWatermark": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
				},
				Required: []string{"url"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/api/resource.Quantity"},
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_ResourceMetricSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.  Only one \"target\" type should be set.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "name is the name of the resource in question.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metricSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "metricSelector is used to identify a specific time series within a given metric.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"highWatermark": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"lowWatermark": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
				},
				Required: []string{"name"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/api/resource.Quantity", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_WatermarkPodAutoscaler(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WatermarkPodAutoscaler is the Schema for the watermarkpodautoscalers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.WatermarkPodAutoscalerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.WatermarkPodAutoscalerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.WatermarkPodAutoscalerSpec", "github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.WatermarkPodAutoscalerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_WatermarkPodAutoscalerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WatermarkPodAutoscalerSpec defines the desired state of WatermarkPodAutoscaler",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"downscaleForbiddenWindowSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "part of HorizontalController, see comments in the k8s repo: pkg/controller/podautoscaler/horizontal.go",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"upscaleForbiddenWindowSeconds": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"scaleUpLimitFactor": {
						SchemaProps: spec.SchemaProps{
							Description: "Percentage of replicas that can be added in an upscale event. Parameter used to be a float, in order to support the transition seamlessly, we validate that it is [0;100] in the code. ScaleUpLimitFactor == 0 means that upscaling will not be allowed for the target.",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"upscaleDelayAboveWatermarkSeconds": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"scaleDownLimitFactor": {
						SchemaProps: spec.SchemaProps{
							Description: "Percentage of replicas that can be removed in an downscale event. Parameter used to be a float, in order to support the transition seamlessly, we validate that it is [0;100[ in the code. ScaleDownLimitFactor == 0 means that downscaling will not be allowed for the target.",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"downscaleDelayBelowWatermarkSeconds": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"replicaScalingAbsoluteModulo": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of replicas to scale by at a time. When set, replicas added or removed must be a multiple of this parameter. Allows for special scaling patterns, for instance when an application requires a certain number of pods in multiple",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"convergeTowardsWatermark": {
						SchemaProps: spec.SchemaProps{
							Description: "Try to make the usage converge towards High Watermark to save resources. This will slowly downscale by `ReplicaScalingAbsoluteModulo` if the predicted usage stays bellow the high watermarks.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tolerance": {
						SchemaProps: spec.SchemaProps{
							Description: "Parameter used to be a float, in order to support the transition seamlessly, we validate that it is ]0;1[ in the code.",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"algorithm": {
						SchemaProps: spec.SchemaProps{
							Description: "computed values take the # of replicas into account",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dryRun": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether planned scale changes are actually applied",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"tolerateZero": {
						SchemaProps: spec.SchemaProps{
							Description: "Zero is a value that can lead to undesired outcomes, unless explicitly set the WPA will not take action if the value retrieved is 0.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"scaleTargetRef": {
						SchemaProps: spec.SchemaProps{
							Description: "part of HorizontalPodAutoscalerSpec, see comments in the k8s-1.10.8 repo: staging/src/k8s.io/api/autoscaling/v1/types.go reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.CrossVersionObjectReference"),
						},
					},
					"metrics": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "specifications that will be used to calculate the desired replica count",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.MetricSpec"),
									},
								},
							},
						},
					},
					"recommender": {
						SchemaProps: spec.SchemaProps{
							Description: "recommender that can be used to request the desired replica count",
							Ref:         ref("github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.RecommenderSpec"),
						},
					},
					"minReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"minAvailableReplicaPercentage": {
						SchemaProps: spec.SchemaProps{
							Description: "MinAvailableReplicaPercentage indicates the minimum percentage of replicas that need to be available in order for the controller to autoscale the target.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"maxReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"readinessDelaySeconds": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"scaleTargetRef"},
			},
		},
		Dependencies: []string{
			"github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.CrossVersionObjectReference", "github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.MetricSpec", "github.com/DataDog/watermarkpodautoscaler/apis/datadoghq/v1alpha1.RecommenderSpec", "k8s.io/apimachinery/pkg/api/resource.Quantity"},
	}
}

func schema_watermarkpodautoscaler_apis_datadoghq_v1alpha1_WatermarkPodAutoscalerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WatermarkPodAutoscalerStatus defines the observed state of WatermarkPodAutoscaler",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"lastScaleTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"scalingEventsCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"currentReplicas": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"desiredReplicas": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"currentMetrics": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/autoscaling/v2beta1.MetricStatus"),
									},
								},
							},
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/autoscaling/v2beta1.HorizontalPodAutoscalerCondition"),
									},
								},
							},
						},
					},
					"lastConditionType": {
						SchemaProps: spec.SchemaProps{
							Description: "LastConditionType correspond to the last condition type updated in the WPA status during the WPA reconcile state.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastConditionState": {
						SchemaProps: spec.SchemaProps{
							Description: "LastConditionType correspond to the last condition state (True,False) updated in the WPA status during the WPA reconcile state.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"currentReplicas", "desiredReplicas"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/autoscaling/v2beta1.HorizontalPodAutoscalerCondition", "k8s.io/api/autoscaling/v2beta1.MetricStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
