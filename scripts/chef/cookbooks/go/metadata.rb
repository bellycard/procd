maintainer       "Belly, Inc."
maintainer_email "sysops@bellycard.com"
license          "Apache v2.0"
description      "Installs/Configures Google Go"
long_description IO.read(File.join(File.dirname(__FILE__), 'README.md'))
version          "1.0.0"

# Operating systems supported
%w{ debian ubuntu }.each do |os|
  supports os
end
