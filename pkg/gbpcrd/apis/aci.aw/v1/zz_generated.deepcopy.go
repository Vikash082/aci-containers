// +build !ignore_autogenerated

/***
Copyright 2019 Cisco Systems Inc. All rights reserved.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Contract) DeepCopyInto(out *Contract) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Contract.
func (in *Contract) DeepCopy() *Contract {
	if in == nil {
		return nil
	}
	out := new(Contract)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Contract) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContractList) DeepCopyInto(out *ContractList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Contract, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContractList.
func (in *ContractList) DeepCopy() *ContractList {
	if in == nil {
		return nil
	}
	out := new(ContractList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContractList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContractSpec) DeepCopyInto(out *ContractSpec) {
	*out = *in
	if in.AllowList != nil {
		in, out := &in.AllowList, &out.AllowList
		*out = make([]WLRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContractSpec.
func (in *ContractSpec) DeepCopy() *ContractSpec {
	if in == nil {
		return nil
	}
	out := new(ContractSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Epg) DeepCopyInto(out *Epg) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Epg.
func (in *Epg) DeepCopy() *Epg {
	if in == nil {
		return nil
	}
	out := new(Epg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Epg) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EpgList) DeepCopyInto(out *EpgList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Epg, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EpgList.
func (in *EpgList) DeepCopy() *EpgList {
	if in == nil {
		return nil
	}
	out := new(EpgList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EpgList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EpgSpec) DeepCopyInto(out *EpgSpec) {
	*out = *in
	if in.ConsContracts != nil {
		in, out := &in.ConsContracts, &out.ConsContracts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProvContracts != nil {
		in, out := &in.ProvContracts, &out.ProvContracts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EpgSpec.
func (in *EpgSpec) DeepCopy() *EpgSpec {
	if in == nil {
		return nil
	}
	out := new(EpgSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GBPServer) DeepCopyInto(out *GBPServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GBPServer.
func (in *GBPServer) DeepCopy() *GBPServer {
	if in == nil {
		return nil
	}
	out := new(GBPServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GBPServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GBPServerList) DeepCopyInto(out *GBPServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GBPServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GBPServerList.
func (in *GBPServerList) DeepCopy() *GBPServerList {
	if in == nil {
		return nil
	}
	out := new(GBPServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GBPServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GBPServerSpec) DeepCopyInto(out *GBPServerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GBPServerSpec.
func (in *GBPServerSpec) DeepCopy() *GBPServerSpec {
	if in == nil {
		return nil
	}
	out := new(GBPServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntRange) DeepCopyInto(out *IntRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntRange.
func (in *IntRange) DeepCopy() *IntRange {
	if in == nil {
		return nil
	}
	out := new(IntRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIF) DeepCopyInto(out *PodIF) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIF.
func (in *PodIF) DeepCopy() *PodIF {
	if in == nil {
		return nil
	}
	out := new(PodIF)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodIF) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIFList) DeepCopyInto(out *PodIFList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodIF, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIFList.
func (in *PodIFList) DeepCopy() *PodIFList {
	if in == nil {
		return nil
	}
	out := new(PodIFList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodIFList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIFStatus) DeepCopyInto(out *PodIFStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIFStatus.
func (in *PodIFStatus) DeepCopy() *PodIFStatus {
	if in == nil {
		return nil
	}
	out := new(PodIFStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WLRule) DeepCopyInto(out *WLRule) {
	*out = *in
	out.Ports = in.Ports
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WLRule.
func (in *WLRule) DeepCopy() *WLRule {
	if in == nil {
		return nil
	}
	out := new(WLRule)
	in.DeepCopyInto(out)
	return out
}
