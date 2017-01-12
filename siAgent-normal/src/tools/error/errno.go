package error

type Error struct {
	LogMsg string
	ErrMsg string
}

const (
	EOK int32 = 0

	EMYSQL        int32 = 1000
	EMYSQL_QUERY  int32 = 1001
	EMYSQL_INSERT int32 = 1002
	EMYSQL_DELETE int32 = 1003

	EREDIS     int32 = 1100
	EREDIS_GET int32 = 1101

	EKUBE                           int32 = 1200
	EKUBE_CLIENT                    int32 = 1201
	EKUBE_CREATE_DEPLOYMENT               = 1202
	EKUBE_LIST_PODS                       = 1203
	EKUBE_CREATE_SERVICE                  = 1206
	EKUBE_LIST_ENDPOINTS                  = 1207
	EKUBE_CREATE_NAMESPACE                = 1204
	EKUBE_CREATE_RESOURCEQUOTA            = 1205
	EKUBE_LIST_SERVICE                    = 1208
	EKUBE_CREATE_ENDPOINTS                = 1209
	EKUBE_LIST_DEPLOYMENTS                = 1210
	EKUBE_ROLLING_DEPLOYMENTS             = 1211
	EKUBE_LABEL_SELECTOR                  = 1212
	EKUBE_GET_DEPLOYMENT                  = 1213
	EKUBE_ROLLBACK_DEPLOYMENT             = 1214
	EKUBE_DELETE_SERVICE                  = 1215
	EKUBE_DELETE_ENDPOINT                 = 1216
	EKUBE_LIST_REPLICASET                 = 1217
	EKUBE_DELETE_REPLICASET               = 1218
	EKUBE_DELETE_DEPLOYMENT               = 1219
	EKUBE_DELETE_POD                      = 1220
	EKUBE_GET_RS_BY_DEPLOYMENT            = 1221
	EKUBE_GET_PODS_BY_RS                  = 1222
	EKUBE_GET_NODE_BY_POD                 = 1223
	EKUBE_GET_SERVICES_BY_NAMESPACE       = 1224
	EKUBE_GET_PODS_BY_SERVICE             = 1225
	EKUBE_LOGS_POD                        = 1226
	EKUBE_SCALE_DEPLOYMENT                = 1227
	EKUBE_FIND_NEW_REPLICASET             = 1228
	EKUBE_DELETE_NAMESPACE                = 1229
	EKUBE_DELETE_RESOURCEQUOTA            = 1230
	EKUBE_UPDATE_RESOURCEQUOTA            = 1231
	EKUBE_GET_ALL_NAMESPACES              = 1232
	EKUBE_UPDATE_SERVICE                  = 1233
	EKUBE_GET_SERVICE_BY_POD              = 1234
	EKUBE_GET_SERVICE_BY_DEPLOYMENT       = 1235
	EKUBE_GET_NODE                        = 1236

	EIRIS int32 = 1300

	EYCE                   int32 = 1400
	EYCE_LOGIN             int32 = 1401
	EYCE_SESSION           int32 = 1402
	EYCE_SESSION_DEL       int32 = 1403
	EYCE_ORG_EXIST         int32 = 1404
	EYCE_NODEPORT_EXIST    int32 = 1405
	EYCE_ORGTODC           int32 = 1406
	EYCE_LIST_EXTENSIONS   int32 = 1407
	EYCE_DELETE_NODEPORT   int32 = 1408
	EYCE_DELETE_DEPLOYMENT int32 = 1409
	EYCE_LOGS_POD          int32 = 1410
	EYCE_LOGOUT            int32 = 1411
	EYCE_DELETE_SERVICE    int32 = 1412
	EYCE_DELETE_ENDPOINTS  int32 = 1413
	EYCE_NOTFOUND          int32 = 1414
	EYCE_EXISTED_NAME      int32 = 1415
	EYCE_INVALID_PORT      int32 = 1416
	EYCE_WRONG_PASSWORD    int32 = 1417
	EYCE_RESPAWN_POD       int32 = 1418

	EREGISTRY     int32 = 1500
	EREGISTRY_GET int32 = 1501

	EJSON          int32 = 1600
	EARGS          int32 = 1601
	EOOM           int32 = 1602
	EOOR           int32 = 1603
	ENULL          int32 = 1604
	EINVALID_PARAM int32 = 1605

	EGDK_IOUTIL int32 = 1700
)

