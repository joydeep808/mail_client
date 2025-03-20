package email

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"html/template"

	"gopkg.in/gomail.v2"
)




var (
	emailDialer *gomail.Dialer
	once         sync.Once
)


func InitializeEmailClient(){
	print("Initializing email client")
	once.Do(func ()  {
		emailDialer =  gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_ID"), os.Getenv("EMAIL_PASSWORD"))
		emailDialer.TLSConfig = nil
	})
}



func SendEmail(to string, subject string, body string) error {
	if emailDialer == nil {
			print("Email client not initialized")
			return fmt.Errorf("email client not initialized")
	}
	fromAddress := fmt.Sprintf("%s <%s>", os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_ID"))

	message := gomail.NewMessage()
	message.SetHeader("From", fromAddress)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	
	// Send the email using the initialized dialer
	if err := emailDialer.DialAndSend(message); err != nil {
			print("Error while sending email" , err.Error())
			return err
	}
	print("Email send successfully done!")
	
	return nil
}




type WelcomeData struct {
	Name        string
	WebsiteLink string
}

const welcomeTemplate = `
<h1 style="text-align:center; color:black;">Hello {{.Name}}</h1>

<h3 style="color:#525f7f">Welcome to our website {{.Name}}! We're very excited to have you on board</h3>
<hr style="width:100%;border:none;border-top:1px solid #eaeaea;border-color:#e6ebf1;margin:20px 0" />
<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">Thanks for submitting your account information. You're now ready to manage your expenses smoothly!</p>

<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">You can access Your Account from the Website Fully</p>

<a href="{{.WebsiteLink}}" style="background-color:#656ee8;border-radius:5px;color:#fff;font-size:16px;font-weight:bold;text-decoration:none;text-align:center;display:inline-block;width:100%;padding:10px 10px 10px 10px;line-height:100%;max-width:100%" target="_blank">
  <span style="max-width:95%;display:inline-block;line-height:120%;">View Your Account</span>
</a>
<hr style="width:100%;border:none;border-top:1px solid #eaeaea;border-color:#e6ebf1;margin:20px 0" />

<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">Once you're verified, you can start</p>
<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">— Joydeep Debnath</p>
<hr style="width:100%;border:none;border-top:1px solid #eaeaea;border-color:#e6ebf1;margin:20px 0" />
<p style="font-size:12px;line-height:16px;margin:16px 0;color:#8898aa">Joydeep Debnath Agartala Tripura West</p>
<h3>Need help, or have questions? Just reply to this email, we'd love to help.</h3>
<p>Thank You</p>
`

	func GenerateWelcomeEmail(name string , website string) string {
		// Prepare the data for template substitution
		data := WelcomeData{
			Name:        name,  // replace with actual name
			WebsiteLink: website, // replace with actual link
		}
	
		// Create a new template and parse the string
		tmpl, err := template.New("welcome").Parse(welcomeTemplate)
		if err != nil {
			fmt.Println("Error parsing template:", err)
		}
	
		// Create a buffer to store the resulting string
		var result bytes.Buffer
	
		// Execute the template and write to the buffer
		err = tmpl.Execute(&result, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	
		// Get the string from the buffer
		outputString := result.String()
		return outputString
	}
func WelcomeMailTemplate (){
	const welcomeTemplate = `
<h1 style="text-align:center; color:black;">Hello {{.Name}}</h1>

<h3 style="color:#525f7f">Welcome to our website {{.Name}}! We're very excited to have you on board</h3>
<hr style="width:100%;border:none;border-top:1px solid #eaeaea;border-color:#e6ebf1;margin:20px 0" />
<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">Thanks for submitting your account information. You're now ready to manage your expenses smoothly!</p>

<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">You can access Your Account from the Website Fully</p>

<a href="{{.WebsiteLink}}" style="background-color:#656ee8;border-radius:5px;color:#fff;font-size:16px;font-weight:bold;text-decoration:none;text-align:center;display:inline-block;width:100%;padding:10px 10px 10px 10px;line-height:100%;max-width:100%" target="_blank">
  <span style="max-width:95%;display:inline-block;line-height:120%;">View Your Account</span>
</a>
<hr style="width:100%;border:none;border-top:1px solid #eaeaea;border-color:#e6ebf1;margin:20px 0" />

<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">Once you're verified, you can start</p>
<p style="font-size:16px;line-height:24px;margin:16px 0;color:#525f7f;text-align:left">— Joydeep Debnath</p>
<hr style="width:100%;border:none;border-top:1px solid #eaeaea;border-color:#e6ebf1;margin:20px 0" />
<p style="font-size:12px;line-height:16px;margin:16px 0;color:#8898aa">Joydeep Debnath Agartala Tripura West</p>
<h3>Need help, or have questions? Just reply to this email, we'd love to help.</h3>
<p>Thank You</p>
`

	// Prepare the data for template substitution
	data := WelcomeData{
		Name:        "John Doe",  // replace with actual name
		WebsiteLink: "http://example.com", // replace with actual link
	}

	// Create a new template and parse the string
	tmpl, err := template.New("welcome").Parse(welcomeTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Create a buffer to store the resulting string
	var result bytes.Buffer

	// Execute the template and write to the buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// Get the string from the buffer
	outputString := result.String()

	// Print the resulting string (or you can return it as needed)
	fmt.Println(outputString)
}