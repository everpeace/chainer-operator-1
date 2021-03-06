// +build !ignore_autogenerated

// Copyright 2018 The Kubeflow Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChainerJob) DeepCopyInto(out *ChainerJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChainerJob.
func (in *ChainerJob) DeepCopy() *ChainerJob {
	if in == nil {
		return nil
	}
	out := new(ChainerJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChainerJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChainerJobList) DeepCopyInto(out *ChainerJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChainerJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChainerJobList.
func (in *ChainerJobList) DeepCopy() *ChainerJobList {
	if in == nil {
		return nil
	}
	out := new(ChainerJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChainerJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChainerJobSpec) DeepCopyInto(out *ChainerJobSpec) {
	*out = *in
	in.Master.DeepCopyInto(&out.Master)
	if in.WorkerSets != nil {
		in, out := &in.WorkerSets, &out.WorkerSets
		*out = make(map[string]*WorkerSetSpec, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(WorkerSetSpec)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChainerJobSpec.
func (in *ChainerJobSpec) DeepCopy() *ChainerJobSpec {
	if in == nil {
		return nil
	}
	out := new(ChainerJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIConfig) DeepCopyInto(out *MPIConfig) {
	*out = *in
	if in.Slots != nil {
		in, out := &in.Slots, &out.Slots
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIConfig.
func (in *MPIConfig) DeepCopy() *MPIConfig {
	if in == nil {
		return nil
	}
	out := new(MPIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterSpec) DeepCopyInto(out *MasterSpec) {
	*out = *in
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.MPIConfig != nil {
		in, out := &in.MPIConfig, &out.MPIConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(MPIConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterSpec.
func (in *MasterSpec) DeepCopy() *MasterSpec {
	if in == nil {
		return nil
	}
	out := new(MasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSetSpec) DeepCopyInto(out *WorkerSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.MPIConfig != nil {
		in, out := &in.MPIConfig, &out.MPIConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(MPIConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSetSpec.
func (in *WorkerSetSpec) DeepCopy() *WorkerSetSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerSetSpec)
	in.DeepCopyInto(out)
	return out
}
