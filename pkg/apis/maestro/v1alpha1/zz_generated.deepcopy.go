// +build !ignore_autogenerated

/*

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(string)
		**out = **in
	}
	if in.Dest != nil {
		in, out := &in.Dest, &out.Dest
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Discovery) DeepCopyInto(out *Discovery) {
	*out = *in
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Discovery.
func (in *Discovery) DeepCopy() *Discovery {
	if in == nil {
		return nil
	}
	out := new(Discovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Framework) DeepCopyInto(out *Framework) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Framework.
func (in *Framework) DeepCopy() *Framework {
	if in == nil {
		return nil
	}
	out := new(Framework)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Framework) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkDependency) DeepCopyInto(out *FrameworkDependency) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkDependency.
func (in *FrameworkDependency) DeepCopy() *FrameworkDependency {
	if in == nil {
		return nil
	}
	out := new(FrameworkDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkList) DeepCopyInto(out *FrameworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Framework, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkList.
func (in *FrameworkList) DeepCopy() *FrameworkList {
	if in == nil {
		return nil
	}
	out := new(FrameworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrameworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkSpec) DeepCopyInto(out *FrameworkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkSpec.
func (in *FrameworkSpec) DeepCopy() *FrameworkSpec {
	if in == nil {
		return nil
	}
	out := new(FrameworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkStatus) DeepCopyInto(out *FrameworkStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkStatus.
func (in *FrameworkStatus) DeepCopy() *FrameworkStatus {
	if in == nil {
		return nil
	}
	out := new(FrameworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkVersion) DeepCopyInto(out *FrameworkVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkVersion.
func (in *FrameworkVersion) DeepCopy() *FrameworkVersion {
	if in == nil {
		return nil
	}
	out := new(FrameworkVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrameworkVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkVersionList) DeepCopyInto(out *FrameworkVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FrameworkVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkVersionList.
func (in *FrameworkVersionList) DeepCopy() *FrameworkVersionList {
	if in == nil {
		return nil
	}
	out := new(FrameworkVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FrameworkVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkVersionSpec) DeepCopyInto(out *FrameworkVersionSpec) {
	*out = *in
	out.Framework = in.Framework
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]FrameworkDependency, len(*in))
		copy(*out, *in)
	}
	if in.UpgradableFrom != nil {
		in, out := &in.UpgradableFrom, &out.UpgradableFrom
		*out = make([]FrameworkVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkVersionSpec.
func (in *FrameworkVersionSpec) DeepCopy() *FrameworkVersionSpec {
	if in == nil {
		return nil
	}
	out := new(FrameworkVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkVersionStatus) DeepCopyInto(out *FrameworkVersionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkVersionStatus.
func (in *FrameworkVersionStatus) DeepCopy() *FrameworkVersionStatus {
	if in == nil {
		return nil
	}
	out := new(FrameworkVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	if in.Cmd != nil {
		in, out := &in.Cmd, &out.Cmd
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVolume) DeepCopyInto(out *HostVolume) {
	*out = *in
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(string)
		**out = **in
	}
	if in.ContainerPath != nil {
		in, out := &in.ContainerPath, &out.ContainerPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVolume.
func (in *HostVolume) DeepCopy() *HostVolume {
	if in == nil {
		return nil
	}
	out := new(HostVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	out.FrameworkVersion = in.FrameworkVersion
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]FrameworkDependency, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.HostPorts != nil {
		in, out := &in.HostPorts, &out.HostPorts
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	if in.ContainerPorts != nil {
		in, out := &in.ContainerPorts, &out.ContainerPorts
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Phase) DeepCopyInto(out *Phase) {
	*out = *in
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(string)
		**out = **in
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]*map[string]*[][]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(map[string]*[][]string)
				if **in != nil {
					in, out := *in, *out
					*out = make(map[string]*[][]string, len(*in))
					for key, val := range *in {
						var outVal *[][]string
						if val == nil {
							(*out)[key] = nil
						} else {
							in, out := &val, &outVal
							*out = new([][]string)
							if **in != nil {
								in, out := *in, *out
								*out = make([][]string, len(*in))
								for i := range *in {
									if (*in)[i] != nil {
										in, out := &(*in)[i], &(*out)[i]
										*out = make([]string, len(*in))
										copy(*out, *in)
									}
								}
							}
						}
						(*out)[key] = outVal
					}
				}
			}
		}
	}
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Phase.
func (in *Phase) DeepCopy() *Phase {
	if in == nil {
		return nil
	}
	out := new(Phase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plan) DeepCopyInto(out *Plan) {
	*out = *in
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(string)
		**out = **in
	}
	if in.Phases != nil {
		in, out := &in.Phases, &out.Phases
		*out = make(map[string]*Phase, len(*in))
		for key, val := range *in {
			var outVal *Phase
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Phase)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plan.
func (in *Plan) DeepCopy() *Plan {
	if in == nil {
		return nil
	}
	out := new(Plan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
	if in.ResourceSets != nil {
		in, out := &in.ResourceSets, &out.ResourceSets
		*out = make(map[string]*ResourceSet, len(*in))
		for key, val := range *in {
			var outVal *ResourceSet
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ResourceSet)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make(map[string]*Network, len(*in))
		for key, val := range *in {
			var outVal *Network
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Network)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.RLimits != nil {
		in, out := &in.RLimits, &out.RLimits
		*out = make(map[string]*RLimit, len(*in))
		for key, val := range *in {
			var outVal *RLimit
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(RLimit)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make(map[string]*Task, len(*in))
		for key, val := range *in {
			var outVal *Task
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Task)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(Volume)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make(map[string]*Volume, len(*in))
		for key, val := range *in {
			var outVal *Volume
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Volume)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PreReservedRole != nil {
		in, out := &in.PreReservedRole, &out.PreReservedRole
		*out = new(string)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]*Secret, len(*in))
		for key, val := range *in {
			var outVal *Secret
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Secret)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.HostVolumes != nil {
		in, out := &in.HostVolumes, &out.HostVolumes
		*out = make(map[string]*HostVolume, len(*in))
		for key, val := range *in {
			var outVal *HostVolume
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(HostVolume)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	if in.EnvKey != nil {
		in, out := &in.EnvKey, &out.EnvKey
		*out = new(string)
		**out = **in
	}
	if in.VIP != nil {
		in, out := &in.VIP, &out.VIP
		*out = new(VIP)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RLimit) DeepCopyInto(out *RLimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RLimit.
func (in *RLimit) DeepCopy() *RLimit {
	if in == nil {
		return nil
	}
	out := new(RLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheck) DeepCopyInto(out *ReadinessCheck) {
	*out = *in
	if in.Cmd != nil {
		in, out := &in.Cmd, &out.Cmd
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheck.
func (in *ReadinessCheck) DeepCopy() *ReadinessCheck {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSet) DeepCopyInto(out *ResourceSet) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(map[string]*Port, len(*in))
		for key, val := range *in {
			var outVal *Port
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Port)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(Volume)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make(map[string]*Volume, len(*in))
		for key, val := range *in {
			var outVal *Volume
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Volume)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSet.
func (in *ResourceSet) DeepCopy() *ResourceSet {
	if in == nil {
		return nil
	}
	out := new(ResourceSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scheduler) DeepCopyInto(out *Scheduler) {
	*out = *in
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(string)
		**out = **in
	}
	if in.Zookeeper != nil {
		in, out := &in.Zookeeper, &out.Zookeeper
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler.
func (in *Scheduler) DeepCopy() *Scheduler {
	if in == nil {
		return nil
	}
	out := new(Scheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	if in.SecretPath != nil {
		in, out := &in.SecretPath, &out.SecretPath
		*out = new(string)
		**out = **in
	}
	if in.EnvKey != nil {
		in, out := &in.EnvKey, &out.EnvKey
		*out = new(string)
		**out = **in
	}
	if in.FilePath != nil {
		in, out := &in.FilePath, &out.FilePath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.WebUrl != nil {
		in, out := &in.WebUrl, &out.WebUrl
		*out = new(string)
		**out = **in
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(Scheduler)
		(*in).DeepCopyInto(*out)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(map[string]*Pod, len(*in))
		for key, val := range *in {
			var outVal *Pod
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Pod)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Plans != nil {
		in, out := &in.Plans, &out.Plans
		*out = make(map[string]*Plan, len(*in))
		for key, val := range *in {
			var outVal *Plan
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Plan)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	if in.Goal != nil {
		in, out := &in.Goal, &out.Goal
		*out = new(string)
		**out = **in
	}
	if in.Cmd != nil {
		in, out := &in.Cmd, &out.Cmd
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make(map[string]*Config, len(*in))
		for key, val := range *in {
			var outVal *Config
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Config)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(map[string]*Port, len(*in))
		for key, val := range *in {
			var outVal *Port
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Port)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessCheck != nil {
		in, out := &in.ReadinessCheck, &out.ReadinessCheck
		*out = new(ReadinessCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(Volume)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make(map[string]*Volume, len(*in))
		for key, val := range *in {
			var outVal *Volume
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Volume)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ResourceSet != nil {
		in, out := &in.ResourceSet, &out.ResourceSet
		*out = new(string)
		**out = **in
	}
	if in.Discovery != nil {
		in, out := &in.Discovery, &out.Discovery
		*out = new(Discovery)
		(*in).DeepCopyInto(*out)
	}
	if in.TransportEncryption != nil {
		in, out := &in.TransportEncryption, &out.TransportEncryption
		*out = make([]*TransportEncryption, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TransportEncryption)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransportEncryption) DeepCopyInto(out *TransportEncryption) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransportEncryption.
func (in *TransportEncryption) DeepCopy() *TransportEncryption {
	if in == nil {
		return nil
	}
	out := new(TransportEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VIP) DeepCopyInto(out *VIP) {
	*out = *in
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VIP.
func (in *VIP) DeepCopy() *VIP {
	if in == nil {
		return nil
	}
	out := new(VIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
