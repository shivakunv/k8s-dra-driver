//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright (c) 2023, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableDevice) DeepCopyInto(out *AllocatableDevice) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(AllocatableGpu)
		**out = **in
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(AllocatableMigDevice)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableDevice.
func (in *AllocatableDevice) DeepCopy() *AllocatableDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatableDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableGpu) DeepCopyInto(out *AllocatableGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableGpu.
func (in *AllocatableGpu) DeepCopy() *AllocatableGpu {
	if in == nil {
		return nil
	}
	out := new(AllocatableGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableMigDevice) DeepCopyInto(out *AllocatableMigDevice) {
	*out = *in
	if in.Placements != nil {
		in, out := &in.Placements, &out.Placements
		*out = make([]MigDevicePlacement, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableMigDevice.
func (in *AllocatableMigDevice) DeepCopy() *AllocatableMigDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatableMigDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedDevice) DeepCopyInto(out *AllocatedDevice) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(AllocatedGpu)
		**out = **in
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(AllocatedMigDevice)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedDevice.
func (in *AllocatedDevice) DeepCopy() *AllocatedDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatedDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AllocatedDevices) DeepCopyInto(out *AllocatedDevices) {
	{
		in := &in
		*out = make(AllocatedDevices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedDevices.
func (in AllocatedDevices) DeepCopy() AllocatedDevices {
	if in == nil {
		return nil
	}
	out := new(AllocatedDevices)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedGpu) DeepCopyInto(out *AllocatedGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedGpu.
func (in *AllocatedGpu) DeepCopy() *AllocatedGpu {
	if in == nil {
		return nil
	}
	out := new(AllocatedGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedMigDevice) DeepCopyInto(out *AllocatedMigDevice) {
	*out = *in
	out.Placement = in.Placement
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedMigDevice.
func (in *AllocatedMigDevice) DeepCopy() *AllocatedMigDevice {
	if in == nil {
		return nil
	}
	out := new(AllocatedMigDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDevicePlacement) DeepCopyInto(out *MigDevicePlacement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDevicePlacement.
func (in *MigDevicePlacement) DeepCopy() *MigDevicePlacement {
	if in == nil {
		return nil
	}
	out := new(MigDevicePlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationState) DeepCopyInto(out *NodeAllocationState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationState.
func (in *NodeAllocationState) DeepCopy() *NodeAllocationState {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateConfig) DeepCopyInto(out *NodeAllocationStateConfig) {
	*out = *in
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(v1.OwnerReference)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateConfig.
func (in *NodeAllocationStateConfig) DeepCopy() *NodeAllocationStateConfig {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateList) DeepCopyInto(out *NodeAllocationStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeAllocationState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateList.
func (in *NodeAllocationStateList) DeepCopy() *NodeAllocationStateList {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateSpec) DeepCopyInto(out *NodeAllocationStateSpec) {
	*out = *in
	if in.AllocatableDevices != nil {
		in, out := &in.AllocatableDevices, &out.AllocatableDevices
		*out = make([]AllocatableDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClaimRequests != nil {
		in, out := &in.ClaimRequests, &out.ClaimRequests
		*out = make(map[string]RequestedDevices, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ClaimAllocations != nil {
		in, out := &in.ClaimAllocations, &out.ClaimAllocations
		*out = make(map[string]AllocatedDevices, len(*in))
		for key, val := range *in {
			var outVal []AllocatedDevice
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(AllocatedDevices, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateSpec.
func (in *NodeAllocationStateSpec) DeepCopy() *NodeAllocationStateSpec {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedDevices) DeepCopyInto(out *RequestedDevices) {
	*out = *in
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(RequestedGpus)
		(*in).DeepCopyInto(*out)
	}
	if in.Mig != nil {
		in, out := &in.Mig, &out.Mig
		*out = new(RequestedMigDevices)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedDevices.
func (in *RequestedDevices) DeepCopy() *RequestedDevices {
	if in == nil {
		return nil
	}
	out := new(RequestedDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedGpu) DeepCopyInto(out *RequestedGpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedGpu.
func (in *RequestedGpu) DeepCopy() *RequestedGpu {
	if in == nil {
		return nil
	}
	out := new(RequestedGpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedGpus) DeepCopyInto(out *RequestedGpus) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]RequestedGpu, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedGpus.
func (in *RequestedGpus) DeepCopy() *RequestedGpus {
	if in == nil {
		return nil
	}
	out := new(RequestedGpus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedMigDevice) DeepCopyInto(out *RequestedMigDevice) {
	*out = *in
	out.Placement = in.Placement
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedMigDevice.
func (in *RequestedMigDevice) DeepCopy() *RequestedMigDevice {
	if in == nil {
		return nil
	}
	out := new(RequestedMigDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedMigDevices) DeepCopyInto(out *RequestedMigDevices) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]RequestedMigDevice, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedMigDevices.
func (in *RequestedMigDevices) DeepCopy() *RequestedMigDevices {
	if in == nil {
		return nil
	}
	out := new(RequestedMigDevices)
	in.DeepCopyInto(out)
	return out
}
