# Author:: Christian Vozar <christian@bellycard.com>
# Cookbook Name:: procd
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

# Installation Attributes
default['belly']['procd']['install_directory']    = '/opt/procd'

# Global Attributes
default['belly']['procd']['ticker_interval']      = 5

# HTTP Interface Attributes
default['belly']['procd']['http']['enabled']      = true
default['belly']['procd']['http']['bind_address'] = '0.0.0.0'
default['belly']['procd']['http']['bind_port']    = 5596

# Heka Client Attributes
default['belly']['procd']['heka']['enabled']      = false
default['belly']['procd']['heka']['server']       = '127.0.0.1'
default['belly']['procd']['heka']['port']         = 5565
default['belly']['procd']['heka']['encoder']      = 'json'
default['belly']['procd']['heka']['sender']       = 'tcp'
default['belly']['procd']['heka']['payload']      = false
