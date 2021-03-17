
Jam 15:00
usecase yang udah kelar (but not tested)
  BuatArisan
  MulaiArisan
  TagihSetoran
  BayarSetoran
  SetoranKadaluwarsa
  KocokUndian

usecase yang sedang dikerjain
  UndangPeserta ?
  JawabUndangan ?

yang belum dikerjain lainnya
  Bikin koneksi ke DB
  Bikin API
  bikin simple BO just for read the result
    akan nampilin
    - tabel arisan (nampilin seluruh arisan yg sedang jalan)
    - tabel peserta
    - tabel akunting (utang, piutang tiap peserta)
    - tabel undian
    - tabel tagihan

  mock server DANA
    API createorder
    webhook call ke ArisanSystem
    API topup
    innerAccount: Saldo Global







Proses pendaftaran peserta arisan

  Admin yang mengundang peserta
    Dengan memberikan link yang akan diklik oleh peserta?
    Dengan memberikan code yang akan diinput peserta?
    code/linknya global atau spesifik untuk user tertentu?
    code/linknya sudah diassign ke jumlah slot tertentu?
    peserta memilih sendiri jumlah slot dan group atau ditentukan admin?

  Peserta yang request ke Admin
    via japri, lalu admin mendaftarkan secara manual dengan no telepon
    admin mendaftarkan sesuai dengan jumlah slot yang diminta peserta



Prinsip dan fakta arisan
  Orang yang jadi peserta haruslah orang yang kredibel
  Jangan coba2 ngundang orang yang tidak dikenal
  salah satu alasan orang mengikuti arisan adalah sebagai alternatif cicilan
  arisan membuka pintu utang
  orang yang terakhir menang undian = menabung
  orang yang pertama menang undian = jackpot!
  yang butuh arisan itu adalah peserta, bukan admin
  admin haya membantu saja. Apakah hanya sekedar membantu? ada feenya?
  admin harus kenal baik peserta dan yang mewakilinya
  admin harus mewawancarai peserta sebelum didaftarkan ke arisan
  peserta harus memberikan identitas pribadi, nomor yg bisa dihubungi dan kontak perwakilan
  Peserta boleh ikut lebih dari satu arisan dalam satu waktu bersamaan?


SingleSlot
  1 peserta berapa rupiah?

MultiSlot
  1 slot berapa rupiah?
  1 peserta maksimal berapa slot?

GroupSlot
  Mau bikin berapa group
  1 group berapa slot?
  1 slot berapa rupiah?
  1 peserta maksimal berapa slot?

BuatArisan
  Admin menginisiasi Arisan
  Tabel yg berpengaruh
    Arisan
      Nama Arisan : ?
      Deskripsi : ?
      Mau bikin berapa group? (pilihannya cuman 0 atau > 1)
      kalo > 1 maka tanya lagi :
        1 group berapa slot?
      1 slot berapa rupiah?
      1 peserta maksimal berapa slot?
      State : Terbuka
      Type : SingleSlot
      Putaran : 1
      TotalPutaran : 0
    Peserta // utk mendaftarkan admin
      NamaAdmin : ?
      JumlahSlot : 1
    SlotPeserta // utk mendaftarkan slot admin
    Undian
      Putaran : 1
      TanggalTagihan : ?
      TanggalUndian : ?
      BiayaAdmin : 0
      BiayaArisan : ?




UndangPeserta
  client sudah jadi admin
  client klik tombol undang peserta
  client akan masukkan nama peserta
  namanya harus ditrim jangan ada spasi, lowercase
  submit




getarisan/arya
getarisan/mirza
getarisan/vira




Peserta buka aplikasi Arisan
  client akan panggil getarisan/:pesertaID
  server akan query utk peserta tsb apakah ada didalam dalam suatu arisan?
  kalo gak ada didalam arisan,
    artinya ya gak diundang, tapi dia bisa bikin arisan
    server akan kasi respon "kamu gak diundang siapapun, tapi kamu bisa buat arisan"
    server akan clear cache arisanID pada user tsb
  kalo ada didalam arisan,
    server akan beri arisanID
    apakah arisan tsb yang masih terbuka atau sudah dimulai?
    - kalo masih terbuka (masih menerima peserta),
      apakah user sudah join?
      - kalo iya,
        server beri respon "silakan menunggu" dan list peserta lainnya
        client tampilkan halaman undangan dan menampilkan list peserta lainnya
      - kalo tidak
        server beri respon "silakan memilih join atau tidak"
        client tampilkan halaman undangan, menampilkan list peserta lainnya beserta tombol TERIMA/TIDAK undangan
        masuk ke usecase JawabUndangan
    - kalo sudah sudah dimulai
      server beri respon list undian dan list tagihan tiap peserta
      client tampilkan halaman undian dan menampilkan list peserta lainnya




