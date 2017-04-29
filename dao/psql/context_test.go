package psql

var db = NewConnect(
	"192.168.1.2",
	"5432",
	"postgres",
	"postgres",
	"TruckMonitor",
	"disable",
)
