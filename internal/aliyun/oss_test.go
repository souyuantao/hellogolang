package aliyun

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	. "github.com/smartystreets/goconvey/convey"
)

var client *oss.Client

func init() {
	var err error
	endpoint, accessKeyID, accessKeySecret, err := LoadOSSConfig()
	if err != nil {
		panic(err)
	}
	client, err = oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}
}

func TestBucket(t *testing.T) {
	Convey("test list bucket", t, func() {
		res, err := client.ListBuckets()
		So(err, ShouldBeNil)

		for _, bucket := range res.Buckets {
			fmt.Println(bucket.Name)
		}
	})
}

func TestObject(t *testing.T) {
	Convey("test object", t, func() {
		bucket, err := client.Bucket("hatlonely-test-bucket")
		So(err, ShouldBeNil)

		Convey("put object", func() {
			err := bucket.PutObjectFromFile("oss_test.go", "oss_test.go")
			So(err, ShouldBeNil)
		})

		Convey("get object", func() {
			fp, err := bucket.GetObject("oss_test.go")
			So(err, ShouldBeNil)
			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				Println(scanner.Text())
			}
		})

		Convey("list object", func() {
			res, err := bucket.ListObjects()
			So(err, ShouldBeNil)

			for _, obj := range res.Objects {
				fmt.Println(obj.Key)
			}
		})

		Convey("del object", func() {
			err := bucket.DeleteObject("oss_test.go")
			So(err, ShouldBeNil)
		})
	})
}