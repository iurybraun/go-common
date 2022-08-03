/*
 * Copyright © 2016 Iury Braun
 * Copyright © 2017 Weyboo
 */

package utils
 
import (
    "log"
    "regexp"
    "crypto/tls"
    "gopkg.in/gomail.v2"
)

func MailSend(Host string, Port int, FromUsername, FromPassword, FromName, ToAddress, Subject, Content string) error {
    d := gomail.NewDialer(Host, Port, FromUsername, FromPassword)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    
    s, err := d.Dial()
    if err != nil {
        log.Println(err)
        return err
    }
    

    m := gomail.NewMessage()
    m.SetHeader("From", m.FormatAddress(FromUsername, FromName))
    m.SetAddressHeader("To", ToAddress, ToAddress)  // ToName
    m.SetHeader("Subject", Subject)
    m.SetBody("text/html", Content)

    if err := gomail.Send(s, m); err != nil {
        log.Println(err)
        return err
    }
    
    m.Reset()
    return nil
}

func ValidateEmail(Email string) bool {
    // Regular expression to validate email address
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    if re.MatchString(Email) == true {
        return true
    } else {
        return false
    }
}


// Efficiently send a customized email to a recipient.
/*func sendSimpleInformativeMail(ToAddress, FullName, Subject, Content string) error {
	
    d := gomail.NewDialer("smtp.weyboo.com", 587, "informativo@*.com", "infor2017")
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    
    s, err := d.Dial()
    if err != nil {
        //panic(err)
        //log.Println(err)
        return err
    }
    

    m := gomail.NewMessage()
    m.SetHeader("From", m.FormatAddress("informativo@*.com", "Renove Life | Vida Saudável"))
    m.SetAddressHeader("To", ToAddress, FullName)
    m.SetHeader("Subject", Subject)
    m.SetBody("text/html", Content)

    if err := gomail.Send(s, m); err != nil {
        //log.Printf("Could not send email to %q: %v", To, err)
        return err
    }
    
    m.Reset()
    return nil
    
}*/
