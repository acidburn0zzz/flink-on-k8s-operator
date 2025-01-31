/*
Copyright 2019 Google LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// Sets default values for unspecified FlinkCluster properties.
func _SetDefault(cluster *FlinkCluster) {
	_SetImageDefault(&cluster.Spec.ImageSpec)
	_SetJobManagerDefault(&cluster.Spec.JobManagerSpec)
	_SetTaskManagerDefault(&cluster.Spec.TaskManagerSpec)
	_SetJobDefault(cluster.Spec.JobSpec)
}

func _SetImageDefault(imageSpec *ImageSpec) {
	if len(imageSpec.PullPolicy) == 0 {
		imageSpec.PullPolicy = corev1.PullAlways
	}
}

func _SetJobManagerDefault(jmSpec *JobManagerSpec) {
	if jmSpec.Replicas == nil {
		jmSpec.Replicas = new(int32)
		*jmSpec.Replicas = 1
	}
	if len(jmSpec.AccessScope) == 0 {
		jmSpec.AccessScope = AccessScope.Cluster
	}
	if jmSpec.Ports.RPC == nil {
		jmSpec.Ports.RPC = new(int32)
		*jmSpec.Ports.RPC = 6123
	}
	if jmSpec.Ports.Blob == nil {
		jmSpec.Ports.Blob = new(int32)
		*jmSpec.Ports.Blob = 6124
	}
	if jmSpec.Ports.Query == nil {
		jmSpec.Ports.Query = new(int32)
		*jmSpec.Ports.Query = 6125
	}
	if jmSpec.Ports.UI == nil {
		jmSpec.Ports.UI = new(int32)
		*jmSpec.Ports.UI = 8081
	}
}

func _SetTaskManagerDefault(tmSpec *TaskManagerSpec) {
	if tmSpec.Ports.Data == nil {
		tmSpec.Ports.Data = new(int32)
		*tmSpec.Ports.Data = 6121
	}
	if tmSpec.Ports.RPC == nil {
		tmSpec.Ports.RPC = new(int32)
		*tmSpec.Ports.RPC = 6122
	}
	if tmSpec.Ports.Query == nil {
		tmSpec.Ports.Query = new(int32)
		*tmSpec.Ports.Query = 6125
	}
}

func _SetJobDefault(jobSpec *JobSpec) {
	if jobSpec == nil {
		return
	}
	if jobSpec.AllowNonRestoredState == nil {
		jobSpec.AllowNonRestoredState = new(bool)
		*jobSpec.AllowNonRestoredState = false
	}
	if jobSpec.Parallelism == nil {
		jobSpec.Parallelism = new(int32)
		*jobSpec.Parallelism = 1
	}
	if jobSpec.NoLoggingToStdout == nil {
		jobSpec.NoLoggingToStdout = new(bool)
		*jobSpec.NoLoggingToStdout = false
	}
	if jobSpec.RestartPolicy == nil {
		jobSpec.RestartPolicy = new(corev1.RestartPolicy)
		*jobSpec.RestartPolicy = corev1.RestartPolicyOnFailure
	}
}
