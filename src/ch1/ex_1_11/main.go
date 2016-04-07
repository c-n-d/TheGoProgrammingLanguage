/*

$ cat data/ch1/sites.dat |\
tr '\n' ' ' |\
xargs go run src/ch1/ex_1_10/main.go

0.06s   84244 http://Bing.com
0.07s  153445 http://Imgur.com
0.16s       0 http://Instagram.com
0.18s   31122 http://Apple.com
0.20s  217419 http://Washingtonpost.com
0.41s  128481 http://Reddit.com
0.41s  165816 http://Ebay.com
0.49s  183899 http://Nytimes.com
0.52s   39259 http://Msn.com
0.52s   11777 http://Wordpress.com
0.53s   30490 http://Office.com
0.54s  372268 http://Weather.com
0.55s  125322 http://Aol.com
0.56s   95108 http://Cnn.com
0.58s  277573 http://Huffingtonpost.com
0.74s  258523 http://Twitter.com
0.79s   40929 http://Linkedin.com
0.85s  210762 http://Pinterest.com
0.86s  513133 http://Yahoo.com
0.87s  314310 http://Yelp.com
0.92s   60687 http://Blogspot.com
0.94s   67382 http://Target.com
0.96s   82759 http://Microsoft.com
1.10s   96366 http://Etsy.com
1.32s  164259 http://Imdb.com
1.57s   73325 http://Tumblr.com
1.85s  737114 http://Zillow.com
2.17s     949 http://Netflix.com
2.18s elapsed
*/

package main

import "fmt"

func main() {
    fmt.Printf("run:\n" +
        "\tcat data/ch1/sites.dat |\\\n" +
        "\ttr '\\n' ' ' |\\\n" +
        "\txargs go run ex_1_10.go\n")
}