JawabUndangan ?
  client akan panggil joinarisan/:arisanID/:pesertaID
  table yg berpengaruh
    Peserta
      Status
    Arisan
      JumlahPeserta





MulaiArisan
  tidak bisa menerima peserta lagi
  Tabel yg berpengaruh
    Arisan
      State : Mulai
      TotalPutaran :
        Jika GroupSlot : Jumlah Total Group
        Else : Jumlah Total Slot
    Tagihan : buat object tagihan untuk setiap peserta
      Status : belumDitagih






TagihSetoran
  Diingatkan secara otomatis oleh system (Scheduler)
  Tabel yg berpengaruh
    Undian:
      tanggal_undian : ambil sebagai tanggal expired
    Tagihan : membuat tagihan untuk tiap peserta
      acquirementID : dari DANA
      nominalToDANA : NilaiSlotPerPeserta + BiayaAdmin + BiayaArisan
      status : menunggu




BayarSetoran
  Ditrigger sama DANA
  Nilai yang disetor adalah sama dengan nominal yang ditagih
  Tabel yg berpengaruh
    Tagihan
      tanggal_pelunasan : tanggal bayar
      status : lunas
    Jurnal:
      Tanggal: tanggalBayar
    AkunBalance:
      Journal : Tambah Modal
      AkunType : Harta +, DEBET
      AkunType : Modal +, CREDIT
    AkunBalance:
      Journal : Setor Arisan
      AkunType : Harta -, CREDIT
      AkunType : Piutang +, DEBET
      AkunType : Biaya +, DEBET
    Peserta
      cendol : +1




SetoranKadaluwarsa
  Ditrigger sama DANA
  Tabel yg berpengaruh
    Tagihan
      status : kadaluwarsa
    Peserta
      bata : +1
    Ditalangi Admin? Buseet daaaah :'(





KocokUndian
  Ditrigger oleh admin (manual) / scheduler (otomatis)
  Nominal undian adalah totalSlot x nilaiPerSlot x jumlahSlotYangDiambilPeserta
  Tabel yg berpengaruh
    SlotPeserta
      TanggalMenang : tanggal hari itu
    AkunBalance:
      Journal : Menang Undian Arisan
      AkunType : Harta, DEBET
      AkunType : Utang, CREDIT
    Arisan:
      if putaranKe+1 < totalPutaran :
        putaranKe = putaranKe+1
      else
        status : Selesai
    Undian : hanya jika belum selesai
      PutaranKe = putaranKe+1
      TanggalTagihan : TanggalTagihan sebelumnya tapi dibulan berikutnya
      TanggalUndian : TanggalUndian sebelumnya tapi dibulan berikutnya
      BiayaAdmin : 0
      BiayaArisan : ?







Undian Arisan Putaran 1

Setoran perorang = 500

A
B
C
D

--------------------------
A
TambahModal
Harta +, DEBET , 500, 500
Modal +, CREDIT, 500, 500

SetorTagihan
Harta -  , CREDIT, 500, 0
Piutang +, DEBET , 500, 500

--------------------------
B
TambahModal
Harta +, DEBET , 500, 500
Modal +, CREDIT, 500, 500

SetorTagihan
Harta -  , CREDIT, 500, 0
Piutang +, DEBET , 500, 500

--------------------------
C
TambahModal
Harta +, DEBET , 500, 500
Modal +, CREDIT, 500, 500

SetorTagihan
Harta -  , CREDIT, 500, 0
Piutang +, DEBET , 500, 500

--------------------------
D
TambahModal
Harta +, DEBET , 500, 500
Modal +, CREDIT, 500, 500

SetorTagihan
Harta -  , CREDIT, 500, 0
Piutang +, DEBET , 500, 500

--------------------------

A
MenangUndian
Harta +,  DEBET , 2000, 2000
Utang +,  CREDIT, 2000, 2000

Penyesuaian
Utang -,  DEBET, 500, 1500
Piutang -, CREDIT, 0, 0



