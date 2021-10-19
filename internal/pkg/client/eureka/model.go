package eureka

type GetAppsResp struct {
	Apps Apps `json:"applications"`
}

type GetAppResp struct {
	App App `json:"application"`
}

type Apps struct {
	App []App `json:"application"`
}

type App struct {
	Name     string     `json:"name"`
	Instance []Instance `json:"instance"`
}

type Instance struct {
	InstanceId string `json:"instanceId"`
	HostName   string `json:"hostName"`
	Host       string `json:"host"`
	IpAddr     string `json:"ipAddr"`
	Status     string `json:"status"`
	HomePageUrl string `json:"homePageUrl"`
	// Port       string `json:"port"`
}
