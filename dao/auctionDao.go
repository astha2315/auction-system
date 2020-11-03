package dao

type faqDao struct{}

type FaqDaoIF interface {
	GetAllFaqs(faqType string) (int, error)
}

func FaqDao() FaqDaoIF {
	return &faqDao{}
}

func (self *faqDao) GetAllFaqs(faqType string) (int, error) {
	// var faq []*models.FaqObject
	// db, ConnectionErrs := db.SqlxConnect()
	// if ConnectionErrs != nil {
	// 	return 1, ConnectionErrs
	// }
	// var sqlStatement string
	// if faqType == "video" {
	// 	sqlStatement = `select question, answer from faq where faq_type=$1`
	// } else {
	// 	sqlStatement = `select faq.question, faq.answer, s.section_name, s.section_image
	// 	from faq  join faq_section faqs on faq.faq_id=faqs.faq_id
	// 	Join section  s on s.section_id=faqs.section_id where faq_type=$1`
	// }

	// err := db.Select(&faq, sqlStatement, faqType)
	// // var err error
	// if err != nil {
	// 	fmt.Println(err)
	// 	return faq, err
	// }
	// defer db.Close()
	// return faq, err

	return 1, nil
}
