# uzakel
Verimerkezleri Arası IP Engelleme Sistemi

![network](uzakel-draw.png)

## Amaç

Kendi IP adresini anons eden network sahipleri, kendilerine ait IP adreslerini sisteme bağlı diğer network sahipleri üzerinde null route yapabilmesini sağlamaktır. 

İş sürekliliğini üst düzeyde tutmak üzere planlanan yapıda, spoof dahil ağlar arası kirli trafik her iki network sahibinde engellenmiş olur.

## Nasıl Çalışır

RIPE Lir olan AS numarasına sahip ve kendi IP adresini anons eden kullanıcılar sisteme kayıt olabilirler. Kayıt esnasında AS numarası girilmesi yeterlidir. RIPE üzerinde kayıtlı **Lir Admin** kontak bilgilerine **secret key** ve kurulum dokümanları gönderilir.

Uygulama kendi lokal ağınıza kurulur. Dışarıdan gelen istekleri sizin belirlediğiniz telnet/ssh komut setleri ile core router cihazınıza istek göndererek null route işlemi yapar. Her marka ve modelde cihaz için kolay bir entegrasyona sahiptir. Template sistemi sayesinde gelen istekleri farklı ACL/AccsesList kurallarında işleyebilirsiniz. 

Sisteme dahil network sahipleri gelecek istekler için genel bir limit ve her kullanıcı için limit tanımı yapabilmektedir. Yapılan tüm işlemler için minimum ve maksimum **expire time** girişi yapılacaktır. Tüm bunlara network sahibi karar verebilmektedir.

AS Numarasına sahip networkler sadece kendi ağına ait IP adresleri için null route isteği gönderebileceği için herhangi bir güvenlik riski yoktur. IP adresi listesi global BGP routing tablosundan anlık olarak çekilecektir. 

Tüm uygulamalar açık kaynak kodlu şekilde sunulur.

## Geliştirme 

- [ ] Private Network
- [ ] Autodiscover
- [ ] Access Control List
- [ ] Packet Capture
- [ ] Execute Trigger
- [ ] RIPE Whois
- [ ] Rate Limiting
