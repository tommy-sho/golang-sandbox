package main

type Message struct {
	InsertId         string      `json:"insertId"`
	JsonPayload      JsonPayload `json:"jsonPayload"`
	ReceiveTimestamp string      `json:"receiveTimestamp"`
	Resource         Resource    `json:"resource"`
	Severity         string      `json:"severity"`
}

type JsonPayload struct {
	Msg string `json:"msg"`
	Ts  int64  `json:"ts"`
}

type Resource struct {
	Labels Labels `json:"labels"`
}

type Labels struct {
	ClusterName   string `json:"cluster_name"`
	ContainerName string `json:"container_name"`
	PodID         string `json:"pod_id"`
}
