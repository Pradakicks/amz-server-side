package main

import (
	Amz "amz-monitors/amz"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Testingz")
	go Amz.NewAmzMonitor("B08G9NKWR9", `Wr0RFj%2BisP6XDoYX7Pa9cykvPmqXPMBz2xsQfKuEDYFmtOC5nhLwdsyiTNEO21AnYke1zdAAU3Vw211fnIBH%2FjR5Dx%2FNsPiNq9zeu6NRq1aFc0dqh5Tsuaa3w7GDdk0Fq57g3QXpbBmx2Ry%2FzqIsTA%3D%3D`, "GPU")
	go Amz.NewAmzMonitor("B08R81J62G", `dPIvIYRHvMgp7%2BKcLoVDXcE8qpp13AKCYwruZDDqR39vETkAKxHMkCUFsl3kyw%2F04zX2f4RprUI7BVuqL%2F4jozYXu6g5wHuk5JIw%2FG%2FRDF7zUXZsZREGGBtKeZC4Fh0llYVk1nywlWsUd%2FkIRglNAw%3D%3D`, "GPU")
	fmt.Scanln()
}
func f() {

	url := "https://www.amazon.com/gp/product/features/dp-fast-track/udp-ajax-handler/get-quantity-update-message.html?ie=UTF8"

	payload := strings.NewReader("asin=B08G9NKWR9&quantity=1&merchantId=A3KJ3YGUCB8ID4")

	req, _ := http.NewRequest("POST", url, payload)
	fmt.Println(url, payload)
	// req.Header.Add("cookie", `session-id=138-0322629-7694230; ubid-main=132-7199499-9402615; x-main="LHE6BnZDP41AAwul9mixw?nocA0On?RG9kXHFLa2KFh2Lr7ST3VYG26VwluEAQOU"; at-main=Atza|IwEBIAUgpGn-OsbKAh8xFHdemqgfqUzqwva175SZrkNcZM6mxiqHvt30-ZCpvd4IDtDfUS6zFffVj_NwJe0VZ2qEBGj9bDv6YrVbaDRdDBWA22hDuwbJMBc1Mzq7unkV6smc5haqXiYcvU_km9e03rtQ-DF8z9HXHO7M-GjkfdpdwR9is7Aq5WTt5x9afWEbmuKzCKY6rf5wvoyJjxuGaKS7hvXu; sess-at-main="FmAlrbPmV0W29DTdKAzieqpgm9IJ4U8tjfitT/uWrU4="; sst-main=Sst1|PQEbA5s4v4W3R1LgdJk50AzPCQlGLeUFsDl49VWuekI0WnDhKFzDoaM7KqvfftXIDAZx9OyhXWJ-P-c7MMYQ2F9zIxRGLPdzKZlT1Z8n6V-P9QjM90PI3ItFLDXcqvzhY156lC5nCxgCao4LqEKtzjfHuEMCJ2tRZH1tlHKLkrQDX9l2DEMtKXoMY8YVUrPiQ3aGGrZvCeqtD17gVEglFekkJAkj1fnInM0z2RguEQkjt2mHDa31IqsV2aXOt9XLaORJA5VLJfYpGhUWpqbkKTY-8ZpIBQ0u1FZ5P8bdvx-gzCc; lc-main=en_US; session-id-time=2082787201l; i18n-prefs=USD; csd-key=eyJ3YXNtVGVzdGVkIjp0cnVlLCJ3YXNtQ29tcGF0aWJsZSI6dHJ1ZSwid2ViQ3J5cHRvVGVzdGVkIjpmYWxzZSwidiI6MSwia2lkIjoiMDQzMTlhIiwia2V5IjoiRXBKQ2ZpY1pienlFazVQWWs3Y2tCZUM0S2NjTGl0RnRiSzlnUHFIdWEvMDAxQTdXeUQ4dmowS3pFczdCeldyRFRqamRSQmFpOVFiYnpvWlhVaHM2OG43ZW1MNVFPblJKQUFZM2JPd2pYN3crZHJUcitHSUxkb2ZlZDFkNHFQYUhaTFc3UXVJTmk1TkJ5aTVsdUdTV0VRbk02Nm9vVXRiRmkwRXU4QmVMNi9MYU05R2ZhTnVjRWdEcGVoQ0dIcitqUEE5OGtlRk14aEM2d1NNdnNoc0s5aHI3S1J3Rms1bUQreU5lb2V3OFV4a3ZuNTl2RUZFL0orckxtZ25XWXAxUWRZMlZ0czIxK2srQVZLanQ5azBYTGxYcERpeTF2dkVEbTg5OTdMMVRSZUVUOWt6TjhFNmhWa3VFU3lJTHB4MUtHUnE2bnNLTjJkQ2pJQTI2cUlqcGF3PT0ifQ==; s_vnum=2055786480819%26vn%3D1; aws_lang=en; s_fid=085844D920902910-0DB536AAEB684C2B; s_cc=true; aws-target-data=%7B%22support%22%3A%221%22%7D; aws-target-visitor-id=1623792154921-241206.34_0; aws-mkto-trk=id%3A112-TZM-766%26token%3A_mch-aws.amazon.com-1623792155111-63494; s_eVar60=ha%7Cacq_aws_takeover-1st-visit-free-tier%7Cawssm-evergreen-1st-visit%7Chero%7Cha_awssm-evergreen-1st-visit; aws-ubid-main=172-3048843-7260845; remember-account=false; regStatus=registered; awsc-color-theme=light; s_sq=%5B%5BB%5D%5D; s_vn=1655867630274%26vn%3D1; s_dslv=1624331690698; s_nr=1624331690710-Repeat; aws-userInfo=%7B%22arn%22%3A%22arn%3Aaws%3Aiam%3A%3A105303053417%3Aroot%22%2C%22alias%22%3A%22%22%2C%22username%22%3A%22PradaKicks%22%2C%22keybase%22%3A%22%22%2C%22issuer%22%3A%22http%3A%2F%2Fsignin.aws.amazon.com%2Fsignin%22%2C%22signinType%22%3A%22PUBLIC%22%7D; session-token="62wsGKl9AUzQdz2YMQt65TM676K1lDaYmNmiux7uT81mHmUqYheBsCdZ6cqjaW5LSlO3oOr7Av/OvO9J/auYimMe7wiX3xj00dRRi/o0npPr5R6SE5y6tJzU7uWD1Ds53h3jEl+XuTMWsVBBQ9MhH25MWF/B61LFOTEhLpdnOGsn5PRIA0ZdO/ToO137W84/w14s2nsDvKtg0lTosvDFog=="; csm-hit=tb:AXPSRB7CD5EH03HZ3X15+s-DXTY1YTEJ3KQH35VQ57E|1625157165330&t:1625157165330&adb:adblk_no`)
	req.Header.Add("authority", "www.amazon.com")
	req.Header.Add("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Add("dnt", "1")
	req.Header.Add("rtt", "50")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Add("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("accept", "text/html, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("downlink", "10")
	req.Header.Add("ect", "4g")
	req.Header.Add("origin", "https://www.amazon.com")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("accept-language", "en-US,en;q=0.9")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println()
	fmt.Println(string(body), res)

}