var Errors = map[int32]*Error{

	EOK: &Error{
		LogMsg: "OK",
		ErrMsg: "操作成功",
	},

	// 1000~1099 MySQL错误
	EMYSQL: &Error{
		LogMsg: "MySQL Error",
		ErrMsg: "MySQL数据库错误",
	},
	EMYSQL_QUERY: &Error{
		LogMsg: "MySQL Query Error",
		ErrMsg: "MySQL查询出错",
	},
	EMYSQL_INSERT: &Error{
		LogMsg: "MySQL Insert Error",
		ErrMsg: "MySQL插入出错",
	},
	EMYSQL_DELETE: &Error{
		LogMsg: "MySQL Delete Error",
		ErrMsg: "MySQL删除出错",
	},

	// 1100~1199 Redis错误
	EREDIS: &Error{
		LogMsg: "Redis Error",
		ErrMsg: "Redis数据库错误",
	},

	// 1200~1299 K8s错误
	EKUBE: &Error{
		LogMsg: "Kubernetes Error",
		ErrMsg: "Kubernetes错误",
	},
	EKUBE_CLIENT: &Error{
		LogMsg: "Kubernetes Create Client Error",
		ErrMsg: "创建K8s客户端出错",
	},
	EKUBE_CREATE_DEPLOYMENT: &Error{
		LogMsg: "Kubernetes Create Deployment Error",
		ErrMsg: "创建Deployment出错",
	},
	EKUBE_LIST_PODS: &Error{
		LogMsg: "Kubernetes List Pods Error",
		ErrMsg: "获取Pod列表出错",
	},
	EKUBE_CREATE_SERVICE: &Error{
		LogMsg: "Kubernetes Create Service Error",
		ErrMsg: "创建Service出错",
	},
	EKUBE_LIST_ENDPOINTS: &Error{
		LogMsg: "Kubernetes List Endpoints Error",
		ErrMsg: "获取Endpoints出错",
	},
	EKUBE_CREATE_NAMESPACE: &Error{
		LogMsg: "Kubernetes create namespace error",
		ErrMsg: "创建命名空间失败",
	},
	EKUBE_CREATE_RESOURCEQUOTA: &Error{
		LogMsg: "Kubernetes create resourceQuota error",
		ErrMsg: "创建资源配额失败",
	},
	EKUBE_LIST_SERVICE: &Error{
		LogMsg: "Kubernetes list Service Error",
		ErrMsg: "获取Service出错",
	},
	EKUBE_CREATE_ENDPOINTS: &Error{
		LogMsg: "Kubernetes create Endpoints Error",
		ErrMsg: "创建Endpoints出错",
	},
	EKUBE_LABEL_SELECTOR: &Error{
		LogMsg: "Kubernetes LabelSelectorAsSelector Error",
		ErrMsg: "通过标签选择出错",
	},
	EKUBE_GET_DEPLOYMENT: &Error{
		LogMsg: "Kubernetes Get Deployment by name Error",
		ErrMsg: "获取Deployment失败",
	},

	EKUBE_LIST_DEPLOYMENTS: &Error{
		LogMsg: "Kubernetes list Deployments Error",
		ErrMsg: "获取Deployments出错",
	},
	EKUBE_ROLLING_DEPLOYMENTS: &Error{
		LogMsg: "Kubernetes RollingUpdate Deployments Error",
		ErrMsg: "滚动升级Deployments出错",
	},
	EKUBE_ROLLBACK_DEPLOYMENT: &Error{
		LogMsg: "Kubernetes Rollback Deployment Error",
		ErrMsg: "回滚失败",
	},
	EKUBE_DELETE_SERVICE: &Error{
		LogMsg: "Kubernetes Delete Service Error",
		ErrMsg: "删除服务失败",
	},
	EKUBE_DELETE_ENDPOINT: &Error{
		LogMsg: "Kubernetes Delete Endpoint Error",
		ErrMsg: "删除访问点失败",
	},
	EKUBE_LIST_REPLICASET: &Error{
		LogMsg: "Kubernetes List ReplicaSet Error",
		ErrMsg: "获取ReplicaSet列表失败",
	},
	EKUBE_DELETE_REPLICASET: &Error{
		LogMsg: "Kubernete Delete ReplicaSet Error",
		ErrMsg: "删除ReplicaSet失败",
	},
	EKUBE_DELETE_DEPLOYMENT: &Error{
		LogMsg: "Kubernetes Delete Deployment Error",
		ErrMsg: "删除应用失败",
	},
	EKUBE_DELETE_POD: &Error{
		LogMsg: "Kubernetes Delete Deployment Error",
		ErrMsg: "删除实例失败",
	},
	EKUBE_GET_RS_BY_DEPLOYMENT: &Error{
		LogMsg: "Kuberetes get Replicaset by Deployment Error",
		ErrMsg: "通过Deployment获取Rs失败",
	},
	EKUBE_GET_PODS_BY_RS: &Error{
		LogMsg: "Kuberentes Get Pods by Replicaset Error",
		ErrMsg: "通过Rs获取Pods失败",
	},
	EKUBE_GET_NODE_BY_POD: &Error{
		LogMsg: "Kubernetes Get Node by Pod Error",
		ErrMsg: "通过Pod获取Node失败",
	},
	EKUBE_GET_SERVICES_BY_NAMESPACE: &Error{
		LogMsg: "Kubernetes Get Services by Namespaces Error",
		ErrMsg: "通过命名空间获取服务列表失败",
	},
	EKUBE_GET_PODS_BY_SERVICE: &Error{
		LogMsg: "Kubernetes Get Pods by Service Error",
		ErrMsg: "通过服务列表获取pods失败",
	},
	EKUBE_LOGS_POD: &Error{
		LogMsg: "Kubernetes Logs Pod Error",
		ErrMsg: "获取Pod日志失败",
	},
	EKUBE_SCALE_DEPLOYMENT: &Error{
		LogMsg: "Kuberentes Scale Deployment Error",
		ErrMsg: "扩容Deployment失败",
	},
	EKUBE_FIND_NEW_REPLICASET: &Error{
		LogMsg: "Kubernetes Find New ReplicaSet Error",
		ErrMsg: "查找最新的ReplicaSet失败",
	},
	EKUBE_GET_ALL_NAMESPACES: &Error{
		LogMsg: "Kubernetes Get all namespaces Error",
		ErrMsg: "获取全部namespace失败",
	},
	EKUBE_DELETE_NAMESPACE: &Error{
		LogMsg: "Kubernetes Delete Namespace Error",
		ErrMsg: "删除组织失败",
	},
	EKUBE_DELETE_RESOURCEQUOTA: &Error{
		LogMsg: "Kubernetes Delete ResourceQuota Error",
		ErrMsg: "删除资源配额失败",
	},
	EKUBE_UPDATE_RESOURCEQUOTA: &Error{
		LogMsg: "Kubernetes Update ResourceQuota Error",
		ErrMsg: "更新资源配额失败",
	},
	EKUBE_UPDATE_SERVICE: &Error{
		LogMsg: "Kubernetes Update Service Error",
		ErrMsg: "更新服务失败",
	},
	EKUBE_GET_SERVICE_BY_POD: &Error{
		LogMsg: "Kubernetes Get Service By Pod Error",
		ErrMsg: "根据Pod获取Service失败",
	},
	EKUBE_GET_SERVICE_BY_DEPLOYMENT: &Error{
		LogMsg: "Kubernetes Get Service By Deployment Error",
		ErrMsg: "根据Deployment获取Service失败",
	},
	EKUBE_GET_NODE: &Error{
		LogMsg: "Kubernetes Get Nodes Error",
		ErrMsg: "获取计算节点信息失败",
	},

	// 1300~1399 Iris错误
	EIRIS: &Error{
		LogMsg: "Iris Error",
		ErrMsg: "Iris服务器错误",
	},

	// 1400~1499 YCE错误
	EYCE: &Error{
		LogMsg: "YCE Internal Error",
		ErrMsg: "YCE内部错误",
	},

	EYCE_LOGIN: &Error{
		LogMsg: "Can't Find the User",
		ErrMsg: "用户名密码错误",
	},

	EYCE_SESSION: &Error{
		LogMsg: "Can't Find the Session",
		ErrMsg: "请重新登录",
	},

	EYCE_SESSION_DEL: &Error{
		LogMsg: "Delete Session Error",
		ErrMsg: "退出遇到问题",
	},
	EYCE_ORG_EXIST: &Error{
		LogMsg: "The organization exists",
		ErrMsg: "组织已经存在",
	},

	EYCE_NODEPORT_EXIST: &Error{
		LogMsg: "The NodePort exists",
		ErrMsg: "NodePort已存在",
	},

	EYCE_ORGTODC: &Error{
		LogMsg: "Get Datacenter by orgId Error",
		ErrMsg: "获取数据中心错误",
	},

	EYCE_LIST_EXTENSIONS: &Error{
		LogMsg: "Get Service and Endpoint list error",
		ErrMsg: "获取服务和访问点列表失败",
	},
	EYCE_DELETE_NODEPORT: &Error{
		LogMsg: "Delete NodePort Error",
		ErrMsg: "删除NodePort失败",
	},

	EYCE_DELETE_DEPLOYMENT: &Error{
		LogMsg: "Delete Deployment Error",
		ErrMsg: "删除应用失败",
	},
	EYCE_LOGS_POD: &Error{
		LogMsg: "Get Pod Logs Error",
		ErrMsg: "获取实例日志失败",
	},
	EYCE_LOGOUT: &Error{
		LogMsg: "Get Logout Params Error",
		ErrMsg: "用户退出参数错误",
	},
	EYCE_DELETE_SERVICE: &Error{
		LogMsg: "Delete Service Error",
		ErrMsg: "删除服务失败",
	},
	EYCE_DELETE_ENDPOINTS: &Error{
		LogMsg: "Delete Endpoints Error",
		ErrMsg: "删除访问点失败",
	},
	EYCE_NOTFOUND: &Error{
		LogMsg: "Not Found",
		ErrMsg: "资源未找到",
	},
	EYCE_EXISTED_NAME: &Error{
		LogMsg: "Duplicated Resources Name",
		ErrMsg: "资源名重复",
	},
	EYCE_INVALID_PORT: &Error{
		LogMsg: "Invalid Port",
		ErrMsg: "端口号无效",
	},
	EYCE_WRONG_PASSWORD: &Error{
		LogMsg: "Wrong Password",
		ErrMsg: "密码错误",
	},
	EYCE_RESPAWN_POD: &Error{
		LogMsg: "Respawn Error",
		ErrMsg: "重启应用实例失败",
	},

	// 1500~1599 Registr错误
	EREGISTRY: &Error{
		LogMsg: "Registry is empty, no docker image",
		ErrMsg: "私有镜像仓库为空",
	},
	EREGISTRY_GET: &Error{
		LogMsg: "Can't Get value from Docker Registry",
		ErrMsg: "不能检索镜像仓库",
	},

	// 1600 Json错误
	EJSON: &Error{
		LogMsg: "Json Marshal/Unmarshal Error",
		ErrMsg: "Json序列化错误",
	},

	// 1601 参数错误
	EARGS: &Error{
		LogMsg: "Invalid arguments",
		ErrMsg: "参数错误",
	},

	// 1602 内存溢出
	EOOM: &Error{
		LogMsg: "Out of memory",
		ErrMsg: "内存溢出",
	},
	EOOR: &Error{
		LogMsg: "Index out of Range",
		ErrMsg: "数组越界",
	},
	ENULL: &Error{
		LogMsg: "Null Pointer Error",
		ErrMsg: "空指针",
	},
	EINVALID_PARAM: &Error{
		LogMsg: "Invalid Param",
		ErrMsg: "无效值",
	},

	EGDK_IOUTIL: &Error{
		LogMsg: "ioutil ReadAll error",
		ErrMsg: "读取HTTP响应体失败",
	},
}
