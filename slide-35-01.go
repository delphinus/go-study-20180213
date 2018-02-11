package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

func top(c *gin.Context) {
	// エラー処理とかはまた後で考えます
	// r は $rhData 的なモノ
	r := getEntities()
	c.HTML(http.StatusOK, "top.tmpl", r)
}

/* top.tmpl
<body>
  <p>hoge: {{ .hoge.Name }}</p>
  <p>fuga: {{ .fuga.Name }}</p>
</body>
*/

func getEntities() map[string]interface{} {
	var wg sync.WaitGroup

	var r map[string]interface{}

	// goroutine 開始前に呼ぶ
	wg.Add(1)

	go func() {
		// エラー処理は後で考えます
		r["hoge"] = getHogeEntity()

		// goroutine 終了前に呼ぶ
		// 本来は defer を使うべき
		wg.Done()
	}()

	// goroutine 開始前に呼ぶ
	wg.Add(1)

	go func() {
		// エラー処理は後で考えます
		r["fuga"] = getFugaEntity()
		
		// goroutine 終了前に呼ぶ
		// 本来は defer を使うべき
		wg.Done()
	}()

	// 終了を待つ
	wg.Wait()

	return r
}

func getHogeEntity() *Hoge {
	...
}

func getFugaEntity() *Fuga {
	...
}nc getFugaEntity() *Fuga {
	...
}
