# Author:: Christian Vozar <christian@bellycard.com>
# Cookbook Name:: go
# Attributes:: default
#
# Copyright 2013, Belly, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Install Method
default["google"]["go"]["install_method"] = "package" # options: package, source

# Source attributes
default["google"]["go"]["architecture"]   = kernel["machine"] =~ /x86_64/ ? "amd64" : "386"
default["google"]["go"]["version"]        = "1.1.2"
default["google"]["go"]["source"]         = "https://go.googlecode.com/files/go#{node["google"]["go"]["version"]}.linux-#{node["google"]["go"]["architecture"]}.tar.gz"
default["google"]["go"]["checksum"]       = node["google"]["go"]["architecture"] =~ /amd64/ ? "42634e25f98a5db1e8a2a8270c3604fcf8fed38d" : "42334112e5ba7fcd1a58de0a85ff2668e446cd0b"
default["google"]["go"]["install_dir"]    = "/usr/local"