package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mq-demo/workers"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	SERVER_IP      = "62.84.125.40"
	RABBITMQ_PORT  = 5672
	RABBITMQ_USER  = "guest"
	RABBITMQ_PASS  = "guest"
	DEMO_POOL_NAME = "demo_pool"
)

func getConfig() []byte {
	endpoint := fmt.Sprintf("http://%s/api/brokers/pool/config?pool=%s", SERVER_IP, DEMO_POOL_NAME)

	response, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	out, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return out
}

type ConfigPart map[string]interface{}

func parseConfig(cfg []byte) (string, map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	var parsed_config map[string]interface{}
	if err := json.Unmarshal(cfg, &parsed_config); err != nil {
		panic(err)
	}

	queue := parsed_config["queue"].(string)
	qcfg := parsed_config["qcfg"].(map[string]interface{})
	rcfg := parsed_config["rcfg"].(map[string]interface{})
	wcfg := parsed_config["wcfg"].(map[string]interface{})

	return queue, qcfg, rcfg, wcfg
}

func readArgs(args interface{}) map[string]interface{} {
	if args == nil {
		return nil
	}
	return args.(map[string]interface{})
}

func main() {
	task_ch := workers.CreateWorkerPool(10)

	fmt.Println("requesting config...")
	cfg := getConfig()
	fmt.Printf("config: %s\n", cfg)
	fmt.Println("done")

	fmt.Println()

	fmt.Println("parsing config...")
	queue, _, rcfg, _ := parseConfig(cfg)
	fmt.Println("done")

	fmt.Println()

	fmt.Println("establishing connection to rabbitmq...")
	conn_str := fmt.Sprintf("amqp://%s:%s@%s:%d/", RABBITMQ_USER, RABBITMQ_PASS, SERVER_IP, RABBITMQ_PORT)

	fmt.Println("using address: ", conn_str)

	conn, err := amqp.Dial(conn_str)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		queue,
		rcfg["consumer"].(string),
		rcfg["auto_ack"].(bool),
		rcfg["exclusive"].(bool),
		rcfg["no_local"].(bool),
		rcfg["no_wait"].(bool),
		readArgs(rcfg["args"]),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("connection established")

	fmt.Println()

	fmt.Println("reading tasks...")
	for task := range msgs {
		task_ch <- string(task.Body)
	}

}
