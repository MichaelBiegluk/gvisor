// Copyright 2018 The gVisor Authors.
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

// Utilities for testing capabilities.

#ifndef GVISOR_TEST_UTIL_CAPABILITY_UTIL_H_
#define GVISOR_TEST_UTIL_CAPABILITY_UTIL_H_

#if defined(__Fuchsia__)
#include "test/util/fuchsia_capability_util.h"
#elif defined(__linux__)
#include "test/util/linux_capability_util.h"
#else
#error "Unhandled platform"
#endif

#endif  // GVISOR_TEST_UTIL_CAPABILITY_UTIL_H_