# softlayer_instance_plugin

softlayer_instance_plugin is an instance plugin for infrakit (see: github.com/docker/infrakit)

The finished plugin should manage Softlayer instances (virtual servers)
and uses the Go Client for the Softlayer API (see: https://github.com/softlayer/softlayer-go)

Note: The plugin is not finished yet!

# Prerequisites

You have to have infrakit installed
See github.com/docker/infrakit on how to install infrakit

You have to have a Softlayer (Bluemix) Account with rights to retrieve info/order/cancel virtual servers
Use your environment variables for API authentication
SL_USERNAME
SL_API_KEY

# Build

go build -o ./build/instance-sl ./plugin/

# Execute


cd
mkdir -p tutorial
./PATH/TO/INFRAKIT_DIR/build/infrakit-group-default
./PATH/TO/INFRAKIT_DIR/build/infrakit-flavor-vanilla
./PATH/TO/SL_PLUGIN_DIR/build/instance-sl--dir ./tutorial
infrakit plugin ls


./PATH/TO/INFRAKIT_DIR/build/infrakit instance describe
./PATH/TO/INFRAKIT_DIR/build/infrakit group commit PATH/TO/SL_PLUGIN_DIR/cattle.json
./PATH/TO/INFRAKIT_DIR/build/infrakit group describe PATH/TO/SL_PLUGIN_DIR/cattle.json
./PATH/TO/INFRAKIT_DIR/build/infrakit instance describe


./PATH/TO/INFRAKIT_DIR/build/infrakit group destroy cattle


# Example

Start the plugin
./go/src/github.com/HerrHuber/softlayer_instance_plugin/build/instance-sl --log 5

Create a new virtual server instance
./go/src/github.com/docker/infrakit/build/infrakit instance --name instance-sl provision go/src/github.com/HerrHuber/softlayer_instance_plugin/config/instance_sl_remote_script.json --log 5

Describe (get Id, LogicalID) all virtual server instances
./go/src/github.com/docker/infrakit/build/infrakit instance describe --name instance-sl

Destroy the virtual server instance with id xxxxxxxxx
./go/src/github.com/docker/infrakit/build/infrakit instance --name instance-sl destroy xxxxxxxxx --log 5


