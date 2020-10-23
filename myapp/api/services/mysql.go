package services

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

const (
	host     = "mysql"
	database = "chat_data"
	user     = "root"
	password = "root"
)

/*資料總筆數 = total * num_times*/
var (
	total     int = 10000 // 每次插入幾筆資料
	num_times int = 400   // go func 會執行幾次
)

type Message struct {
	room_id  int
	type_tag string
	sender   string
	content  string
	hash     string
	read_tab int
}

type Job struct {
	db    *sql.DB
	ch    chan int
	total int
	n     int
}

/*插入大量資料進入資料庫*/
func Insertmysql() string {
	//建立連線
	var connString = fmt.Sprintf("%s:%s@tcp(%s)/%s?&charset=utf8mb4&collation=utf8mb4_unicode_ci", user, password, host, database)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("資料庫連線錯誤")
	}
	defer db.Close()

	//db設定
	db.SetConnMaxLifetime(time.Second * 500)
	db.SetMaxOpenConns(100)

	fmt.Println("====start=====")
	start := time.Now()

	//插入資料庫，每次最多20個goroutine
	jobChan := make(chan Job, 20)
	go worker(jobChan)

	//統計使用次數
	ch := make(chan int, num_times)
	for n := 0; n < num_times; n++ {
		job := Job{
			db:    db,
			ch:    ch,
			total: total,
			n:     n,
		}
		jobChan <- job
	}

	//signal 等待執行緒完成
	i := 0
	for {
		<-ch
		i++
		if i >= num_times {
			break
		}
	}

	end := time.Now()
	curr := end.Sub(start)
	fmt.Println("run time:", curr)

	return curr.String()
}

func worker(jobChan <-chan Job) {
	for job := range jobChan {
		go sqlExec(job)
	}
}

func sqlExec(job Job) {
	buf := make([]byte, 0, job.total+1)
	buf = append(buf, "INSERT INTO `messages` (`room_id`, `type`, `sender`, `content`, `hash`, `read_tab`) VALUES"...)

	for i := 0; i < job.total; i++ {
		tmp_name := RandName()

		msg := Message{
			room_id:  g.Intn(1000),
			type_tag: RandType_tag(),
			sender:   tmp_name,
			content:  uuid.NewV4().String(),
			hash:     Namehash(tmp_name),
			read_tab: rand.Intn(2),
		}
		str := fmt.Sprintf("(%d ,'%s' ,'%s', '%s', '%s', %d)", msg.room_id, msg.type_tag, msg.sender, msg.content, msg.hash, msg.read_tab)
		buf = append(buf, str+","...)
	}

	if len(buf) == 0 {
		return
	} else {
		fmt.Println("---開始" + strconv.Itoa(job.n) + "次插入total條！")
		_, err := job.db.Exec(strings.Trim(string(buf), ","))
		checkErr(err)
		fmt.Println("完成---" + strconv.Itoa(job.n) + "次插入total條！")
		job.ch <- 1
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
