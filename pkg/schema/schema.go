package schema

type CatHealthForm struct {
	Cluster             string `json:"cluster"`
	Status              string `json:"status"`
	NodeTotal           string `json:"node.total"`
	NodeData            string `json:"node.data"`
	Shards              string `json:"shards"`
	PriShards           string `json:"pri"`
	ReloShards          string `json:"relo"`
	InitShards          string `json:"init"`
	UnassignedShards    string `json:"unassign"`
	PendingTasks        string `json:"pending_tasks"`
	ActiveShardsPercent string `json:"active_shards_percent"`
}

type CatIndexForm struct {
	Health        string `json:"health"`
	Status        string `json:"status"`
	Index         string `json:"Index"`
	UUID          string `json:"uuid"`
	PrimaryShards string `json:"pri"`
	ReplicaShards string `json:"rep"`
	DocsCount     string `json:"docs.count"`
	DocsDeleted   string `json:"docs.deleted"`
	StoreSize     string `json:"store.size"`
	PriStoreSize  string `json:"pri.store.size"`
}

type CatMasterForm struct {
	ID        string `json:"id"`
	Host        string `json:"host"`
	Ip         string `json:"index"`
	Node          string `json:"node"`
}

type Node struct {
	IP              string `json:"ip"`
	NodeRole        string `json:"role"`
	Name            string `json:"name"`
	DiskUsedPercent string `json:"disk.used_percent"`
	Load1M          string `json:"load_1m"`
	Load5M          string `json:"load_5m"`
	Load15M         string `json:"load_15m"`
	Uptime          string `json:"uptime"`
}

type Shard struct {
	Index            string `json:"index"`
	Shard            string `json:"shard"`
	PriRep           string `json:"prirep"`
	State            string `json:"state"`
	Docs             string `json:"docs"`
	Store            string `json:"store"`
	IP               string `json:"ip"`
	Node             string `json:"node"`
	UnassignedReason string `json:"unassigned.reason"`
}

type Repository struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

