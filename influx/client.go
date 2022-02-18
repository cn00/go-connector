package influx

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type Connector struct {
	proxy influxdb2.Client
}

func New(host, token string) (connector *Connector, err error) {

	var c influxdb2.Client
	c = influxdb2.NewClient(host, token)
	if err != nil {
		return connector, err
	}
	connector = &Connector{
		proxy: c,
	}
	return connector, err
}

func (connector Connector) Save(database string, measurement string, tags map[string]string, fields map[string]interface{}) (err error) {
	// get non-blocking write influxdb2
	writeAPI := connector.proxy.WriteAPI("tp", database)
	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()

	return err
}

func (connector Connector) Close() {
	if connector.proxy != nil {
		connector.proxy.Close()
	}
}
