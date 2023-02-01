/**
** @创建时间: 2021/4/11 12:57 上午
** @作者　　: return
** @描述　　:
 */
package base

import (
	"github.com/gincmf/xpyunSdk"
	"github.com/gincmf/xpyunSdk/model"
	"github.com/gincmf/xpyunSdk/service"
	"github.com/gincmf/xpyunSdk/util"
)

type Printer struct{}

/**
 * @Author return <1140444693@qq.com>
 * @Description 新增打印机
 * @Date 2021/4/11 1:24:42
 * @Param
 * @return
 **/
func (rest *Printer) Add(snList []*model.AddPrinterRequestItem) *model.XPYunResp {

	options := xpyunSdk.Options()

	request := model.AddPrinterRequest{}

	request.User = options.User
	request.UserKey = options.Ukey
	request.Timestamp = util.GetMillisecond()

	request.GenerateSign()
	request.Items = snList

	result := service.XpYunAddPrinters(&request)

	return result
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 删除打印机
 * @Date 2021/4/11 1:24:35
 * @Param
 * @return
 **/

func (rest *Printer) Delete(snList string) *model.XPYunResp {

	options := xpyunSdk.Options()

	request := model.DelPrinterRequest{}

	request.User = options.User
	request.UserKey = options.Ukey
	request.Timestamp = util.GetMillisecond()

	request.GenerateSign()

	request.SnList = []string{snList}

	result := service.XpYunDelPrinters(&request)

	return result
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 打印订单
 * @Date 2021/4/11 1:30:3
 * @Param
 * @return
 **/

func (rest *Printer) Printer(sn string, content string, times int) *model.XPYunResp {

	options := xpyunSdk.Options()

	request := model.PrintRequest{}
	request.User = options.User
	request.UserKey = options.Ukey

	//*必填*：打印机编号
	request.Sn = sn
	request.GenerateSign()

	//*必填*：打印内容,不能超过12K
	request.Content = content

	//打印份数，默认为1
	request.Copies = times

	//声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，默认为 2 来单播放模式
	request.Voice = 2

	//打印模式：
	//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
	//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
	request.Mode = 1

	result := service.XpYunPrint(&request)

	return result

}
