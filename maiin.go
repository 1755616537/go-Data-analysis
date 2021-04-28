package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/encoding/gjson"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("启动中")
	var _time=time.Now()

	fmt.Println(Run("D:\\1755616537\\OneDrive\\桌面/新建文件夹"))

	fmt.Println("耗时",time.Since(_time) / time.Millisecond)
	fmt.Println("完成")
}

//启动
func Run(path string) error {
	//扫描目录->获取文件名
	fmt.Println("扫描文件中")
	var fileList = make(map[string]string)
	err := filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				//fmt.Println("dir:", path)
				return nil
			}
			//fmt.Println("file:", path)

			//只获取json类型文件
			namearr:=strings.Split(f.Name(),".")
			if len(namearr)==2 {
				if namearr[1]=="json" {
					fileList[path] = f.Name()
				}
			}else if len(namearr)>2 {
				if namearr[len(namearr)-1]=="json" {
					fileList[path] = f.Name()
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	if fileList == nil {
		return errors.New("该目录底下扫描文件数量为0")
	}
	if len(fileList) == 0 {
		return errors.New("该目录底下扫描文件数量为0")
	}

	//批量更名
	//for i, _ := range fileList {
	//	_ = os.Rename(i, strings.Replace(i, ".txt", ".json", 1))
	//}
	//return nil

	fmt.Println("读取文件内容中")
	//解析每一个文件的数据
	type abstracts struct {
		Text            string `json:"text"`
		BackgroundColor string `json:"backgroundColor"`
		BorderColor     string `json:"borderColor"`
		FontColor       string `json:"fontColor"`
		Icon            string `json:"icon"`
		Message         string `json:"message"`
	}
	type searchResult struct {
		Id                 int           `json:"id"`
		Template           string        `json:"template"`
		ImageUrl           string        `json:"imageUrl"`
		Title              string        `json:"title"`
		Address            string        `json:"address"`
		Lowestprice        int           `json:"lowestprice"`
		Avgprice           int           `json:"avgprice"`
		Latitude           string           `json:"latitude"`
		Longitude          string           `json:"longitude"`
		ShowType           string        `json:"showType"`
		Avgscore           int           `json:"avgscore"`
		Comments           int           `json:"comments"`
		HistoryCouponCount int           `json:"historyCouponCount"`
		BackCateName       string        `json:"backCateName"`
		Areaname           string        `json:"areaname"`
		Tag                []interface{} `json:"tag"`
		Cate               []int         `json:"cate"`
		RecentScreen       string        `json:"recentScreen"`
		Abstracts          []abstracts   `json:"abstracts"`
		DangleAbstracts    string        `json:"dangleAbstracts"`
		TitleTags          []interface{} `json:"titleTags"`
		IUrl               string        `json:"iUrl"`
		Deals              string        `json:"deals"`
		Posdescr           string        `json:"posdescr"`
		Ct_poi             string        `json:"ct_poi"`
		Trace              string        `json:"trace"`
		LandmarkDistance   string        `json:"landmarkDistance"`
		HasAds             string        `json:"hasAds"`
		AdsClickUrl        string        `json:"adsClickUrl"`
		AdsShowUrl         string        `json:"adsShowUrl"`
		Distance           string        `json:"distance"`
		CityId             int           `json:"cityId"`
		City               string        `json:"city"`
		Phone              string        `json:"phone"`
		Full               bool          `json:"full"`
	}
	var searchResultList []searchResult
	for i, _ := range fileList {
		//读取文件内容
		_byte, err := ioutil.ReadFile(i)
		if err != nil {
			fmt.Println(i, "读取文件失败")
			continue
		}
		//转换json数据类型
		_json := gjson.New(string(_byte))
		o := "data.searchResult"
		for i := 0; i < _json.Len(o); i++ {
			_searchResult := searchResult{
				_json.GetInt(fmt.Sprint(o, ".", i, ".id")),
				_json.GetString(fmt.Sprint(o, ".", i, ".template")),
				_json.GetString(fmt.Sprint(o, ".", i, ".imageUrl")),
				_json.GetString(fmt.Sprint(o, ".", i, ".title")),
				_json.GetString(fmt.Sprint(o, ".", i, ".address")),
				_json.GetInt(fmt.Sprint(o, ".", i, ".lowestprice")),
				_json.GetInt(fmt.Sprint(o, ".", i, ".avgprice")),
				_json.GetString(fmt.Sprint(o, ".", i, ".latitude")),
				_json.GetString(fmt.Sprint(o, ".", i, ".longitude")),
				_json.GetString(fmt.Sprint(o, ".", i, ".showType")),
				_json.GetInt(fmt.Sprint(o, ".", i, ".avgscore")),
				_json.GetInt(fmt.Sprint(o, ".", i, ".comments")),
				_json.GetInt(fmt.Sprint(o, ".", i, ".historyCouponCount")),
				_json.GetString(fmt.Sprint(o, ".", i, ".backCateName")),
				_json.GetString(fmt.Sprint(o, ".", i, ".areaname")),
				_json.GetArray(fmt.Sprint(o, ".", i, ".tag")),
				_json.GetInts(fmt.Sprint(o, ".", i, ".cate")),
				_json.GetString(fmt.Sprint(o, ".", i, ".recentScreen")),
				nil,
				_json.GetString(fmt.Sprint(o, ".", i, ".dangleAbstracts")),
				_json.GetArray(fmt.Sprint(o, ".", i, ".titleTags")),
				_json.GetString(fmt.Sprint(o, ".", i, ".iUrl")),
				_json.GetString(fmt.Sprint(o, ".", i, ".deals")),
				_json.GetString(fmt.Sprint(o, ".", i, ".posdescr")),
				_json.GetString(fmt.Sprint(o, ".", i, ".ct_poi")),
				_json.GetString(fmt.Sprint(o, ".", i, ".trace")),
				_json.GetString(fmt.Sprint(o, ".", i, ".landmarkDistance")),
				_json.GetString(fmt.Sprint(o, ".", i, ".hasAds")),
				_json.GetString(fmt.Sprint(o, ".", i, ".adsClickUrl")),
				_json.GetString(fmt.Sprint(o, ".", i, ".adsShowUrl")),
				_json.GetString(fmt.Sprint(o, ".", i, ".distance")),
				_json.GetInt(fmt.Sprint(o, ".", i, ".cityId")),
				_json.GetString(fmt.Sprint(o, ".", i, ".city")),
				_json.GetString(fmt.Sprint(o, ".", i, ".phone")),
				_json.GetBool(fmt.Sprint(o, ".", i, ".full")),
			}

			var _abstracts []abstracts
			for i2 := 0; i2 < _json.Len(fmt.Sprint(o, ".", i, ".abstracts")); i2++ {
				_abstracts = append(_abstracts, abstracts{
					_json.GetString(fmt.Sprint(o, ".", i, ".abstracts", ".", i2, ".text")),
					_json.GetString(fmt.Sprint(o, ".", i, ".abstracts", ".", i2, ".backgroundColor")),
					_json.GetString(fmt.Sprint(o, ".", i, ".abstracts", ".", i2, ".borderColor")),
					_json.GetString(fmt.Sprint(o, ".", i, ".abstracts", ".", i2, ".fontColor")),
					_json.GetString(fmt.Sprint(o, ".", i, ".abstracts", ".", i2, ".icon")),
					_json.GetString(fmt.Sprint(o, ".", i, ".abstracts", ".", i2, ".message")),
				})
			}
			_searchResult.Abstracts = _abstracts

			searchResultList = append(searchResultList, _searchResult)
		}
	}

	fmt.Println("数据去重中")
	//去重
	var ChongFuList =make(map[int]int)
	{
		var _searchResultList []searchResult
		for i, i2 := range searchResultList {
			if ChongFuList[i2.Id]==0 {
				ChongFuList[i2.Id]=i+1
			}
		}
		//跳过重复数据
		for i, i2 := range searchResultList {
			if ChongFuList[i2.Id]==i+1 {
				_searchResultList= append(_searchResultList, i2)
			}
		}

		fmt.Println("重复数量", len(searchResultList)-len(_searchResultList))
		fmt.Println("有效数量",len(_searchResultList))
		searchResultList=_searchResultList
	}

	fmt.Println("生成文件中")
	var Phone []int=[]int{0,0,0}
	//创建Excel
	excel := excelize.NewFile()
	//创建一个工作表
	excel.NewSheet("数据")
	//设置单元格值
	if err := excel.SetCellValue("数据", fmt.Sprint("A", 1), "ID"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("B", 1), "Template"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("C", 1), "ImageUrl"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("D", 1), "标题"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("E", 1), "地址"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("F", 1), "最低价"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("G", 1), "平均价"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("H", 1), "纬度"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("I", 1), "经度"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("J", 1), "经度"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("K", 1), "显示类型"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("L", 1), "Avgscore"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("M", 1), "评论"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("N", 1),"历史优惠券计数"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("O", 1), "BackCateName"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("P", 1), "区域名称"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("Q", 1), "标签"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("R", 1), "Cate"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("S", 1), "RecentScreen"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("T", 1), "Abstracts"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("U", 1), "DangleAbstracts"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("V", 1), "TitleTags"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("W", 1), "IUrl"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("X", 1), "Deals"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("Y", 1), "Posdescr"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("Z", 1), "Ct_poi"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AA", 1), "Trace"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AB", 1), "LandmarkDistance"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AC", 1), "HasAds"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AD", 1), "AdsClickUrl"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AE", 1), "AdsShowUrl"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AF", 1), "Distance"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AG", 1), "CityId"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AH", 1), "City"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AI", 1), "号码"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AJ", 1), "Full"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据", fmt.Sprint("AK", 1),"号码类型"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	for i, i2 := range searchResultList {
		//设置单元格值
		if err := excel.SetCellValue("数据", fmt.Sprint("A", i+2), i2.Id); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("B", i+2), i2.Template); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("C", i+2), i2.ImageUrl); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("D", i+2), i2.Title); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("E", i+2), i2.Address); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("F", i+2), i2.Lowestprice); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("G", i+2), i2.Avgprice); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("H", i+2), i2.Latitude); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("I", i+2), i2.Longitude); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("J", i+2), i2.Longitude); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("K", i+2), i2.ShowType); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("L", i+2), i2.Avgscore); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("M", i+2), i2.Comments); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("N", i+2), i2.HistoryCouponCount); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("O", i+2), i2.BackCateName); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("P", i+2), i2.Areaname); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("Q", i+2), i2.Tag); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("R", i+2), i2.Cate); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("S", i+2), i2.RecentScreen); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("T", i+2), i2.Abstracts); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("U", i+2), i2.DangleAbstracts); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("V", i+2), i2.TitleTags); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("W", i+2), i2.IUrl); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("X", i+2), i2.Deals); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("Y", i+2), i2.Posdescr); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("Z", i+2), i2.Ct_poi); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AA", i+2), i2.Trace); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AB", i+2), i2.LandmarkDistance); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AC", i+2), i2.HasAds); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AD", i+2), i2.AdsClickUrl); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AE", i+2), i2.AdsShowUrl); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AF", i+2), i2.Distance); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AG", i+2), i2.CityId); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AH", i+2), i2.City); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AI", i+2), i2.Phone); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AJ", i+2), i2.Full); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}

		//识别手机号类型
		PhoneOn:=strings.Index(i2.Phone, "0771")
		PhoneType:="座机号"
		if PhoneOn==-1 {
			PhoneType="手机号"
			Phone[0]++
		}
		if PhoneOn!=-1 {
			Phone[1]++
		}
		if i2.Phone=="" {
			PhoneType="无号码"
			Phone[2]++
		}
		if err := excel.SetCellValue("数据", fmt.Sprint("AK", i+2),PhoneType); err != nil {
			fmt.Println(i, "设置内容失败", err.Error())
		}
	}


	//创建一个工作表
	excel.NewSheet("数据2")
	//创建一个工作表
	if err := excel.SetCellValue("数据2", fmt.Sprint("A", 1), "姓名"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	if err := excel.SetCellValue("数据2", fmt.Sprint("B", 1), "手机号"); err != nil {
		fmt.Println( "设置内容失败", err.Error())
	}
	//设置单元格值
	i:=0
	a:=""
	for _, i2 := range searchResultList {
		if i2.Phone=="" {
			continue
		}
		//识别手机号类型
		PhoneOn:=strings.Index(i2.Phone, "0771")
		if PhoneOn==-1 {
			//是否是多手机号类型
			PhoneArr:=strings.Split(i2.Phone, "/")
			if len(PhoneArr)>=2 {
				for i3, i4 := range PhoneArr {
					//设置单元格值
					if err := excel.SetCellValue("数据2", fmt.Sprint("A", i+2), fmt.Sprint(i2.Title,i3)); err != nil {
						fmt.Println(i, "设置内容失败", err.Error())
					}
					//设置单元格值
					if err := excel.SetCellValue("数据2", fmt.Sprint("B", i+2), i4); err != nil {
						fmt.Println(i, "设置内容失败", err.Error())
					}
					i++

					a=fmt.Sprint(a,
						"BEGIN:VCARD\n",
						"VERSION:2.1\n",
						"N:;",fmt.Sprint(i2.Title,i3),";;;\n",
						"FN:",fmt.Sprint(i2.Title,i3),"\n",
						"TEL;CELL:",i4,"\n",
						"END:VCARD\n",
					)

				}
			}else {
				//设置单元格值
				if err := excel.SetCellValue("数据2", fmt.Sprint("A", i+2), i2.Title); err != nil {
					fmt.Println(i, "设置内容失败", err.Error())
				}
				//设置单元格值
				if err := excel.SetCellValue("数据2", fmt.Sprint("B", i+2), i2.Phone); err != nil {
					fmt.Println(i, "设置内容失败", err.Error())
				}
				i++

				a=fmt.Sprint(a,
					"BEGIN:VCARD\n",
					"VERSION:2.1\n",
					"N:;",i2.Title,";;;\n",
					"FN:",i2.Title,"\n",
					"TEL;CELL:",i2.Phone,"\n",
					"END:VCARD\n",
				)
			}
		}
	}


	fmt.Println(ioutil.WriteFile("D:\\1755616537\\OneDrive\\桌面/kaka.txt", []byte(a),0666))





	fmt.Println("数理统计:手机号[",Phone[0],"]","座机号[",Phone[1],"]","无号码[",Phone[2],"]")
	//保存文件
	if err := excel.SaveAs("数据.xlsx"); err != nil {
		return err
	}

	return nil
}
