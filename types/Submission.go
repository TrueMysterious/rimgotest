package types

type Submission struct {
	Id        string
	Title     string
	Link      string
	Cover     Media
	Points    int64
	Upvotes   int64
	Downvotes int64
	Comments  int64
	Views     int64
	IsAlbum   bool
}