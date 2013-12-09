maintainer       'Belly, Inc.'
maintainer_email 'sysops@bellycard.com'
license          'Apache v2.0'
description      'Installs & Configures Belly Procd'
long_description IO.read(File.join(File.dirname(__FILE__), 'README.md'))
version          '1.0.0'

# Cookbook dependancies
%w{ go }.each do |cookbooks|
  depends cookbooks
end
