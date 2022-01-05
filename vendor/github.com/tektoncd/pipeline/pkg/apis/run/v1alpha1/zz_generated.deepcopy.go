// +build !ignore_autogenerated

/*
Copyright 2020 The Tekton Authors

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

package v1alpha1

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunStatus) DeepCopyInto(out *RunStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.RunStatusFields.DeepCopyInto(&out.RunStatusFields)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunStatus.
func (in *RunStatus) DeepCopy() *RunStatus {
	if in == nil {
		return nil
	}
	out := new(RunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunStatusFields) DeepCopyInto(out *RunStatusFields) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]RunResult, len(*in))
		copy(*out, *in)
	}
	if in.RetriesStatus != nil {
		in, out := &in.RetriesStatus, &out.RetriesStatus
		*out = make([]RunStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ExtraFields.DeepCopyInto(&out.ExtraFields)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunStatusFields.
func (in *RunStatusFields) DeepCopy() *RunStatusFields {
	if in == nil {
		return nil
	}
	out := new(RunStatusFields)
	in.DeepCopyInto(out)
	return out
}
