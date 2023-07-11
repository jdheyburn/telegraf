//go:build !custom || outputs || outputs.datadog_statsd

package all

import _ "github.com/influxdata/telegraf/plugins/outputs/datadog_statsd" // register plugin
