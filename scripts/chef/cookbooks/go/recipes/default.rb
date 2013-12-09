# Author:: Christian Vozar <christian@bellycard.com>
# Cookbook Name:: go
# Recipe:: default
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

if node["google"]["go"]["install_method"] == "package"

  case node["platform_family"]
  when "debian"
    package "golang"
  else
    package "golang"
  end

else

  remote_file "#{Chef::Config[:file_cache_path]}/go#{node["google"]["go"]["version"]}.linux-#{node["google"]["go"]["architecture"]}.tar.gz" do
    checksum node["google"]["go"]["checksum"]
    source node["google"]["go"]["source"]
    mode "0644"
  end

  execute "Install Google Go v#{node["google"]["go"]["version"]}" do
    cwd Chef::Config[:file_cache_path]
    command "tar -C #{node["google"]["go"]["install_dir"]} -xzf go#{node["google"]["go"]["version"]}.linux-#{node["google"]["go"]["architecture"]}.tar.gz"
    not_if "go version |grep #{node["google"]["go"]["version"]}"
  end
  
  if node["platform_family"] == "debian"
    # Update PATHadd /usr/local/go/bin to PATH in /etc/environment
    execute "Verify PATH" do
      command "grep -q PATH=$PATH:/usr/local/go/bin /etc/environment || echo PATH=$PATH:/usr/local/go/bin >> /etc/environment"
    end
  end
end