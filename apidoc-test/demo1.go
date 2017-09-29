package main
import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
) 

/**
 * @api get /users 获取所有的用户信息
 * @apiGroup users
 * @apiQuery page int 显示第几页的内容
 * @apiQuery size int 每页显示的数量
 *
 * @apiSuccess 200 ok
 * @apiParam count int 符合条件的所有用户数量
 * @apiParam users array 用户列表。
 * @apiExample json
 * {
 *     "count": 500,
 *     "users": [
 *         {"id":1, "username": "admin1", "name": "管理员2"},
 *         {"id":2, "username": "admin2", "name": "管理员2"}
 *     ],
 * }
 * @apiExample xml
 * <users count="500">
 *     <user id="1" username="admin1" name="管理员1" />
 *     <user id="2" username="admin2" name="管理员2" />
 * </users>
 */
func HandleHi(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Hi")
}

func main() {
	router := httprouter.New()
	router.ServeFiles("/doc/*filepath", http.Dir("./doc/"))
	router.GET("/hi", HandleHi)
	http.ListenAndServe(":8080", router)
}
