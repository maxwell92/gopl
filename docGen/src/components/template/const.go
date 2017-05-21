package template

const (
	TIMEFORMAT = `2006-01-02 03:04:05`

	WORD = ``
	MARKDOWN = `<img src="http://kubernetes.io/kubernetes/img/warning.png" alt="WARNING" width="25" height="25">

#### 修改请谨慎

<Topic>
==============

作者: [<AuthorName>](<AuthorGithubLink>)

最后修订: <LastRevised>
修订说明: <RevisedDescription>

目录
--------------

### 目的

<Purpose>

### 请求

请求方法: <RequestMethod>
请求URL: <RequestUrl>
请求头: <RequestHeader>
请求参数: <RequestParameters>

### 页面设计

<TableDescription>
<TableDesign>


### 程序实现逻辑

<KeyLogic>

### 响应数据结构:

<Response>


### 备注

<Comments>
`
)
