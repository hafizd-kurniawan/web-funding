sturktur folder
internal/
identity/
domain/
application/
infrastructure/
delivery/

campaign/
domain/
application/
infrastructure/
delivery/

funding/
domain/
application/
infrastructure/
delivery/

payment/
infrastructure/ // pure adapter (midtrans)

shared/
eventbus/
outbox/ (optional)

sebelum create harus analisa ini

1. Bounded context
   agregat root: user
   VO: email, password
   invariant: email, password
   domain behavior
   semua perubahan domain harus dibuat dengan domain behavior
   ngk ada repository di domain
   janganfokus: jwt, domain lain
   domainLain hanya pakai USERID domain

arti:
VO: value object
invariant: aturan yg hanya bisa diubah di domain

contoh salah mencampurkan domain
type Campaign struct {
User User.User
}

contoh benar agregat root
type Campaign struct {
id CampaignID
ownerID UserID
goalAmount Money
currentAmount Money
backerCount int
}
Campaign TIDAK TAHU user itu apa, cuma UserID

contoh lain agregat root terhubung sama domain lain
type Transaction struct {
id TransactionID
userID UserID
campaignID CampaignID
amount Money
status Status
}

```

func (t *Transaction) MarkPaid() TransactionPaid {
  t.status = Paid
  return TransactionPaid{
    TransactionID: t.id,
    CampaignID:    t.campaignID,
    Amount:        t.amount,
  }
}
```

efek diatas:

1. ubah state agregat
2. keluarkan fakta/evnt
3. ngk publis apapun

kemudian flow dari setelah itu akan seperti ini ketika ada event:

```


type MarkTransactionPaid struct {
	repo       TransactionRepository
	outboxRepo outbox.Repository
}

func (uc *MarkTransactionPaid) Execute(
	txID TransactionID,
) error {

	// 1. load aggregate
	tx, err := uc.repo.FindByID(txID)
	if err != nil {
		return err
	}

	// 2. call domain behavior
	event, err := tx.MarkPaid()
	if err != nil {
		return err
	}

	// 3. save aggregate
	if err := uc.repo.Save(tx); err != nil {
		return err
	}

	// 4. simpan EVENT ke outbox
	payload, eventType := Serialize(event)

	outboxEvent := &outbox.OutboxEvent{
		AggregateType: "Transaction",
		AggregateID:   int64(txID),
		EventType:     eventType,
		Payload:       payload,
	}

	return uc.outboxRepo.Save(outboxEvent)
}
```

KENAPA TIDAK LANGSUNG bus.Publish() DI SINI?
Karena:
kalau publish langsung
lalu crash
event hilang
Makanya:
simpan ke outbox dulu
worker yang publish
📌 Application layer = tempat ORKESTRASI

setelah itu flow nya seperti ini
outbox table
↓
outbox worker
↓
EventBus.Publish(TransactionPaid)
apa itu event?
“Ada transaksi yang SUDAH DIBAYAR.”
📌 Event BUKAN perintah
📌 Event TIDAK bilang “tolong update campaign”
📌 Event cuma bilang: INI TERJADI

↓
Campaign handler

HTTP Handler
↓
Application Use Case
↓
Domain (User, Password, Rule)
↑
Application bikin TOKEN
↓
Handler kirim response
