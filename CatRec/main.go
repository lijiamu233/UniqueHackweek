package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var Nowfile string
var result string

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	timeUnix := time.Now().Unix()

	filename := strconv.FormatInt(timeUnix, 10) + header.Filename

	out, err := os.Create("upload/" + filename)

	if err != nil {
		log.Fatal(err)
	}
	//defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	out.Close()
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
	Nowfile = "./upload/" + filename
	result = RunPy()
	_ = os.Remove(Nowfile)
}

func RunPy() string {
	args := []string{"run.py", Nowfile}
	out, _ := exec.Command("python", args...).Output()

	Num1 := string(out[len(out)-4])
	Num2 := string(out[len(out)-3])

	if Num1 == "1" {
		Num1 += Num2
	} else {
		Num1 = Num2
	}
	return Num1
}
func main() {
	CatsType := make(map[string]string)
	CatsType["0"] = "新加坡猫"
	CatsType["1"] = "豹猫"
	CatsType["2"] = "伯曼猫"
	CatsType["3"] = "孟买猫"
	CatsType["4"] = "英国短毛猫"
	CatsType["5"] = "埃及猫"
	CatsType["6"] = "缅因猫"
	CatsType["7"] = "波斯猫"
	CatsType["8"] = "雪鞋猫"
	CatsType["9"] = "俄罗斯蓝猫"
	CatsType["10"] = "泰国暹罗猫"
	CatsType["11"] = "斯芬克斯猫"
	Links := make(map[string]string)
	Links["0"] = "https://baike.baidu.com/item/%E6%96%B0%E5%8A%A0%E5%9D%A1%E7%8C%AB/671184?fr=aladdin#2"
	Links["1"] = "https://baike.baidu.com/item/%E8%B1%B9%E7%8C%AB/4921581?fr=aladdin"
	Links["2"] = "https://baike.baidu.com/item/%E4%BC%AF%E6%9B%BC%E7%8C%AB/641852?fr=aladdin"
	Links["3"] = "https://baike.baidu.com/item/%E5%AD%9F%E4%B9%B0%E7%8C%AB/4510178?fr=aladdin"
	Links["4"] = "https://baike.baidu.com/item/%E8%8B%B1%E5%9B%BD%E7%9F%AD%E6%AF%9B%E7%8C%AB/672846?fr=aladdin"
	Links["5"] = "https://baike.baidu.com/item/%E5%9F%83%E5%8F%8A%E7%8C%AB/386503?fr=aladdin"
	Links["6"] = "https://baike.baidu.com/item/%E7%BC%85%E5%9B%A0%E7%8C%AB/647590?fr=aladdin"
	Links["7"] = "https://baike.baidu.com/item/%E6%B3%A2%E6%96%AF%E7%8C%AB/585?fr=aladdin"
	Links["8"] = "https://baike.baidu.com/item/%E9%9B%AA%E9%9E%8B%E7%8C%AB/4513092"
	Links["9"] = "https://baike.baidu.com/item/%E4%BF%84%E7%BD%97%E6%96%AF%E8%93%9D%E7%8C%AB/643065?fr=aladdin"
	Links["10"] = "https://baike.baidu.com/item/%E6%9A%B9%E7%BD%97%E7%8C%AB/556082?fr=aladdin"
	Links["11"] = "https://baike.baidu.com/item/%E5%8A%A0%E6%8B%BF%E5%A4%A7%E6%97%A0%E6%AF%9B%E7%8C%AB/643507?fromtitle=%E6%96%AF%E8%8A%AC%E5%85%8B%E6%96%AF%E7%8C%AB&fromid=8028531&fr=aladdin"

	router := gin.Default()
	//router.LoadHTMLGlob("template/*")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "upload.html", gin.H{})
	// })
	router.POST("/upload", upload)
	//result := RunPy()
	//log.Fatal(result)
	router.GET("/run", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"CatType": CatsType[result],
			"Link":    Links[result],
		})
	})
	router.StaticFS("/file", http.Dir("upload"))
	router.Run()
}
