package model

// User struct
type User struct {
  Username    string  `json:"username"`
  Firstname   string  `json:"firstname"`
  Lastname    string  `json:"lastname"`
  Password    string  `json:"password"`
  Token       string  `json:"token"`
}

// Response Struct
type Response struct {
  Result  interface{}   `json:"result"`
  Error   string        `json:"error"`
}

// Application Struct
type Application struct {
  ID                  string  `bson:"_id" json:"id"`
  Username            string  `json:"username"`
  CompanyName         string  `json:"companyname"`
  Status              int     `json:"status"`
  JobLink             string  `json:"joblink"`
  Recruiter1FirstName string  `json:"recruiter1firstname"`
  Recruiter1LastName  string  `json:"recruiter1lastname"`
  Recruiter1Email     string  `json:"recruiter1email"`
  Recruiter1Phone     string  `json:"recruiter1phone"`
  Referred            bool    `json:"referred"`
  LastContact         string  `json:"lastcontact"`
  Notes               string  `json:"notes"`
}
