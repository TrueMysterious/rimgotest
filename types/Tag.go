package types

type Tag struct {
	Tag	        string
	Display			string
	Sort				string
	PostCount	  int64
	Posts				[]Submission
	Background  string
}