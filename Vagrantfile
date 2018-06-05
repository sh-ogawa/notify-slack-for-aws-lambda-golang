# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|

  config.vm.box = "hbsmith/awslinux"  
  config.vm.synced_folder ".", "/home/vagrant/notfy-slack", owner: "vagrant", group: "vagrant"
  
  config.vm.define :master do |node|
    node.vm.network :private_network, ip:"192.168.172.11"
    node.vm.hostname = "aws-lambda-local"
  end

  config.vm.provider "virtualbox" do |vb|
    vb.cpus = 1
    vb.memory = 1024
  end

end
