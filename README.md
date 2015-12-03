# Mongolian version of Taobao.com

* taobao.com search
* On demand translation: mongol <-> chinese
* Facebook Integration
* Responsive, Mobile friendly web
* Fast product search (and syncing)
* Search Ranking


# Ерөнхий

bubble view, list view

* зүүн талын шүүлтүүр
* хуудаслалт, more
* metabase маягийн query interface хэрэгтэй


# Home хуудас

* сурталчилгаа
* шинэ бараанууд

# Барааны дэлгэрэнгүй мэдээлэл

* үндсэн дэлгэц зурах
* Facebook comment, like (facebook-с энэ мэдээллүүдийг дараа нь боловсруулж чадах уу?)
* Холбоотой мэдээллүүд, бараанууд


# Орчуулга - Google Translate
https://www.googleapis.com/language/translate/v2?key=INSERT-YOUR-KEY&source=en&target=de&q=Hello%20world
https://www.translate.com/
http://kupinatao.com/


1	淘宝网	taobao
Coding=gbk
Product name=/<title>(?P<this>.*)-(.*)<\/title>/U
Price=/<strong id=\"J_StrPrice\" >(?P<this>.*)<\/strong>/U
/<em class=\"price left\">(?P<this>.*)<\/em>/U



# TODO

* on-demand орчуулга шийдэх: google translate, translate.com. Монголоос хятад, хятадаас монгол руу
* taobao goods API судлах (open.taobao.com)
* taobao сайтаас бараа хайх, барааг өөрийн бааз руу sync хийх (эрэлттэй tag-аар, idle цагаар)
* sync хийх job-г hosting server (digital ocean) дээр тавьж ажиллуулж эхлэх, өгөгдөл цуглуулж эхлэх
* цугларсан бараа, каталоги дээр анализ хийх
* realtime image optimizer оруулах
* хайлт, filter хэсгийг оновчлох
* facebook intregration хийх: comment, like, view
* Hosting сервер дээр байршуулж тест хийх
* сагсанд бараа нэмэх хэсэг
* төлбөрийн шийдэл холбох. Эхэндээ банкны карт холбоно. Дараа нь хялбар шийдэл нэвтрүүлнэ. MostMoney, Credit Cards
* төлбөрийн маргаан, буцаалт шийдэх хэсэг
* бараа track хийх хэсэг: taobao logistic API
* монгол дотор хүргэлт хийх компанитай API тохирох (Монгол шуудан, TGB г.м)
* admin UI
* барааны үнийн margin засах
* нууцлалын нэмэлтүүд
* fiverrs - log design харах
* Loayalty - Candy, GG